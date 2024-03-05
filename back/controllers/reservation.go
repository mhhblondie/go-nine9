package controllers

import (
	"fmt"

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/helper"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetReservationById(c *fiber.Ctx) error {
	slotId := c.Params("slotId")
	var reservation []models.Reservation
	result := database.DB.Db.
		Preload("User").
		Preload("Slot").
		Preload("Slot.HairdressingStaff").
		Preload("Slot.HairdressingStaff.Salon").
		Where("reservations.slot_id = ?", slotId).
		Find(&reservation)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(reservation)
}

func CreateReservation(c *fiber.Ctx) error {

	customerId, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reservation := new(models.Reservation)
	if err := c.BodyParser(reservation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	reservation.ID, _ = models.GenerateUUID()

	customerIDString := customerId["id"].(string)
	customerID, err := uuid.Parse(customerIDString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reservation.CustomerID = customerID
	customerFirstname := customerId["firstname"].(string)
	result := database.DB.Db.Create(&reservation)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	slot := new(models.Slot)
	resultSlot := database.DB.Db.
		Preload("HairdressingStaff").
		Preload("HairdressingStaff.Salon").
		Where("id = ?", reservation.SlotID).
		First(&slot)

	if resultSlot.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": resultSlot.Error,
		})
	}

	salonName := slot.HairdressingStaff.Salon.Name

	hairdressingStaffName := slot.HairdressingStaff.Firstname
	date := slot.SlotTime
	dateStr := date.Format("2 janvier 2006 à 15h")

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	fmt.Println("Sending email to", customerId["email"].(string))
	body := helper.CreateConfirmationEmailBody(customerFirstname, dateStr, hairdressingStaffName, salonName)
	helper.SendConfirmationEmail(body, customerId["email"].(string), "Confirmation de réservation")
	return c.JSON(reservation)
}

func DeleteReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation models.Reservation
	result := database.DB.Db.Where("id = ?", id).Delete(&reservation)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Reservation successfully deleted",
	})
}
