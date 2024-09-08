package biz_apphub

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
    biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
)

type BizCloudFunctionService struct {}

// CreateBizCloudFunction 创建云函数记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCloudFunctionService *BizCloudFunctionService) CreateBizCloudFunction(bizCloudFunction *biz_apphub.BizCloudFunction) (err error) {
	err = global.GVA_DB.Create(bizCloudFunction).Error
	return err
}

// DeleteBizCloudFunction 删除云函数记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCloudFunctionService *BizCloudFunctionService)DeleteBizCloudFunction(ID string) (err error) {
	err = global.GVA_DB.Delete(&biz_apphub.BizCloudFunction{},"id = ?",ID).Error
	return err
}

// DeleteBizCloudFunctionByIds 批量删除云函数记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCloudFunctionService *BizCloudFunctionService)DeleteBizCloudFunctionByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]biz_apphub.BizCloudFunction{},"id in ?",IDs).Error
	return err
}

// UpdateBizCloudFunction 更新云函数记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCloudFunctionService *BizCloudFunctionService)UpdateBizCloudFunction(bizCloudFunction biz_apphub.BizCloudFunction) (err error) {
	err = global.GVA_DB.Model(&biz_apphub.BizCloudFunction{}).Where("id = ?",bizCloudFunction.ID).Updates(&bizCloudFunction).Error
	return err
}

// GetBizCloudFunction 根据ID获取云函数记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCloudFunctionService *BizCloudFunctionService)GetBizCloudFunction(ID string) (bizCloudFunction biz_apphub.BizCloudFunction, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bizCloudFunction).Error
	return
}

// GetBizCloudFunctionInfoList 分页获取云函数记录
// Author [piexlmax](https://github.com/piexlmax)
func (bizCloudFunctionService *BizCloudFunctionService)GetBizCloudFunctionInfoList(info biz_apphubReq.BizCloudFunctionSearch) (list []biz_apphub.BizCloudFunction, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&biz_apphub.BizCloudFunction{})
    var bizCloudFunctions []biz_apphub.BizCloudFunction
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CnName != "" {
        db = db.Where("cn_name LIKE ?","%"+ info.CnName+"%")
    }
    if info.CodeName != "" {
        db = db.Where("code_name LIKE ?","%"+ info.CodeName+"%")
    }
    if info.Classify != "" {
        db = db.Where("classify = ?",info.Classify)
    }
    if info.ExecMode != "" {
        db = db.Where("exec_mode = ?",info.ExecMode)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Content != "" {
        db = db.Where("content LIKE ?","%"+ info.Content+"%")
    }
    if info.ContentType != "" {
        db = db.Where("content_type = ?",info.ContentType)
    }
    if info.IsPublic != "" {
        db = db.Where("is_public = ?",info.IsPublic)
    }
    if info.Tags != "" {
        db = db.Where("tags LIKE ?","%"+ info.Tags+"%")
    }
        if info.StartViews != nil && info.EndViews != nil {
            db = db.Where("views BETWEEN ? AND ? ",info.StartViews,info.EndViews)
        }
        if info.StartExecCount != nil && info.EndExecCount != nil {
            db = db.Where("exec_count BETWEEN ? AND ? ",info.StartExecCount,info.EndExecCount)
        }
        if info.StartColl != nil && info.EndColl != nil {
            db = db.Where("coll BETWEEN ? AND ? ",info.StartColl,info.EndColl)
        }
        if info.StartLike != nil && info.EndLike != nil {
            db = db.Where("like BETWEEN ? AND ? ",info.StartLike,info.EndLike)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["views"] = true
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
	
	err = db.Find(&bizCloudFunctions).Error
	return  bizCloudFunctions, total, err
}