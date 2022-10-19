package api

import (
	"gin_template/internal/app/api/handler"
	"gin_template/internal/app/api/middleware"
	"gin_template/internal/app/config"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
}

type Router struct {
	Casbin       *casbin.Enforcer
	HelloHandler *handler.HelloHandler
	UserHandler  *handler.UserHandler
}

func (r *Router) Register(app *gin.Engine) error {
	r.RegisterApi(app)
	return nil
}

func (r *Router) RegisterApi(engine *gin.Engine) {
	gBase := engine.Group(config.C.System.GlobalPrefix)

	gPublic := gBase.Group("")
	{
		gPublic.GET("/hello", r.HelloHandler.Hello)

		// 用户模块
		gPublic.POST("/user/login", r.UserHandler.Login)
		gPublic.POST("/user/list", r.UserHandler.List)
	}

	gPrivate := gBase.Group("")
	gPrivate.Use(middleware.CheckAuth(), middleware.CheckPermission(r.Casbin))
	{
		// 用户模块
		gUser := gPrivate.Group("user")
		{
			gUser.POST("register", r.UserHandler.Register)

		}
	}
}

// 初始化总路由
//func New() *gin.Engine {
//	// 依赖注入 handler
//	injector, _, _ := app.BuildHandlerInjector()
//
//	Router := gin.New()
//	Router.Use(gin.Logger(), middleware.Recovery(), middleware.Cors())
//
//	// swagger 文档地址
//	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//
//	apiRouter := Router.Group(global.CONFIG.System.GlobalPrefix)
//
//	// 公共路由
//	publicRoutes := apiRouter.Group("")
//	{
//		publicRoutes.GET("/hello", injector.HelloHandler.Hello)
//
//		// 用户模块
//		publicRoutes.POST("/user/login", injector.UserHandler.Login)
//		publicRoutes.POST("/user/list", injector.UserHandler.List)
//	}
//
//	// 鉴权认证路由
//	privateRoutes := apiRouter.Group("")
//	privateRoutes.Use(middleware.CheckAuth(), middleware.CheckPermission())
//	{
//		// 用户模块
//		gUser := privateRoutes.Group("user")
//		{
//			gUser.POST("register", injector.UserHandler.Register)
//
//		}
//	}
//
//	return Router
//}
