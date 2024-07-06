package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/AaronDennis07/electrum/internals/database"
	"github.com/AaronDennis07/electrum/internals/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type StudentRegisterRequest struct {
	USN      string `json:"usn"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentLoginRequest struct {
	USN      string `json:"usn"`
	Password string `json:"password"`
}

func isPasswordEmpty(student models.Student) bool {

	return student.Password == nil || *student.Password == ""
}

func RegisterStudent(c *fiber.Ctx) error {
	request := new(StudentRegisterRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db := database.DB.Db
	var student models.Student
	result := db.Where("USN = ?", request.USN).First(&student)

	if !isPasswordEmpty(student) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Student already registered"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No record found please contact the administrator"})
	}
	fmt.Println(student)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot hash password"})
	}
	hashedPasswordString := string(hashedPassword)

	result = db.Model(&student).Update("password", hashedPasswordString)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot register student"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Registration successful"})
}

func LoginStudent(c *fiber.Ctx) error {
	loginStudent := new(StudentLoginRequest)
	secret := os.Getenv("JWT_SECRET")

	if err := c.BodyParser(loginStudent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var student models.Student
	db := database.DB.Db

	result := db.Where("USN = ?", loginStudent.USN).First(&student)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if isPasswordEmpty(student) {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Not registered"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*student.Password), []byte(loginStudent.Password)); err == nil {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["usn"] = student.Usn
		claims["name"] = student.Name
		claims["is_admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(secret))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": t})
	}

	return c.SendStatus(fiber.StatusUnauthorized)
}

func AuthMiddlewareStudent(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	secret := os.Getenv("JWT_SECRET")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["is_admin"] == true {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Route only accessible to students"})
	}
	c.Locals("name", claims["name"])
	c.Locals("usn", claims["usn"])
	c.Locals("is_admin", claims["is_admin"])
	return c.Next()
}

func AuthMiddlewareAdmin(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	secret := os.Getenv("JWT_SECRET")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["is_admin"] == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Route only accessible to Admin"})
	}
	c.Locals("name", claims["name"])
	c.Locals("usn", claims["usn"])
	c.Locals("is_admin", claims["is_admin"])
	return c.Next()
}
func AuthMiddlewareGeneral(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	secret := os.Getenv("JWT_SECRET")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims := token.Claims.(jwt.MapClaims)

	c.Locals("name", claims["name"])
	c.Locals("usn", claims["usn"])
	c.Locals("is_admin", claims["is_admin"])
	return c.Next()
}
