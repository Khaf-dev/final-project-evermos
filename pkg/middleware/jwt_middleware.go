package middleware

import (
	"final-project/internal/repository"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var jwtSecret = []byte("secretkey_final_project") // Harus sama jwt.go di utils

func JWTProtected(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or Invalid Authorization header",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or Expired Token",
			})
		}

		// Ngambil claim token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Token Claims",
			})
		}

		userID := uint(claims["user_id"].(float64))

		uidFloat, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid user_id in token",
			})
		}

		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid role in token",
			})
		}
		c.Locals("user_id", userID)
		c.Locals("user_id", uint(uidFloat))
		c.Locals("role", role)

		// Ambil store_id jika rolenya bukan admin
		if role != "admin" {
			storeRepo := repository.NewStoreRepository(db)
			store, err := storeRepo.FindByUserID(userID)

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Store ID tidak ditemukan untuk user",
				})
			}
			c.Locals("store_id", store.ID)
			fmt.Printf("[JWT-MIDDLEWARE] user_id=%d, role=%s, store_id=%v\n", userID, role, c.Locals("store_id"))

		}
		return c.Next()
	}
}
