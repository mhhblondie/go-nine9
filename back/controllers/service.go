package controllers

import (
	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ServiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SalonID     string `json:"salon_id"`
}

func CreateService(c *fiber.Ctx) error {

	var serviceRequest ServiceRequest
	if err := c.BodyParser(&serviceRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	salonID, err := uuid.Parse(serviceRequest.SalonID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid salon ID",
		})
	}

	service := models.Service{
		ID:          uuid.New(),
		Name:        serviceRequest.Name,
		Description: serviceRequest.Description,
		SalonID:     &salonID,
	}

	service.ID, _ = models.GenerateUUID()

	result := database.DB.Db.Create(&service)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	response := fiber.Map{
		"message": "Service créé avec succès",
		"service": service,
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func UpdateService(c *fiber.Ctx) error {
	var serviceRequest ServiceRequest
	if err := c.BodyParser(&serviceRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	serviceID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid service ID",
		})
	}

	var existingService models.Service
	result := database.DB.Db.First(&existingService, "id = ?", serviceID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Service not found",
		})
	}

	existingService.Name = serviceRequest.Name
	existingService.Description = serviceRequest.Description

	result = database.DB.Db.Save(&existingService)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	response := fiber.Map{
		"message":         "Service mis à jour avec succès",
		"updated_service": existingService,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func DeleteService(c *fiber.Ctx) error {
	serviceID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid service ID",
		})
	}

	var service models.Service
	result := database.DB.Db.First(&service, "id = ?", serviceID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Service not found",
		})
	}

	result = database.DB.Db.Delete(&service)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	response := fiber.Map{
		"message": "Service supprimé avec succès",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
