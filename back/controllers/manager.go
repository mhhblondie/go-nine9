package controllers

import "github.com/gofiber/fiber/v2"

func GetAllOwnSlots(c *fiber.Ctx) error {
	return c.SendString("GetAllOwnSlots")
}
