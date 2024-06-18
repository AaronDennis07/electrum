package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AaronDennis07/electrum/internals/cache"
	"github.com/AaronDennis07/electrum/internals/ctx"
	"github.com/AaronDennis07/electrum/internals/database"
	"github.com/AaronDennis07/electrum/internals/models"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Session struct {
	Courses  map[string]string
	Students map[string]string
}
type SessionRequest struct {
	Session models.Session `json:"session"`
	// Courses  []CourseRequest `json:"courses"`
	// Students []string        `json:"students"`
}
type CourseRequest struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Department string `json:"department"`
}

func CreateSession(c *fiber.Ctx) error {
	db := database.DB.Db

	uploadedFile, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in uploading file",
			"err":     err,
		})
	}
	students, CourseData, err := parseExcel(uploadedFile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in parsing file",
			"err":     err,
		})
	}

	request := new(SessionRequest)
	//session
	//courses
	//students
	re := c.FormValue("data")
	err = json.Unmarshal([]byte(re), &request)
	// err = c.BodyParser(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data recieved",
			"err":     err,
		})
	}

	//adding courses
	for _, reqCourse := range CourseData {
		var department models.Department
		err := db.Where("name=?", reqCourse.Department).First(&department).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "error in finding department in db",
				"err":     err,
			})
		}
		course := models.Course{
			Name:       reqCourse.Name,
			Code:       reqCourse.Code,
			Seats:      &reqCourse.Seats,
			Department: department,
		}
		err = db.Create(&course).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "error in creating course in db",
				"err":     err,
			})
		}
		// var cc models.Course
		// db.Preload("Department").Where("name=?", course.Name).First(&cc)
		// log.Println("just checking")
		// log.Println(*cc.DepartmentID)
		// log.Println(cc.ID)
		// log.Println(*cc.Name)
		// log.Println(*cc.Department.Name)
		request.Session.Courses = append(request.Session.Courses, course)
	}

	//creating session
	err = db.Create(&request.Session).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error in creating session in db",
			"err":     err,
		})
	}

	//checking if students exist
	var notFound []string = []string{}
	for _, usn := range students {
		var student models.Student
		err = db.Where("usn=?", usn).First(&student).Error
		if err != nil {
			notFound = append(notFound, usn)
			continue
		}
		err = db.Create(&models.Enrollment{
			StudentID: &usn,
			SessionID: &request.Session.ID,
		}).Error
		if err != nil {
			log.Println("creating enrollment: ", err)
		}
	}

	var createdSession models.Session
	err = db.Preload("Courses").Find(&createdSession, request.Session.ID).Error
	if err != nil {
		log.Println("loading session:", err)
	}
	var enrolledStudents []models.Enrollment
	err = db.Preload("Student").Where("session_id=?", request.Session.ID).Find(&enrolledStudents).Error
	if err != nil {
		log.Println("loading enrollment:", err)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"session":  createdSession,
		"enrolled": enrolledStudents,
		"notFound": notFound,
	})
}

func StartSession(c *fiber.Ctx) error {

	db := database.DB.Db
	// var session models.Session
	// db.First(&session, c.Params("session"))
	// if session.ID == 0 {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Session does not exist",
	// 	})
	// }

	sessionName := c.Params("session")

	var sessionDb models.Session
	err := db.Preload("Courses").Where("name=?", sessionName).First(&sessionDb).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Session does not exist",
			"err":     err,
		})
	}
	courseKey := sessionName + ":courses"
	studentKey := sessionName + ":students"

	// var session Session
	// err = c.BodyParser(&session)
	// if err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Something went wrong",
	// 		"err":     err,
	// 	})
	// }

	if sessionExists(courseKey, studentKey) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Session: " + sessionName + " already exists",
		})
	}

	for _, course := range sessionDb.Courses {
		err = cache.Client.Redis.HSet(ctx.Ctx, courseKey, *course.Code, *course.Seats).Err()
		if err != nil {
			log.Println("populating session redis", err)
		}
	}

	var enrollments []models.Enrollment
	db.Preload("Student").Preload("Session").Where("Session_ID = ?", sessionDb.ID).Find(&enrollments)
	for _, enrollment := range enrollments {
		err = cache.Client.Redis.HSet(ctx.Ctx, studentKey, *enrollment.StudentID, "").Err()
		if err != nil {
			log.Println("populating students redis", err)
		}
	}

	// cache.Client.Redis.HSet(ctx.Ctx, courseKey, session.Courses)
	// cache.Client.Redis.HSet(ctx.Ctx, studentKey, session.Students)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"session":  sessionName,
		"courses":  cache.Client.Redis.HGetAll(ctx.Ctx, courseKey).Val(),
		"students": cache.Client.Redis.HGetAll(ctx.Ctx, studentKey).Val(),
	})
}

func SubscribeToSession(c *websocket.Conn) {

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
		fmt.Println("payload:" + res.Payload)

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

func sessionExists(key1 string, key2 string) bool {
	return cache.Client.Redis.Exists(ctx.Ctx, key1, key2).Val() > 0
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
	if !sessionExists(courseKey, studentKey) {
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
	if err == redis.Nil {
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

	go func() {
		db := database.DB.Db
		var session models.Session
		db.Where("name=?", channel).First(&session)
		var course models.Course
		db.Where("code=?", req.Course).First(&course)
		db.Model(&models.Enrollment{}).Where("session_id=? AND student_id=?", session.ID, req.ID).Update("course1_id", course.ID)
	}()

	jsonCourses, _ := json.Marshal(courses)
	cache.Client.Redis.Publish(ctx.Ctx, channel, string(jsonCourses))

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Successfully enrolled",
	})
}

func StopSession(c *fiber.Ctx) error {
	sessionName := c.Params("session")
	isDeleted := cache.Client.Redis.Del(ctx.Ctx, sessionName+":courses", sessionName+":students").Val()

	if isDeleted == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Session does not exist",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Session stopped",
	})
}

type CourseData struct {
	Id         uint
	Name       string
	Code       string
	Seats      uint
	Department string
}

func GetSession(c *fiber.Ctx) error {
	db := database.DB.Db
	sessionName := c.Params("session")
	var session models.Session
	db.Where("name=?", sessionName).First(&session)
	sessionId := session.ID
	var courses []models.Course
	db.Preload("Department").Where("session_id=?", sessionId).Find(&courses)
	courseData := []CourseData{}
	for _, course := range courses {
		courseData = append(courseData, CourseData{
			Id:         course.ID,
			Name:       *course.Name,
			Code:       *course.Code,
			Seats:      *course.Seats,
			Department: *course.Department.Name,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		// "session": session,
		"courses": courseData,
	})
}
