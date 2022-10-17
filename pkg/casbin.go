package pkg

import (
	"gorm.io/gorm"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 初始化casbin策略管理器
func NewCasbinEnforcer(db *gorm.DB) (e *casbin.Enforcer, err error) {
	e, err = mysqlCasbin(db)
	return e, err
}

func mysqlCasbin(db *gorm.DB) (*casbin.Enforcer, error) {
	rootPath, _ := os.Getwd()
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(rootPath+"/resources/rbac_model.conf", a)
	if err != nil {
		return nil, err
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}
