package initialize

import (
	"fmt"
	"go-learn/global"
	"go-learn/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Gorm Connection
func Gorm() *gorm.DB {
	db := GormPG()
	AutoMigrations(db)

	return db
}

// AutoMigrations DB
func AutoMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Movie{},
		&model.User{},
		&model.Article{},
	)
	if err != nil {
		os.Exit(0)
	}
}

// GormPG Connection
func GormPG() *gorm.DB {
	c := global.CONFIG.PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s TimeZone=%s", c.Host, c.Port, c.Username, c.Password, c.Dbname, c.Timezone)
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger})

	if err != nil {
		os.Exit(0)
		return nil
	}
	return db
}
