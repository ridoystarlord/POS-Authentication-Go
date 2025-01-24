package middleware

import (
	"authentication/Responses"
	"authentication/config"
	"authentication/utils"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthenticated(c *fiber.Ctx) error { //
	// Get the token from the Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return Responses.Unauthenticated(c)
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	// Load configuration
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	// Parse the token
	token, err := utils.VerifyToken(tokenString, config.JWTAccessSecret)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Attach token claims to the context for future use
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Responses.Unauthorized(c)
	}

	userId := claims["id"]
	if userId == nil {
		return Responses.Unauthorized(c)
	}
	c.Locals("payload", userId)

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
	// Load configuration
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	// Parse the token
	token, err := utils.VerifyToken(tokenString, config.JWTRefreshSecret)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Attach token claims to the context for future use

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Responses.Unauthorized(c)
	}

	userId := claims["id"]
	if userId == nil {
		return Responses.Unauthorized(c)
	}
	c.Locals("payload", userId)

	return c.Next()
}
