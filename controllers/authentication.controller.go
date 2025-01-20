package controllers

import (
	"authentication/models"
	"authentication/storage"
	"authentication/utils"

	"github.com/gofiber/fiber/v2"
)

//  ListBooks godoc
//	@Summary		Login
//	@Description	Login
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	models.Credential	true	"Login"
//	@Success		200
//	@Router			/api/auth/login [post]
func  Login(c *fiber.Ctx) error {
	var credentials models.Credential
	err := c.BodyParser(&credentials)
	if  err != nil {
		 c.JSON(&fiber.Map{"message":"Unable to parse body"})
		 return err
	}
	err=storage.DB.Where("username = ?", credentials.Username).First(&credentials).Error
	if err != nil {
		c.JSON(&fiber.Map{"message":"Unable to find user"})
		 return err
	}
	token, err := utils.CreateToken(credentials.UserID,30)
	if err != nil {
		c.JSON(&fiber.Map{"message":"Unable to create token"})
		 return err
	}
	c.JSON(&fiber.Map{"message":"Login successful","token":token})
	return nil
}

//  ListBooks godoc
//	@Summary		Create user
//	@Description	Create a new user
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	models.User	true	"Register"
//	@Success		200
//	@Router			/api/auth/register [post]
func  Register(c *fiber.Ctx) error {
	var payload models.User
	err := c.BodyParser(&payload)
	if  err != nil {
		 c.JSON(&fiber.Map{"message":"Unable to parse body"})
		 return err
	}
	payload.Credential.Password = utils.HashPassword(payload.Credential.Password)
	err=storage.DB.Create(&payload).Error
	if err!= nil {
		c.JSON(&fiber.Map{"message":"Unable to create user"})
		 return err
	}
	c.JSON(&fiber.Map{"message":"User created successfully"})
	return nil
}