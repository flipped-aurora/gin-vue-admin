package timer

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type GVA_Timer interface {
	Timer
	FindTaskList() map[string]*taskManager
	AddTaskByFuncWithSecond(taskName string, spec string, fun func(), Desc string, option ...cron.Option) (cron.EntryID, error)
	AddTaskByJobWithSeconds(taskName string, spec string, job interface{ Run() }, Desc string, option ...cron.Option) (cron.EntryID, error)
}

type Timer interface {
	AddTaskByFunc(taskName string, spec string, task func(), Desc string, option ...cron.Option) (cron.EntryID, error)
	AddTaskByJob(taskName string, spec string, job interface{ Run() }, Desc string, option ...cron.Option) (cron.EntryID, error)
	FindCron(taskName string) (*taskManager, bool)
	StartTask(taskName string)
	StopTask(taskName string)
	Remove(taskName string, id int)
	Clear(taskName string)
	Close()
}

type task struct {
	EntryID cron.EntryID
	Spec    string
	Desc    string
}

type taskManager struct {
	corn  *cron.Cron
	tasks map[cron.EntryID]*task
}

// timer 定时任务管理
type timer struct {
	taskList map[string]*taskManager
	sync.Mutex
}

// AddTaskByFunc 通过函数的方法添加任务
func (t *timer) AddTaskByFunc(taskName string, spec string, fun func(), Desc string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.taskList[taskName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.taskList[taskName].corn.AddFunc(spec, fun)
	t.taskList[taskName].corn.Start()
	t.taskList[taskName].tasks[id] = &task{
		EntryID: id,
		Spec:    spec,
		Desc:    Desc,
	}
	return id, err
}

// AddTaskByFuncWithSeconds 通过函数的方法使用WithSeconds添加任务
func (t *timer) AddTaskByFuncWithSecond(taskName string, spec string, fun func(), Desc string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.taskList[taskName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.taskList[taskName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.taskList[taskName].corn.AddFunc(spec, fun)
	t.taskList[taskName].corn.Start()
	t.taskList[taskName].tasks[id] = &task{
		EntryID: id,
		Spec:    spec,
		Desc:    Desc,
	}
	return id, err
}

// AddTaskByJob 通过接口的方法添加任务
func (t *timer) AddTaskByJob(taskName string, spec string, job interface{ Run() }, Desc string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.taskList[taskName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.taskList[taskName].corn.AddJob(spec, job)
	t.taskList[taskName].corn.Start()
	t.taskList[taskName].tasks[id] = &task{
		EntryID: id,
		Spec:    spec,
		Desc:    Desc,
	}
	return id, err
}

// AddTaskByJobWithSeconds 通过接口的方法添加任务
func (t *timer) AddTaskByJobWithSeconds(taskName string, spec string, job interface{ Run() }, Desc string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.taskList[taskName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.taskList[taskName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.taskList[taskName].corn.AddJob(spec, job)
	t.taskList[taskName].corn.Start()
	t.taskList[taskName].tasks[id] = &task{
		EntryID: id,
		Spec:    spec,
		Desc:    Desc,
	}
	return id, err
}

// FindCron 获取对应taskName的cron 可能会为空
func (t *timer) FindCron(taskName string) (*taskManager, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.taskList[taskName]
	return v, ok
}

// FindTaskList 获取所有的任务列表
func (t *timer) FindTaskList() map[string]*taskManager {
	t.Lock()
	defer t.Unlock()
	return t.taskList
}

// StartTask 开始任务
func (t *timer) StartTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.corn.Start()
	}
}

// StopTask 停止任务
func (t *timer) StopTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.corn.Stop()
	}
}

// Remove 从taskName 删除指定任务
func (t *timer) Remove(taskName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.corn.Remove(cron.EntryID(id))
		delete(v.tasks, cron.EntryID(id))
	}
}

// Clear 清除任务
func (t *timer) Clear(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.corn.Stop()
		delete(t.taskList, taskName)
	}
}

// Close 释放资源
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.taskList {
		v.corn.Stop()
	}
}

func NewTimerTask() GVA_Timer {
	return &timer{taskList: make(map[string]*taskManager)}
}
