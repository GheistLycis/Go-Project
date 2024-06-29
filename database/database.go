package database

import (
	"fmt"
	"log"
	"os"

	user_structs "github.com/go-project/app/user/structs"
	godotenv "github.com/joho/godotenv"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

var DB *gorm.DB

func getDSN() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)

		return ""
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}

func connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

		return nil
	}

	return db
}

func migrate(database *gorm.DB) {
	err := database.AutoMigrate(&user_structs.User{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func Init() {
	dsn := getDSN()

	DB = connect(dsn)

	migrate(DB)
}
