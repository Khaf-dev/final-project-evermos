package handler

import (
	"final-project/dto/request"
	"final-project/internal/service"
	"final-project/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	Service *service.ProductService
}

// Fungsi Handler Product
func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: s}
}

// Create Method Product Handler
func (h *ProductHandler) Create(c *fiber.Ctx) error {

	// Dikarenakan saya menggunakan integrasi upload gambar langsung pada Create method--
	// --maka kode ini mesti di update kembali

	// Kode untuk ambil data form
	name := c.FormValue("name")
	priceStr := c.FormValue("price")
	stockStr := c.FormValue("stock")
	categoryIDStr := c.FormValue("category_id")

	// Validasi dan konversi
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Price"})
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Stock"})
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Category ID"})
	}

	// Upload Gambar
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dibutuhkan file gambar"})
	}

	openedFile, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuka file gambar"})
	}
	defer openedFile.Close()

	imageURL, err := utils.UploadToCloudinary(openedFile, file)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengupload image"})
	}

	// Buat Produk

	storeIDVal := c.Locals("store_id")
	storeID, ok := storeIDVal.(uint)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Store ID tidak ditemukan atau tidak valid"})
	}

	input := request.CreateProductRequest{
		Name:       name,
		Price:      float64(price),
		Stock:      stock,
		CategoryID: uint(categoryID),
		ImageURL:   imageURL,
	}

	product, err := h.Service.CreateProduct(input, storeID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Produk berhasil dibuat",
		"data":    product,
	})
}

// GetAll Method Product Handler
func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	name := c.Query("name")
	categoryStr := c.Query("category")
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	var categoryID uint
	if categoryStr != "" {
		cid, err := strconv.Atoi(categoryStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid category ID"})
		}
		categoryID = uint(cid)
	}

	products, err := h.Service.GetAllProductFiltered(name, categoryID, page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

// GetByID Method Product Handler
func (h *ProductHandler) GetByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product, err := h.Service.GetProductByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product tidak ditemukan"})
	}
	return c.JSON(product)
}

// Update Method Product Handler
func (h *ProductHandler) Update(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	storeID := c.Locals("store_id").(uint)

	isOwner, err := h.Service.IsOwner(uint(id), storeID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	if !isOwner {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden: Anda tidak berhak mengubah produk ini"})
	}

	var input request.UpdateProductRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := utils.Validator.Struct(input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.Service.UpdateProduct(uint(id), input, storeID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"message": "Product telah di perbarui",
	})
}

// Delete Method Product Handler
func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	storeID := c.Locals("store_id").(uint)

	isOwner, err := h.Service.IsOwner(uint(id), storeID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal servel error"})
	}
	if !isOwner {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden: Anda tidak berhak mengubah produk"})
	}

	if err := h.Service.DeleteProduct(uint(id), storeID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Product telah dihapus",
	})
}

// Cloudinary Handler
func (h *ProductHandler) UploadImage(c *fiber.Ctx) error {

	// Kode untuk ambil file
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Tidak ada gambar"})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menampilkan gambar"})
	}
	defer file.Close()

	imageURL, err := utils.UploadToCloudinary(file, fileHeader)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengupload gambar"})
	}

	return c.JSON(fiber.Map{
		"message":   "Gambar berhasil di upload",
		"image_url": imageURL,
	})
}

// Khusus SELLER
func (h *ProductHandler) GetByStore(c *fiber.Ctx) error {
	storeID := c.Locals("store_id").(uint)

	products, err := h.Service.GetProductByStore(storeID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Produk dari toko",
		"data":    products,
	})
}
