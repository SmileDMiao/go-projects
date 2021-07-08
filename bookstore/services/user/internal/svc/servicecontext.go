package svc

import (
	"bookstore/services/user/internal/config"
	"bookstore/services/user/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(postgres.Open(c.Postgresql.DataSource), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&model.User{},
	)

	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}
