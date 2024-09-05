package biz_apphub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner/response"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func NewCaller(types string) Caller {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir = strings.Replace(dir, "\\", "/", -1)
	return &SoftCall{AppPath: dir + "/soft_cmd"}
}

type CallResponse struct {
	Type       string      `json:"type"`
	FilePath   string      `json:"path"`
	DeleteFile bool        `json:"delete_file"`
	Data       interface{} `json:"data"`
}

func sCall(req request.Call) (*response.CallResponse, error) {
	return nil, nil
}

type Caller interface {
	Call(req request.Call) (*response.CallResponse, error)
	CallerPath() string
}

type SoftCall struct {
	AppPath string
}

func (p *SoftCall) CallerPath() string {
	return p.AppPath
}

type CallResult struct {
	ContentType string
	Data        string
	HasFile     bool
	FilePath    string
	DeleteFile  bool `json:"delete_file"`
}

func (p *SoftCall) Call(req request.Call) (*response.CallResponse, error) {
	softPath := p.AppPath + fmt.Sprintf("/%s/%s/%s",
		req.User, req.Soft, req.Soft)

	//execSoft := "./" + req.Soft
	//softHome := p.AppPath + fmt.Sprintf("/%s/%s",
	//	req.User, req.Soft)
	if runtime.GOOS == "windows" {
		softPath = p.AppPath + fmt.Sprintf("/%s/%s/%s.exe",
			req.User, req.Soft, req.Soft)
		//execSoft = execSoft + ".exe"
	}

	//err := os.Chdir(softHome)
	//if err != nil {
	//	panic(err)
	//}
	cmd := exec.Command(softPath, req.Command, req.RequestJsonPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	cmdStr := fmt.Sprintf("%s %s %s", softPath, req.Command, req.RequestJsonPath)
	fmt.Println(cmdStr)
	if err != nil {
		return nil, err
	}
	s := out.String()
	if s == "" {
		panic("")
		//return &CallResult{ContentType: "text"}, nil
	}
	split := strings.Split(s, "<Response>")[1]
	split1 := strings.Split(split, "</Response>")[0]
	ss := strings.ReplaceAll(split1, "</Response>", "")
	//mp := make(map[string]interface{})
	var res response.CallResponse
	err = json.Unmarshal([]byte(ss), &res)
	if err != nil {
		return nil, err
	}
	//if res.Type == "file" {
	//	return &runner.CallResponse{ContentType: "file", HasFile: true, FilePath: res.FilePath, DeleteFile: res.DeleteFile}, nil
	//}
	//
	//if res.Type == "text" {
	//	return &runner.CallResponse{ContentType: "text", Data: res.Data}, nil
	//}
	//if res.Type == "json" {
	//	return &runner.CallResponse{ContentType: "json", Data: res.Data}, nil
	//}
	return &res, nil
}

type Python struct {
}

func (p *Python) Call(req request.Call) (*response.CallResponse, error) {
	return nil, nil
}
func (p *Python) CallerPath() string {
	return ""
}
