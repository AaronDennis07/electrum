package main

import (
	"log"
	"net/http"

	"github.com/AaronDennis07/electrum/internals/database"
	"github.com/AaronDennis07/electrum/internals/hub"
	"github.com/AaronDennis07/electrum/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func test(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "this is a test"})
}

func main() {
	err := godotenv.Load("./config/.env")

	if err != nil {
		log.Fatal("could not load config file ")
	}
	database.ConnectDB()

	app := fiber.New()
	app.Use(logger.New())
	go hub.RunHub()
	routers.SetupCourseRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
