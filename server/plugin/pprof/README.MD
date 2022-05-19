## GVA Pprof&Trace 功能插件
By using Go’s profiling tools to identify and correct specific bottlenecks, we can make programm faster and shrink resource occupation.
#### 开发者：trumanhe0@gmail.com

### 使用步骤

#### 1. 前往GVA主程序下的initialize/router.go 在Routers 方法最末尾按照你需要的及安全模式添加本插件
    例：
    本插件可以采用gva的配置文件 也可以直接写死内容作为配置 建议为gva添加配置文件结构 然后将配置传入:
	PluginInit(PrivateGroup, pprof.CreatePprofPlug(global.GVA_CONFIG.Pprof.Prefix))

    同样也可以在传入时写死:
    PluginInit(PrivateGroup, pprof.CreatePprofPlug("debug/pprof"))

### 2. 配置说明

#### 2-1 全局配置结构体说明

    type Pprof struct {
        Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"` // path prefix
    }


### 3. 方法API
All routers are listed in `plugin\pprof\router\pprof.go`
* pprofRouter.GET("/", pprofHandler(pprof.Index)) 
Index responds with the pprof-formatted profile named by the request. For example, "/debug/pprof/heap" serves the "heap" profile.

* pprofRouter.GET("/cmdline", pprofHandler(pprof.Cmdline)) 
Cmdline responds with the running program's command line, with arguments separated by NUL bytes. The package initialization registers it as /debug/pprof/cmdline.

* pprofRouter.GET("/profile", pprofHandler(pprof.Profile)) 
Profile responds with the pprof-formatted cpu profile. Profiling lasts for duration specified in seconds GET parameter, or for 30 seconds if not specified. The package initialization registers it as /debug/pprof/profile.

* pprofRouter.POST("/symbol", pprofHandler(pprof.Symbol)) 
Symbol looks up the program counters listed in the request, responding with a table mapping program counters to function names. The package initialization registers it as /debug/pprof/symbol.

* pprofRouter.GET("/symbol", pprofHandler(pprof.Symbol))
Symbol looks up the program counters listed in the request, responding with a table mapping program counters to function names. The package initialization registers it as /debug/pprof/symbol.

* pprofRouter.GET("/trace", pprofHandler(pprof.Trace))
Trace responds with the execution trace in binary form. Tracing lasts for duration specified in seconds GET parameter, or for 1 second if not specified. The package initialization registers it as /debug/pprof/trace.

* pprofRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
The endpoint "/allocs" responds with the memory allocations of your running programm.

* pprofRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
The endpoint "/block" responds with the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program.

* pprofRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
The endpoint "/gorotine" responds with the gorotine profile of your running programm.

* pprofRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
The endpoint "/heap" responds with the heap profile of your running programm.

* pprofRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
The endpoint "/mutex" responds with the holders of contended mutexes, after calling runtime.SetMutexProfileFraction in your program.

* pprofRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
The endpoint "/threadcreate" responds with the created threads profile of your programm.

### Reference:  
* Package Doc, https://pkg.go.dev/net/http/pprof   
* Profiling Go Programs-The go blog, https://go.dev/blog/pprof
