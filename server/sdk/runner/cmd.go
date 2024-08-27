package runner

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Call struct {
	User    string   `json:"user"`    //软件所属的用户
	Soft    string   `json:"soft"`    //软件名
	Command string   `json:"command"` //命令
	Files   []string `json:"files"`

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

type FileResponse struct {
	Type     string `json:"type"`
	FilePath string `json:"path"`
}

func init() {
	logger.Setup()
}

type Context struct {
	Cmd      string   `json:"cmd"`
	Request  string   `json:"request"`
	FileList []string `json:"file_list"`

	Req *Call
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
	//if ca.Data != nil {
	//	marshal, err := json.Marshal(ca.Data)
	//	if err != nil {
	//		return err
	//	}
	//	err = json.Unmarshal(marshal, jsonBody)
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func (c *Context) BindJSON(jsonBody interface{}) error {
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

func (c *Context) ResponseJSON(res interface{}) error {
	marshal, err := json.Marshal(res)
	if err != nil {
		return err
	}
	c.Response(string(marshal))
	return nil
}

func (c *Context) ResponseFile(filePath string) error {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	//这里需要转换成绝对路径
	marshal, err := json.Marshal(&FileResponse{Type: "file", FilePath: abs})
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
	fn, ok := r.CmdMap[command]
	if ok {
		context := &Context{Request: jsonFileName}
		bind(context)
		fn(context)
	} else {
	}
}
