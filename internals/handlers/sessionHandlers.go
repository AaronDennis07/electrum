package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AaronDennis07/electrum/internals/cache"
	"github.com/AaronDennis07/electrum/internals/ctx"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Session struct {
	Courses  map[string]string
	Students map[string]string
}

func StartSession(c *fiber.Ctx) error {
	sessionName  := c.Params("session")
	courseKey := sessionName + ":courses"
	studentKey := sessionName + ":students"

	var session Session
	err := c.BodyParser(&session)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Something went wrong",
			"err":     err,
		})
	}
	


	if sessionExists(courseKey, studentKey) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Session: "+sessionName +" already exists",
		})
	}


	cache.Client.Redis.HSet(ctx.Ctx, courseKey, session.Courses)
	cache.Client.Redis.HSet(ctx.Ctx, studentKey, session.Students)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"session": sessionName,
		"courses":  cache.Client.Redis.HGetAll(ctx.Ctx, courseKey).Val(),
		"students": cache.Client.Redis.HGetAll(ctx.Ctx, studentKey).Val(),
	})
}

func SubscribeToSession(c *websocket.Conn) {

	// id := c.Params("id")

	// if exists := cache.Client.Redis.HGet(ctx.Ctx, "students", id); exists != nil {
	// 	return
	// }

	channel := c.Params("session")
	courseKey := channel + ":courses"
	pubsub := cache.Client.Redis.Subscribe(ctx.Ctx, channel)

	ch := pubsub.Channel()
	courses := cache.Client.Redis.HGetAll(ctx.Ctx, courseKey).Val()
	if len(courses) != 0 {
		jsonCourses, _ := json.Marshal(courses)
		if err := c.WriteMessage(websocket.TextMessage, jsonCourses); err != nil {
			return
		} 
	}
	for msg := range ch {
		res := msg
		fmt.Println("payload:"+res.Payload)

		var payloadMap map[string]interface{}
		err := json.Unmarshal([]byte(res.Payload), &payloadMap)
		if err != nil {
			return
		}

		jsonMessage, err := json.Marshal(payloadMap)
		if err != nil {
			return
		}

		if err := c.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
			return
		} else {
			fmt.Println(string(jsonMessage))
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

func sessionExists(key1 string, key2 string) bool {
	return cache.Client.Redis.Exists(ctx.Ctx, key1,key2).Val() > 0 
}

func EnrollToCourse(c *fiber.Ctx) error {
	channel := c.Params("session")
	courseKey := channel + ":courses"
	studentKey := channel + ":students"
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
	if !sessionExists(courseKey, studentKey){
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "session does not exist",
		})
	}

	course := cache.Client.Redis.HGet(ctx.Ctx, courseKey, req.Course).Val()
	if course == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Course does not exist",
		})
	}

	student, err := cache.Client.Redis.HGet(ctx.Ctx, studentKey, req.ID).Result()
	if err == redis.Nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Student does not exist",
		})
	} else if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	if student != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Student already enrolled in course:" + student,
		})
	}

	courseInt, err := strconv.Atoi(course)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid course",
		})
	}

	if courseInt <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Course is full",
		})
	}

	cache.Client.Redis.HSet(ctx.Ctx, studentKey, req.ID, req.Course)
	cache.Client.Redis.HIncrBy(ctx.Ctx, courseKey, req.Course, -1)

	courses := cache.Client.Redis.HGetAll(ctx.Ctx, courseKey).Val()

	jsonCourses, _ := json.Marshal(courses)
	fmt.Println("jsoncourses"+string(jsonCourses))
	cache.Client.Redis.Publish(ctx.Ctx, channel, string(jsonCourses))

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Successfully enrolled",
	})
}

func StopSession(c *fiber.Ctx) error {
	sessionName := c.Params("session")
	isDeleted := cache.Client.Redis.Del(ctx.Ctx, sessionName+":courses",sessionName+":students").Val()
	
	if isDeleted == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Session does not exist",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Session stopped",
	})
}
