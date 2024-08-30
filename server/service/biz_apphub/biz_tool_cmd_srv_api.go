package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"gorm.io/gorm"
)

type BizToolCmdSrvApiService struct{}

// CreateBizToolCmdSrvApi 创建后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizToolCmdSrvApiService *BizToolCmdSrvApiService) CreateBizToolCmdSrvApi(bizToolCmdSrvApi *biz_apphub.BizToolCmdSrvApi) (err error) {
	err = global.GVA_DB.Create(bizToolCmdSrvApi).Error
	return err
}

// DeleteBizToolCmdSrvApi 删除后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizToolCmdSrvApiService *BizToolCmdSrvApiService) DeleteBizToolCmdSrvApi(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&biz_apphub.BizToolCmdSrvApi{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&biz_apphub.BizToolCmdSrvApi{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBizToolCmdSrvApiByIds 批量删除后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizToolCmdSrvApiService *BizToolCmdSrvApiService) DeleteBizToolCmdSrvApiByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&biz_apphub.BizToolCmdSrvApi{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&biz_apphub.BizToolCmdSrvApi{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBizToolCmdSrvApi 更新后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizToolCmdSrvApiService *BizToolCmdSrvApiService) UpdateBizToolCmdSrvApi(bizToolCmdSrvApi biz_apphub.BizToolCmdSrvApi) (err error) {
	err = global.GVA_DB.Model(&biz_apphub.BizToolCmdSrvApi{}).Where("id = ?", bizToolCmdSrvApi.ID).Updates(&bizToolCmdSrvApi).Error
	return err
}

// GetBizToolCmdSrvApi 根据ID获取后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizToolCmdSrvApiService *BizToolCmdSrvApiService) GetBizToolCmdSrvApi(ID string) (bizToolCmdSrvApi biz_apphub.BizToolCmdSrvApi, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bizToolCmdSrvApi).Error
	return
}

// GetBizToolCmdSrvApiInfoList 分页获取后端工具指令api记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizToolCmdSrvApiService *BizToolCmdSrvApiService) GetBizToolCmdSrvApiInfoList(info biz_apphubReq.BizToolCmdSrvApiSearch) (list []biz_apphub.BizToolCmdSrvApi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&biz_apphub.BizToolCmdSrvApi{})
	var bizToolCmdSrvApis []biz_apphub.BizToolCmdSrvApi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AppName != "" {
		db = db.Where("app_name = ?", info.AppName)
	}
	if info.AppCode != "" {
		db = db.Where("app_code = ?", info.AppCode)
	}
	if info.ToolType != "" {
		db = db.Where("tool_type = ?", info.ToolType)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	if info.Classify != "" {
		db = db.Where("classify = ?", info.Classify)
	}
	if info.Version != "" {
		db = db.Where("version = ?", info.Version)
	}
	if info.Mode != "" {
		db = db.Where("mode = ?", info.Mode)
	}
	if info.Tags != "" {
		db = db.Where("tags LIKE ?", "%"+info.Tags+"%")
	}
	if info.Video != "" {
		db = db.Where("video = ?", info.Video)
	}
	if info.ApiPath != "" {
		db = db.Where("api_path LIKE ?", "%"+info.ApiPath+"%")
	}
	if info.IsPublic != "" {
		db = db.Where("is_public = ?", info.IsPublic)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["version"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&bizToolCmdSrvApis).Error
	return bizToolCmdSrvApis, total, err
}
