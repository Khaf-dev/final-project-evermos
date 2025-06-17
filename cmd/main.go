package main

import (
	configs "final-project/configs"
	"final-project/internal/domain"
	"final-project/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.InitDB()

	// Auto Migrate
	db := configs.DB
	db.AutoMigrate(&domain.User{}, &domain.Store{}, &domain.Category{})

	// Routes
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
