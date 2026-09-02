package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	autoModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(
		new(autoModel.SysAIWorkflowSession),
		new(autoModel.SysAutoCodeHistory),
		new(autoModel.SysAutoCodePackage),
	)
	if err != nil {
		err = errors.Wrap(err, "register auto plugin tables failed")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
