package router

import (
	"gin_template/internal/global"
	helloHandlers "gin_template/internal/handlers/hello"
	userHandlers "gin_template/internal/handlers/user"

	_ "gin_template/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由
func New() *gin.Engine {
	Router := gin.Default()

	// swagger 文档地址
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := Router.Group(global.CONFIG.System.GlobalPrefix)

	publicRoutes := apiRouter.Group("")
	{
		publicRoutes.GET("/hello", helloHandlers.Hello)
		publicRoutes.POST("/user/register", userHandlers.Register)
	}

	// privateRoutes := apiRouter.Group("")
	// {

	// }

	global.LOG.Info("路由注册成功")
	return Router
}
