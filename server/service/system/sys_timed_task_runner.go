// server/service/system/sys_timed_task_runner.go
package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/datascope"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/sse"
)

// 超时用 var 而非 const: 单测需收窄
var (
	defaultMethodTimeout = 5 * time.Minute
	defaultHTTPTimeout   = 30 * time.Second
)

const (
	maxHTTPRespBytes = 1 << 20 // HTTP 响应体读取上限 1MB
	maxLogTextLen    = 4000    // error/output 落库截断长度
	alertAuthorityID = 888     // 失败告警接收角色
	alertEventName   = "timedTask:alert"
)

// errTaskTimeout 超时哨兵: Runner 据此把状态记为 timeout 而非 fail
var errTaskTimeout = errors.New("任务执行超时")

func truncateText(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "...(截断)"
}

// RunTask 统一执行入口(自动调度与手动触发共用):
// panic 兜底、起止/耗时/状态/错误落 sys_timed_task_logs、失败经 SSE 告警。
// 阻塞执行; 调度器回调与手动触发均应在独立 goroutine 中调用。
func (s *TimedTaskService) RunTask(t system.SysTimedTask, trigger string) {
	started := time.Now()
	var output string
	var runErr error
	switch t.ExecutorType {
	case system.TimedTaskExecutorMethod:
		output, runErr = s.runMethod(t)
	case system.TimedTaskExecutorHTTP:
		output, runErr = s.runHTTP(t)
	default:
		runErr = fmt.Errorf("未知执行器类型: %s", t.ExecutorType)
	}
	finished := time.Now()

	status := system.TimedTaskStatusSuccess
	errMsg := ""
	if runErr != nil {
		if errors.Is(runErr, errTaskTimeout) {
			status = system.TimedTaskStatusTimeout
		} else {
			status = system.TimedTaskStatusFail
		}
		errMsg = truncateText(runErr.Error(), maxLogTextLen)
	}

	logRow := system.SysTimedTaskLog{
		TaskId:      t.ID,
		TaskName:    t.Name,
		TriggerType: trigger,
		StartedAt:   started,
		FinishedAt:  finished,
		DurationMs:  finished.Sub(started).Milliseconds(),
		Status:      status,
		ErrorMsg:    errMsg,
		Output:      truncateText(output, maxLogTextLen),
	}
	ctx := datascope.WithSystem(context.Background())
	if err := global.GVA_DB.WithContext(ctx).Create(&logRow).Error; err != nil {
		logger.Bg().Mod("timedTask").Err(err).Error("定时任务执行日志落库失败: " + t.Name)
	}
	if runErr != nil {
		logger.Bg().Mod("timedTask").Err(runErr).Error("定时任务执行失败: " + t.Name)
		s.alertFailure(t, errMsg)
	}
}

// runMethod 执行已注册本体方法。
// 超时语义: 只能标记状态, goroutine 无法强杀; 任务函数应响应 ctx 取消。
func (s *TimedTaskService) runMethod(t system.SysTimedTask) (string, error) {
	fn, ok := task.Get(t.MethodName)
	if !ok {
		return "", fmt.Errorf("方法 %s 未注册(需在 initialize/timer.go 经 task.Register 注册)", t.MethodName)
	}
	ctx, cancel := context.WithTimeout(datascope.WithSystem(context.Background()), defaultMethodTimeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				done <- fmt.Errorf("panic: %v", r)
			}
		}()
		done <- fn(ctx, json.RawMessage(t.Params))
	}()
	select {
	case err := <-done:
		if err != nil && errors.Is(err, context.DeadlineExceeded) {
			return "", errTaskTimeout
		}
		return "", err
	case <-ctx.Done():
		return "", errTaskTimeout
	}
}

// runHTTP 执行 HTTP 回调(SSRF 防护见 sys_timed_task_http.go)
func (s *TimedTaskService) runHTTP(t system.SysTimedTask) (string, error) {
	u, err := url.Parse(t.HttpUrl)
	if err != nil {
		return "", fmt.Errorf("URL 非法: %w", err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", fmt.Errorf("仅允许 http/https, 实际为 %q", u.Scheme)
	}
	method := strings.ToUpper(strings.TrimSpace(t.HttpMethod))
	if method == "" {
		method = http.MethodGet
	}
	var body io.Reader
	if t.HttpBody != "" {
		body = strings.NewReader(t.HttpBody)
	}
	req, err := http.NewRequest(method, t.HttpUrl, body)
	if err != nil {
		return "", fmt.Errorf("构造请求失败: %w", err)
	}
	if len(t.HttpHeader) > 0 {
		var hdr map[string]string
		if err := json.Unmarshal(t.HttpHeader, &hdr); err != nil {
			return "", fmt.Errorf("http_header 必须是 JSON 对象: %w", err)
		}
		for k, v := range hdr {
			req.Header.Set(k, v)
		}
	}

	client := newTimedTaskHTTPClient(t.HttpAllowPrivate, defaultHTTPTimeout)
	resp, err := client.Do(req)
	if err != nil {
		var uerr *url.Error
		if errors.As(err, &uerr) && uerr.Timeout() {
			return "", errTaskTimeout
		}
		return "", err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(io.LimitReader(resp.Body, maxHTTPRespBytes))
	out := fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(data))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return out, fmt.Errorf("非 2xx 响应: %d", resp.StatusCode)
	}
	return out, nil
}

// alertFailure 失败告警: 查 888 角色用户, 经本体 SSE Hub 定向推送(离线静默丢弃, 不阻塞)
func (s *TimedTaskService) alertFailure(t system.SysTimedTask, errMsg string) {
	var ids []uint
	if err := global.GVA_DB.Model(&system.SysUserAuthority{}).
		Where("sys_authority_authority_id = ?", alertAuthorityID).
		Pluck("sys_user_id", &ids).Error; err != nil {
		logger.Bg().Mod("timedTask").Err(err).Error("查询告警接收人失败")
		return
	}
	if len(ids) == 0 {
		return
	}
	payload, _ := json.Marshal(map[string]interface{}{
		"taskId": t.ID,
		"name":   t.Name,
		"error":  errMsg,
		"time":   time.Now().Format(time.RFC3339),
	})
	sse.Default().PublishToUsers(ids, sse.Event{Name: alertEventName, Data: string(payload)})
}
