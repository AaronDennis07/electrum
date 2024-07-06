package routers

import (
	"github.com/AaronDennis07/electrum/internals/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupSessionhRoutes(app *fiber.App) {
	api := app.Group("/session")
	api.Get("/ws/:session", websocket.New(handlers.SubscribeToSession)) //studnet
	api.Post("", handlers.CreateSession)
	api.Post("/:session/start", handlers.StartSession)
	api.Post("/:session/enroll", handlers.AuthMiddlewareStudent, handlers.EnrollToCourse) //student
	api.Post("/:session/stop", handlers.StopSession)
	api.Get("", handlers.GetAllSessions)      // Add this line student
	api.Get("/:session", handlers.GetSession) //student
	api.Get("/details/:session", handlers.GetSessionDetails)
	api.Get("/:session/excel", handlers.SendEnrollmentsExcel)
	api.Post("/:session/courseupload", handlers.UploadCourse)
	api.Post("/:session/upload", handlers.UploadData)

}
