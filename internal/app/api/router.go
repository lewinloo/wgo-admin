package api

import (
	_ "gin_template/docs"
	"gin_template/internal/app"
	"gin_template/internal/app/api/middleware"
	"gin_template/internal/app/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由
func New() *gin.Engine {
	// 依赖注入 handler
	injector, _, _ := app.BuildHandlerInjector()

	Router := gin.New()
	Router.Use(gin.Logger(), middleware.Recovery(), middleware.Cors())

	// swagger 文档地址
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := Router.Group(global.CONFIG.System.GlobalPrefix)

	// 公共路由
	publicRoutes := apiRouter.Group("")
	{
		publicRoutes.GET("/hello", injector.HelloHandler.Hello)

		// 用户模块
		publicRoutes.POST("/user/login", injector.UserHandler.Login)
		publicRoutes.POST("/user/list", injector.UserHandler.List)
	}

	// 鉴权认证路由
	privateRoutes := apiRouter.Group("")
	privateRoutes.Use(middleware.CheckAuth(), middleware.CheckPermission())
	{
		// 用户模块
		gUser := privateRoutes.Group("user")
		{
			gUser.POST("register", injector.UserHandler.Register)

		}
	}

	return Router
}
