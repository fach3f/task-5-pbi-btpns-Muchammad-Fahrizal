package database

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tugasakhir/models"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDb() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error setting up connection pool: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	fmt.Println("Connection successful!")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})

	DB = db
}
