package runner

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/logger"
	logger2 "github.com/flipped-aurora/gin-vue-admin/server/sdk/runner/logger"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner/request"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	logger.Setup()
}

type Context struct {
	IsDebug  bool
	Cmd      string   `json:"cmd"`
	Request  string   `json:"request"`
	FileList []string `json:"file_list"`

	Req    *request.Call
	runner *Runner `json:"-"`
}

func (c *Context) Logger() *logrus.Entry {
	return logrus.StandardLogger().WithFields(logrus.Fields{
		"user":      c.Req.User,
		"soft":      c.Req.Soft,
		"command":   c.Req.Command,
		"user_soft": fmt.Sprintf("/%s/%s", c.Req.User, c.Req.Soft),
	})
}
func (c *Context) GetLogger() *logger2.Logger {

	mp := make(map[string]interface{})
	mp["a_tenant"] = c.Req.User
	mp["a_soft"] = c.Req.Soft
	mp["a_command"] = c.Req.Command
	if c.runner != nil {
		if c.runner.Version != "" {
			mp["a_version"] = c.runner.Version
		}
	}

	mp["a_soft_classify"] = fmt.Sprintf("/%s/%s", c.Req.User, c.Req.Soft)
	return &logger2.Logger{
		DataMap: mp,
	}
}

func (r *Runner) Notfound(fn func(ctx *Context)) {
	r.NotFound = fn
}

func (c *Context) DebugPrintf(format string, args ...interface{}) {
	if c.IsDebug {
		fmt.Printf(format, args...)
	}
}

func (r *Runner) DebugPrintf(format string, args ...interface{}) {
	if r.IsDebug {
		fmt.Printf(format, args...)
	}
}

func bind(ctx *Context) error {
	var ca request.Call
	err := jsonx.UnmarshalFromFile(ctx.Request, &ca)
	ctx.DebugPrintf("bind err:%v", err)
	if err != nil {
		return err
	}
	ctx.Req = &ca
	return nil
}

func (c *Context) ShouldBindJSON(jsonBody interface{}) error {
	if c.Req != nil {
		marshal, err := json.Marshal(c.Req.Data)
		if err != nil {
			return err
		}
		err = json.Unmarshal(marshal, jsonBody)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *Context) ReqMap() map[string]interface{} {
	if c.Req != nil {
		return c.Req.Data
	}
	return nil
}

func New() *Runner {
	r := &Runner{
		CmdMapHandel: make(map[string]*Worker),
	}
	for path, fn := range GetMap {
		r.Get(path, fn)
	}
	return r
}

type Worker struct {
	Handel []func(*Context)
	Path   string
	Method string
	Config *Config
}

type Runner struct {
	IsDebug      bool
	Version      string
	CmdMapHandel map[string]*Worker
	NotFound     func(ctx *Context)
}

func (r *Runner) SetVersion(version string) {
	r.Version = version
}
func (r *Runner) Post(commandName string, handelFunc func(ctx *Context), config ...*Config) {
	_, ok := r.CmdMapHandel[commandName]
	if !ok {
		worker := &Worker{
			Handel: []func(*Context){handelFunc},
			Method: "POST",
			Path:   commandName,
		}
		if len(config) > 0 {
			worker.Config = config[0]
		}
		r.CmdMapHandel[commandName+".POST"] = worker
	} else {
		r.CmdMapHandel[commandName].Handel = append(r.CmdMapHandel[commandName].Handel, handelFunc)
	}

}
func (r *Runner) Get(commandName string, handelFunc func(ctx *Context), config ...*Config) {
	_, ok := r.CmdMapHandel[commandName]
	if !ok {
		worker := &Worker{
			Handel: []func(*Context){handelFunc},
			Method: "GET",
			Path:   commandName,
		}
		if len(config) > 0 {
			worker.Config = config[0]
		}
		r.CmdMapHandel[commandName+".GET"] = worker
	} else {
		r.CmdMapHandel[commandName].Handel = append(r.CmdMapHandel[commandName].Handel, handelFunc)
	}
}

func (r *Runner) Run() {
	command := os.Args[1]
	jsonFileName := os.Args[2]
	if len(os.Args) > 3 {
		r.IsDebug = true
	}
	r.DebugPrintf("run ....")

	context := &Context{Request: jsonFileName, IsDebug: r.IsDebug}
	err := bind(context)
	if err != nil {
		context.ResponseFailJSONWithCode(http.StatusBadRequest, map[string]interface{}{
			"msg": "参数解析失败: " + err.Error(),
		})
		return
	}

	method := strings.ToUpper(context.Req.Method)
	//todo
	if command == "_docs_info_text" && method == "GET" { //获取接口文档
		var s []string
		for _, worker := range r.CmdMapHandel {
			if worker.Config == nil {
				continue
			}
			if !worker.Config.IsPublicApi {
				continue
			}
			s = append(s, fmt.Sprintf("%s\t %s \t %s", worker.Path, worker.Method, worker.Config.ApiDesc))
		}
		//res := append([]string{"请求地址 \t 请求方式 \t 接口描述"}, s...)
		context.ResponseOkWithText(strings.Join(s, "\n"))
		return
	}

	worker, ok := r.CmdMapHandel[command+"."+method]
	if ok {
		if context.Req == nil {
			panic("context.Req == nil")
		}
		handelList := worker.Handel
		//if context.Req.Method == "GET" {
		//	handelList = worker.Handel
		//}
		if len(handelList) == 0 {
			context.ResponseFailTextWithCode(http.StatusBadRequest, "bad request: method not handel")
			return
		}
		for _, fn := range handelList {
			now := time.Now()
			fn(context)
			t := time.Since(now)
			fmt.Println(fmt.Sprintf("<UserCost>%s</UserCost>", t.String()))
		}

	} else { //not found
		if r.NotFound != nil {
			r.NotFound(context)
		} else {
			context.ResponseFailTextWithCode(http.StatusNotFound, "command not found")
		}
	}
}
