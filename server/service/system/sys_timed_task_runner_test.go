// server/service/system/sys_timed_task_runner_test.go
package system

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/internal/testutil"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
)

// setupTimedTaskTestDB 内存库 + 必需表(runner 失败路径会查 sys_user_authority 发告警)。
// 复用 testutil.NewMemoryDB：内部完成 nop logger 兜底、sqlite :memory:、AutoMigrate、
// 赋值 global.GVA_DB 并在 t.Cleanup 还原，替代原先手写的整套样板。
func setupTimedTaskTestDB(t *testing.T) {
	t.Helper()
	testutil.NewMemoryDB(t, &sysModel.SysTimedTask{}, &sysModel.SysTimedTaskLog{}, &sysModel.SysUserAuthority{})
}

func lastLog(t *testing.T) sysModel.SysTimedTaskLog {
	t.Helper()
	var row sysModel.SysTimedTaskLog
	if err := global.GVA_DB.Order("id desc").First(&row).Error; err != nil {
		t.Fatalf("读日志失败: %v", err)
	}
	return row
}

func TestRunTaskMethodSuccess(t *testing.T) {
	setupTimedTaskTestDB(t)
	var gotParams string
	task.Register("rt_ok", "测试", func(ctx context.Context, p json.RawMessage) error {
		gotParams = string(p)
		return nil
	})
	tk := sysModel.SysTimedTask{Name: "ok任务", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "rt_ok", Params: []byte(`{"days":30}`)}
	tk.ID = 1
	TimedTaskServiceApp.RunTask(tk, sysModel.TimedTaskTriggerManual)

	if gotParams != `{"days":30}` {
		t.Fatalf("params 未透传: %s", gotParams)
	}
	row := lastLog(t)
	if row.Status != sysModel.TimedTaskStatusSuccess || row.TriggerType != sysModel.TimedTaskTriggerManual || row.TaskName != "ok任务" {
		t.Fatalf("日志不符: %+v", row)
	}
}

func TestRunTaskMethodNotRegistered(t *testing.T) {
	setupTimedTaskTestDB(t)
	tk := sysModel.SysTimedTask{Name: "missing", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "rt_not_exist"}
	tk.ID = 2
	TimedTaskServiceApp.RunTask(tk, sysModel.TimedTaskTriggerAuto)
	row := lastLog(t)
	if row.Status != sysModel.TimedTaskStatusFail || !strings.Contains(row.ErrorMsg, "未注册") {
		t.Fatalf("应记 fail 未注册: %+v", row)
	}
}

func TestRunTaskMethodPanicRecovered(t *testing.T) {
	setupTimedTaskTestDB(t)
	task.Register("rt_panic", "测试", func(ctx context.Context, _ json.RawMessage) error {
		panic("boom!")
	})
	tk := sysModel.SysTimedTask{Name: "panic任务", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "rt_panic"}
	tk.ID = 3
	TimedTaskServiceApp.RunTask(tk, sysModel.TimedTaskTriggerAuto) // 不得崩进程
	row := lastLog(t)
	if row.Status != sysModel.TimedTaskStatusFail || !strings.Contains(row.ErrorMsg, "panic") {
		t.Fatalf("panic 应记 fail: %+v", row)
	}
}

func TestRunTaskMethodTimeout(t *testing.T) {
	setupTimedTaskTestDB(t)
	oldTimeout := defaultMethodTimeout
	defaultMethodTimeout = 50 * time.Millisecond
	defer func() { defaultMethodTimeout = oldTimeout }()

	task.Register("rt_slow", "测试", func(ctx context.Context, _ json.RawMessage) error {
		<-ctx.Done() // 规范任务: 响应取消
		return ctx.Err()
	})
	tk := sysModel.SysTimedTask{Name: "slow任务", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "rt_slow"}
	tk.ID = 4
	TimedTaskServiceApp.RunTask(tk, sysModel.TimedTaskTriggerAuto)
	row := lastLog(t)
	if row.Status != sysModel.TimedTaskStatusTimeout {
		t.Fatalf("应记 timeout: %+v", row)
	}
}

func TestRunTaskHTTP(t *testing.T) {
	setupTimedTaskTestDB(t)
	var gotHeader, gotBody, gotMethod string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotHeader = r.Header.Get("X-Biz")
		b := make([]byte, r.ContentLength)
		r.Body.Read(b)
		gotBody = string(b)
		w.WriteHeader(200)
		w.Write([]byte("pong"))
	}))
	defer srv.Close()

	tk := sysModel.SysTimedTask{
		Name: "http任务", ExecutorType: sysModel.TimedTaskExecutorHTTP,
		HttpUrl: srv.URL, HttpMethod: "POST",
		HttpHeader: []byte(`{"X-Biz":"gva"}`), HttpBody: `{"ping":1}`,
		HttpAllowPrivate: true, // httptest 在 127.0.0.1, 需豁免
	}
	tk.ID = 5
	TimedTaskServiceApp.RunTask(tk, sysModel.TimedTaskTriggerManual)
	row := lastLog(t)
	if row.Status != sysModel.TimedTaskStatusSuccess || !strings.Contains(row.Output, "200") || !strings.Contains(row.Output, "pong") {
		t.Fatalf("http 成功日志不符: %+v", row)
	}
	if gotMethod != "POST" || gotHeader != "gva" || gotBody != `{"ping":1}` {
		t.Fatalf("请求未按配置发出: %s %s %s", gotMethod, gotHeader, gotBody)
	}
}

func TestRunTaskHTTPNon2xxAndSSRF(t *testing.T) {
	setupTimedTaskTestDB(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(500)
	}))
	defer srv.Close()

	// 非 2xx → fail 且记录状态码
	tk := sysModel.SysTimedTask{Name: "e500", ExecutorType: sysModel.TimedTaskExecutorHTTP, HttpUrl: srv.URL, HttpAllowPrivate: true}
	tk.ID = 6
	TimedTaskServiceApp.RunTask(tk, sysModel.TimedTaskTriggerAuto)
	row := lastLog(t)
	if row.Status != sysModel.TimedTaskStatusFail || !strings.Contains(row.ErrorMsg, "500") {
		t.Fatalf("非2xx应记fail: %+v", row)
	}

	// SSRF: 未豁免时打 127.0.0.1 被拒
	tk2 := sysModel.SysTimedTask{Name: "ssrf", ExecutorType: sysModel.TimedTaskExecutorHTTP, HttpUrl: srv.URL, HttpAllowPrivate: false}
	tk2.ID = 7
	TimedTaskServiceApp.RunTask(tk2, sysModel.TimedTaskTriggerAuto)
	row2 := lastLog(t)
	if row2.Status != sysModel.TimedTaskStatusFail || !strings.Contains(row2.ErrorMsg, "SSRF") {
		t.Fatalf("SSRF 应被拒: %+v", row2)
	}
}
