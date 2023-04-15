package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"
	"maftuhin/kpop-api/tools/paging"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LatestPost(c *fiber.Ctx) error {
	db := database.DBConn
	var posts []models.PostGet
	columns := []string{"posts.id", "posts.uid", "title", "artist", "language", "username", "image"}
	db.Table("posts").Order("posts.id desc").Select(columns).Joins("join users on users.uid=posts.uid").Limit(6).Find(&posts)
	return c.JSON(posts)
}

func PostSearch(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	q := c.Query("query")
	db := database.DBConn

	columns := []string{"posts.id", "posts.uid", "title", "artist", "language", "username", "image"}
	var data []models.PostGet

	sql := db.Select(columns).Table("posts").Select(columns).Where("MATCH(artist, title) AGAINST(?)", q).Joins("join users on users.uid=posts.uid")
	paginator := paging.Paging(&paging.Param{
		DB:    sql,
		Limit: 10,
		Page:  page,
	}, data)
	if paginator.TotalRecord > 0 {
		return c.JSON(paginator)
	}
	return c.Status(500).JSON(&fiber.Map{"message": "no data"})
}

func PostDetail(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	var post models.PostGet
	db.Table("posts").Select("posts.id, users.uid, artist, title, content, language, username, image").Joins("join users on users.uid=posts.uid").Where("posts.id=?", id).First(&post)
	return c.JSON(post)
}

func MyPost(c *fiber.Ctx) error {
	db := database.DBConn
	page, _ := strconv.Atoi(c.Query("page", "1"))
	uid := c.Get("uid")
	var posts []models.Post

	sql := db.Select("id, title, artist,language").Where("uid=?", uid)
	paginator := paging.Paging(&paging.Param{
		DB:    sql,
		Limit: 20,
		Page:  page,
	}, posts)
	if paginator.TotalRecord > 0 {
		return c.JSON(paginator)
	}
	return c.Status(500).JSON(&fiber.Map{"message": "no data"})
}

func PostCreate(c *fiber.Ctx) error {
	db := database.DBConn
	uid := c.Get("uid")
	artist := c.FormValue("artist")
	title := c.FormValue("title")
	content := c.FormValue("content")
	lang := c.FormValue("language")
	hashtag := c.FormValue("hashtag")

	post := models.Post{
		UID:      uid,
		Artist:   artist,
		Title:    title,
		Content:  content,
		Language: lang,
		Hashtag:  hashtag,
	}

	if insert := db.Create(&post); insert.Error != nil {
		return c.Status(500).JSON(insert.Error)
	}
	return c.JSON(post)
}

func PostUpdate(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.FormValue("id")
	title := c.FormValue("title")
	artist := c.FormValue("artist")
	content := c.FormValue("content")
	language := c.FormValue("language")

	data := models.Post{
		Title:    title,
		Artist:   artist,
		Content:  content,
		Language: language,
	}

	if update := db.Model(&models.Post{}).Where("id=?", id).Updates(&data); update.Error != nil {
		return c.Status(500).JSON(&fiber.Map{"message": "something wrong"})
	}
	return c.JSON(data)
}

func PostDelete(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.FormValue("id")
	if delete := db.Where("id=?", id).Delete(&models.Post{}); delete.Error != nil {
		c.Status(500).JSON(delete.Error)
	}
	return c.JSON(&fiber.Map{"message": "Success"})
}
