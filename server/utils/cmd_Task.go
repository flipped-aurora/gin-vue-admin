package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

type RunTask interface {
	AddTask()
	RunTask()
}

// T: Task任务
type T struct {
	sync.Mutex

	// ch: 获取事件channel
	ch chan struct{}

	closeChan chan struct{}

	// 记录process对象
	p *os.Process

	// f: 执行任务
	f func(chan struct{}) error
}

// NewT: 实例化方法
func NewT() *T {
	return newT(nil)
}
func newT(f func(chan struct{}) error) *T {
	t := &T{
		Mutex:     sync.Mutex{},
		ch:        make(chan struct{}, 1),
		closeChan: make(chan struct{}),
		f:         f,
	}
	if f == nil {
		t.f = t.DefaultF
	}
	return t
}

func (t *T) AddTask() {
	if len(t.ch) == 1 {
		return
	}
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
	fmt.Println("进入")
	// 这里做的make 是用于关闭上一个执行的任务
	ch := make(chan struct{})
	// 先run服务
	go t.f(ch)
	for {
		_, ok := <-t.ch
		if !ok {
			return
		}
		ch <- struct{}{}
		// 等待上一个关闭
		<-t.closeChan
		go t.f(ch)
	}

}

// DefaultF: 默认的StartFunction
func (t *T) DefaultF(ch chan struct{}) error {
	var buildCmd *exec.Cmd
	var cmd *exec.Cmd

	// 检测系统是否有编译环境
	_, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	// build

	switch runtime.GOOS {
	case "windows":
		buildCmd = exec.Command("go", "build", "-o", "gva.exe", "main.go")
	default:
		buildCmd = exec.Command("go", "build", "-o", "gva", "main.go")
	}
	//cmd = exec.Command("go", "run", "main.go")
	err = buildCmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("build 执行完成")
	// 执行
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("gva.exe")
	default:
		cmd = exec.Command("./gva")
	}

	// 开始执行任务
	ctx, cancel := context.WithCancel(context.Background())
	err = t.echo(cmd, ctx)
	<-ch
	// 回收资源
	fmt.Println("pid:", t.p.Pid, "->Kill")
	err = t.p.Kill()
	cancel()
	// 发送关闭完成信号
	t.closeChan <- struct{}{}
	return err
}

// echo: 封装回显
func (t *T) echo(cmd *exec.Cmd, ctx context.Context) error {
	var stdoutBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	var errStdout error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	err := cmd.Start()
	if err != nil {
		return err
	}
	go func(ctx context.Context) {
		_, errStdout = io.Copy(stdout, stdoutIn)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}(ctx)
	t.p = cmd.Process
	fmt.Println("pid", t.p.Pid)
	go func() {
		_ = cmd.Wait()
		if errStdout != nil {
			fmt.Printf("failed to capture stdout or stderr\n")
		}
		fmt.Printf("%s\n", string(stdoutBuf.Bytes()))
		select {
		case <-ctx.Done():
			_ = os.Stdout.Close()
			return
		default:
		}
	}()
	return nil
}
