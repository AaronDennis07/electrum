package handlers

import (
	"net/http"

	"github.com/AaronDennis07/electrum/internals/cache"
	"github.com/gofiber/fiber/v2"
)

type Session struct {
	Courses  map[string]string
	Students map[string]string
}

func StartSession(c *fiber.Ctx) error {
	var session Session
	err := c.BodyParser(&session)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Something went wrong",
			"err":     err,
		})
	}

	cache.Client.Redis.HSet(Ctx, "courses", session.Courses)
	cache.Client.Redis.HSet(Ctx, "students", session.Students)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"courses":  cache.Client.Redis.HGetAll(Ctx, "courses").Val(),
		"students": cache.Client.Redis.HGetAll(Ctx, "students").Val(),
	})
}
