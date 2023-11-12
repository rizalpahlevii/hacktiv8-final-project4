package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hacktiv8-final-project4/model"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func handleDatabaseConnection() {
	//err := godotenv.Load()
	//helper.PanicIfError(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)

	fmt.Println(connectionString)
	dbInstance, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Error connecting to database. Error : ", err)
	}

	db = dbInstance
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

	// Seed Admin User
	if database.Where("email = ?", "admin@mail.com").First(&model.User{}).RowsAffected > 0 {
		return
	}

	admin := model.User{
		FullName: "Admin",
		Email:    "admin@mail.com",
		Password: "admin",
		Role:     model.AdminRole,
	}
	database.Create(&admin)

	if err != nil {
		panic(err)
	}
}
