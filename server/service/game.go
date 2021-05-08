package service

import (
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func OpenConnection(connection request.CreateConnection) (err error, token model.ConnectionToken) {
	var (
		ok      bool
		t       uuid.UUID
		list    interface{}
		plugin  *model.ProductPlugin
		plugins []model.ProductPlugin
		info    = request.ProductPluginSearch{
			PageInfo: request.PageInfo{Page: 1, PageSize: 1000},
		}
	)

	if err, list, _ = GetProductPluginInfoList(info); err != nil {
		return
	}

	if list == nil {
		err = errors.New("游戏插件列表为空")
		return
	}

	if plugins, ok = list.([]model.ProductPlugin); !ok {
		err = errors.New("游戏插件配置错误")
		return
	}

	for idx, _ := range plugins {
		if connection.ID == plugins[idx].ProductID {
			plugin = &plugins[idx]
			break
		}
	}

	if plugin == nil {
		err = errors.Errorf("游戏插件未配置: %s", connection.ID)
		return
	}

	if t, err = uuid.NewV4(); err != nil {
		err = errors.New("生成token失败")
		return
	}

	token = model.ConnectionToken{Token: t.String()}

	return
}

func CloseConnection(connection request.CloseConnection) (err error) {
	return
}

func GameRequest(request request.GameRequest) (err error, data interface{}) {
	return
}
