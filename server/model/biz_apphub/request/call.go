package request

import (
	"encoding/json"
	"fmt"
	"time"
)

type Call struct {
	User    string   `json:"user"`    //软件所属的用户
	Soft    string   `json:"soft"`    //软件名
	Command string   `json:"command"` //命令
	Method  string   `json:"method"`  //请求方式
	Files   []string `json:"files"`

	UpdateVersion   bool                   `json:"update_version"`    //此时正处于版本更新的状态
	RequestJsonPath string                 `json:"request_json_path"` //请求参数存储路径
	Data            map[string]interface{} `json:"data"`              //请求json
	ReqBody         string
}

func (c *Call) RequestJSON() (string, error) {
	j, err := json.Marshal(c.Data)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func (c *Call) GetRequestFilePath(callerPath string) string {
	reqJson := callerPath + fmt.Sprintf("\\%s\\%s\\%v_%v.json",
		c.User, c.Soft, c.Soft, time.Now().UnixNano())
	return reqJson
}
