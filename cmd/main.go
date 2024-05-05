package main

import (
	"log"

	"github.com/AaronDennis07/electrum/internals/cache"
	"github.com/AaronDennis07/electrum/internals/database"
	"github.com/AaronDennis07/electrum/internals/handlers"
	"github.com/AaronDennis07/electrum/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./config/.env")

	if err != nil {
		log.Fatal("could not load config file ")
	}
	database.ConnectDB()
	cache.SetupCache()
	app := fiber.New()
	app.Use(logger.New())

	routers.SetupCourseRoutes(app)

	app.Get("/ws/session", websocket.New(handlers.SubscribeToSession))
	app.Post("/session/start", handlers.StartSession)
	app.Post("/session/enroll", handlers.EnrollToCourse)
	log.Fatal(app.Listen(":8000"))
}
