package routers

import (
	"github.com/AaronDennis07/electrum/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupCourseRoutes(app *fiber.App) {
	api := app.Group("/api/v1/courses")

	api.Get("", handlers.AllCourses)
	api.Get("/:id", handlers.GetCourse)
	api.Post("", handlers.CreateCourse)
	api.Put("/:id", handlers.UpdateCourse)
	api.Delete("/:id", handlers.DeleteCourse)
	//app.Get("/ws", websocket.New(handlers.WSHandler))
}
