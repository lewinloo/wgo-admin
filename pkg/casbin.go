package pkg

import (
	"fmt"
	"gin_template/internal/global"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 初始化casbin策略管理器
func InitCasbinEnforcer() {
	e, err := mysqlCasbin()
	if err != nil {
		global.LOG.Error(fmt.Sprintf("初始化Casbin失败：%v", err))
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}

	global.CASBIN = e
	global.LOG.Info("初始化Casbin完成!")

}

func mysqlCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(global.DB)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("rbac_model.conf", a)
	if err != nil {
		return nil, err
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}