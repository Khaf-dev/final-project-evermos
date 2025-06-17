package handler

import (
	"final-project/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *CategoryHandler {
	return &CategoryHandler{s}
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	type req struct {
		Name string `json:"name"`
	}
	var body req
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := h.categoryService.Create(body.Name)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Category created"})
}

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	categories, err := h.categoryService.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(categories)
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	type req struct {
		Name string `json:"name"`
	}
	var body req
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := h.categoryService.Update(uint(id), body.Name)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Category updated"})
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := h.categoryService.Delete(uint(id))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Category deleted"})
}
