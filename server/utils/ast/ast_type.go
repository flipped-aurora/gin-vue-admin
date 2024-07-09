package ast

type Type string

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
		return "Api"
	case TypePluginRouterEnter:
		return "Router"
	case TypePluginServiceEnter:
		return "Service"
	default:
		return ""
	}
}

const (
	TypePackageApiEnter           = "PackageEnter"              // server/api/v1/enter.go
	TypePackageRouterEnter        = "PackageRouterEnter"        // server/router/enter.go
	TypePackageServiceEnter       = "PackageServiceEnter"       // server/service/enter.go
	TypePackageApiModuleEnter     = "PackageApiModuleEnter"     // server/api/v1/{package}/enter.go
	TypePackageRouterModuleEnter  = "PackageRouterModuleEnter"  // server/router/{package}/enter.go
	TypePackageServiceModuleEnter = "PackageServiceModuleEnter" // server/service/{package}/enter.go
	TypeInitializeGorm            = "InitializeGorm"            // server/initialize/gorm_biz.go
	TypeInitializeRouter          = "InitializeRouter"          // server/initialize/router_biz.go
	TypeInitializePlugin          = "InitializePlugin"          // server/initialize/plugin_biz.go
	TypePluginApiEnter            = "PluginApiEnter"            // server/plugin/{package}/enter.go
	TypePluginRouterEnter         = "PluginRouterEnter"         // server/plugin/{package}/enter.go
	TypePluginServiceEnter        = "PluginServiceEnter"        // server/plugin/{package}/enter.go
)
