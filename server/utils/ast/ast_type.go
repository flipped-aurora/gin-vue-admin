package ast

type Type string

func (r Type) String() string {
	return string(r)
}

func (r Type) Group() string {
	switch r {
	case TypePackageApiEnter:
		return "ApiGroup"
	case TypePackageRouterEnter:
		return "RouterGroup"
	case TypePackageServiceEnter:
		return "ServiceGroup"
	case TypePackageApiModuleEnter:
		return "ApiGroup"
	case TypePackageRouterModuleEnter:
		return "RouterGroup"
	case TypePackageServiceModuleEnter:
		return "ServiceGroup"
	case TypePluginApiEnter:
		return "api"
	case TypePluginRouterEnter:
		return "router"
	case TypePluginServiceEnter:
		return "service"
	default:
		return ""
	}
}

const (
	TypePackageApiEnter           = "PackageApiEnter"           // server/api/v1/enter.go
	TypePackageRouterEnter        = "PackageRouterEnter"        // server/router/enter.go
	TypePackageServiceEnter       = "PackageServiceEnter"       // server/service/enter.go
	TypePackageApiModuleEnter     = "PackageApiModuleEnter"     // server/api/v1/{package}/enter.go
	TypePackageRouterModuleEnter  = "PackageRouterModuleEnter"  // server/router/{package}/enter.go
	TypePackageServiceModuleEnter = "PackageServiceModuleEnter" // server/service/{package}/enter.go
	TypePackageInitializeGorm     = "PackageInitializeGorm"     // server/initialize/gorm_biz.go
	TypePackageInitializeRouter   = "PackageInitializeRouter"   // server/initialize/router_biz.go
	TypePluginGen                 = "PluginGen"                 // server/plugin/{package}/gen/main.go
	TypePluginApiEnter            = "PluginApiEnter"            // server/plugin/{package}/enter.go
	TypePluginInitializeV1        = "PluginInitializeV1"        // server/initialize/plugin_biz_v1.go
	TypePluginInitializeV2        = "PluginInitializeV2"        // server/initialize/plugin_biz_v2.go
	TypePluginRouterEnter         = "PluginRouterEnter"         // server/plugin/{package}/enter.go
	TypePluginServiceEnter        = "PluginServiceEnter"        // server/plugin/{package}/enter.go
	TypePluginInitializeApi       = "PluginInitializeApi"       // server/plugin/{package}/initialize/api.go
	TypePluginInitializeGorm      = "PluginInitializeGorm"      // server/plugin/{package}/initialize/gorm.go
	TypePluginInitializeMenu      = "PluginInitializeMenu"      // server/plugin/{package}/initialize/menu.go
	TypePluginInitializeRouter    = "PluginInitializeRouter"    // server/plugin/{package}/initialize/router.go
)
