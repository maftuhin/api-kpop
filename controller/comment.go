package controller

import (
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/models"
	"maftuhin/kpop-api/tools/paging"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CommentList(c *fiber.Ctx) error {
	db := database.DBConn
	page, _ := strconv.Atoi(c.Query("page", "1"))
	id := c.Query("pid")
	column := []string{"comment", "username as uid", "timestamp"}
	var comments []models.Comment

	sql := db.Select(column).Where("post_id=?", id).Joins("JOIN users ON users.uid=comments.uid")
	paginator := paging.Paging(&paging.Param{
		DB:      sql,
		Page:    page,
		OrderBy: []string{"comments.id DESC"},
	}, comments)
	if paginator.TotalRecord > 0 {
		return c.JSON(paginator)
	}
	return c.Status(500).JSON(&fiber.Map{"message": "No Comment"})
}

func CommentCreate(c *fiber.Ctx) error {
	db := database.DBConn
	uid := c.Get("uid")
	id := c.FormValue("post_id")
	comment := c.FormValue("comment")
	timestamp := c.FormValue("timestamp")

	dataInsert := models.Comment{
		PostId:    id,
		UID:       uid,
		Comment:   comment,
		Timestamp: timestamp,
	}

	if insert := db.Create(&dataInsert); insert.Error != nil {
		return c.JSON(insert.Error)
	}
	return c.JSON(dataInsert)
}
