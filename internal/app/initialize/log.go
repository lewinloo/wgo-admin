package initialize

import (
	"gin_template/pkg"
	"go.uber.org/zap"
)

func Log() (*zap.Logger, func(), error) {
	log := pkg.Zap()
	cleanFunc := func() {}
	zap.ReplaceGlobals(log)

	return log, cleanFunc, nil
}
