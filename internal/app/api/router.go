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
	MenuHandler  *handler.MenuHandler
	RoleHandler  *handler.RoleHandler
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
	}

	gPrivate := gBase.Group("")
	gPrivate.Use(middleware.CheckAuth(), middleware.CheckPermission(r.Casbin))
	{
		// 用户模块
		gUser := gPrivate.Group("user")
		{
			gUser.POST("register", r.UserHandler.Register)
			gUser.DELETE("delete", r.UserHandler.Delete)
			gUser.POST("list", r.UserHandler.List)

		}

		// 菜单模块
		gMenu := gPrivate.Group("menu")
		{
			gMenu.POST("", r.MenuHandler.Create)
			gMenu.GET("tree", r.MenuHandler.FindMenuTree)
			gMenu.DELETE(":id", r.MenuHandler.DeleteById)
		}

		// 角色模块
		gRole := gPrivate.Group("role")
		{
			gRole.POST("", r.RoleHandler.Create)
			gRole.POST("list", r.RoleHandler.FindList)
			gRole.DELETE(":id", r.RoleHandler.DeleteById)
		}

	}
}
