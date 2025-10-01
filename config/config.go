package config

import (
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lensz-server-web/internal/model"
)

type Config struct {
	DBUrl      string
	ServerPort string
}

func Load() Config {
	return Config{
		DBUrl:      os.Getenv("DATABASE_URL"),
		ServerPort: os.Getenv("PORT"),
	}
}

func InitDB(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate all models here
	err = db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
