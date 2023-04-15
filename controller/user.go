package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"
	"maftuhin/kpop-api/tools/paging"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserSearch(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	db := database.DBConn
	var users []models.User
	q := c.Query("query")
	columns := []string{"uid, username, image, description"}
	orders := "username ASC"
	if q == "" {
		orders = "RAND()"
	}

	sql := db.Select(columns).Where("username LIKE ?", "%"+q+"%")
	paginator := paging.Paging(&paging.Param{
		DB:      sql,
		Limit:   10,
		Page:    page,
		OrderBy: []string{orders},
	}, users)
	if paginator.TotalRecord > 0 {
		return c.JSON(paginator)
	}
	return c.Status(500).JSON(&fiber.Map{"message": "no data"})
}

func UserProfile(c *fiber.Ctx) error {
	db := database.DBConn
	uid := c.Params("uid")
	var user models.User
	result := db.Where("uid=?", uid).First(&user)
	if result.Error == nil {
		return c.JSON(user)
	}
	return c.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "User Not Found"})
}

func UserDetail(c *fiber.Ctx) error {
	db := database.DBConn
	uid := c.Get("uid")
	var user models.User
	result := db.Where("uid=?", uid).First(&user)
	if result.Error == nil {
		return c.JSON(user)
	}
	return c.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "User Not Found"})
}

func UserUpdate(c *fiber.Ctx) error {
	db := database.DBConn
	uid := c.Get("uid")
	username := c.FormValue("username")
	description := c.FormValue("description")
	dataUpdate := models.User{
		Username:    username,
		Description: description,
	}

	if update := db.Model(&models.User{}).Where("uid=?", uid).Updates(&dataUpdate); update.Error != nil {
		return c.Status(500).JSON(update.Error)
	}
	return c.JSON(dataUpdate)
}

func UserPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	uid := c.Params("uid")
	db := database.DBConn
	var posts []models.Post

	sql := db.Select("id,uid,artist,title,language").Where("uid=?", uid)
	paginator := paging.Paging(&paging.Param{
		DB:    sql,
		Limit: 10,
		Page:  page,
	}, posts)
	if paginator.TotalRecord > 0 {
		return c.JSON(paginator)
	}
	return c.Status(500).JSON(&fiber.Map{"message": "No Data"})
}
