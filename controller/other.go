package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"

	"github.com/gofiber/fiber/v2"
)

func Language(c *fiber.Ctx) error {
	db := database.DBConn

	var langs []models.Language
	db.Find(&langs)
	return c.JSON(langs)
}

func SendRequest(c *fiber.Ctx) error {
	db := database.DBConn
	uid := c.Get("uid")
	artist := c.FormValue("artist")
	title := c.FormValue("title")

	dataRequest := models.Request{
		UID:    uid,
		Artist: artist,
		Title:  title,
	}
	if insert := db.Create(&dataRequest); insert.Error != nil {
		c.JSON(insert.Error)
	}
	return c.JSON(&fiber.Map{"message": "Thank you"})
}
