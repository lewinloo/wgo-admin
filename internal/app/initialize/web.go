package initialize

import (
	_ "gin_template/docs"
	"gin_template/internal/app/api"
	"gin_template/internal/app/api/middleware"
	"gin_template/internal/app/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GinEngine(r api.IRouter) *gin.Engine {
	// 设置生产模式
	if config.C.System.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()

	app.Use(gin.Logger(), middleware.Recovery(), middleware.Cors())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = r.Register(app)

	return app
}
