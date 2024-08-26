package biz_apphub

import (
	"bytes"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func NewCaller(types string) Caller {
	return &SoftCall{}
}

type Caller interface {
	Call(req request.Call) (string, error)
}

type SoftCall struct {
	AppPath string
}

func (p *SoftCall) Call(req request.Call) (string, error) {
	softPath := fmt.Sprintf("D:\\code\\github.com\\apphub\\server\\soft_cmd\\%s\\%s\\%s",
		req.User, req.Soft, req.Soft)
	reqJson := fmt.Sprintf("D:\\code\\github.com\\apphub\\server\\soft_cmd\\%s\\%s\\%v_%v.json",
		req.User, req.Soft, req.Soft, time.Now().UnixNano())

	if runtime.GOOS == "windows" {
		softPath = fmt.Sprintf("D:\\code\\github.com\\apphub\\server\\soft_cmd\\%s\\%s\\%s.exe",
			req.User, req.Soft, req.Soft)
	}

	err := jsonx.SaveFile(reqJson, req.Data)
	if err != nil {
		return "", err
	}
	cmd := exec.Command(softPath, req.Command, reqJson)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	s := out.String()
	split := strings.Split(s, "<Response>")[1]
	split1 := strings.Split(split, "</Response>")[0]
	ss := strings.ReplaceAll(split1, "</Response>", "")
	return ss, nil
}

type Python struct {
}

func (p *Python) Call(req request.Call) ([]byte, error) {
	return nil, nil
}
