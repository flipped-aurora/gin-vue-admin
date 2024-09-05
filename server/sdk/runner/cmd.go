package runner

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/logger"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner/request"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner/response"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	logger.Setup()
}

type Context struct {
	IsDebug  bool
	Cmd      string   `json:"cmd"`
	Request  string   `json:"request"`
	FileList []string `json:"file_list"`

	Req *request.Call
}

func (c *Context) Logger() *logrus.Entry {
	return logrus.StandardLogger().WithFields(logrus.Fields{
		"user":      c.Req.User,
		"soft":      c.Req.Soft,
		"command":   c.Req.Command,
		"user_soft": fmt.Sprintf("/%s/%s", c.Req.User, c.Req.Soft),
	})
}

//func (c *Context) FileList() []string {
//	return nil
//}

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

func (c *Context) ResponseFailJSONWithCode(code int, jsonEl interface{}) {
	c.Response(jsonx.JSONString(&response.CallResponse{StatusCode: code, ContentType: "json", Data: jsonEl}))
}

func (c *Context) ResponseFailDefaultJSONWithMsg(errMsg string) {
	c.Response(jsonx.JSONString(&response.CallResponse{StatusCode: 200, ContentType: "json", Data: map[string]interface{}{"msg": errMsg}}))
}

func (c *Context) ResponseFailTextWithCode(code int, text string) {
	c.Response(jsonx.JSONString(&response.CallResponse{StatusCode: code, ContentType: "text", Data: text}))
}
func (c *Context) ResponseOkWithJSON(jsonEl interface{}) {
	c.Response(jsonx.JSONString(&response.CallResponse{StatusCode: 200, ContentType: "json", Data: jsonEl}))
}

func (c *Context) ResponseOkWithFile(filePath string, deleteFile bool) error {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	//这里需要转换成绝对路径
	marshal, err := json.Marshal(&response.CallResponse{HasFile: true, StatusCode: 200, ContentType: "file", FilePath: abs, DeleteFile: deleteFile})
	if err != nil {
		return err
	}
	c.Response(string(marshal))
	return nil
}

func (c *Context) ResponseOkWithText(text string) error {
	marshal, err := json.Marshal(&response.CallResponse{StatusCode: 200, ContentType: "text", Data: text})
	if err != nil {
		return err
	}
	c.Response(string(marshal))
	return nil
}

func (c *Context) Response(text string) {
	fmt.Println("<Response>" + text + "</Response>")
}

func New() *Runner {
	//logger.Setup()

	r := &Runner{
		//CmdMap: make(map[string]func(*Context)),
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
	//GetHandel  []func(*Context)
	//PostHandel []func(*Context)
}

type Runner struct {
	IsDebug      bool
	CmdMapHandel map[string]*Worker
	NotFound     func(ctx *Context)
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
	fmt.Println(command)
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
			fn(context)
		}

	} else { //not found
		if r.NotFound != nil {
			r.NotFound(context)
		} else {
			context.ResponseFailTextWithCode(http.StatusNotFound, "command not found")
		}
	}
}
