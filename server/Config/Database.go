package Config

import (
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	"github.com/lol-iris/aglearning/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getEnvVar(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func dbUrl() string {

	dbHost := getEnvVar("DB_HOST")
	dbPort := getEnvVar("DB_PORT")
	dbName := getEnvVar("DB_NAME")
	dbUser := getEnvVar("DB_USER")
	dbPassword := getEnvVar("DB_PASSWORD")
	fmt.Println("The host from env:", dbHost)

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

	database.AutoMigrate(&Models.User{}, &Models.Course{}, &Models.Announcement{}, &Models.Tag{}, &Models.Section{}, &Models.SectionAttachment{}, &Models.SectionComment{}, &Models.Assignment{}, &Models.AssignmentSubmission{})
	DB = database
	log.Println("Database initialized.")

}
