package routers

import (
	"github.com/AaronDennis07/electrum/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupStudentRoutes(app *fiber.App) {
	api := app.Group("api/v1/student")

	api.Post("/upload", handlers.UploadStudent)

}
