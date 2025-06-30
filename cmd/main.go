package main

import (
	configs "final-project/configs"
	"final-project/internal/domain"
	"final-project/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	os.MkdirAll("uploads", os.ModePerm)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
				"debug": err.Error(),
			})
		},
	})

	configs.LoadEnv()
	configs.InitDB()

	// Auto Migrate
	db := configs.DB

	db.AutoMigrate(
		&domain.User{},
		&domain.Store{},
		&domain.Category{},
		&domain.Product{},
		&domain.Transaction{},
		&domain.TransactionDetail{},
		&domain.ProductLog{},
		&domain.Address{},
	)

	// Routes
	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":8080"))
}
