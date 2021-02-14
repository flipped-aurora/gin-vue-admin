package request

import "gin-vue-admin/model"

type ExcelInfo struct {
	FileName string `json:"fileName"`
	InfoList []model.SysBaseMenu `json:"infoList"`
}