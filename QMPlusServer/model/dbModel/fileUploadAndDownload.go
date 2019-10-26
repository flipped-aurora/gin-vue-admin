package dbModel

import (
	"github.com/jinzhu/gorm"
	"main/controller/servers"
	"main/init/qmsql"
	"main/model/modelInterface"
)

type FileUploadAndDownload struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
	Tag  string `json:"tag"`
}

func (f *FileUploadAndDownload) Upload() error {
	err := qmsql.DEFAULTDB.Create(f).Error
	return err
}

// 分页获取数据  需要分页实现这个接口即可
func (f *FileUploadAndDownload) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(f, info)
	if err != nil {
		return
	} else {
		var fileLists []FileUploadAndDownload
		err = db.Order("updated_at desc").Find(&fileLists).Error
		return err, fileLists, total
	}
}
