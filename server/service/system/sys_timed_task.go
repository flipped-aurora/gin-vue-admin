// server/service/system/sys_timed_task.go
package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/logger"
	"github.com/robfig/cron/v3"
)

// TimedTaskService 定时任务服务(全局单例 + global.GVA_DB, 遵项目既有约定)
type TimedTaskService struct{}

var TimedTaskServiceApp = new(TimedTaskService)

// timedTaskCronName 一任务一 cronName: robfig/cron 无单任务 pause,
// 启停语义 = Clear(cronName) + 按 DB 重加, 状态以 DB enabled 为准。
func timedTaskCronName(id uint) string { return fmt.Sprintf("timedTask/%d", id) }

// ValidateSpec 服务端校验 cron 表达式(含 @daily/@hourly 等描述符)
func (s *TimedTaskService) ValidateSpec(spec string, withSeconds bool) error {
	var err error
	if withSeconds {
		_, err = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor).Parse(spec)
	} else {
		_, err = cron.ParseStandard(spec)
	}
	if err != nil {
		return fmt.Errorf("cron 表达式非法: %w", err)
	}
	return nil
}

// validateTask 落库前统一校验(创建/更新共用)
func (s *TimedTaskService) validateTask(t *system.SysTimedTask) error {
	if t.Name == "" {
		return errors.New("任务名不能为空")
	}
	if err := s.ValidateSpec(t.Spec, t.WithSeconds); err != nil {
		return err
	}
	switch t.ExecutorType {
	case system.TimedTaskExecutorMethod:
		if _, ok := task.Get(t.MethodName); !ok {
			return fmt.Errorf("方法 %s 未注册", t.MethodName)
		}
		if len(t.Params) > 0 && !json.Valid(t.Params) {
			return errors.New("params 必须是合法 JSON")
		}
	case system.TimedTaskExecutorHTTP:
		u, err := url.Parse(t.HttpUrl)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
			return errors.New("httpUrl 必须是合法的 http/https 地址")
		}
		if len(t.HttpHeader) > 0 {
			var hdr map[string]string
			if err := json.Unmarshal(t.HttpHeader, &hdr); err != nil {
				return errors.New(`httpHeader 必须是 {"Key":"Value"} 形式的 JSON 对象`)
			}
		}
	default:
		return fmt.Errorf("executorType 必须为 %s 或 %s", system.TimedTaskExecutorMethod, system.TimedTaskExecutorHTTP)
	}
	return nil
}

// checkNameUnique 软删除下不建 DB 唯一索引, 由服务层保证活跃行内唯一
func (s *TimedTaskService) checkNameUnique(ctx context.Context, name string, excludeID uint) error {
	var count int64
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysTimedTask{}).Where("name = ?", name)
	if excludeID > 0 {
		db = db.Where("id <> ?", excludeID)
	}
	if err := db.Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("任务名 %s 已存在", name)
	}
	return nil
}

// ScheduleTask 幂等调度: 先 Clear 再按 enabled 重新注册
func (s *TimedTaskService) ScheduleTask(t system.SysTimedTask) error {
	name := timedTaskCronName(t.ID)
	global.GVA_Timer.Clear(name)
	if !t.Enabled {
		return nil
	}
	taskCopy := t // 值拷贝, 供回调闭包持有(更新任务时会整体重调度, 不会读到旧配置)
	fn := func() { s.RunTask(taskCopy, system.TimedTaskTriggerAuto) }
	var err error
	if t.WithSeconds {
		_, err = global.GVA_Timer.AddTaskByFuncWithSecond(name, t.Spec, fn, t.Name)
	} else {
		_, err = global.GVA_Timer.AddTaskByFunc(name, t.Spec, fn, t.Name)
	}
	return err
}

// LoadAll 启动/重载/initdb 后按 DB 恢复调度(仅 enabled 任务注册; 幂等)
func (s *TimedTaskService) LoadAll(ctx context.Context) error {
	var tasks []system.SysTimedTask
	if err := global.GVA_DB.WithContext(ctx).Find(&tasks).Error; err != nil {
		return err
	}
	for i := range tasks {
		if err := s.ScheduleTask(tasks[i]); err != nil {
			logger.Bg().Mod("timedTask").Err(err).Error("任务恢复调度失败: " + tasks[i].Name)
		}
	}
	logger.Bg().Mod("timedTask").Info(fmt.Sprintf("定时任务加载完成, 共 %d 条", len(tasks)))
	return nil
}

// CreateTimedTask 创建并按 enabled 调度
func (s *TimedTaskService) CreateTimedTask(ctx context.Context, t *system.SysTimedTask) error {
	if err := s.validateTask(t); err != nil {
		return err
	}
	if err := s.checkNameUnique(ctx, t.Name, 0); err != nil {
		return err
	}
	if err := global.GVA_DB.WithContext(ctx).Create(t).Error; err != nil {
		return err
	}
	return s.ScheduleTask(*t)
}

// UpdateTimedTask 更新并重新调度。Select 显式列 + Updates:
// 允许 enabled=false 等零值写入, 同时不碰 created_at/deleted_at。
func (s *TimedTaskService) UpdateTimedTask(ctx context.Context, t *system.SysTimedTask) error {
	if t.ID == 0 {
		return errors.New("缺少任务 ID")
	}
	if err := s.validateTask(t); err != nil {
		return err
	}
	if err := s.checkNameUnique(ctx, t.Name, t.ID); err != nil {
		return err
	}
	err := global.GVA_DB.WithContext(ctx).Model(&system.SysTimedTask{}).Where("id = ?", t.ID).
		Select("name", "description", "spec", "with_seconds", "executor_type", "method_name",
			"params", "http_url", "http_method", "http_header", "http_body", "http_allow_private", "enabled").
		Updates(t).Error
	if err != nil {
		return err
	}
	return s.ScheduleTask(*t)
}

// DeleteTimedTask 先移出调度再软删
func (s *TimedTaskService) DeleteTimedTask(ctx context.Context, id uint) error {
	global.GVA_Timer.Clear(timedTaskCronName(id))
	return global.GVA_DB.WithContext(ctx).Delete(&system.SysTimedTask{}, "id = ?", id).Error
}

// ToggleTimedTask 启停(状态以 DB 为准, 调度随之增删)
func (s *TimedTaskService) ToggleTimedTask(ctx context.Context, id uint, enabled bool) error {
	var row system.SysTimedTask
	if err := global.GVA_DB.WithContext(ctx).First(&row, id).Error; err != nil {
		return err
	}
	if err := global.GVA_DB.WithContext(ctx).Model(&system.SysTimedTask{}).
		Where("id = ?", id).Update("enabled", enabled).Error; err != nil {
		return err
	}
	row.Enabled = enabled
	return s.ScheduleTask(row)
}

// TriggerTimedTask 手动触发: goroutine 内跑一次 Runner, 不进调度不影响计划
func (s *TimedTaskService) TriggerTimedTask(ctx context.Context, id uint) error {
	var row system.SysTimedTask
	if err := global.GVA_DB.WithContext(ctx).First(&row, id).Error; err != nil {
		return err
	}
	go s.RunTask(row, system.TimedTaskTriggerManual)
	return nil
}

// GetTimedTaskList 分页列表 + 从调度器快照实时补 nextRunAt
func (s *TimedTaskService) GetTimedTaskList(ctx context.Context, info systemReq.SysTimedTaskSearch) (list []systemRes.SysTimedTaskRow, total int64, err error) {
	limit, offset := info.LimitOffset()
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysTimedTask{})
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.ExecutorType != "" {
		db = db.Where("executor_type = ?", info.ExecutorType)
	}
	if info.Enabled != nil {
		db = db.Where("enabled = ?", *info.Enabled)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	var tasks []system.SysTimedTask
	if err = db.Order("id desc").Find(&tasks).Error; err != nil {
		return
	}

	next := make(map[string]time.Time)
	for _, d := range global.GVA_Timer.Snapshot() {
		next[d.CronName] = d.NextRun
	}
	list = make([]systemRes.SysTimedTaskRow, 0, len(tasks))
	for _, t := range tasks {
		row := systemRes.SysTimedTaskRow{SysTimedTask: t}
		if nr, ok := next[timedTaskCronName(t.ID)]; ok && !nr.IsZero() {
			nrCopy := nr
			row.NextRunAt = &nrCopy
		}
		list = append(list, row)
	}
	return
}

// GetTimedTaskLogList 执行日志分页(按任务/状态过滤)
func (s *TimedTaskService) GetTimedTaskLogList(ctx context.Context, info systemReq.SysTimedTaskLogSearch) (list []system.SysTimedTaskLog, total int64, err error) {
	limit, offset := info.LimitOffset()
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysTimedTaskLog{})
	if info.TaskId > 0 {
		db = db.Where("task_id = ?", info.TaskId)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&list).Error
	return
}

// GetRegisteredMethods 已注册方法下拉(名称+说明)
func (s *TimedTaskService) GetRegisteredMethods() []task.TaskMeta {
	return task.List()
}
