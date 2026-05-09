package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wubi/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(
		new(model.Score),
	)
	if err != nil {
		err = errors.Wrap(err, "wubi 插件注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
