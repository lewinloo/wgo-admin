//go:build wireinject
// +build wireinject

package app

import (
	"gin_template/internal/app/api"
	"gin_template/internal/app/api/handler"
	"gin_template/internal/app/domain/repository"
	"gin_template/internal/app/domain/service"
	"gin_template/internal/app/initialize"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine *gin.Engine
}

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		initialize.GormDB,
		initialize.GinEngine,
		initialize.Casbin,
		initialize.Redis,
		initialize.Log,
		repository.ProvideSet,
		service.ProvideSet,
		handler.ProvideSet,
		api.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
