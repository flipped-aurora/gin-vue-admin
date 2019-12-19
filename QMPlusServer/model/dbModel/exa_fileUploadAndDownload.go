package dbModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type ExaFileUploadAndDownload struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
	Tag  string `json:"tag"`
	Key  string `json:"key"`
}

func (f *ExaFileUploadAndDownload) Upload() error {
	err := qmsql.DEFAULTDB.Create(f).Error
	return err
}

func (f *ExaFileUploadAndDownload) DeleteFile() error {
	err := qmsql.DEFAULTDB.Where("id = ?", f.ID).Unscoped().Delete(f).Error
	return err
}

func (f *ExaFileUploadAndDownload) FindFile() (error, ExaFileUploadAndDownload) {
	var file ExaFileUploadAndDownload
	err := qmsql.DEFAULTDB.Where("id = ?", f.ID).First(&file).Error
	return err, file
}

// 分页获取数据  需要分页实现这个接口即可
func (f *ExaFileUploadAndDownload) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(f, info)
	if err != nil {
		return
	} else {
		var fileLists []ExaFileUploadAndDownload
		err = db.Order("updated_at desc").Find(&fileLists).Error
		return err, fileLists, total
	}
}
