package initialize

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{
			Path:        "/info/createInfo",
			Description: global.Translate("system.api.desc.newAnnouncement"),
			ApiGroup:    global.Translate("system.api.group.announcement"),
			Method:      "POST",
		},
		{
			Path:        "/info/deleteInfo",
			Description: global.Translate("system.api.desc.deleteAnnouncement"),
			ApiGroup:    global.Translate("system.api.group.announcement"),
			Method:      "DELETE",
		},
		{
			Path:        "/info/deleteInfoByIds",
			Description: global.Translate("system.api.desc.batchDeleteAnnouncement"),
			ApiGroup:    global.Translate("system.api.group.announcement"),
			Method:      "DELETE",
		},
		{
			Path:        "/info/updateInfo",
			Description: global.Translate("system.api.desc.updateAnnouncement"),
			ApiGroup:    global.Translate("system.api.group.announcement"),
			Method:      "PUT",
		},
		{
			Path:        "/info/findInfo",
			Description: global.Translate("system.api.desc.getAnnouncementByID"),
			ApiGroup:    global.Translate("system.api.group.announcement"),
			Method:      "GET",
		},
		{
			Path:        "/info/getInfoList",
			Description: global.Translate("system.api.desc.getAnnouncementList"),
			ApiGroup:    global.Translate("system.api.group.announcement"),
			Method:      "GET",
		},
	}
	utils.RegisterApis(entities...)
}
