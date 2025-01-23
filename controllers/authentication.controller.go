package controllers

import (
	DBManager "authentication/Database"
	"authentication/Responses"
	"authentication/models"
	"authentication/utils"

	"github.com/gofiber/fiber/v2"
)

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
		return Responses.BadRequest(c, "Unable to parse body")
	}

	err = DBManager.DB.Where("username = ?", credentials.Username).First(&user).Error
	if err != nil {
		return Responses.Unauthorized(c)
	}
	if ok := utils.ComparePassword(user.Password, credentials.Password); !ok {
		return Responses.Unauthorized(c)
	}

	data := map[string]interface{}{}
	accessToken, refreshToken, err := utils.GenerateTokenPair(credentials.UserID)
	if err != nil {
		return Responses.SomethingGoneWrong(c)
	}

	data["accessToken"] = accessToken
	data["refreshToken"] = refreshToken
	data["warehouseId"] = ""
	data["organizationId"] = ""
	data["type"] = ""

	Responses.Response(c, 200, true, "Login Successfully", data)
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
	Responses.Created(c, "user", nil)
	return nil
}

func RefreshToken(c *fiber.Ctx) error {
	payload := c.Locals("payload")

	userId := payload.(float64)
	data := map[string]interface{}{}

	accessToken, refreshToken, err := utils.GenerateTokenPair(uint(userId))
	if err != nil {
		return Responses.InternalServerError(c)
	}
	data["accessToken"] = accessToken
	data["refreshToken"] = refreshToken
	data["warehouseId"] = ""
	data["organizationId"] = ""
	data["type"] = ""
	Responses.Response(c, 200, true, "Access Token Refreshed", data)
	return nil
}
