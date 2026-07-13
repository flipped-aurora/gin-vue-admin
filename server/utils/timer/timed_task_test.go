package timer

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var job = mockJob{}

type mockJob struct{}

func (job mockJob) Run() {
	mockFunc()
}

func mockFunc() {
	time.Sleep(time.Second)
	fmt.Println("1s...")
}

func TestNewTimerTask(t *testing.T) {
	tm := NewTimerTask()
	_tm := tm.(*timer)

	{
		_, err := tm.AddTaskByFunc("func", "@every 1s", mockFunc, "测试mockfunc")
		assert.Nil(t, err)
		_, ok := _tm.cronList["func"]
		if !ok {
			t.Error("no find func")
		}
	}

	{
		_, err := tm.AddTaskByJob("job", "@every 1s", job, "测试job mockfunc")
		assert.Nil(t, err)
		_, ok := _tm.cronList["job"]
		if !ok {
			t.Error("no find job")
		}
	}

	{
		_, ok := tm.FindCron("func")
		if !ok {
			t.Error("no find func")
		}
		_, ok = tm.FindCron("job")
		if !ok {
			t.Error("no find job")
		}
		_, ok = tm.FindCron("none")
		if ok {
			t.Error("find none")
		}
	}
	{
		tm.Clear("func")
		_, ok := tm.FindCron("func")
		if ok {
			t.Error("find func")
		}
	}
	{
		a := tm.FindCronList()
		b, c := tm.FindCron("job")
		fmt.Println(a, b, c)
	}
}

func TestSnapshot(t *testing.T) {
	tm := NewTimerTask()
	defer tm.Close()
	_, err := tm.AddTaskByFunc("snapCron", "@hourly", func() {}, "快照任务")
	if err != nil {
		t.Fatal(err)
	}
	snap := tm.Snapshot()
	if len(snap) != 1 {
		t.Fatalf("期望 1 条, got %d", len(snap))
	}
	d := snap[0]
	if d.CronName != "snapCron" || d.TaskName != "快照任务" || d.Spec != "@hourly" {
		t.Fatalf("快照字段不符: %+v", d)
	}
	if !d.NextRun.After(time.Now()) {
		t.Fatalf("NextRun 应在未来: %v", d.NextRun)
	}
}

func TestSnapshotConcurrent(t *testing.T) { // 配合 -race: 快照与增删并发安全
	tm := NewTimerTask()
	defer tm.Close()
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(2)
		go func(n int) {
			defer wg.Done()
			_, _ = tm.AddTaskByFunc(fmt.Sprintf("c%d", n), "@daily", func() {}, "并发")
		}(i)
		go func() {
			defer wg.Done()
			for _, d := range tm.Snapshot() {
				_ = d.NextRun
			}
		}()
	}
	wg.Wait()
}
