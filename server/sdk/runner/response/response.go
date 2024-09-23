package response

import "time"

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
	Body        interface{} `json:"data"`

	Header map[string]string `json:"header"` // response header

	//meta data
	CallCostTime     time.Duration `json:"-"` //
	ResponseMetaData string        `json:"-"`
}

func (r *CallResponse) GetContentType() string {
	if r.Header == nil {
		return ""
	}
	return r.Header["content-type"]
}
