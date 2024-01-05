package utils

import (
	"fmt"
	"net/url"
	"strings"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func routerprefix() string {
	prefix := global.GVA_CONFIG.System.RouterPrefix
	if prefix != "" && !strings.HasPrefix(prefix, "/") {
		prefix = fmt.Sprintf("/%s", prefix)
	}
	return prefix
}

func FixRouterPrefix_CasbinRule(ruleList []adapter.CasbinRule) {
	routerprefix := routerprefix()
	for idx, rule := range ruleList {
		ruleList[idx].V1, _ = url.JoinPath(routerprefix, rule.V1)
	}
}

func FixRouterPrefix_SysApi(apiList []sysModel.SysApi) {
	routerprefix := routerprefix()
	for idx, api := range apiList {
		apiList[idx].Path, _ = url.JoinPath(routerprefix, api.Path)
	}
}
