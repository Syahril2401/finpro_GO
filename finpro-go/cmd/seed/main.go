package main

import (
	"finpro/config"
	"finpro/model"
	"finpro/utils"
	"log"

	"github.com/google/uuid"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	// 1. Create the Admins table if it doesn't exist
	err := config.DB.AutoMigrate(&model.Admin{})
	if err != nil {
		log.Fatalf("Failed to create admins table: %v", err)
	}
	log.Println("Admins table ensured to exist.")

	// 2. Check if admin already exists
	var count int64
	config.DB.Model(&model.Admin{}).Count(&count)
	if count > 0 {
		log.Println("Admin already exists. No need to seed.")
		return
	}

	// 3. Create a default admin
	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	admin := model.Admin{
		AdminID:      uuid.New().String(),
		Name:         "Super Admin",
		Email:        "admin@example.com",
		PasswordHash: hashedPassword,
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to create default admin: %v", err)
	}

	log.Println("=====================================")
	log.Println("Default Admin Account Created!")
	log.Println("Email: admin@example.com")
	log.Println("Password: admin123")
	log.Println("=====================================")
}
