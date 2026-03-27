package system

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	commonResp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/request"
	"github.com/goccy/go-json"
)

func (s *AutoCodeService) LLMAuto(ctx context.Context, llm common.JSONMap) (interface{}, error) {
	path, err := buildLLMAutoPath(llm)
	if err != nil {
		return nil, err
	}

	res, err := request.HttpRequestWithContextAndTimeout(
		ctx,
		path,
		http.MethodPost,
		nil,
		nil,
		llm,
	)
	if err != nil {
		return nil, fmt.Errorf("调用上游大模型服务失败: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取大模型响应失败: %w", err)
	}

	bodyPreview := previewResponseBody(body)
	contentType := res.Header.Get("Content-Type")
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("上游大模型服务返回非 2xx: status=%d content-type=%s body=%s", res.StatusCode, contentType, bodyPreview)
	}

	var resStruct commonResp.Response
	if err = json.Unmarshal(body, &resStruct); err != nil {
		return nil, fmt.Errorf("解析大模型响应失败: status=%d content-type=%s body=%s err=%w", res.StatusCode, contentType, bodyPreview, err)
	}

	if resStruct.Code != commonResp.SUCCESS {
		return nil, fmt.Errorf("大模型服务返回业务错误: code=%d msg=%s body=%s", resStruct.Code, resStruct.Msg, bodyPreview)
	}

	return resStruct.Data, nil
}

func (s *AutoCodeService) LLMAutoStream(ctx context.Context, llm common.JSONMap) (*http.Response, error) {
	path, err := buildLLMAutoPath(llm)
	if err != nil {
		return nil, err
	}

	payload := cloneLLMAutoJSONMap(llm)
	responseMode := strings.ToLower(strings.TrimSpace(fmt.Sprintf("%v", payload["response_mode"])))
	if responseMode == "" {
		payload["response_mode"] = "streaming"
	}

	res, err := request.HttpRequestWithContextAndTimeout(
		ctx,
		path,
		http.MethodPost,
		map[string]string{
			"Accept":          "text/event-stream",
			"Accept-Encoding": "identity", // 禁止 gzip，避免 SSE 流被压缩导致缓冲卡住
			"Cache-Control":   "no-cache",
		},
		nil,
		payload,
		-1, // 不设置 client.Timeout，SSE 流的生命周期由 ctx 控制
	)
	if err != nil {
		return nil, fmt.Errorf("调用上游大模型流式服务失败: %w", err)
	}
	return res, nil
}

func buildLLMAutoPath(llm common.JSONMap) (string, error) {
	if global.GVA_CONFIG.AutoCode.AiPath == "" {
		return "", errors.New("请先前往插件市场个人中心获取 AiPath 并填写到 config.yaml 中")
	}

	mode := strings.TrimSpace(fmt.Sprintf("%v", llm["mode"]))
	if mode == "" {
		return "", errors.New("llmAuto 缺少 mode 参数")
	}

	return strings.ReplaceAll(global.GVA_CONFIG.AutoCode.AiPath, "{FUNC}", mode), nil
}

func cloneLLMAutoJSONMap(src common.JSONMap) common.JSONMap {
	dst := make(common.JSONMap, len(src))
	for key, value := range src {
		dst[key] = value
	}
	return dst
}

func previewResponseBody(body []byte) string {
	text := strings.TrimSpace(string(body))
	text = strings.ReplaceAll(text, "\r", " ")
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.Join(strings.Fields(text), " ")
	if text == "" {
		return "<empty>"
	}
	runes := []rune(text)
	if len(runes) > 300 {
		return string(runes[:300]) + "..."
	}
	return text
}
