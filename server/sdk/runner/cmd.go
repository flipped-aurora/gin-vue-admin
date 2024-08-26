package runner

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func init() {
	logger.Setup()
}

type Context struct {
	Cmd     string `json:"cmd"`
	Request string
}

func (r *Runner) Notfound(fn func(ctx *Context)) {
	r.NotFound = fn
}

func (c *Context) BindJSON(v interface{}) error {
	logrus.Infof(c.Request)
	return jsonx.UnmarshalFromFile(c.Request, v)
}

func (c *Context) ResponseJSON(res interface{}) error {
	marshal, err := json.Marshal(res)
	if err != nil {
		return err
	}
	//fmt.Println(string(marshal))
	c.Response(string(marshal))
	return nil
}

func (c *Context) Response(text string) {
	fmt.Println("<Response>" + text + "</Response>")
}

func New() *Runner {
	return &Runner{
		CmdMap: make(map[string]func(*Context)),
	}
}

type Runner struct {
	CmdMap   map[string]func(*Context)
	NotFound func(ctx *Context)
}

func (r *Runner) AddCmd(name string, fn func(ctx *Context)) {
	r.CmdMap[name] = fn
}

func (r *Runner) Run() {

	command := os.Args[1]
	jsonFileName := os.Args[2]
	logrus.Info(fmt.Sprintf("command:%s,args:%s", command, jsonFileName))
	f, ok := r.CmdMap[command]
	if ok {
		f(&Context{
			Request: jsonFileName,
		})
	} else {
	}

	//for name, fn := range r.CmdMap {
	//	req := flag.String(name, "", "")
	//
	//	if req != nil {
	//		logrus.Info(fmt.Sprintf("cmd:%s,args:%s", name, *req))
	//		fn(&Context{Request: *req})
	//		return
	//	}
	//}
}
