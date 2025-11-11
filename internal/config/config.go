package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"lensz-server-web/internal/model"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	JWTSecret  string
}

func Load() Config {
	// Try to load .env if available
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, relying on system environment variables")
	}

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("SERVER_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

func InitDB(cfg Config) (*gorm.DB, error) {
	// DSN for PostgreSQL
	dsn := "host=" + cfg.DBHost +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" port=" + cfg.DBPort +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate ALL models
	err = db.AutoMigrate(
		&model.Drawer{},
		&model.Brand{},
		&model.Company{},
		&model.Glasses{},
		&model.StatusHistory{},
		&model.Role{},
		&model.Organization{},
		&model.User{},
		&model.Scanner{},
	)
	if err != nil {
		return nil, err
	}

	log.Println("✅ Database migration complete")
	return db, nil
}
