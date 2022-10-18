package api

import (
	_ "gin_template/docs"
	"gin_template/internal/app/api/handler"
	"gin_template/internal/app/api/middleware"
	"gin_template/internal/app/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化 handler
var (
	helloHandler = handler.NewHello()
	userHandler  = handler.NewUser()
)

// 初始化总路由
func New() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Logger(), middleware.Recovery(), middleware.Cors())

	// swagger 文档地址
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := Router.Group(global.CONFIG.System.GlobalPrefix)

	// 公共路由
	publicRoutes := apiRouter.Group("")
	{
		publicRoutes.GET("/hello", helloHandler.Hello)

		// 用户模块
		publicRoutes.POST("/user/login", userHandler.Login)
	}

	// 鉴权认证路由
	privateRoutes := apiRouter.Group("")
	privateRoutes.Use(middleware.CheckAuth(), middleware.CheckPermission())
	{
		// 用户模块
		privateRoutes.POST("/user/register", userHandler.Register)
	}

	return Router
}
