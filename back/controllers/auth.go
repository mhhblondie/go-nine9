package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/helper"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = os.Getenv("JWT_SECRET")

func CreateNewUser(c *fiber.Ctx) error {
	var body map[string]interface{}
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		return err
	}
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user.Roles == "manager" {
		salon := new(models.Salon)
		salon.ID, _ = models.GenerateUUID()
		salonInfo, ok := body["SalonInfo"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("SalonInfo is not a map")
		}
		address, ok := salonInfo["address"].(string)
		if !ok {
			return fmt.Errorf("address is not a string")
		}
		phone, ok := salonInfo["phone"].(string)
		if !ok {
			return fmt.Errorf("phone is not a string")
		}
		name, ok := salonInfo["name"].(string)
		if !ok {
			return fmt.Errorf("name is not a string")
		}
		description, ok := salonInfo["description"].(string)
		if !ok {
			return fmt.Errorf("description is not a string")
		}
		salon.Description = description
		salon.Address = address
		salon.Phone = phone
		salon.Name = name
		database.DB.Db.Create(&salon)
		user.SalonID = &salon.ID

		users, ok := body["Users"].([]interface{})
		if !ok {
			return fmt.Errorf("Users is not a slice")
		}
		for _, user := range users {
			employee := new(models.User)
			employee.ID, _ = models.GenerateUUID()
			employeeInfo, ok := user.(map[string]interface{})
			if !ok {
				return fmt.Errorf("User is not a map")
			}
			employee.Firstname, ok = employeeInfo["firstname"].(string)
			if !ok {
				return fmt.Errorf("firstname is not a string")
			}
			employee.Lastname, ok = employeeInfo["lastname"].(string)
			if !ok {
				return fmt.Errorf("lastname is not a string")
			}
			employee.Email, ok = employeeInfo["email"].(string)
			if !ok {
				return fmt.Errorf("email is not a string")
			}
			employee.Password = helper.GeneratePassword(8)
			body := helper.CreateMailBody(employee.Firstname, employee.Email, employee.Password, salon.Name)
			helper.SendConfirmationEmail(body, employee.Email, "Confirmation de création de compte")

			employee.Roles = "staff"
			hashedPassword, err := models.HashPassword(employee.Password)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Failed to hash password",
				})
			}
			employee.Password = hashedPassword
			employee.SalonID = &salon.ID
			database.DB.Db.Create(&employee)
		}
	}

	hashedPassword, err := models.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = hashedPassword

	user.ID, err = models.GenerateUUID()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate UUID",
		})
	}

	if user.Roles == "" {
		user.Roles = "users"
	}

	database.DB.Db.Create(&user)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 1 jour
		"role":      user.Roles,
		"firstname": user.Firstname,
		"email":     user.Email,
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}
	userWithToken := fiber.Map{
		"jwt": token,
	}

	return c.Status(200).JSON(userWithToken)
}

func LoginUser(c *fiber.Ctx) error {

	loginUser := new(models.User)

	if err := c.BodyParser(loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var existingUser models.User
	if err := database.DB.Db.Where("email = ?", loginUser.Email).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := models.VerifyPassword(existingUser.Password, loginUser.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        existingUser.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 1 jour
		"role":      existingUser.Roles,
		"firstname": existingUser.Firstname,
		"email":     existingUser.Email,
		"salonID":   existingUser.SalonID,
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}
	response := fiber.Map{
		"jwt": token,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
