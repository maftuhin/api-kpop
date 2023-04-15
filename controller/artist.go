package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"
	"maftuhin/kpop-api/tools/paging"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ArtistDetail(c *fiber.Ctx) error {
	db := database.DBConn
	code := c.Params("code")
	var artist models.Artist

	result := db.Where("code=?", code).First(&artist)
	if result.RowsAffected > 0 {
		return c.JSON(artist)
	}
	return c.Status(500).JSON(&fiber.Map{
		"message": "Not Found",
	})
}

func ArtistSearch(c *fiber.Ctx) error {
	db := database.DBConn
	page, _ := strconv.Atoi(c.Query("page", "1"))
	query := c.Query("query")
	var artist []models.Artist

	sql := db.Where("name LIKE ?", "%"+query+"%").Order("name ASC")
	paginator := paging.Paging(&paging.Param{
		Page:    page,
		DB:      sql,
		Limit:   10,
		ShowSQL: false,
	}, artist)
	return c.JSON(paginator)
}
