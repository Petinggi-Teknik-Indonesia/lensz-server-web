package main

import (
	"lensz-server-web/config" // adjust to your module name

	"gorm.io/gen"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load env file
	configs.LoadEnv()

	// connect DB
	db, err := gorm.Open(postgres.Open(configs.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("❌ failed to connect database")
	}

	// initialize generator
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/domain/models", // folder for models
	})

	g.UseDB(db)

	// generate models for all tables
	g.GenerateAllTable()

	// execute
	g.Execute()
}
