package model

type ExcelInfo struct {
	FileName string        `json:"fileName"`
	InfoList []SysBaseMenu `json:"infoList"`
}
