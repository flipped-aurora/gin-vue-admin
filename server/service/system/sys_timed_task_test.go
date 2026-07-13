// server/service/system/sys_timed_task_test.go
package system

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
)

func snapshotHas(cronName string) bool {
	for _, d := range global.GVA_Timer.Snapshot() {
		if d.CronName == cronName {
			return true
		}
	}
	return false
}

func TestValidateSpec(t *testing.T) {
	svc := TimedTaskServiceApp
	if err := svc.ValidateSpec("@daily", false); err != nil {
		t.Fatalf("@daily 应合法: %v", err)
	}
	if err := svc.ValidateSpec("*/5 * * * *", false); err != nil {
		t.Fatalf("五位表达式应合法: %v", err)
	}
	if err := svc.ValidateSpec("61 * * * *", false); err == nil {
		t.Fatal("61 分钟应非法")
	}
	if err := svc.ValidateSpec("*/30 * * * * *", true); err != nil {
		t.Fatalf("六位含秒应合法: %v", err)
	}
}

func TestTimedTaskCRUDAndScheduleLifecycle(t *testing.T) {
	setupTimedTaskTestDB(t)
	ctx := context.Background()
	task.Register("lc_noop", "生命周期测试", func(ctx context.Context, _ json.RawMessage) error { return nil })
	svc := TimedTaskServiceApp

	// 创建(enabled) → 进入调度
	tk := &sysModel.SysTimedTask{Name: "生命周期", Spec: "@hourly", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "lc_noop", Enabled: true}
	if err := svc.CreateTimedTask(ctx, tk); err != nil {
		t.Fatalf("创建失败: %v", err)
	}
	cn := timedTaskCronName(tk.ID)
	if !snapshotHas(cn) {
		t.Fatal("创建后应已调度")
	}
	defer global.GVA_Timer.Clear(cn)

	// 同名再建被拒
	dup := &sysModel.SysTimedTask{Name: "生命周期", Spec: "@daily", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "lc_noop"}
	if err := svc.CreateTimedTask(ctx, dup); err == nil || !strings.Contains(err.Error(), "已存在") {
		t.Fatalf("同名应被拒: %v", err)
	}

	// 非法 cron 被拒
	bad := &sysModel.SysTimedTask{Name: "bad", Spec: "61 * * * *", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "lc_noop"}
	if err := svc.CreateTimedTask(ctx, bad); err == nil {
		t.Fatal("非法 cron 应被拒")
	}

	// 未注册方法被拒
	badm := &sysModel.SysTimedTask{Name: "badm", Spec: "@daily", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "no_such"}
	if err := svc.CreateTimedTask(ctx, badm); err == nil || !strings.Contains(err.Error(), "未注册") {
		t.Fatalf("未注册方法应被拒: %v", err)
	}

	// 停用 → 移出调度; DB enabled=false
	if err := svc.ToggleTimedTask(ctx, tk.ID, false); err != nil {
		t.Fatalf("停用失败: %v", err)
	}
	if snapshotHas(cn) {
		t.Fatal("停用后不应在调度中")
	}
	var row sysModel.SysTimedTask
	global.GVA_DB.First(&row, tk.ID)
	if row.Enabled {
		t.Fatal("DB enabled 应为 false")
	}

	// 启用 → 回到调度
	if err := svc.ToggleTimedTask(ctx, tk.ID, true); err != nil {
		t.Fatalf("启用失败: %v", err)
	}
	if !snapshotHas(cn) {
		t.Fatal("启用后应回到调度")
	}

	// 更新(改 spec) → 重新调度且仍只有一份
	// 重新读取最新行(启用后 enabled=true), 避免用停用期的旧快照把任务改回停用
	if err := global.GVA_DB.First(&row, tk.ID).Error; err != nil {
		t.Fatalf("重读任务失败: %v", err)
	}
	row.Spec = "@daily"
	if err := svc.UpdateTimedTask(ctx, &row); err != nil {
		t.Fatalf("更新失败: %v", err)
	}
	count := 0
	for _, d := range global.GVA_Timer.Snapshot() {
		if d.CronName == cn {
			count++
			if d.Spec != "@daily" {
				t.Fatalf("更新后 spec 应生效: %s", d.Spec)
			}
		}
	}
	if count != 1 {
		t.Fatalf("重复调度: %d 份", count)
	}

	// 删除 → 移出调度 + 软删
	if err := svc.DeleteTimedTask(ctx, tk.ID); err != nil {
		t.Fatalf("删除失败: %v", err)
	}
	if snapshotHas(cn) {
		t.Fatal("删除后不应在调度中")
	}
	if err := global.GVA_DB.First(&sysModel.SysTimedTask{}, tk.ID).Error; err == nil {
		t.Fatal("应已软删")
	}
}

func TestLoadAllOnlyEnabled(t *testing.T) {
	setupTimedTaskTestDB(t)
	ctx := context.Background()
	task.Register("la_noop", "loadall", func(ctx context.Context, _ json.RawMessage) error { return nil })

	on := sysModel.SysTimedTask{Name: "on", Spec: "@hourly", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "la_noop", Enabled: true}
	off := sysModel.SysTimedTask{Name: "off", Spec: "@hourly", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "la_noop", Enabled: false}
	global.GVA_DB.Create(&on)
	global.GVA_DB.Create(&off)
	defer global.GVA_Timer.Clear(timedTaskCronName(on.ID))

	if err := TimedTaskServiceApp.LoadAll(ctx); err != nil {
		t.Fatalf("LoadAll: %v", err)
	}
	if !snapshotHas(timedTaskCronName(on.ID)) {
		t.Fatal("enabled 任务应被调度")
	}
	if snapshotHas(timedTaskCronName(off.ID)) {
		t.Fatal("disabled 任务不应被调度")
	}
	// 幂等: 再跑一遍不产生重复
	if err := TimedTaskServiceApp.LoadAll(ctx); err != nil {
		t.Fatalf("LoadAll twice: %v", err)
	}
	count := 0
	for _, d := range global.GVA_Timer.Snapshot() {
		if d.CronName == timedTaskCronName(on.ID) {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("LoadAll 应幂等, got %d", count)
	}
}

func TestGetTimedTaskListWithNextRun(t *testing.T) {
	setupTimedTaskTestDB(t)
	ctx := context.Background()
	task.Register("ls_noop", "list", func(ctx context.Context, _ json.RawMessage) error { return nil })
	tk := &sysModel.SysTimedTask{Name: "列表任务", Spec: "@hourly", ExecutorType: sysModel.TimedTaskExecutorMethod, MethodName: "ls_noop", Enabled: true}
	if err := TimedTaskServiceApp.CreateTimedTask(ctx, tk); err != nil {
		t.Fatal(err)
	}
	defer global.GVA_Timer.Clear(timedTaskCronName(tk.ID))

	list, total, err := TimedTaskServiceApp.GetTimedTaskList(ctx, systemReq.SysTimedTaskSearch{})
	if err != nil || total != 1 || len(list) != 1 {
		t.Fatalf("列表: err=%v total=%d len=%d", err, total, len(list))
	}
	if list[0].NextRunAt == nil {
		t.Fatal("enabled 任务应有 nextRunAt")
	}
}
