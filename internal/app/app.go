package app

import (
  "fmt"
  "gin_template/internal/app/config"
  "gin_template/internal/app/initialize"
  "gin_template/pkg"
  "github.com/fvbock/endless"
  "github.com/gin-gonic/gin"
  "go.uber.org/zap"
  "time"
)

type Core struct {
  Log *zap.Logger
}

func LoadResource(path string) *Core {
  c := new(Core)
  initialize.Config(path, config.C)
  c.Log = pkg.Zap()

  return c
}

func (c *Core) Run() {

  // 依赖注入
  injector, clearFunc, _ := BuildInjector()
  clearFunc()

  address := fmt.Sprintf(":%d", config.C.System.Port)
  s := initServer(address, injector.Engine)

  time.Sleep(10 * time.Microsecond)
  c.Log.Info("server run success on ", zap.String("port", address))

  fmt.Printf(`
	Swagger文档地址:http://127.0.0.1%s/swagger/index.html
	服务地址:http://127.0.0.1%s%s
	`, address, address, config.C.System.GlobalPrefix)
  c.Log.Error(s.ListenAndServe().Error())
}

type server interface {
  ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
  s := endless.NewServer(address, router)
  s.ReadHeaderTimeout = 20 * time.Second
  s.WriteTimeout = 20 * time.Second
  s.MaxHeaderBytes = 1 << 20
  return s
}
