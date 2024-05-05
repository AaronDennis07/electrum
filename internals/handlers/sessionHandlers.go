package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AaronDennis07/electrum/internals/cache"
	"github.com/AaronDennis07/electrum/internals/ctx"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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

	cache.Client.Redis.HSet(ctx.Ctx, "courses", session.Courses)
	cache.Client.Redis.HSet(ctx.Ctx, "students", session.Students)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"courses":  cache.Client.Redis.HGetAll(ctx.Ctx, "courses").Val(),
		"students": cache.Client.Redis.HGetAll(ctx.Ctx, "students").Val(),
	})
}

func SubscribeToSession(c *websocket.Conn) {

	// id := c.Params("id")

	// if exists := cache.Client.Redis.HGet(ctx.Ctx, "students", id); exists != nil {
	// 	return
	// }

	channel := "enroll"
	pubsub := cache.Client.Redis.Subscribe(ctx.Ctx, channel)

	ch := pubsub.Channel()

	for msg := range ch {
		res := msg
		fmt.Println(res.Payload)
		jsonMessage, err := json.Marshal(res.Payload)
		if err != nil {
			return
		}

		if err := c.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
			return
		} else {
			fmt.Println(jsonMessage)
		}
	}
}

// for {
// 	_, msg, err := c.ReadMessage()
// 	if err != nil {
// 		return
// 	}

// 	var message Message
// 	err = json.Unmarshal(msg, &message)
// 	if err != nil {
// 		return
// 	}

// 	cache.Client.Redis.Publish(ctx.Ctx, channel, message.Text)
// }

func EnrollToCourse(c *fiber.Ctx) error {
	channel := "enroll"
	req := struct {
		ID     string
		Course string
	}{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data recieved",
		})
	}
	course := cache.Client.Redis.HGet(ctx.Ctx, "courses", req.Course)
	if course == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "course does not exist",
		})
	}
	fmt.Println(req)
	cache.Client.Redis.HSet(ctx.Ctx, "students", req.ID, req.Course)
	cache.Client.Redis.HIncrBy(ctx.Ctx, "courses", req.Course, -1)
	courses, _ := cache.Client.Redis.HGetAll(ctx.Ctx, "courses").Result()
	jsonCourses, _ := json.Marshal(courses)
	fmt.Println(string(jsonCourses))
	cache.Client.Redis.Publish(ctx.Ctx, channel, string(jsonCourses))
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Successfully enrolled",
	})
}
