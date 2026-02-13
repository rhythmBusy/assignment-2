package config

import (
	"fmt"
	"os"

	"assignment2/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	ssl := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, pass, name, port, ssl,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("✅ Postgres connected")

	// ✅ AUTO MIGRATE ALL MODELS
	err = db.AutoMigrate(
		&models.User{},
		&models.Bank{},
		&models.Branch{},
		&models.Account{},
		&models.Transaction{},
		&models.Loan{},
		&models.LoanPayment{},
	)

	if err != nil {
		panic("Migration failed")
	}

	fmt.Println(" Migration done")

	return db
}
