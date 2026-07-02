package media

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FileUploadAndDownload struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:文件名"`                                // 文件名
	ClassId int    `json:"classId" form:"classId" gorm:"default:0;type:int;column:class_id;comment:分类id;"` // 分类id
	Url     string `json:"url" form:"url" gorm:"column:url;comment:文件地址"`                                  // 文件地址
	Tag     string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签"`                                  // 文件标签
	Key     string `json:"key" form:"key" gorm:"column:key;comment:编号"`                                    // 编号
	Size    int64  `json:"size" form:"size" gorm:"column:size;comment:文件大小(字节);default:0"`                // 文件大小(字节)
	Mime    string `json:"mime" form:"mime" gorm:"column:mime;type:varchar(255);comment:MIME类型"`            // MIME类型
	Md5     string `json:"md5" form:"md5" gorm:"column:md5;type:varchar(64);index;comment:文件MD5"`           // 文件MD5
	UserID  uint   `json:"userId" form:"userId" gorm:"column:user_id;index;comment:上传者ID"`                 // 上传者ID
}

func (FileUploadAndDownload) TableName() string {
	return "media_file_upload_and_downloads"
}
