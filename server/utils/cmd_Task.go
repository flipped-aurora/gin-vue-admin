package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type RunTask interface {
	AddTask()
	RunTask()
}

// T: Task任务
type T struct {
	sync.Mutex

	// ch: 获取时间channel
	ch chan struct{}

	// 记录pid 用于kill
	pid int

	// 执行shell命令
	exec.Cmd
}

// NewT: 实例化方法
func NewT(path string, args []string, environment ...[]string) *T {
	env := os.Environ()
	if len(environment) > 0 {
		for k, v := range environment[0] {
			env[k] = v
		}
	}
	t := &T{
		Mutex: sync.Mutex{},
		ch:    make(chan struct{}),
		Cmd: exec.Cmd{
			Path:       path,
			Args:       []string{path},
			Env:        env,
			Stdin:      os.Stdin,
			Stdout:     os.Stdout,
			Stderr:     os.Stderr,
			ExtraFiles: make([]*os.File, 0),
		},
		pid: os.Getpid(),
	}
	t.Dir, _ = os.Getwd()
	if len(args) > 0 {
		// Exclude of current binary path.
		start := 0
		if strings.EqualFold(path, args[0]) {
			start = 1
		}
		t.Args = append(t.Args, args[start:]...)
	}
	return t
}

func (t *T) AddTask() {
	t.Lock()
	defer t.Unlock()
	if len(t.ch) == 1 {
		// 代表已经有任务了
		// 直接丢弃这次任务
		return
	}
	t.ch <- struct{}{}
}

func (t *T) RunTask() {
	// 先run服务
	err := t.Run()
	if err != nil {
		return
	}
	for {
		_, ok := <-t.ch
		if !ok {
			return
		}
		// todo 执行任务
		// 先编译新的文件 然后 kill 然后执行
		fmt.Println(kill(t.Cmd.Process.Pid))
		err = t.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

// kill:
func kill(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Kill()
}
