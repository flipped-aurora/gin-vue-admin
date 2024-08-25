package runner

import (
	"encoding/json"
	"flag"
	"fmt"
)

type Context struct {
	Request string
}

func (c *Context) BindJSON(v interface{}) error {
	return json.Unmarshal([]byte(c.Request), v)
}

func (c *Context) ResponseJSON(res interface{}) error {
	marshal, err := json.Marshal(res)
	if err != nil {
		return err
	}
	fmt.Println(string(marshal))
	return nil
}

func (c *Context) Response(text string) {
	fmt.Println(text)
}

func New() *Runner {
	return &Runner{
		CmdMap: make(map[string]func(*Context)),
	}
}

type Runner struct {
	CmdMap map[string]func(*Context)
}

func (r *Runner) AddCmd(name string, fn func(ctx *Context)) {
	r.CmdMap[name] = fn
}

func (r *Runner) Run() {
	for name, fn := range r.CmdMap {
		req := flag.String(name, "", "")
		if req != nil {
			fn(&Context{Request: *req})
			return
		}
	}
}
