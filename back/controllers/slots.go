package controllers

import (
	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/helper"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetAllSlots(c *fiber.Ctx) error {
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	salonId := claims["salonID"].(string)
	role := claims["role"].(string)
	if role == "admin" {
		salon := new(models.Salon)
		if err := c.BodyParser(salon); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		salonId = salon.ID.String()
	}
	var slots []models.Slot
	result := database.DB.Db.
		Preload("HairdressingStaff").
		Find(&slots, "salon_id = ?", salonId)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(slots)
}

func CreateSlot(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	slot := new(models.Slot)
	if err := c.BodyParser(slot); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	role := claims["role"].(string)
	if role == "manager" {
		salonIdToken, err := uuid.Parse(claims["salonID"].(string))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to parse salonID",
			})
		}
		SalonID := salonIdToken
		slot.SalonID = SalonID
	}

	slot.ID, _ = uuid.NewUUID()
	database.DB.Db.Create(&slot)
	return c.JSON(slot)
}

func UpdateSlot(c *fiber.Ctx) error {
	id := c.Params("id")
	slot := new(models.Slot)
	if err := c.BodyParser(slot); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result := database.DB.Db.Where("id = ?", id).Updates(&slot)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update slot",
		})
	}
	return c.JSON(slot)
}

func DeleteSlot(c *fiber.Ctx) error {
	id := c.Params("id")
	var slot models.Slot
	result := database.DB.Db.Unscoped().Where("id = ?", id).Delete(&slot)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(slot)
}
