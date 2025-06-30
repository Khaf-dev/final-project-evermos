package routes

import (
	_ "final-project/configs"
	"final-project/internal/handler"
	"final-project/internal/repository"
	"final-project/internal/service"
	"final-project/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	app.Use("/api/v1/products", middleware.JWTProtected(db))

	// Product Routes
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(&productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Store Routes

	// Auth Routes
	userRepo := repository.NewUserRepository(db)
	storeRepo := repository.NewStoreRepository(db)
	authService := service.NewAuthService(userRepo, storeRepo)
	authHandler := handler.NewAuthHandler(authService)

	// Category Routes
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// User Routes
	userRepo = repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepo)
	userHandler := handler.NewUserHandler(userService)

	// API Handler
	api := app.Group("/api/v1")
	api.Post("/products/upload", productHandler.UploadImage)
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// Handler
	product := api.Group("/products", middleware.JWTProtected(db))
	admin := api.Group("/admin", middleware.JWTProtected(db), middleware.AdminOnly())
	user := api.Group("/user", middleware.JWTProtected(db))
	auth := api.Group("/auth", middleware.JWTProtected(db))
	transaction := api.Group("/transactions", middleware.JWTProtected(db))
	store := api.Group("/store", middleware.JWTProtected(db))

	// Autentikasi
	auth.Get("/me", userHandler.GetProfile)
	auth.Get("/me", userHandler.GetProfile)
	user.Put("/me", userHandler.UpdateProfile)
	user.Get("/my-products", productHandler.GetByStore)

	// Admin (Kategori)
	admin.Post("/", categoryHandler.Create)
	admin.Get("/", categoryHandler.GetAll)
	admin.Put("/:id", categoryHandler.Update)
	admin.Delete("/:id", categoryHandler.Delete)

	// Product
	admin.Post("/products/upload-image", productHandler.UploadImage)
	product.Post("/", productHandler.Create)
	product.Get("/", productHandler.GetAll)
	product.Get("/:id", productHandler.GetByID)
	product.Put("/:id", productHandler.Update)
	product.Delete("/:id", productHandler.Delete)

	//Transaction
	log := api.Group("/logs/product", middleware.JWTProtected(db))

	productLogRepo := repository.NewProductLogRepository(db)
	productLogService := service.NewProductLogService(productLogRepo, productRepo, *storeRepo)
	productLogHandler := handler.NewProductLogHandler(productLogService)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionDetailRepo := repository.NewTransactionDetailRepository(db)

	transactionService := service.NewTransactionService(
		transactionRepo,
		productRepo,
		*transactionDetailRepo,
		productLogRepo,
	)

	transactionHandler := handler.NewTransactionHandler(
		transactionService,
		storeRepo,
	)

	transaction.Post("/", transactionHandler.CreateTransaction)
	api.Post("/transactions/admin", transactionHandler.GetAllTransactions) // KHSUUS ADSMIN
	transaction.Get("/:id", transactionHandler.GetByID)
	transaction.Get("/store/transactions", transactionHandler.GetStoreTransactions)
	api.Get("/transactions/user", transactionHandler.GetUserTransactions)
	log.Get("/", productLogHandler.GetAll)
	log.Get("/:product_id", productLogHandler.GetLogsByProductID)
	store.Get("/transactions", transactionHandler.GetByStore)
	api.Get("/transactions", transactionHandler.GetUserTransactions)

	// Wilayah routes
	wilayahService := service.NewWilayahService()
	wilayahHandler := handler.NewWilayahHandler(wilayahService)

	wilayah := api.Group("/wilayah")
	wilayah.Get("/provinsi", wilayahHandler.GetProvinces)
	wilayah.Get("/kabupaten-kota/:province_id", wilayahHandler.GetRegencies)
	wilayah.Get("/kecamatan/:regency_id", wilayahHandler.GetDistricts)
	wilayah.Get("/kelurahan/:district_id", wilayahHandler.GetVillages)

	//Alamat routes
	addressRepo := repository.NewAddressRepository(db)
	addressService := service.NewAddressService(addressRepo)
	addressHandler := handler.NewAddressHandler(addressService)

	address := api.Group("/address", middleware.JWTProtected(db))
	address.Post("/", addressHandler.Create)
	address.Get("/", addressHandler.GetByUser)

}
