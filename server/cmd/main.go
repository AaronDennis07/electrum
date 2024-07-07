package main

import (
	"log"

	"github.com/AaronDennis07/electrum/internals/cache"
	"github.com/AaronDennis07/electrum/internals/database"
	"github.com/AaronDennis07/electrum/internals/handlers"
	"github.com/AaronDennis07/electrum/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Use(logger.New())

	routers.SetupCourseRoutes(app)
	routers.SetupAuthRoutes(app)
	routers.SetupSessionhRoutes(app)
	routers.SetupStudentRoutes(app)
	// app.Get("/ws/session/:session", websocket.New(handlers.SubscribeToSession)) //studnet
	// app.Post("/session", handlers.CreateSession)
	// app.Post("/session/:session/start", handlers.StartSession)
	// app.Post("/session/:session/enroll", handlers.EnrollToCourse) //student
	// app.Post("/session/:session/stop", handlers.StopSession)
	// app.Get("/session", handlers.GetAllSessions) // Add this line
	// app.Get("/session/:session", handlers.GetSession) //student
	// app.Get("/sessiondetails/:session", handlers.GetSessionDetails)
	// app.Get("/session/:session/excel", handlers.SendEnrollmentsExcel)
	// app.Post("/session/:session/courseupload", handlers.UploadCourse)
	// app.Post("/session/:session/upload", handlers.UploadData)
	app.Post("/students/upload", handlers.UploadStudent)
	log.Fatal(app.Listen(":8000"))
}
