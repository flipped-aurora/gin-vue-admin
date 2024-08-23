package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type GetDeployList struct {
	request.PageInfo
	AppId       string `json:"appId" form:"appId"`
	Sort        string `json:"sort" form:"sort"`
	Order       string `json:"order" form:"order"`
	OperateUser string
}

type RollbackVersion struct {
	Appid       string `json:"appId" form:"appId"`
	RecordID    string `json:"recordId" form:"recordId"`
	OperateUser string
}
