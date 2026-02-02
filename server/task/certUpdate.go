package task

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cert_manager/service"
	"go.uber.org/zap"
)

func UpdateAllCertificates() error {
	err := service.ServiceGroupApp.CertCertificate.UpdateAllCertificates()
	if err != nil {
		global.GVA_LOG.Error("批量更新证书失败!", zap.Error(err))
		return fmt.Errorf("批量更新证书失败: %v", err)
	}
	global.GVA_LOG.Info("批量更新证书成功")
	return nil
}
