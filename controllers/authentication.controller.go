package controllers

import (
	DBManager "authentication/Database"
	"authentication/Responses"
	"authentication/config"
	"authentication/models"
	"authentication/utils"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenResponse struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

//	 ListBooks godoc
//		@Summary		Login
//		@Description	Login
//		@Tags			Authentication
//		@Accept			json
//		@Produce		json
//		@Param			request	body	models.Credential	true	"Login"
//		@Success		200
//		@Router			/api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	var credentials models.Credential
	var user models.Credential
	err := c.BodyParser(&credentials)
	if err != nil {
		c.JSON(&fiber.Map{"message": "Unable to parse body"})
		return err
	}

	err = DBManager.DB.Where("username = ?", credentials.Username).First(&user).Error
	if err != nil {
		c.JSON(&fiber.Map{"message": "Invalid Credentials"})
		return err
	}
	if ok := utils.ComparePassword(user.Password, credentials.Password); !ok {
		return c.JSON(&fiber.Map{"message": "Invalid Credentials"})
	}

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	accessTokenDuration, err := strconv.Atoi(config.JWTAccessExp)
	if err != nil {
		log.Fatalf("Invalid env format: %v", err)
	}
	fmt.Println(accessTokenDuration)
	accessToken, err := utils.CreateToken(credentials.UserID, accessTokenDuration, config.JWTAccessSecret)
	if err != nil {
		c.JSON(&fiber.Map{"message": "Unable to create token"})
		return err
	}

	refreshTokenDuration, err := strconv.Atoi(config.JWTRefreshExp)
	if err != nil {
		log.Fatalf("Invalid env format: %v", err)
	}
	refreshToken, err := utils.CreateToken(credentials.UserID, refreshTokenDuration, config.JWTRefreshSecret)
	if err != nil {
		c.JSON(&fiber.Map{"message": "Unable to create token"})
		return err
	}

	token := &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	// c.JSON(&fiber.Map{"message": "Login successful", "token": token})
	Responses.Response(c, 200, true, "Login Successfully", token)
	return nil
}

//	 ListBooks godoc
//		@Summary		Create user
//		@Description	Create a new user
//		@Tags			Authentication
//		@Accept			json
//		@Produce		json
//		@Param			request	body	models.User	true	"Register"
//		@Success		200
//		@Router			/api/v1/auth/register [post]
func Register(c *fiber.Ctx) error {
	var payload models.User
	err := c.BodyParser(&payload)
	if err != nil {
		c.JSON(&fiber.Map{"message": "Unable to parse body"})
		return err
	}
	payload.Credential.Password = utils.HashPassword(payload.Credential.Password)
	err = DBManager.DB.Create(&payload).Error
	if err != nil {
		c.JSON(&fiber.Map{"message": "Unable to create user"})
		return err
	}
	c.JSON(&fiber.Map{"message": "User created successfully"})
	return nil
}

func Refresh(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")[7:]
	user := c.Locals("user")
	if user == nil {
		return Responses.Unauthorized(c)
	}
	claims, ok := user.(jwt.MapClaims)
	if !ok {
		return Responses.Unauthorized(c)
	}

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	accessTokenDuration, err := strconv.Atoi(config.JWTAccessExp)
	if err != nil {
		log.Fatalf("Invalid env format: %v", err)
	}

	id := claims["id"].(float64)
	newAccessToken, err := utils.CreateToken(uint(id), accessTokenDuration, config.JWTAccessSecret)
	if err != nil {
		c.JSON(&fiber.Map{"message": "Unable to create user"})
		return err
	}
	token := &TokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: tokenString,
	}
	c.JSON(&fiber.Map{"message": "Refresh successful", "token": token})
	return nil
}
