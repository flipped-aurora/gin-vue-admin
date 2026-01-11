package initialize

import (
	"context"
	model "{{.Module}}/model/system"
	"{{.Module}}/plugin/plugin-tool/utils"
)

func Dictionary(ctx context.Context) {
	entities := []model.SysDictionary{}
	utils.RegisterDictionaries(entities...)
}
