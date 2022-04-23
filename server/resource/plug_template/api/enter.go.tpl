package api

type ApiGroup struct {
	{{ .PlugName}}Api
}

var ApiGroupApp = new(ApiGroup)
