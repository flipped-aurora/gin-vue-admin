package system

import (
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// SysAutoCodeHistory 自动迁移代码记录,用于回滚,重放使用
type SysAutoCodeHistory struct {
	global.GVA_MODEL
	Package       string `json:"package"`
	BusinessDB    string `json:"businessDB"`
	TableName     string `json:"tableName"`
	MenuID        uint   `json:"menuID"`
	RequestMeta   string `gorm:"type:text" json:"requestMeta,omitempty"`   // 前端传入的结构化信息
	AutoCodePath  string `gorm:"type:text" json:"autoCodePath,omitempty"`  // 其他meta信息 path;path
	InjectionMeta string `gorm:"type:text" json:"injectionMeta,omitempty"` // 注入的内容 RouterPath@functionName@RouterString;
	StructName    string `json:"structName"`
	StructCNName  string `json:"structCNName"`
	ApiIDs        string `json:"apiIDs,omitempty"` // api表注册内容
	Flag          int    `json:"flag"`             // 表示对应状态 0 代表创建, 1 代表回滚 ...
}

// ToRequestIds ApiIDs 转换 request.IdsReq
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *SysAutoCodeHistory) ToRequestIds() request.IdsReq {
	if m.ApiIDs == "" {
		return request.IdsReq{}
	}
	slice := strings.Split(m.ApiIDs, ";")
	ids := make([]int, 0, len(slice))
	length := len(slice)
	for i := 0; i < length; i++ {
		id, _ := strconv.ParseInt(slice[i], 10, 32)
		ids = append(ids, int(id))
	}
	return request.IdsReq{Ids: ids}
}

// 记录删除文件的路径
type RecordsDeleteCode struct {
	Path       string    `gorm:"type:text;comment:已删除文件保存路径"`
	File       string    `gorm:"type:text;comment:新生成文件路径"`
	UpdateTime time.Time `gorm:"type:datetime;comment:更新时间"`
}

func (RecordsDeleteCode) TableName() string {
	return "records"
}
