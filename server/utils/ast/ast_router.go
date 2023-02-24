package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

const src = `
package initialize

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)



func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GVA_LOG.Info("use middleware cors")
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                 // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)                // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)              // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
		systemRouter.InitAutoCodeRouter(PrivateGroup)            // 创建自动化代码
		systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)     // 自动化代码历史
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)  // 字典详情管理

		exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	}

	{
		testRouterRouter := router.RouterGroupApp.TestRouter
		testRouterRouter.InitGVA222StructRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}

`

func AppendNodeToList(stmts []ast.Stmt, stmt ast.Stmt, index int) []ast.Stmt {
	return append(stmts[:index], append([]ast.Stmt{stmt}, stmts[index:]...)...)
}

func AddRouterCode(path, funcName, pk, model string) {
	_, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", src, parser.ParseComments)

	if err != nil {
		fmt.Println(err)
	}

	FuncNode := FindFunction(astFile, funcName)

	pkName := strings.ToUpper(pk[:1]) + pk[1:]
	routerName := fmt.Sprintf("%sRouter", pk)
	modelName := fmt.Sprintf("Init%sRouter", model)
	var bloctPre *ast.BlockStmt
	for i := len(FuncNode.Body.List) - 1; i >= 0; i-- {
		if block, ok := FuncNode.Body.List[i].(*ast.BlockStmt); ok {
			bloctPre = block
		}
	}
	ast.Print(fileSet, FuncNode)
	if ok, b := needAppendRouter(FuncNode, pk); ok {
		routerNode :=
			&ast.BlockStmt{
				List: []ast.Stmt{
					&ast.AssignStmt{
						Lhs: []ast.Expr{
							&ast.Ident{Name: routerName},
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.SelectorExpr{
								X: &ast.SelectorExpr{
									X:   &ast.Ident{Name: "router"},
									Sel: &ast.Ident{Name: "RouterGroupApp"},
								},
								Sel: &ast.Ident{Name: pkName},
							},
						},
					},
				},
			}

		FuncNode.Body.List = AppendNodeToList(FuncNode.Body.List, routerNode, len(FuncNode.Body.List)-2)
		bloctPre = routerNode
	} else {
		bloctPre = b
	}

	if needAppendInit(FuncNode, routerName, modelName) {
		bloctPre.List = append(bloctPre.List,
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: routerName},
						Sel: &ast.Ident{Name: modelName},
					},
					Args: []ast.Expr{
						&ast.Ident{
							Name: "PrivateGroup",
						},
					},
				},
			})
	}
	var out []byte
	bf := bytes.NewBuffer(out)
	printer.Fprint(bf, fileSet, astFile)
	fmt.Println(bf.String())
}

func needAppendRouter(funcNode ast.Node, pk string) (bool, *ast.BlockStmt) {
	flag := true
	var block *ast.BlockStmt
	ast.Inspect(funcNode, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.BlockStmt:
			for i := range n.List {
				if assignNode, ok := n.List[i].(*ast.AssignStmt); ok {
					if identNode, ok := assignNode.Lhs[0].(*ast.Ident); ok {
						if identNode.Name == fmt.Sprintf("%sRouter", pk) {
							flag = false
							block = n
							return false
						}
					}
				}
			}

		}
		return true
	})
	return flag, block
}

func needAppendInit(funcNode ast.Node, routerName string, modelName string) bool {
	flag := true
	ast.Inspect(funcNode, func(node ast.Node) bool {
		switch n := funcNode.(type) {
		case *ast.CallExpr:
			if selectNode, ok := n.Fun.(*ast.SelectorExpr); ok {
				x, xok := selectNode.X.(*ast.Ident)
				if xok && x.Name == routerName && selectNode.Sel.Name == modelName {
					flag = false
					return false
				}
			}
		}
		return true
	})
	return flag
}
