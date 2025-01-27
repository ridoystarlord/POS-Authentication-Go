package controllers

import (
	DBManager "authentication/Database"
	"authentication/Responses"
	"authentication/dto"
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
	var userCredential models.Credential
	var inputCredentials dto.InputCredentials
	err := c.BodyParser(&inputCredentials)
	if err != nil {
		return Responses.BadRequest(c, "Unable to parse body")
	}

	err = DBManager.DB.Where("username = ?", inputCredentials.Username).First(&userCredential).Error
	if err != nil {
		return Responses.Unauthorized(c)
	}
	if ok := utils.ComparePassword(userCredential.Password, inputCredentials.Password); !ok {
		return Responses.Unauthorized(c)
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(userCredential.UserID)
	if err != nil {
		return Responses.InternalServerError(c)
	}
	result := DBManager.DB.Model(&userCredential).Update("refresh_token", refreshToken)
	if result.Error != nil {
		return Responses.InternalServerError(c)
	}

	var data dto.AuthResponse
	data.AccessToken = accessToken
	data.RefreshToken = refreshToken
	data.WarehouseId = ""
	data.OrganizationId = ""
	data.Type = ""

	Responses.Response(c, 200, true, "Login Successfully", data)
	return nil
}

func RefreshToken(c *fiber.Ctx) error {
	var credentials models.Credential
	payload := c.Locals("payload")
	userId := payload.(float64)
	credentials.UserID = uint(userId)

	accessToken, refreshToken, err := utils.GenerateTokenPair(credentials.UserID)
	if err != nil {
		return Responses.InternalServerError(c)
	}

	result := DBManager.DB.Model(&credentials).Where("user_id = ?", credentials.UserID).Update("refresh_token", refreshToken)
	if result.Error != nil {
		return Responses.InternalServerError(c)
	}

	var data dto.AuthResponse
	data.AccessToken = accessToken
	data.RefreshToken = refreshToken
	data.WarehouseId = ""
	data.OrganizationId = ""
	data.Type = ""

	Responses.Response(c, 200, true, "Access Token Refreshed", data)
	return nil
}

func Logout(c *fiber.Ctx) error {
	var credentials models.Credential
	payload := c.Locals("payload")

	userId := payload.(float64)
	credentials.UserID = uint(userId)
	result := DBManager.DB.Model(&credentials).Where("user_id = ?", credentials.UserID).Update("refresh_token", nil)
	if result.Error != nil {
		return Responses.InternalServerError(c)
	}
	Responses.Response(c, 200, true, "Logout Successfully", nil)
	return nil
}
