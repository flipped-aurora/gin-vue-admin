package dto

import "time"

type RunnerRunResp struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
	//ContentType    string      `json:"content_type"`
	//HasFile        bool        `json:"has_file"`
	FilePath string `json:"path"`
	//DeleteFile     bool        `json:"delete_file"`
	DeleteFileTime int         `json:"delete_file_time"`
	Body           interface{} `json:"data"`

	Header map[string]string `json:"header"` // response header

	//meta data
	CallCostTime     time.Duration `json:"-"` //
	ResponseMetaData string        `json:"-"`
}

type RunnerRunReq struct {
	StatusCode       int               `json:"status_code"`
	Msg              string            `json:"msg"`
	FilePath         string            `json:"path"`
	DeleteFileTime   int               `json:"delete_file_time"`
	Body             interface{}       `json:"data"`
	Header           map[string]string `json:"header"` // response header
	CallCostTime     time.Duration     `json:"-"`      //
	ResponseMetaData string            `json:"-"`
}
