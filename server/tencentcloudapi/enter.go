package tencentcloudapi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

func getClient() *common.Credential {
	credential := common.NewCredential(
		global.GVA_CONFIG.TxCloud.SecretId,
		global.GVA_CONFIG.TxCloud.SecretKey,
	)
	return credential
}
