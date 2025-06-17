package routes

import (
	configs "final-project/configs"
	"final-project/internal/handler"
	"final-project/internal/repository"
	"final-project/internal/service"
	"final-project/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := configs.DB

	userRepo := repository.NewUserRepository(db)
	storeRepo := repository.NewStoreRepository(db)
	authService := service.NewAuthService(userRepo, storeRepo)
	authHandler := handler.NewAuthHandler(authService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	userService := service.NewUserService(*userRepo)
	userHandler := handler.NewUserHandler(userService)

	api := app.Group("/api/v1")
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	auth := api.Group("/auth", middleware.JWTProtected())

	admin := api.Group("/admin", middleware.JWTProtected(), middleware.AdminOnly())
	admin.Post("/categories", categoryHandler.Create)
	admin.Get("/categories", categoryHandler.GetAll)
	admin.Put("/categories/:id", categoryHandler.Update)
	admin.Delete("/categories/:id", categoryHandler.Delete)

	auth.Get("/me", userHandler.GetProfile)
	auth.Put("/me", userHandler.UpdateProfile)
	auth.Get("/me", userHandler.GetProfile)
}
