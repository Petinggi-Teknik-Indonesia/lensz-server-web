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

	// ------------------------------
	// 2️⃣ Seed Organizations
	// ------------------------------
	orgNames := []string{"Optic Gembira", "Asean Baru"}
	orgMap := make(map[string]model.Organization)

	for _, name := range orgNames {
		var org model.Organization
		err := db.WithContext(ctx).First(&org, "name = ?", name).Error
		if err == gorm.ErrRecordNotFound {
			org = model.Organization{Name: name}
			if err := db.WithContext(ctx).Create(&org).Error; err != nil {
				return fmt.Errorf("failed to seed organization %s: %w", name, err)
			}
			log.Println("🏢 Organization:", name)
		}
		orgMap[name] = org
	}

	// === 3️⃣ Seed Admin User (Backdoor) ===
	var backdoorRole model.Role
	if err := db.WithContext(ctx).First(&backdoorRole, "name = ?", "Backdoor").Error; err != nil {
		return fmt.Errorf("backdoor role not found: %w", err)
	}

	var existingUser model.User
	err := db.WithContext(ctx).First(&existingUser, "email = ?", "deswandy88@gmail.com").Error
	if err == gorm.ErrRecordNotFound {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)

		adminUser := model.User{
			Name:           cfg.AdminName,
			Email:          cfg.AdminEmail,
			Password:       string(hashed),
			RoleID:         backdoorRole.ID,
			OrganizationID: orgMap["Optic Gembira"].ID,
		}

		if err := db.WithContext(ctx).Create(&adminUser).Error; err != nil {
			return fmt.Errorf("failed to create backdoor admin: %w", err)
		}

		log.Println("✅ Backdoor admin created: deswandy88@gmail.com (password: admin123)")
	} else {
		log.Println("ℹ️ Backdoor admin already exists, skipping.")
	}

		// ------------------------------
	// 4️⃣ Seed Drawers
	// ------------------------------
	drawers := []string{"A", "B", "C", "D", "E", "F", "G"}
	drawerMap := make(map[string]model.Drawer)

	for _, name := range drawers {
		var d model.Drawer
		err := db.WithContext(ctx).First(&d, "name = ?", name).Error
		if err == gorm.ErrRecordNotFound {
			d = model.Drawer{Name: name}
			if err := db.WithContext(ctx).Create(&d).Error; err != nil {
				return fmt.Errorf("failed to seed drawer %s: %w", name, err)
			}
			log.Println("📦 Drawer:", name)
		}
		drawerMap[name] = d
	}

	// ------------------------------
	// 5️⃣ Seed Brands
	// ------------------------------
	brands := []string{"RayBan", "Cartier", "Adidas", "Nike", "Oakley"}
	brandMap := make(map[string]model.Brand)

	for _, name := range brands {
		var b model.Brand
		err := db.WithContext(ctx).First(&b, "name = ?", name).Error
		if err == gorm.ErrRecordNotFound {
			b = model.Brand{Name: name}
			if err := db.WithContext(ctx).Create(&b).Error; err != nil {
				return fmt.Errorf("failed to seed brand %s: %w", name, err)
			}
			log.Println("🕶️ Brand:", name)
		}
		brandMap[name] = b
	}

	// ------------------------------
	// 6️⃣ Seed Companies
	// ------------------------------
	companies := []string{
		"PT Kacamata Nusa",
		"PT Vision Abadi",
		"PT Lensa Optik",
		"PT Cahaya Sejahtera",
	}
	companyMap := make(map[string]model.Company)

	for _, name := range companies {
		var c model.Company
		err := db.WithContext(ctx).First(&c, "name = ?", name).Error
		if err == gorm.ErrRecordNotFound {
			c = model.Company{Name: name}
			if err := db.WithContext(ctx).Create(&c).Error; err != nil {
				return fmt.Errorf("failed to seed company %s: %w", name, err)
			}
			log.Println("🏭 Company:", name)
		}
		companyMap[name] = c
	}

	// ------------------------------
	// 7️⃣ Seed Sample Glasses (RFID blank)
	// ------------------------------
	samples := []model.Glasses{
		{
			Name:        "Aviator Classic",
			Type:        "Sunglasses",
			Color:       "Gold",
			Description: nil,
			RFID:        nil,
			Status:      model.Tersedia,
			DrawerID:    drawerMap["A"].ID,
			BrandID:     brandMap["RayBan"].ID,
			CompanyID:   companyMap["PT Kacamata Nusa"].ID,
		},
		{
			Name:        "Cartier Premium",
			Type:        "Optical",
			Color:       "Silver",
			Description: nil,
			RFID:        nil,
			Status:      model.Tersedia,
			DrawerID:    drawerMap["B"].ID,
			BrandID:     brandMap["Cartier"].ID,
			CompanyID:   companyMap["PT Vision Abadi"].ID,
		},
		{
			Name:        "Nike Vision Sport",
			Type:        "Optical",
			Color:       "Black",
			Description: nil,
			RFID:        nil,
			Status:      model.Tersedia,
			DrawerID:    drawerMap["C"].ID,
			BrandID:     brandMap["Nike"].ID,
			CompanyID:   companyMap["PT Lensa Optik"].ID,
		},
	}

	for _, g := range samples {
		var existing model.Glasses
		err := db.WithContext(ctx).First(&existing, "name = ?", g.Name).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.WithContext(ctx).Create(&g).Error; err != nil {
				return fmt.Errorf("failed to seed sample glasses %s: %w", g.Name, err)
			}
			log.Println("🧿 Sample glasses:", g.Name)
		}
	}

	log.Println("🌳 Seeding complete.")
	return nil
}
