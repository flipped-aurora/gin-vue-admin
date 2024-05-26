package system

import (
	"errors"
	"fmt"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"gorm.io/gorm"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"go.uber.org/zap"
)

var RepeatErr = errors.New("重复创建")

type AutoCodeHistoryService struct{}

var AutoCodeHistoryServiceApp = new(AutoCodeHistoryService)

// CreateAutoCodeHistory 创建代码生成器历史记录
// RouterPath : RouterPath@RouterString;RouterPath2@RouterString2
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) CreateAutoCodeHistory(meta, structName, structCNName, autoCodePath string, injectionMeta string, tableName string, apiIds string, Package string, BusinessDB string, menuID uint) error {
	return global.GVA_DB.Create(&system.SysAutoCodeHistory{
		Package:       Package,
		RequestMeta:   meta,
		AutoCodePath:  autoCodePath,
		InjectionMeta: injectionMeta,
		StructName:    structName,
		StructCNName:  structCNName,
		TableName:     tableName,
		ApiIDs:        apiIds,
		BusinessDB:    BusinessDB,
		MenuID:        menuID,
	}).Error
}

// First 根据id获取代码生成器历史的数据
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) First(info *request.GetById) (string, error) {
	var meta string
	return meta, global.GVA_DB.Model(system.SysAutoCodeHistory{}).Select("request_meta").Where("id = ?", info.Uint()).First(&meta).Error
}

// Repeat 检测重复
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) Repeat(businessDB, structName, Package string) bool {
	var count int64
	global.GVA_DB.Model(&system.SysAutoCodeHistory{}).Where("business_db = ? and struct_name = ? and package = ? and flag = 0", businessDB, structName, Package).Count(&count)
	return count > 0
}

// RollBack 回滚
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) RollBack(info *systemReq.RollBack) error {
	md := system.SysAutoCodeHistory{}
	if err := global.GVA_DB.Where("id = ?", info.ID).First(&md).Error; err != nil {
		return err
	}
	// 清除API表

	ids := request.IdsReq{}
	idsStr := strings.Split(md.ApiIDs, ";")
	for i := range idsStr[0 : len(idsStr)-1] {
		id, err := strconv.Atoi(idsStr[i])
		if err != nil {
			return err
		}
		ids.Ids = append(ids.Ids, id)
	}
	err := ApiServiceApp.DeleteApisByIds(ids)

	if err != nil {
		global.GVA_LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
	}

	err = BaseMenuServiceApp.DeleteBaseMenu(int(md.MenuID))

	if err != nil {
		global.GVA_LOG.Error("ClearTag DeleteBaseMenu:", zap.Error(err))
	}

	// 删除表
	if info.DeleteTable {
		if err = AutoCodeServiceApp.DropTable(md.BusinessDB, md.TableName); err != nil {
			global.GVA_LOG.Error("ClearTag DropTable:", zap.Error(err))
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
		delteSavePath := filepath.Join(global.GVA_CONFIG.AutoCode.Root,
			"rm_file", time.Now().Format("20060102"), filepath.Base(filepath.Dir(filepath.Dir(path))), filepath.Base(filepath.Dir(path)), filepath.Base(path))
		// 判断目标文件是否存在
		for utils.FileExist(delteSavePath) {
			fmt.Println("文件已存在:", delteSavePath)
			delteSavePath += fmt.Sprintf("_%d", time.Now().Nanosecond())
		}
		parts := strings.Split(path, "gin-vue-admin")
		//从gin-vue-admin文件开始 例如serve\
		FilePath := strings.TrimPrefix(parts[1], "\\")
		var record system.RecordsDeleteCode
		result := global.GVA_DB.Table("records").Where("file = ?", FilePath).First(&record)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			record = system.RecordsDeleteCode{
				Path:       delteSavePath,
				File:       FilePath,
				UpdateTime: time.Now(),
			}
			if err := global.GVA_DB.Create(&record).Error; err != nil {
				log.Fatal()
			}

		} else if result.Error != nil {
			log.Fatal(result.Error)
		} else {
			//更新已有记录
			record.Path = delteSavePath
			record.UpdateTime = time.Now()
			if err := global.GVA_DB.Model(&system.RecordsDeleteCode{}).Where("file = ?", FilePath).Updates(&record).Error; err != nil {
				log.Fatal(err)
			}
		}
		err = utils.FileMove(path, delteSavePath)
		if err != nil {
			global.GVA_LOG.Error("file move err ", zap.Error(err))
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

	ast.RollBackAst(md.Package, md.StructName)

	md.Flag = 1
	return global.GVA_DB.Save(&md).Error
}

// Delete 删除历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) Delete(info *request.GetById) error {
	return global.GVA_DB.Where("id = ?", info.Uint()).Delete(&system.SysAutoCodeHistory{}).Error
}

// GetList 获取系统历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) GetList(info request.PageInfo) (list []response.AutoCodeHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAutoCodeHistory{})
	var entities []response.AutoCodeHistory
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&entities).Error
	return entities, total, err
}
