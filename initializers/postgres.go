package initializers

import (
	"fmt"
	"github.com/ibanmarco/gin-golang-postgres/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func PostgresConnection() {
	var err error
	loadEnvVariables()
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
		os.Getenv("DB_TZ"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	if os.Getenv("DB_MIGRATE") == "True" {
		err := DB.AutoMigrate(&models.Books{})
		if err != nil {
			log.Fatal("Failed to migrate schema")
		}
	}

}
