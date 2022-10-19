package initialize

import (
	"gin_template/pkg"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

func Casbin(db *gorm.DB) (*casbin.Enforcer, func(), error) {
	enforcer, err := pkg.NewCasbinEnforcer(db)
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {}

	return enforcer, cleanFunc, nil
}
