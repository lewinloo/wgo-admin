//go:build wireinject
// +build wireinject

package app

import (
	"gin_template/internal/app/api/handler"
	"gin_template/internal/app/domain/repository"
	"gin_template/internal/app/domain/service"
	"github.com/google/wire"
)

var HandlerInjectorSet = wire.NewSet(wire.Struct(new(HandlerInjector), "*"))

type HandlerInjector struct {
	HelloHandler *handler.HelloHandler
	UserHandler  *handler.UserHandler
}

func BuildHandlerInjector() (*HandlerInjector, func(), error) {
	wire.Build(
		repository.ProvideSet,
		service.ProvideSet,
		handler.ProvideSet,
		HandlerInjectorSet)

	return new(HandlerInjector), nil, nil
}
