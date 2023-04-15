package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"

	"github.com/gofiber/fiber/v2"
)

func SocialLogin(c *fiber.Ctx) error {
	db := database.DBConn
	email := c.FormValue("email")
	image := c.FormValue("image")
	uid := c.FormValue("uid")
	username := c.FormValue("username")

	var user models.User

	result := db.Where("uid=?", uid).First(&user)
	if result.RowsAffected > 0 {
		return c.JSON(user)
	} else {
		dataRegister := models.User{
			UID:      uid,
			Image:    image,
			Email:    email,
			Username: username,
		}
		if register := db.Create(&dataRegister); register.Error != nil {
			return c.Status(500).JSON(register.Error)
		}
		return c.JSON(dataRegister)
	}
}
