package config

import (
	"context"
	"fmt"
	"log"

	"lensz-server-web/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedDatabase(ctx context.Context, db *gorm.DB, cfg Config) error {
	log.Println("🌱 Starting database seeding...")

	// === 1️⃣ Seed Roles ===
	roles := []model.Role{
		{Name: "Admin"},
		{Name: "Backdoor"},
		{Name: "User"},
	}

	for _, role := range roles {
		var existing model.Role
		err := db.WithContext(ctx).First(&existing, "name = ?", role.Name).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.WithContext(ctx).Create(&role).Error; err != nil {
				return fmt.Errorf("failed to create role %s: %w", role.Name, err)
			}
			log.Printf("✅ Role created: %s\n", role.Name)
		}
	}

	// === 2️⃣ Seed Organization ===
	org := model.Organization{Name: "Optic Gembira"}
	var existingOrg model.Organization
	err := db.WithContext(ctx).First(&existingOrg, "name = ?", org.Name).Error
	if err == gorm.ErrRecordNotFound {
		if err := db.WithContext(ctx).Create(&org).Error; err != nil {
			return fmt.Errorf("failed to create organization: %w", err)
		}
		log.Println("✅ Organization created: Optic Gembira")
	} else {
		org = existingOrg
	}

	// === 3️⃣ Seed Admin User (Backdoor) ===
	var backdoorRole model.Role
	if err := db.WithContext(ctx).First(&backdoorRole, "name = ?", "Backdoor").Error; err != nil {
		return fmt.Errorf("backdoor role not found: %w", err)
	}

	var existingUser model.User
	err = db.WithContext(ctx).First(&existingUser, "email = ?", "deswandy88@gmail.com").Error
	if err == gorm.ErrRecordNotFound {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)

		adminUser := model.User{
			Name:           cfg.AdminName,
			Email:          cfg.AdminEmail,
			Password:       string(hashed),
			RoleID:         backdoorRole.ID,
			OrganizationID: org.ID,
		}

		if err := db.WithContext(ctx).Create(&adminUser).Error; err != nil {
			return fmt.Errorf("failed to create backdoor admin: %w", err)
		}

		log.Println("✅ Backdoor admin created: deswandy88@gmail.com (password: admin123)")
	} else {
		log.Println("ℹ️ Backdoor admin already exists, skipping.")
	}

	log.Println("🌳 Seeding complete.")
	return nil
}
