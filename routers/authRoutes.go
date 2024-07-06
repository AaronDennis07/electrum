package routers

import (
	"github.com/AaronDennis07/electrum/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	api := app.Group("")

	api.Post("/auth/student/register", handlers.RegisterStudent)
	api.Post("/auth/student/login", handlers.LoginStudent)
	api.Post("/auth/admin/register", handlers.RegisterAdmin)
	api.Post("/auth/admin/login", handlers.LoginAdmin)

}
