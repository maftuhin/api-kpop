package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"
	"maftuhin/kpop-api/tools/paging"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SearchSong(c *fiber.Ctx) error {
	db := database.DBConn
	page, _ := strconv.Atoi(c.Query("page", "1"))
	query := c.Query("q")

	var song []models.Song
	sql := db.Select("uid, title,language, artists.name as artist, view").Joins("join artists ON artists.code=songs.artist").Where("language=? AND title LIKE ?", "original", "%"+query+"%").Order("title ASC")
	paginator := paging.Paging(&paging.Param{
		DB:    sql,
		Limit: 10,
		Page:  page,
	}, song)
	return c.JSON(paginator)
}

func SongDetail(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var song []models.Song
	db.Raw("SELECT songs.id, uid, title, language, lyric, artists.name as artist FROM songs JOIN artists ON artists.code=songs.artist WHERE uid=? LIMIT 10", id).Scan(&song)
	return c.JSON(song)
}

func SongByCode(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	db := database.DBConn
	code := c.Params("code")

	var song []models.Lyric
	sql := db.Table("songs").
		Select("uid, title, language, view").
		Where("artist=? AND language=?", code, "original").
		Order("title ASC")

	paginator := paging.Paging(&paging.Param{
		DB:    sql,
		Page:  page,
		Limit: 10,
	}, song)
	return c.JSON(paginator)
}

func LastUpdate(c *fiber.Ctx) error {
	db := database.DBConn

	var song []models.Song
	db.Raw("SELECT title, uid, language, artists.name, view as artist FROM songs JOIN artists ON artists.code = songs.artist WHERE language='original' ORDER BY songs.id DESC LIMIT 7").
		Scan(&song)
	return c.JSON(song)
}

func MostVisited(c *fiber.Ctx) error {
	db := database.DBConn

	var song []models.Song
	db.Raw("SELECT title, uid, language, view, artists.name as artist FROM songs JOIN artists ON artists.code = songs.artist ORDER BY view DESC LIMIT 10").Scan(&song)
	return c.JSON(song)
}

func Counter(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var song models.Song

	db.Where("id=?", id).First(&song)
	var plus = song.View + 1
	song.View = plus
	db.Save(&song)
	return c.SendString("done")
}
