package biz_apphub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"os/exec"
	"runtime"
	"strings"
)

func NewCaller(types string) Caller {
	return &SoftCall{AppPath: "D:\\code\\github.com\\apphub\\server\\soft_cmd"}
}

type Caller interface {
	Call(req request.Call) (*CallResult, error)
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
}

func (p *SoftCall) Call(req request.Call) (*CallResult, error) {
	softPath := p.AppPath + fmt.Sprintf("\\%s\\%s\\%s",
		req.User, req.Soft, req.Soft)
	if runtime.GOOS == "windows" {
		softPath = p.AppPath + fmt.Sprintf("\\%s\\%s\\%s.exe",
			req.User, req.Soft, req.Soft)
	}
	cmd := exec.Command(softPath, req.Command, req.RequestJsonPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	s := out.String()
	split := strings.Split(s, "<Response>")[1]
	split1 := strings.Split(split, "</Response>")[0]
	ss := strings.ReplaceAll(split1, "</Response>", "")
	mp := make(map[string]interface{})

	err = json.Unmarshal([]byte(ss), &mp)
	if err != nil {
		return nil, err
	}
	if _, ok := mp["type"]; ok {
		return &CallResult{ContentType: "file", HasFile: true, FilePath: mp["path"].(string)}, nil
	} else {
		return &CallResult{ContentType: "json", Data: ss}, nil
	}
}

type Python struct {
}

func (p *Python) Call(req request.Call) (*CallResult, error) {
	return nil, nil
}
func (p *Python) CallerPath() string {
	return ""
}
