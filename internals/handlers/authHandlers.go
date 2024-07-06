package handlers

import (
	"fmt"
	"os"
	"strings"
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

type AdminRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AdminLoginRequest struct {
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
func RegisterAdmin(c *fiber.Ctx) error {
	request := new(AdminRegisterRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	db := database.DB.Db
	var admin models.Admin
	result := db.Where("name = ?", request.Name).Or("email = ?", request.Email).First(&admin)

	if result.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name or Email already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot hash password"})
	}
	hashedPasswordString := string(hashedPassword)
	admin.Name = request.Name
	admin.Email = &request.Email
	admin.Password = &hashedPasswordString
	result = db.Create(&admin)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot register admin"})
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
func LoginAdmin(c *fiber.Ctx) error {
	loginAdmin := new(AdminLoginRequest)
	secret := os.Getenv("JWT_SECRET")

	if err := c.BodyParser(loginAdmin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var admin models.Admin
	db := database.DB.Db

	result := db.Where("email = ?", loginAdmin.Email).First(&admin)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*admin.Password), []byte(loginAdmin.Password)); err == nil {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = admin.Email
		claims["name"] = admin.Name
		claims["is_admin"] = true
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
	authHeader := c.Get("Authorization")
	secret := os.Getenv("JWT_SECRET")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	// Check for "Bearer " prefix and remove it
	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
	if bearerToken == authHeader {
		// If the token didn't have "Bearer " prefix, return an error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token: " + err.Error()})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	if claims["is_admin"] == true {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Route only accessible to Students"})
	}

	c.Locals("name", claims["name"])
	c.Locals("usn", claims["usn"])
	c.Locals("is_admin", claims["is_admin"])

	return c.Next()
}

func AuthMiddlewareAdmin(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	secret := os.Getenv("JWT_SECRET")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	// Check for "Bearer " prefix and remove it
	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
	if bearerToken == authHeader {
		// If the token didn't have "Bearer " prefix, return an error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token: " + err.Error()})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	if claims["is_admin"] == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Route only accessible to Admins"})
	}

	c.Locals("name", claims["name"])
	c.Locals("email", claims["email"])
	c.Locals("is_admin", claims["is_admin"])

	return c.Next()
}
