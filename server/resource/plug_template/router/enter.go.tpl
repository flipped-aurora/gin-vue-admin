package router

type RouterGroup struct {
	{{ .PlugName}}Router
}

var RouterGroupApp = new(RouterGroup)
