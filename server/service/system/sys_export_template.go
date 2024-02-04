package system

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"mime/multipart"
	"strings"
)

type SysExportTemplateService struct {
}

// CreateSysExportTemplate 创建导出模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) CreateSysExportTemplate(sysExportTemplate *system.SysExportTemplate) (err error) {
	err = global.GVA_DB.Create(sysExportTemplate).Error
	return err
}

// DeleteSysExportTemplate 删除导出模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) DeleteSysExportTemplate(sysExportTemplate system.SysExportTemplate) (err error) {
	err = global.GVA_DB.Delete(&sysExportTemplate).Error
	return err
}

// DeleteSysExportTemplateByIds 批量删除导出模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) DeleteSysExportTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysExportTemplate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSysExportTemplate 更新导出模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) UpdateSysExportTemplate(sysExportTemplate system.SysExportTemplate) (err error) {
	err = global.GVA_DB.Save(&sysExportTemplate).Error
	return err
}

// GetSysExportTemplate 根据id获取导出模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) GetSysExportTemplate(id uint) (sysExportTemplate system.SysExportTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysExportTemplate).Error
	return
}

// GetSysExportTemplateInfoList 分页获取导出模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) GetSysExportTemplateInfoList(info systemReq.SysExportTemplateSearch) (list []system.SysExportTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysExportTemplate{})
	var sysExportTemplates []system.SysExportTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.TableName != "" {
		db = db.Where("table_name = ?", info.TableName)
	}
	if info.TemplateID != "" {
		db = db.Where("template_id = ?", info.TemplateID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysExportTemplates).Error
	return sysExportTemplates, total, err
}

// ExportExcel 导出Excel
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) ExportExcel(templateID string) (file *bytes.Buffer, name string, err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return nil, "", err
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var templateInfoMap = make(map[string]string)
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return nil, "", err
	}
	var columns []string
	var tableTitle []string
	for key := range templateInfoMap {
		columns = append(columns, key)
		tableTitle = append(tableTitle, templateInfoMap[key])
	}
	selects := strings.Join(columns, ", ")
	var tableMap []map[string]interface{}
	err = global.GVA_DB.Select(selects).Table(template.TableName).Find(&tableMap).Error
	if err != nil {
		return nil, "", err
	}
	var rows [][]string
	rows = append(rows, tableTitle)
	for _, table := range tableMap {
		var row []string
		for _, column := range columns {
			row = append(row, fmt.Sprintf("%v", table[column]))
		}
		rows = append(rows, row)
	}
	for i, row := range rows {
		for j, colCell := range row {
			err := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", getColumnName(j), i+1), colCell)
			if err != nil {
				return nil, "", err
			}
		}
	}
	f.SetActiveSheet(index)
	file, err = f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	return file, template.Name, nil
}

// ExportTemplate 导出Excel模板
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) ExportTemplate(templateID string) (file *bytes.Buffer, name string, err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return nil, "", err
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var templateInfoMap = make(map[string]string)
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return nil, "", err
	}
	var tableTitle []string
	for key := range templateInfoMap {
		tableTitle = append(tableTitle, templateInfoMap[key])
	}
	for i := range tableTitle {
		fErr := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", getColumnName(i), 1), tableTitle[i])
		if fErr != nil {
			return nil, "", fErr
		}
	}
	f.SetActiveSheet(index)
	file, err = f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	return file, template.Name, nil
}

// ImportExcel 导入Excel
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) ImportExcel(templateID string, file *multipart.FileHeader) (err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	var templateInfoMap = make(map[string]string)
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return err
	}

	var titleKeyMap = make(map[string]string)
	for key, title := range templateInfoMap {
		titleKeyMap[title] = key
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]
		for _, row := range values {
			var item = make(map[string]interface{})
			for ii, value := range row {
				key := titleKeyMap[excelTitle[ii]]
				item[key] = value
			}
			cErr := tx.Table(template.TableName).Create(&item).Error
			if cErr != nil {
				return cErr
			}
		}
		return nil
	})
}

func getColumnName(n int) string {
	columnName := ""
	for n > 0 {
		n--
		columnName = string(rune('A'+n%26)) + columnName
		n /= 26
	}
	return columnName
}
