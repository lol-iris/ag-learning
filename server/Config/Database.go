package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/lol-iris/aglearning/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	database, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	log.Println("Database connected.")

	if err != nil {
		log.Fatalln(err)
	}

	database.AutoMigrate(&Models.User{})
	DB = database
	log.Println("Database initialized.")
}
