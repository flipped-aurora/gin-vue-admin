package runner

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/logger"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	logger.Setup()
}

type Call struct {
	User            string                 `json:"user"`    //软件所属的用户
	Soft            string                 `json:"soft"`    //软件名
	Command         string                 `json:"command"` //命令
	Files           []string               `json:"files"`
	Method          string                 `json:"method"`            //请求方式
	UpdateVersion   bool                   `json:"update_version"`    //此时正处于版本更新的状态
	RequestJsonPath string                 `json:"request_json_path"` //请求参数存储路径
	Data            map[string]interface{} `json:"data"`              //请求json
	ReqBody         string
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CallResponse struct {
	StatusCode  int         `json:"status_code"`
	Msg         string      `json:"msg"`
	ContentType string      `json:"content_type"`
	HasFile     bool        `json:"has_file"`
	FilePath    string      `json:"path"`
	DeleteFile  bool        `json:"delete_file"`
	Data        interface{} `json:"data"`
}

type Context struct {
	Cmd      string   `json:"cmd"`
	Request  string   `json:"request"`
	FileList []string `json:"file_list"`

	Req *Call
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

func bind(ctx *Context) error {
	var ca Call
	err := jsonx.UnmarshalFromFile(ctx.Request, &ca)
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
	c.Response(jsonx.JSONString(&CallResponse{StatusCode: code, ContentType: "json", Data: jsonEl}))
}

func (c *Context) ResponseFailDefaultJSONWithMsg(errMsg string) {
	c.Response(jsonx.JSONString(&CallResponse{StatusCode: 200, ContentType: "json", Data: map[string]interface{}{
		"msg": errMsg,
	}}))
}

func (c *Context) ResponseFailTextWithCode(code int, text string) {
	c.Response(jsonx.JSONString(&CallResponse{StatusCode: code, ContentType: "text", Data: text}))
}
func (c *Context) ResponseOkWithJSON(jsonEl interface{}) {
	c.Response(jsonx.JSONString(&CallResponse{StatusCode: 200, ContentType: "json", Data: jsonEl}))
}

func (c *Context) ResponseOkWithFile(filePath string, deleteFile bool) error {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	//这里需要转换成绝对路径
	marshal, err := json.Marshal(&CallResponse{HasFile: true, StatusCode: 200, ContentType: "file", FilePath: abs, DeleteFile: deleteFile})
	if err != nil {
		return err
	}
	c.Response(string(marshal))
	return nil
}

func (c *Context) ResponseOkWithText(text string) error {
	marshal, err := json.Marshal(&CallResponse{StatusCode: 200, ContentType: "text", Data: text})
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
	return &Runner{
		//CmdMap: make(map[string]func(*Context)),
		CmdMapHandel: make(map[string]*Worker),
	}
}

type Worker struct {
	GetHandel  []func(*Context)
	PostHandel []func(*Context)
}

type Runner struct {
	CmdMapHandel map[string]*Worker
	NotFound     func(ctx *Context)
}

func (r *Runner) Post(commandName string, handelList ...func(ctx *Context)) {
	_, ok := r.CmdMapHandel[commandName]
	if !ok {

		r.CmdMapHandel[commandName] = &Worker{
			PostHandel: handelList,
		}
	} else {
		r.CmdMapHandel[commandName].PostHandel = append(r.CmdMapHandel[commandName].PostHandel, handelList...)
	}

}
func (r *Runner) Get(commandName string, handelList ...func(ctx *Context)) {
	_, ok := r.CmdMapHandel[commandName]
	if !ok {
		r.CmdMapHandel[commandName] = &Worker{
			GetHandel: handelList,
		}
	} else {
		r.CmdMapHandel[commandName].GetHandel = append(r.CmdMapHandel[commandName].GetHandel, handelList...)
	}
}

func (r *Runner) Run() {
	command := os.Args[1]
	jsonFileName := os.Args[2]
	worker, ok := r.CmdMapHandel[command]
	context := &Context{Request: jsonFileName}
	err := bind(context)
	if err != nil {
		context.ResponseFailJSONWithCode(http.StatusBadRequest, map[string]interface{}{
			"msg": "参数解析失败",
		})
		return
	}
	if ok {
		if context.Req == nil {
			panic("context.Req == nil")
		}
		handelList := worker.PostHandel
		if context.Req.Method == "GET" {
			handelList = worker.GetHandel
		}
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
