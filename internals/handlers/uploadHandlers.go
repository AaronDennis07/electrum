package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func UploadCourse(c *fiber.Ctx) error {
	courseFile, err := c.FormFile("courses")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in uploading file",
			"err":     err,
		})
	}

	file, err := courseFile.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in opening file",
			"err":     err,
		})
	}
	defer file.Close()

	f, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Println(err)
	}

	cols, err := f.GetCols("Sheet1")

	if err != nil {
		fmt.Println(err)

	}

	data := map[string]interface{}{
		"course": cols[0],
	}
	out, _ := json.Marshal(data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "file uploaded successfully",
		"data":    string(out),
	})
}

func UploadStudent(c *fiber.Ctx) error {
	courseFile, err := c.FormFile("student")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in uploading file",
			"err":     err,
		})
	}

	file, err := courseFile.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in opening file",
			"err":     err,
		})
	}
	defer file.Close()

	f, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Println(err)
	}

	cols, err := f.GetCols("Sheet1")

	if err != nil {
		fmt.Println(err)

	}

	data := map[string]interface{}{
		"students": cols[0],
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "file uploaded successfully",
		"data":    data,
	})
}

func UploadData(c *fiber.Ctx) error {
	uploadedFile, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in uploading file",
			"err":     err,
		})
	}

	//save the file from c.formfile
	file, err := uploadedFile.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in opening file",
			"err":     err,
		})
	}
	defer file.Close()

	f, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Println(err)
	}

	students, err := f.GetCols("Sheet1")

	if err != nil {
		fmt.Println(err)

	}
	coursesRow, err := f.GetRows("Sheet2")
	if err != nil {
		fmt.Println(err)

	}

	var courses []map[string]interface{}
	for _, row := range coursesRow {
		mapRow := map[string]interface{}{
			row[0]: map[string]interface{}{
				"name":  row[1],
				"seats": row[2],
			},
		}
		courses = append(courses, mapRow)
	}

	data := map[string]interface{}{
		"students": students[0],
		"courses":  courses,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "file uploaded successfully",
		"data":    data,
	})
}
