package biz_apphub

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/httpx"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/oss"
	"os"
)

// GetDeployList 发布历史
func (bizAppHubService *BizAppHubService) GetDeployList(req request.GetDeployList) (res []biz_apphub.BizAppHubRecord, totalCount int64, err error) {
	db := global.GVA_DB.Model(&biz_apphub.BizAppHubRecord{}).Where("app_id = ?", req.AppId)
	err = db.Offset(req.GetOffset()).Limit(req.GetLimit()).Order("created_at desc").Find(&res).Error
	if err != nil {
		return nil, 0, err
	}
	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

// RollbackVersion 版本回滚
func (bizAppHubService *BizAppHubService) RollbackVersion(req request.RollbackVersion) (err error) {
	var record biz_apphub.BizAppHubRecord
	var app biz_apphub.BizAppHub
	err = global.GVA_DB.Model(&biz_apphub.BizAppHubRecord{}).Where("id =? ", req.RecordID).First(&record).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Model(&biz_apphub.BizAppHub{}).Where("id =? ", req.Appid).First(&app).Error
	if err != nil {
		return err
	}

	tempPath := fmt.Sprintf("./temp/%s/%s/", req.OperateUser, app.AppCode)
	err = os.MkdirAll(tempPath, 0755)
	if err != nil {
		return err
	}
	file := tempPath + "index.html"
	//下载index文件
	url := "http://cdn.geeleo.com/" + record.GetOssStorePath()
	err = httpx.DownloadFile(url, file)
	if err != nil {
		return err
	}
	defer os.ReadFile(file)

	store := oss.NewDefaultQiNiu()
	deleteOss := app.GetUseOssPathV2()
	err = store.DeleteFile(deleteOss)
	if err != nil {
		fmt.Println(err)
	}

	_, err = store.UploadLocalFile(file, deleteOss) //版本回滚后，更新app版本
	if err != nil {
		return err
	}

	ups := map[string]interface{}{
		"app_name":     record.AppName,
		"title":        record.Title,
		"desc":         record.Desc,
		"classify":     record.Classify,
		"version":      record.Version,
		"mode":         record.Mode,
		"develop_mode": record.DevelopMode,
		"cover":        record.Cover,
		"tags":         record.Tags,
		"video":        record.Video,
	}

	err = global.GVA_DB.Model(&biz_apphub.BizAppHub{}).Where("id = ?", req.Appid).Updates(ups).Error
	if err != nil {
		return err
	}
	return nil
}
