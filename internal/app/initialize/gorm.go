package initialize

import (
	"gin_template/internal/app/config"
	"gin_template/internal/app/domain/model"
	"gin_template/pkg"
	"gorm.io/gorm"
)

func GormDB() (*gorm.DB, func(), error) {
	cfg := config.C
	db := pkg.NewDB(*cfg)

	cleanFunc := func() {}

	err := model.AutoMigrate(db)
	if err != nil {
		return nil, cleanFunc, err
	}

	return db, cleanFunc, nil
}
