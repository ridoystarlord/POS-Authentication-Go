package middleware

import (
	"authentication/Responses"
	"authentication/storage"
	"authentication/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthenticated(c *fiber.Ctx) error { //
	// Get the token from the Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	// Parse the token
	secret := storage.Config.JWTAccessSecret
	token, err := utils.VerifyToken(tokenString, storage.Config.JWTAccessSecret)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Attach token claims to the context for future use
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Locals("user", claims)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}
	// fmt.Println("Token is valid")

	return c.Next()
}

func IsAuthorized(c *fiber.Ctx) error {
	// Get the token from the Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return Responses.Unauthorized(c)
	}

	// Remove "Bearer " prefix if present
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	// Parse the token
	token, err := utils.VerifyRefreshToken(tokenString)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Attach token claims to the context for future use
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Locals("user", claims)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}
	// fmt.Println("Token is valid")

	return c.Next()
}
