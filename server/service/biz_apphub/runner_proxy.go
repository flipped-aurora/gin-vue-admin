package biz_apphub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/stringsx"
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner/response"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
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

// GetSoftLogs softRunTime 程序执行耗时，callCostTime：调用+执行的耗时
func GetSoftLogs(callResponse *response.CallResponse) {
	softRunTime := ""
	softRunTimeList := stringsx.ParserHtmlTagContent(callResponse.ResponseMetaData, "UserCost")
	if len(softRunTimeList) > 0 {
		softRunTime = softRunTimeList[0]
	}
	logs := stringsx.ParserHtmlTagContent(callResponse.ResponseMetaData, "Logger")
	for _, v := range logs {
		mp := make(map[string]interface{})
		err := json.Unmarshal([]byte(v), &mp)
		if err != nil {
			continue
		}
		msg := mp["msg"]
		delete(mp, "msg")
		mp["soft_run_time"] = softRunTime
		mp["call_cost_time"] = callResponse.CallCostTime.String()
		logrus.WithFields(mp).Info(msg)
	}
}

func (p *SoftCall) Call(req request.Call) (*response.CallResponse, error) {
	now := time.Now()
	softPath := p.AppPath + fmt.Sprintf("/%s/%s/%s",
		req.User, req.Soft, req.Soft)
	if runtime.GOOS == "windows" {
		softPath = p.AppPath + fmt.Sprintf("/%s/%s/%s.exe",
			req.User, req.Soft, req.Soft)
	}
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
		//todo
		panic("")
	}
	resList := stringsx.ParserHtmlTagContent(s, "Response")
	if len(resList) == 0 {
		//todo 请使用sdk开发软件
		return nil, fmt.Errorf("soft call err 请使用sdk开发软件")
	}
	//split := strings.Split(s, "<Response>")[1]
	//split1 := strings.Split(split, "</Response>")[0]
	//ss := strings.ReplaceAll(split1, "</Response>", "")
	//mp := make(map[string]interface{})
	var res response.CallResponse
	err = json.Unmarshal([]byte(resList[0]), &res)
	if err != nil {
		return nil, err
	}
	since := time.Since(now)
	res.CallCostTime = since
	res.ResponseMetaData = s
	//p.printSoftLogs(s, since)
	logrus.WithFields(logrus.Fields{})
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
