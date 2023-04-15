package main

import (
	"database/sql"
	"log"
	"maftuhin/kpop-api/database"
	"maftuhin/kpop-api/routes"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	// sqlDB, _ := sql.Open("mysql", "remote:8Belas0694@tcp(limaefdua.com:3306)/kpop")
	sqlDB, _ := sql.Open("mysql", "root:118806@tcp(127.0.0.1:3306)/kpop")
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalln("s", err)
	}
	return gormDB
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	database.DBConn = initDatabase()

	app.Use(logger.New())
	app.Use(cors.New(cors.ConfigDefault))
	// app.Use(ApiKey())
	routes.SetUpRoutes(app)
	log.Fatal(app.Listen(":8000"))
}

func ApiKey() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Get("apiKey")
		if key != "MtxgcHHvcf" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}
