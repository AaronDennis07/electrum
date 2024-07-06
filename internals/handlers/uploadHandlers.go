package handlers

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/AaronDennis07/electrum/internals/database"
	"github.com/AaronDennis07/electrum/internals/models"
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
	db := database.DB.Db
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

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		fmt.Println(err)

	}
	nCreated := []string{}
	for _, row := range rows {
		var department models.Department
		err := db.Where("name=?", row[4]).First(&department).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error in finding department in db",
				"err":     err,
			})
		}

		student := models.Student{
			Usn:        row[0],
			Name:       &row[1],
			Email:      &row[2],
			Department: department,
		}
		err = db.Create(&student).Error
		if err != nil {
			nCreated = append(nCreated, row[0])
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "file uploaded successfully",
		"notCreated": nCreated,
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
	students, courses, err := parseExcel(uploadedFile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in parsing file",
			"err":     err,
		})
	}

	// //save the file from c.formfile
	// file, err := uploadedFile.Open()
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "error in opening file",
	// 		"err":     err,
	// 	})
	// }
	// defer file.Close()

	// f, err := excelize.OpenReader(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// students, err := f.GetCols("Sheet1")

	// if err != nil {
	// 	fmt.Println(err)

	// }
	// coursesRow, err := f.GetRows("Sheet2")
	// if err != nil {
	// 	fmt.Println(err)

	// }

	// var courses []map[string]interface{}
	// for _, row := range coursesRow {
	// 	mapRow := map[string]interface{}{
	// 		"code":       row[0],
	// 		"name":       row[1],
	// 		"seats":      row[2],
	// 		"department": row[3],
	// 	}
	// 	courses = append(courses, mapRow)
	// }

	// data := map[string]interface{}{
	// 	"students": students[0],
	// 	"courses":  courses,
	// }
	data := map[string]interface{}{
		"students": students[0],
		"courses":  courses,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "file uploaded successfully",
		"data":    data,
	})
}
func parseExcel(uploadedFile *multipart.FileHeader) ([]string, []models.CourseData, error) {
	file, err := uploadedFile.Open()
	if err != nil {
		return nil, nil, err
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

	var courses []models.CourseData
	for _, row := range coursesRow {
		seats, _ := strconv.ParseUint(row[2], 10, 32)
		mapRow := models.CourseData{
			Code:       &row[0],
			Name:       &row[1],
			Seats:      uint(seats),
			Department: &row[3],
		}
		courses = append(courses, mapRow)
	}

	// data := map[string]interface{}{
	// 	"students": students[0],
	// 	"courses":  courses,
	// }

	return students[0], courses, nil
}
