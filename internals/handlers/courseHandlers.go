package handlers

import (
	"net/http"

	"github.com/AaronDennis07/electrum/internals/database"
	"github.com/AaronDennis07/electrum/internals/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCourse(c *fiber.Ctx) error {
	db := database.DB.Db
	course := new(models.Course)
	err := c.BodyParser(course)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid data recieved",
			"err":     err,
		})
	}

	err = db.Create(&course).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
			"err":     err,
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": course,
	})
}

func AllCourses(c *fiber.Ctx) error {
	var courses []models.Course
	db := database.DB.Db

	db.Find(&courses)

	return c.JSON(fiber.Map{
		"data": courses,
	})
}

func GetCourse(c *fiber.Ctx) error {
	db := database.DB.Db
	var course models.Course
	id := c.Params("id")

	err := db.Where("id=?", id).First(&course).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Course not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": course,
	})
}

func UpdateCourse(c *fiber.Ctx) error {
	db := database.DB.Db
	type UpdateCourse struct {
		Name string
		Code string
	}

	var course models.Course

	id := c.Params("id")

	err := db.Where("id=?", id).First(&course).Error

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Course not found",
		})
	}

	var updatedCourse UpdateCourse

	err = c.BodyParser(&updatedCourse)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid data recieved",
			"err":     err.Error(),
		})
	}

	course.Code = &updatedCourse.Code
	course.Name = &updatedCourse.Name

	db.Save(&course)

	return c.JSON(fiber.Map{
		"data": course,
	})
}

func DeleteCourse(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var course models.Course

	err := db.Where("id=?", id).First(&course).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Course not found",
		})
	}

	err = db.Delete(&course).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
			"err":     err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": "Course Deleted",
	})
}
