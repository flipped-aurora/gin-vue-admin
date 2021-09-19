package system

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"go.uber.org/zap"
)

var RepeatErr = errors.New("重复创建")

type AutoCodeHistoryService struct {
}

var AutoCodeHistoryServiceApp = new(AutoCodeHistoryService)

func (autoCodeHistoryService *AutoCodeHistoryService) Repeat(structName string) bool {

	var count int64
	global.GVA_DB.Model(&system.SysAutoCodeHistory{}).Where("struct_name = ? and flag = 0", structName).Count(&count)
	return count > 0
}

// CreateAutoCodeHistory RouterPath : RouterPath@RouterString;RouterPath2@RouterString2
func (autoCodeHistoryService *AutoCodeHistoryService) CreateAutoCodeHistory(meta, structName, structCNName, autoCodePath string, injectionMeta string, tableName string, apiIds string) error {
	return global.GVA_DB.Create(&system.SysAutoCodeHistory{
		RequestMeta:   meta,
		AutoCodePath:  autoCodePath,
		InjectionMeta: injectionMeta,
		StructName:    structName,
		StructCNName:  structCNName,
		TableName:     tableName,
		ApiIDs:        apiIds,
	}).Error
}

// RollBack 回滚
func (autoCodeHistoryService *AutoCodeHistoryService) RollBack(id uint) error {
	md := system.SysAutoCodeHistory{}
	if err := global.GVA_DB.First(&md, id).Error; err != nil {
		return err
	}
	// 清除API表
	err := ApiServiceApp.DeleteApiByIds(strings.Split(md.ApiIDs, ";"))
	if err != nil {
		global.GVA_LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
	}
	// 获取全部表名
	err, dbNames := AutoCodeServiceApp.GetTables(global.GVA_CONFIG.Mysql.Dbname)
	if err != nil {
		global.GVA_LOG.Error("ClearTag GetTables:", zap.Error(err))
	}
	// 删除表
	for _, name := range dbNames {
		if strings.Contains(strings.ToUpper(strings.Replace(name.TableName, "_", "", -1)), strings.ToUpper(md.TableName)) {
			// 删除表
			if err = AutoCodeServiceApp.DropTable(name.TableName); err != nil {
				global.GVA_LOG.Error("ClearTag DropTable:", zap.Error(err))

			}
		}
	}
	// 删除文件

	for _, path := range strings.Split(md.AutoCodePath, ";") {

		// 增加安全判断补丁:
		_path, err := filepath.Abs(path)
		if err != nil || _path != path {
			continue
		}

		// 迁移
		nPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root,
			"rm_file", time.Now().Format("20060102"), filepath.Base(filepath.Dir(filepath.Dir(path))), filepath.Base(filepath.Dir(path)), filepath.Base(path))
		err = utils.FileMove(path, nPath)
		if err != nil {
			fmt.Println(">>>>>>>>>>>>>>>>>>>", err)
		}
		//_ = utils.DeLFile(path)
	}
	// 清除注入
	for _, v := range strings.Split(md.InjectionMeta, ";") {
		// RouterPath@functionName@RouterString
		meta := strings.Split(v, "@")
		if len(meta) == 3 {
			_ = utils.AutoClearCode(meta[0], meta[2])
		}
	}
	md.Flag = 1
	return global.GVA_DB.Save(&md).Error
}

func (autoCodeHistoryService *AutoCodeHistoryService) GetMeta(id uint) (string, error) {
	var meta string
	return meta, global.GVA_DB.Model(system.SysAutoCodeHistory{}).Select("request_meta").First(&meta, id).Error
}

// GetSysHistoryPage  获取系统历史数据
func (autoCodeHistoryService *AutoCodeHistoryService) GetSysHistoryPage(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAutoCodeHistory{})
	var fileLists []system.SysAutoCodeHistory
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Select("id,created_at,updated_at,struct_name,struct_cn_name,flag,table_name").Find(&fileLists).Error
	return err, fileLists, total
}

// DeletePage 删除历史数据
func (autoCodeHistoryService *AutoCodeHistoryService) DeletePage(id uint) error {
	return global.GVA_DB.Delete(system.SysAutoCodeHistory{}, id).Error
}
