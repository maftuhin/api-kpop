package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateArtist(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Query("id")
	var artists []models.Artist

	if result := db.Where("id > ?", id).Find(&artists); result.RowsAffected > 0 {
		return c.JSON(artists)
	}
	return c.Status(500).JSON(&fiber.Map{"message": "Up To Date"})
}

func UpdateSong(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Query("id")
	var songs []models.Song

	if result := db.Where("id > ?", id).Limit(50).Find(&songs); result.RowsAffected > 0 {
		return c.JSON(songs)
	}
	return c.Status(500).JSON(&fiber.Map{"message": "Up To Date"})
}

func UpdateSoundTrack(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.FormValue("id")

	var tracks []models.Track
	db.Where("id > ?", id).Find(&tracks)
	return c.JSON(tracks)
}

func UpdateLink(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.FormValue("id")

	var links []models.Link
	db.Where("id > ?", id).Find(&links)
	return c.JSON(links)
}
