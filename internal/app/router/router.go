package router

import (
	_ "gin_template/docs"
	"gin_template/internal/app/api"
	"gin_template/internal/app/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化 handler
var (
	helloApi = api.NewHello()
	userApi  = api.NewUser()
)

// 初始化总路由
func New() *gin.Engine {
	Router := gin.Default()

	// 设置生产模式
	if global.CONFIG.System.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// swagger 文档地址
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := Router.Group(global.CONFIG.System.GlobalPrefix)

	publicRoutes := apiRouter.Group("")
	{
		publicRoutes.GET("/hello", helloApi.Hello)
		publicRoutes.POST("/user/register", userApi.Register)
		publicRoutes.POST("/user/login", userApi.Login)
	}

	// privateRoutes := apiRouter.Group("")
	// {

	// }

	global.LOG.Info("路由注册成功")
	return Router
}
