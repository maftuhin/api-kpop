package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"
	"maftuhin/kpop-api/tools/paging"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SearchSoundtrack(c *fiber.Ctx) error {
	db := database.DBConn
	q := c.Query("q")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	var ost []models.Track

	sql := db.Where("title LIKE ?", "%"+q+"%")
	paging := paging.Paging(&paging.Param{
		DB:      sql,
		Page:    page,
		Limit:   10,
		OrderBy: []string{"title ASC"},
	}, ost)

	return c.JSON(paging)
}

func SoundtrackPlaylist(c *fiber.Ctx) error {
	db := database.DBConn
	uid := c.Params("uid")

	var song []models.Song
	sql := db.Table("tracklinks").
		Select("songs.uid, songs.title, artists.name as artist").
		Joins("JOIN songs ON songs.uid = tracklinks.song").
		Joins("JOIN artists ON artists.code=songs.artist").
		Where("soundtrack=? AND songs.language=?", uid, "original")
	paging := paging.Paging(&paging.Param{
		DB:      sql,
		Limit:   10,
		OrderBy: []string{"songs.title ASC"},
	}, song)

	return c.JSON(paging)
}
