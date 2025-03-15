package timer

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type Timer interface {
	// 寻找所有Cron
	FindCronList() map[string]*taskManager
	// 添加Task 方法形式以秒的形式加入
	AddTaskByFuncWithSecond(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) // 添加Task Func以秒的形式加入
	// 添加Task 接口形式以秒的形式加入
	AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	// 通过函数的方法添加任务
	AddTaskByFunc(cronName string, spec string, task func(), taskName string, option ...cron.Option) (cron.EntryID, error)
	// 通过接口的方法添加任务 要实现一个带有 Run方法的接口触发
	AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	// 获取对应taskName的cron 可能会为空
	FindCron(cronName string) (*taskManager, bool)
	// 指定cron开始执行
	StartCron(cronName string)
	// 指定cron停止执行
	StopCron(cronName string)
	// 查找指定cron下的指定task
	FindTask(cronName string, taskName string) (*task, bool)
	// 根据id删除指定cron下的指定task
	RemoveTask(cronName string, id int)
	// 根据taskName删除指定cron下的指定task
	RemoveTaskByName(cronName string, taskName string)
	// 清理掉指定cronName
	Clear(cronName string)
	// 停止所有的cron
	Close()
}

type task struct {
	EntryID  cron.EntryID
	Spec     string
	TaskName string
}

type taskManager struct {
	corn  *cron.Cron
	tasks map[cron.EntryID]*task
}

// timer 定时任务管理
type timer struct {
	cronList map[string]*taskManager
	sync.Mutex
}

// AddTaskByFunc 通过函数的方法添加任务
func (t *timer) AddTaskByFunc(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddFunc(spec, fun)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByFuncWithSecond 通过函数的方法使用WithSeconds添加任务
func (t *timer) AddTaskByFuncWithSecond(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddFunc(spec, fun)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByJob 通过接口的方法添加任务
func (t *timer) AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddJob(spec, job)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByJobWithSeconds 通过接口的方法添加任务
func (t *timer) AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddJob(spec, job)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// FindCron 获取对应cronName的cron 可能会为空
func (t *timer) FindCron(cronName string) (*taskManager, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.cronList[cronName]
	return v, ok
}

// FindTask 获取对应cronName的cron 可能会为空
func (t *timer) FindTask(cronName string, taskName string) (*task, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.cronList[cronName]
	if !ok {
		return nil, ok
	}
	for _, t2 := range v.tasks {
		if t2.TaskName == taskName {
			return t2, true
		}
	}
	return nil, false
}

// FindCronList 获取所有的任务列表
func (t *timer) FindCronList() map[string]*taskManager {
	t.Lock()
	defer t.Unlock()
	return t.cronList
}

// StartCron 开始任务
func (t *timer) StartCron(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Start()
	}
}

// StopCron 停止任务
func (t *timer) StopCron(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Stop()
	}
}

// RemoveTask 从cronName 删除指定任务
func (t *timer) RemoveTask(cronName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Remove(cron.EntryID(id))
		delete(v.tasks, cron.EntryID(id))
	}
}

// RemoveTaskByName 从cronName 使用taskName 删除指定任务
func (t *timer) RemoveTaskByName(cronName string, taskName string) {
	fTask, ok := t.FindTask(cronName, taskName)
	if !ok {
		return
	}
	t.RemoveTask(cronName, int(fTask.EntryID))
}

// Clear 清除任务
func (t *timer) Clear(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Stop()
		delete(t.cronList, cronName)
	}
}

// Close 释放资源
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.cronList {
		v.corn.Stop()
	}
}

func NewTimerTask() Timer {
	return &timer{cronList: make(map[string]*taskManager)}
}
