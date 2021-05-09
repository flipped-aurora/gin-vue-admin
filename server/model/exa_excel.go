package model

type ExcelInfo struct {
	FileName string        `json:"fileName"` // 文件名
	InfoList []SysBaseMenu `json:"infoList"`
}
