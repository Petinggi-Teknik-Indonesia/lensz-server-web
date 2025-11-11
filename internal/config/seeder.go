package config

import (
	"log"
	"lensz-server-web/internal/model"
	"gorm.io/gorm"
)

// Seed inserts initial data (roles, admin user, etc.)
func Seed(db *gorm.DB) {
	// Seed roles
	roles := []model.Role{
		{Name: "Admin"},
		{Name: "User"},
	}

	for _, role := range roles {
		var existing model.Role
		if err := db.FirstOrCreate(&existing, model.Role{Name: role.Name}).Error; err != nil {
			log.Printf("❌ Failed to seed role %s: %v", role.Name, err)
		}
	}

	// Seed organization
	org := model.Organization{Name: "Default Organization"}
	if err := db.FirstOrCreate(&org, model.Organization{Name: org.Name}).Error; err != nil {
		log.Printf("❌ Failed to seed organization: %v", err)
	}

	// Seed admin user
	var admin model.User
	if err := db.Where("email = ?", "admin@example.com").First(&admin).Error; err != nil {
		admin = model.User{
			Name:           "Admin",
			Email:          "admin@example.com",
			Phone:          "08123456789",
			RoleID:         1,
			OrganizationID: org.ID,
		}
		_ = admin.HashPassword("password123")
		if err := db.Create(&admin).Error; err != nil {
			log.Printf("❌ Failed to seed admin user: %v", err)
		}
	}

	log.Println("🌱 Database seeding complete")
}
