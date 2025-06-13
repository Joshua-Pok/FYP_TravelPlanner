package main

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func InitializeDB() {

	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/travelapp?charset=utf8mb4&parseTime=True&loc=Local"

	// Configure GORM with MySQL driver
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable detailed logging
		// Optional: Configure connection settings
		NowFunc: func() time.Time {
			return time.Now().UTC() // Ensure consistent timezone
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		err = DB.AutoMigrate(&User{}, &Personality{})
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}

		log.Println("Database connected and migrated successfully")
	}

}
