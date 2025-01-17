package database

import (
	"fmt"
	"instalasi-pro/configs"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connection() {
	godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		configs.AppConfig.Database.Host,
		configs.AppConfig.Database.User,
		configs.AppConfig.Database.Password,
		configs.AppConfig.Database.DBName,
		configs.AppConfig.Database.Port,
		configs.AppConfig.Database.SSLMode,
		configs.AppConfig.Database.Timezone,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection successful")
}
