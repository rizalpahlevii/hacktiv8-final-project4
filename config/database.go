package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/model"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func handleDatabaseConnection() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Error connecting to database", err)
	}

	log.Println("Database connection established")
}

func DatabaseConnection() *gorm.DB {
	if db == nil {
		handleDatabaseConnection()
	}

	return db
}

func StartDatabase() {
	handleDatabaseConnection()
	database := DatabaseConnection()
	err := database.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{}, &model.TransactionHistory{})
	if err != nil {
		panic(err)
	}
}
