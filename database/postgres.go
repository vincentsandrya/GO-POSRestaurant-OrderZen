package database

import (
	"fmt"
	"log"
	"os"

	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func Connect() (*gorm.DB, error) {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
		return nil, err
	}

	fmt.Println("Connected to the database")

	gd := &GormDB{DB: db}
	gd.MigrationDatabase()

	return db, nil
}

func (gd *GormDB) MigrationDatabase() {

	err := gd.DB.AutoMigrate(
		&models.Role{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	err = gd.DB.AutoMigrate(
		&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	err = gd.DB.AutoMigrate(
		&models.MenuCategory{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	err = gd.DB.AutoMigrate(
		&models.Menu{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	err = gd.DB.AutoMigrate(
		&models.Promo{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	err = gd.DB.AutoMigrate(
		&models.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	err = gd.DB.AutoMigrate(
		&models.OrderItem{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	err = gd.DB.AutoMigrate(
		&models.PaymentMidtrans{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	roles := []models.Role{
		{RoleId: 1, Name: "admin"},
		{RoleId: 2, Name: "cashier"},
		{RoleId: 3, Name: "kitchen"},
		{RoleId: 4, Name: "customer"},
	}

	for _, role := range roles {
		var count int64
		gd.DB.Model(&models.Role{}).Where("role_id = ?", role.RoleId).Count(&count)
		if count == 0 {
			gd.DB.Create(&role)
		}
	}

	var userCount int64
	gd.DB.Model(&models.User{}).Where("email = ?", "admin@gmail.com").Count(&userCount)
	if userCount == 0 {
		adminUser := models.User{
			Email:     "admin@gmail.com",
			Name:      "Admin",
			Password:  "$2a$10$y4GXKf1p6cHUOFpBi2VF8efKzq8sHeWtyup9NAxSKCFJbAuITyXcC", //12345
			RoleId:    1,
			CreatedBy: "system",
			UpdatedBy: "system",
		}
		gd.DB.Create(&adminUser)
	}
}
