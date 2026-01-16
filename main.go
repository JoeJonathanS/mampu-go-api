package main

import (
	"fmt"
	"log"
	"mampu-go-api/handlers"
	"mampu-go-api/models"
	"mampu-go-api/repositories"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.User{}, &models.Transaction{})
}

func seedDB() {
	// Truncate tables
	db.Exec("TRUNCATE TABLE users")
	db.Exec("TRUNCATE TABLE transactions")

	// Seed 10 users
	for i := 1; i <= 10; i++ {
		user := models.User{
			Name:    fmt.Sprintf("User %d", i),
			Email:   fmt.Sprintf("user%d@example.com", i),
			Balance: float64(i * 100),
		}
		db.Create(&user)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initDB()
	seedDB()

	userRepo := repositories.NewUserRepository(db)
	transRepo := repositories.NewTransactionRepository(db)
	walletHandler := handlers.NewWalletHandler(userRepo, transRepo)

	app := fiber.New()

	app.Get("/balance/:userId", walletHandler.GetBalance)
	app.Post("/withdraw", walletHandler.Withdraw)

	log.Fatal(app.Listen(":3000"))
}
