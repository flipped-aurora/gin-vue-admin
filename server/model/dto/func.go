package dto

import "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"

type SyncFunctionReq struct {
	ID uint `json:"id" form:"id"`

	Runner *biz_apphub.BizToolCmdSrvApi
}
type Api struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Config
}

type Config struct {
	ApiDesc     string `json:"api_desc"`
	IsPublicApi bool   `json:"is_public_api"`

	ChineseName string      `json:"chinese_name"`
	EnglishName string      `json:"english_name"`
	Classify    string      `json:"classify"`
	Tags        string      `json:"tags"`
	Params      []FuncParam `json:"params"`

	Request  interface{} `json:"request,omitempty"`
	Response interface{} `json:"response,omitempty"`
}

type FuncParam struct {
	Code          string `json:"code,omitempty"`
	Desc          string `json:"desc,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Type          string `json:"type,omitempty"`
	Value         string `json:"value,omitempty"`
	Options       string `json:"options,omitempty"`
	Required      string `json:"required,omitempty"`
	MockData      string `json:"mock_data,omitempty"`
	InputMode     string `json:"input_mode,omitempty"`
	TextLimit     string `json:"text_limit,omitempty"`
	NumberLimit   string `json:"number_limit,omitempty"`
	SelectOptions string `json:"select_options,omitempty"`
	FileSizeLimit string `json:"file_size_limit,omitempty"`
	FileTypeLimit string `json:"file_type_limit,omitempty"`
}

type SyncFunctionResp struct {
	Apis []Api `json:"apis"`
}

type Response[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}
