package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/article/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(
		new(model.Article),
		new(model.Category),
	)
	if err != nil {
		err = errors.Wrap(err, "article 插件注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
