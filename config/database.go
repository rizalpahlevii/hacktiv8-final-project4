package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hacktiv8-final-project4/models"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func handleDatabaseConnection() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "root"
	dbname := "hacktiv8_final_project4"

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
	err := database.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.TransactionHistory{})
	if err != nil {
		panic(err)
	}
}
