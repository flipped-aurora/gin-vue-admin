package router

import (
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

type PprofRouter struct{}

func (s *PprofRouter) InitPprofRouter(Router *gin.RouterGroup) {
	pprofRouter := Router // .Use(middleware.OperationRecord())
	{
		pprofRouter.GET("/", pprofHandler(pprof.Index))
		pprofRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
		pprofRouter.GET("/profile", pprofHandler(pprof.Profile))
		pprofRouter.POST("/symbol", pprofHandler(pprof.Symbol))
		pprofRouter.GET("/symbol", pprofHandler(pprof.Symbol))
		pprofRouter.GET("/trace", pprofHandler(pprof.Trace))
		pprofRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		pprofRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		pprofRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		pprofRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		pprofRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		pprofRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}
}

// pprofHandler, transfer `http.HandlerFunc` to `gin.HandlerFunc`

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
