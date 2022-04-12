package example_plugin

// 我们为您准备了一个需要提供api的插件模板 您只需要按照此模板注册路由即可 插件使用请把此包放置在plugin下
import (
	"github.com/gin-gonic/gin"
)

var ExamplePlugin = new(pluginExample)

type pluginExample struct{}

func NewPluginExample() *pluginExample {
	// 此处为注册生命周期 可以在此处写初始化内容
	return &pluginExample{}
}

func (*pluginExample) Register(group *gin.RouterGroup) {
	//如需细分权限 可以在此处use中间件 gva项目包名已改为github模式
	//所以整个plugin可以直接独立到外层开启为新的项目 然后用包的形式导入也是可以完整运行的

	group.GET("hello",
		func(context *gin.Context) {
			//	此处请填写handle函数
			//	您依然可以模仿gva分层进行插件制作 当然您也可以按照您所习惯的分层模式开发
			context.JSON(200, "hello world")
		})
}

func (*pluginExample) RouterPath() string {
	// 此处为您插件的总路由path 录入本插件安装进入项目后 会自动产生路由 /[主项目跟路由（如果有的话）]/group/xxxx
	return "group"
}
