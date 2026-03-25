package system

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

func (autoApi *AutoCodeApi) LLMAutoSSE(c *gin.Context) {
	var llm common.JSONMap
	if err := c.ShouldBindJSON(&llm); err != nil {
		global.GVA_LOG.Error("LLMAutoSSE 参数绑定失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if llm == nil {
		llm = common.JSONMap{}
	}
	llm["response_mode"] = "streaming"
	global.GVA_LOG.Info("LLMAutoSSE 收到请求", zap.Any("mode", llm["mode"]))

	if err := autoApi.streamLLMAsSSE(c, llm); err != nil {
		global.GVA_LOG.Error("大模型 SSE 代理失败!", zap.Error(err))
		if c.Writer.Written() {
			writeLLMStreamError(c, err)
			return
		}
		response.FailWithMessage(err.Error(), c)
	}
}

func (autoApi *AutoCodeApi) streamLLMAsSSE(c *gin.Context, llm common.JSONMap) error {
	res, err := autoCodeService.LLMAutoStream(c.Request.Context(), llm)
	if err != nil {
		return fmt.Errorf("调用上游大模型失败: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return fmt.Errorf("上游大模型流式服务返回非 2xx: status=%d content-type=%s read-body-err=%w", res.StatusCode, res.Header.Get("Content-Type"), readErr)
		}
		return fmt.Errorf("上游大模型流式服务返回非 2xx: status=%d content-type=%s body=%s", res.StatusCode, res.Header.Get("Content-Type"), previewResponseBody(body))
	}

	ct := res.Header.Get("Content-Type")
	global.GVA_LOG.Info("LLMAutoSSE 上游返回成功，开始 SSE 流式转发",
		zap.Int("status", res.StatusCode),
		zap.String("content-type", ct))

	// 如果上游返回的不是 SSE 流（可能是 blocking 模式返回的 JSON），直接读取并转发
	if !strings.Contains(ct, "text/event-stream") && !strings.Contains(ct, "text/plain") {
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return fmt.Errorf("读取上游非流式响应失败: %w", readErr)
		}
		global.GVA_LOG.Warn("LLMAutoSSE 上游返回非 SSE 流，Content-Type: "+ct+", 将以单次事件转发",
			zap.String("body_preview", previewResponseBody(body)))

		flusher, ok := c.Writer.(http.Flusher)
		if !ok {
			return errors.New("当前响应不支持流式输出")
		}
		prepareSSEHeaders(c)
		c.Status(http.StatusOK)

		var payload any
		if err := json.Unmarshal(body, &payload); err != nil {
			payload = string(body)
		}
		if err := renderSSE(c, sse.Event{Event: "message", Data: payload}); err != nil {
			return err
		}
		if err := renderSSE(c, sse.Event{Event: "done", Data: gin.H{"done": true}}); err != nil {
			return err
		}
		flusher.Flush()
		return nil
	}

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		return errors.New("当前响应不支持流式输出")
	}

	prepareSSEHeaders(c)
	c.Status(http.StatusOK)
	flusher.Flush()

	reader := bufio.NewReader(res.Body)
	lines := make([]string, 0, 8)
	blockCount := 0

	global.GVA_LOG.Info("LLMAutoSSE 开始读取上游流数据...")

	for {
		global.GVA_LOG.Debug("LLMAutoSSE 等待读取下一行...")
		line, readErr := reader.ReadString('\n')
		if readErr != nil && !errors.Is(readErr, io.EOF) {
			global.GVA_LOG.Error("LLMAutoSSE 读取上游流失败", zap.Int("已转发块数", blockCount), zap.Error(readErr))
			return fmt.Errorf("读取上游流式响应失败: %w", readErr)
		}

		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			if len(lines) > 0 {
				blockCount++
				if blockCount <= 3 {
					global.GVA_LOG.Debug("LLMAutoSSE 转发 SSE 块", zap.Int("block", blockCount), zap.Strings("lines", lines))
				}
			}
			if err := emitSSEBlock(c, lines); err != nil {
				return err
			}
			lines = lines[:0]
		} else {
			lines = append(lines, line)
		}

		if errors.Is(readErr, io.EOF) {
			if err := emitSSEBlock(c, lines); err != nil {
				return err
			}
			if err := renderSSE(c, sse.Event{
				Event: "done",
				Data:  gin.H{"done": true},
			}); err != nil {
				return err
			}
			flusher.Flush()
			global.GVA_LOG.Info("LLMAutoSSE 流式转发完成", zap.Int("总块数", blockCount))
			return nil
		}
	}
}

func prepareSSEHeaders(c *gin.Context) {
	header := c.Writer.Header()
	header.Set("Content-Type", "text/event-stream; charset=utf-8")
	header.Set("Cache-Control", "no-cache, no-transform")
	header.Set("Connection", "keep-alive")
	header.Set("X-Accel-Buffering", "no")
}

func emitSSEBlock(c *gin.Context, lines []string) error {
	if len(lines) == 0 {
		return nil
	}

	eventName := "message"
	eventID := ""
	dataLines := make([]string, 0, len(lines))

	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "event:"):
			eventName = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
		case strings.HasPrefix(line, "id:"):
			eventID = strings.TrimSpace(strings.TrimPrefix(line, "id:"))
		case strings.HasPrefix(line, "data:"):
			dataLines = append(dataLines, strings.TrimSpace(strings.TrimPrefix(line, "data:")))
		}
	}

	rawData := strings.TrimSpace(strings.Join(dataLines, "\n"))
	if rawData == "" {
		return nil
	}
	if rawData == "[DONE]" {
		return renderSSE(c, sse.Event{
			Id:    eventID,
			Event: "done",
			Data:  gin.H{"done": true},
		})
	}

	var payload interface{}
	if err := json.Unmarshal([]byte(rawData), &payload); err != nil {
		payload = rawData
	}

	return renderSSE(c, sse.Event{
		Id:    eventID,
		Event: eventName,
		Data:  payload,
	})
}

func renderSSE(c *gin.Context, event sse.Event) error {
	if err := event.Render(c.Writer); err != nil {
		return fmt.Errorf("写入 SSE 事件失败: %w", err)
	}
	if flusher, ok := c.Writer.(http.Flusher); ok {
		flusher.Flush()
	}
	return nil
}
