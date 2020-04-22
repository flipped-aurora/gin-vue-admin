package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCoffeeRouter(Router *gin.RouterGroup) {
	coffeeRouter := Router.Group("/coffee").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		coffeeRouter.POST("/add", api.AddCoffee)
		coffeeRouter.PATCH("/update", api.UpdateCoffee)
		coffeeRouter.DELETE("/delete", api.DeleteCoffee)
		coffeeRouter.POST("getbyid", api.GetCoffeeByUUID)
		coffeeRouter.POST("/type", api.ChangeCoffeeType)
	}

	coffeeTypeRouter := Router.Group("/coffeetype").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		coffeeTypeRouter.POST("/add", api.AddCoffeeType)
		coffeeTypeRouter.PATCH("/update", api.UpdateCoffeeType)
		coffeeTypeRouter.DELETE("/delete", api.DeleteCoffeeType)
		coffeeTypeRouter.POST("/getbycode", api.GetCoffeeByCode)
	}

	coffeeSpecRouter := Router.Group("/coffeespec").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		coffeeSpecRouter.POST("/getbycoffeeid", api.GetCoffeeSpecByCoffeeId)
		coffeeSpecRouter.POST("/add", api.AddCoffeeSpec)
		coffeeSpecRouter.PATCH("/update", api.UpdateCoffeeSpec)
		coffeeSpecRouter.DELETE("/delete", api.DeleteCoffeeSpec)
		coffeeSpecRouter.POST("/getbyid", api.GetCoffeeSpecById)
		coffeeSpecRouter.POST("/detail", api.GetCoffeeSpecDetail)
		coffeeSpecRouter.POST("/addDetail", api.AddCoffeeSpecDetail)
	}
}
