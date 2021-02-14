package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

func ParseInfoList2Excel(infoList []model.SysBaseMenu, filePath string) error {
	excel := excelize.NewFile()
		excel.SetSheetRow("Sheet1","A1",&[]string{"ID","路由Name","路由Path","是否隐藏","父节点","排序","文件名称"})
    	for i, menu := range infoList {
    		axis := fmt.Sprintf("A%d",i+2)
    		excel.SetSheetRow("Sheet1",axis,&[]interface{}{
    			menu.ID,
    			menu.Name,
    			menu.Path,
    			menu.Hidden,
    			menu.ParentId,
    			menu.Sort,
    			menu.Component,
    		})
		}
	excel.SaveAs(filePath)
	return nil
}

func ParseExcel2InfoList() ([]model.SysBaseMenu, error) {
	skipHeader := true
	fixedHeader := []string{"ID","路由Name","路由Path","是否隐藏","父节点","排序","文件名称"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir+"ExcelImport.xlsx")
	if err != nil {
        return nil, err
	}
	menus := make([]model.SysBaseMenu, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		hidden, _ := strconv.ParseBool(row[3])
		sort, _ := strconv.Atoi(row[5])
		menu := model.SysBaseMenu{
			GVA_MODEL: global.GVA_MODEL{
				ID: uint(id),
			},
			Name: row[1],
			Path: row[2],
			Hidden: hidden,
			ParentId: row[4],
			Sort: sort,
			Component: row[6],
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}