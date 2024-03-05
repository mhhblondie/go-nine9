package controllers

import (
	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/helper"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetMe(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := claims["id"].(string)
	var user models.User
	database.DB.Db.
		Preload("Salon").
		Where("id = ?", id).
		First(&user)
	return c.JSON(user)
}

func UpdateMe(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := claims["id"].(string)
	role := claims["role"].(string)

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	if role != "admin" {
		user.Roles = role
	}

	result := database.DB.Db.Model(&models.User{}).Where("id = ?", id).Omit("password").Updates(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update the user",
		})
	}

	return c.JSON(user)
}

func UpdateMePassword(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id := claims["id"].(string)

	user := new(models.User)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := models.HashPassword(user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = hashedPassword
	result := database.DB.Db.Where("id = ?", id).Updates(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return c.SendString("User successfully updated")
}

func GetAllUsers(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	role := claims["role"].(string)
	var users []models.User

	if role == "admin" {
		database.DB.Db.
			Preload("Salon").
			Find(&users)
		return c.JSON(users)
	}

	salonID, err := uuid.Parse(claims["salonID"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse salonID",
		})
	}
	database.DB.Db.
		Preload("Salon").
		Where("salon_id = ?", salonID).
		Find(&users)
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	role := claims["role"].(string)
	if role == "manager" {
		salonID, err := uuid.Parse(claims["salonID"].(string))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to parse salonID",
			})
		}
		database.DB.Db.Where("id = ? AND salon_id = ?", id, salonID).First(&user)
		return c.JSON(user)
	}
	database.DB.Db.
		Preload("Salon").
		Where("id = ?", id).
		First(&user)
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {

	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	role := claims["role"].(string)
	roleUser := "user"
	var salonId uuid.UUID

	if role == "manager" {
		salonID, err := uuid.Parse(claims["salonID"].(string))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to parse salonID",
			})
		}
		salonId = salonID
		roleUser = "employee"
		user.Roles = roleUser
		if salonId != uuid.Nil {
			user.SalonID = &salonId
		}
	}
	user.ID, _ = models.GenerateUUID()
	hashedPassword, err := models.HashPassword(helper.GeneratePassword(8))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = hashedPassword
	database.DB.Db.Create(&user)
	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := models.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = hashedPassword
	role := user.Roles
	if role == "manager" {
		salonID, err := uuid.Parse(user.SalonID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to parse salonID",
			})
		}
		user.SalonID = &salonID
		user.Roles = "employee"
		result := database.DB.Db.Where("id = ? AND salon_id = ?", id, salonID).Updates(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update user",
			})
		}
		return c.JSON(user)
	}

	result := database.DB.Db.Where("id = ?", id).Updates(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	role := claims["role"].(string)
	if role == "manager" {
		salonID, err := uuid.Parse(claims["salonID"].(string))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to parse salonID",
			})
		}
		database.DB.Db.Where("id = ? AND salon_id = ?", id, salonID).Delete(&user)
		return c.JSON(user)
	}
	database.DB.Db.Where("id = ?", id).Delete(&user)
	return c.JSON(user)
}
