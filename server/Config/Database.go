package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/lol-iris/aglearning/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func dbUrl() string {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
}

func Init() {
	log.Println("Initializing database...")
	url := dbUrl()
	logMode := logger.Warn
	// Set logging level
	if os.Getenv("DEBUG") == "true" {
		logMode = logger.Info
		log.Println("Database logging level set to Info.")
	}

	database, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	log.Println("Database connected.")

	if err != nil {
		log.Fatalln(err)
	}

	database.AutoMigrate(&Models.User{})
	DB = database
	log.Println("Database initialized.")
}
