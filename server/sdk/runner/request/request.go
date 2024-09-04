package request

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
