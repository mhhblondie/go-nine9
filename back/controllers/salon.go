package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/helper"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetMySalons(c *fiber.Ctx) error {
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	if claims["salonID"] == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "noon",
		})
	}

	salonId := claims["salonID"].(string)

	var salons []models.Salon
	result := database.DB.Db.
		Preload("User").
		Preload("User.Slots").
		Preload("Service.Prestation").
		Find(&salons, "id = ?", salonId)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(salons)
}

func GetSalons(c *fiber.Ctx) error {
	var salons []models.Salon
	result := database.DB.Db.
		Preload("User").
		Preload("User.Salon").
		Preload("User.Slots").
		Preload("Hours").
		Preload("Service").
		Find(&salons)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(salons)
}

type SalonRequest struct {
	User        []models.User `json:"user"`
	Phone       string        `json:"phone"`
	Name        string        `json:"name"`
	Address     string        `json:"address"`
	Description string        `json:"description"`
}

func CreateSalon(c *fiber.Ctx) error {
	// Vérifie si les revendications de l'utilisateur existent
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	// Parse la requete
	var salonRequest SalonRequest
	if err := c.BodyParser(&salonRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	salon := models.Salon{
		Name:        salonRequest.Name,
		Address:     salonRequest.Address,
		Phone:       salonRequest.Phone,
		Description: salonRequest.Description,
	}

	//génère un UUID pour le salon
	salon.ID, _ = models.GenerateUUID()

	// récupère l'id de la personne qui créé le salon
	userID := claims["id"].(string)

	// Convert the manager ID from string to UUID
	managerID, _ := uuid.Parse(userID)

	// Find the manager by ID
	var manager models.User
	result := database.DB.Db.First(&manager, "id = ?", managerID)

	// Check if there is an error and the manager doesn't exist
	if result.Error != nil && result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Manager not found",
		})
	}

	if manager.SalonID != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Vous avez déja un salon",
		})
	}

	role := claims["role"].(string)
	// Associate the salon to manager or create a new manager if not found
	manager.SalonID = &salon.ID

	//crée le salon
	result = database.DB.Db.Create(&salon)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	// Vérifie si l'utilisateur est déjà un employé d'un salon
	if role == "employee" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "You already have a salon",
		})
	}

	// Vérifie si l'utilisateur a un rôle utilisateur
	if role == "users" {
		manager.Roles = "manager"

	}

	// Met à jour l'utilisateur dans la base de données avec le nouveau rôle et l'ID du salon
	if err := database.DB.Db.Save(&manager).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messageHooo": err.Error(),
		})
	}

	signedToken, ok := helper.GenerateToken(manager.ID, manager.Roles, manager.Firstname, manager.Email, salon.ID)

	if ok != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not ",
		})
	}
	// Create staff users
	users := salonRequest.User
	for i := 0; i < len(users); i++ {
		var user models.User
		user = users[i]
		passwordToSendEMail := users[i].Password
		// Assign the salonId to staff
		user.SalonID = &salon.ID
		user.Roles = "staff"
		user.ID, _ = models.GenerateUUID()

		hashedPassword, err := models.HashPassword(user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to hash password",
			})
		}
		user.Password = hashedPassword

		if err := database.DB.Db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"messageHooo": err.Error(),
			})
		}

		// Send Email to the staff
		now := time.Now()
		year := fmt.Sprintf("%d", now.Year())
		month := fmt.Sprintf("%02d", now.Month())
		day := fmt.Sprintf("%02d", now.Day())

		dateStr := year + "-" + month + "-" + day

		fmt.Println("Sending email to", user.Email)
		body := helper.CreateNewStaffBody(user.Firstname, dateStr, salon.Name, passwordToSendEMail)
		helper.SendConfirmationEmail(body, user.Email, "Nouveau compte Planity")

	}

	response := fiber.Map{
		"jwt": signedToken,
	}

	return c.Status(fiber.StatusOK).JSON(response)

}

func CreateSalonAdmin(c *fiber.Ctx) error {

	// Parse la requete
	var salonRequest SalonRequest
	if err := c.BodyParser(&salonRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	salon := models.Salon{
		Name:        salonRequest.Name,
		Address:     salonRequest.Address,
		Phone:       salonRequest.Phone,
		Description: salonRequest.Description,
	}

	//génère un UUID pour le salon
	salon.ID, _ = models.GenerateUUID()

	//crée le salon
	result := database.DB.Db.Create(&salon)

	// Create staff users
	users := salonRequest.User
	for i := 0; i < len(users); i++ {
		var user models.User
		user = users[i]
		passwordToSendEMail := users[i].Password
		// Assign the salonId to staff
		user.SalonID = &salon.ID

		user.ID, _ = models.GenerateUUID()

		hashedPassword, err := models.HashPassword(user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to hash password",
			})
		}
		user.Password = hashedPassword

		if i == 0 {
			user.Roles = "manager"

		} else {
			user.Roles = "staff"
		}

		if err := database.DB.Db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"erreur de création": err.Error(),
			})
		}

		// Send Email to the staff
		now := time.Now()
		year := fmt.Sprintf("%d", now.Year())
		month := fmt.Sprintf("%02d", now.Month())
		day := fmt.Sprintf("%02d", now.Day())

		dateStr := year + "-" + month + "-" + day

		fmt.Println("Sending email to", user.Email)
		body := helper.CreateNewStaffBody(user.Firstname, dateStr, salon.Name, passwordToSendEMail)
		helper.SendConfirmationEmail(body, user.Email, "Nouveau compte Planity")

	}

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	response := fiber.Map{
		"message": "Salon crée",
	}

	return c.Status(fiber.StatusOK).JSON(response)

}

func GetSalonById(c *fiber.Ctx) error {
	id := c.Params("id")
	var salon models.Salon
	result := database.DB.Db.
		Preload("Hours").
		Preload("User").
		Preload("User.Slots").
		Preload("Service").
		Preload("Service.Prestation").
		Where("id = ?", id).First(&salon)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(salon)
}

func AddStaff(c *fiber.Ctx) error {
	// Retrieve user claims
	userClaims, ok := c.Locals("userClaims").(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	// Parse salonID from claims
	salonIDClaim, ok := userClaims["salonID"]
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "SalonID not found in claims",
		})
	}

	salonID, err := uuid.Parse(salonIDClaim.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse salonID",
		})
	}

	var salon models.Salon
	result := database.DB.Db.Where("id = ?", salonID).First(&salon)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Couldn't find salon",
		})
	}

	var users []models.User
	if err := c.BodyParser(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Create user for each iteration of array
	for i := 0; i < len(users); i++ {
		var user models.User
		user = users[i]
		passwordToSendEMail := users[i].Password
		// Assign the salonId to staff
		user.SalonID = &salonID
		user.Roles = "staff"
		user.ID, _ = models.GenerateUUID()

		hashedPassword, err := models.HashPassword(user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to hash password",
			})
		}
		user.Password = hashedPassword

		database.DB.Db.Create(&user)

		// Send Email to the staff
		now := time.Now()
		year := fmt.Sprintf("%d", now.Year())
		month := fmt.Sprintf("%02d", now.Month())
		day := fmt.Sprintf("%02d", now.Day())

		dateStr := year + "-" + month + "-" + day

		fmt.Println("Sending email to", user.Email)
		body := helper.CreateNewStaffBody(user.Firstname, dateStr, salon.Name, passwordToSendEMail)
		helper.SendConfirmationEmail(body, user.Email, "Nouveau compte Planity")

	}

	return c.SendString("Salon successfully updated")
}

func UpdateSalon(c *fiber.Ctx) error {
	id := c.Params("id")
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}
	salon := new(models.Salon)

	if err := c.BodyParser(salon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// If manager we get the id in the token, if admin, we retrieve the id in the path
	role := claims["role"].(string)
	if role == "manager" {
		salon.ID, _ = uuid.Parse(claims["salonID"].(string))
	} else {
		salon.ID, _ = uuid.Parse(id)
	}

	result := database.DB.Db.Updates(&salon)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update salon",
		})
	}
	return c.SendString("Salon successfully updated")
}

func DeleteSalon(c *fiber.Ctx) error {

	id := c.Params("id")

	// Convert the Salon Id string into UUID
	salonId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid salon ID format",
		})
	}
	// Find the manager of the salon
	var manager models.User
	result := database.DB.Db.Where("salon_id = ? AND roles = ?", salonId, "manager").First(&manager)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot find manager",
		})
	}

	// Set the SalonId to null to not be deleted after
	manager.SalonID = nil
	result = database.DB.Db.Save(&manager)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot set null to manager",
		})
	}

	// Send Message to deleted staff
	var users []models.User
	result = database.DB.Db.Where("salon_id = ? AND roles = ?", salonId, "staff").Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot find staff",
		})
	}
	for i := 0; i < len(users); i++ {
		var user models.User
		user = users[i]
		// Send Email to the staff
		now := time.Now()
		year := fmt.Sprintf("%d", now.Year())
		month := fmt.Sprintf("%02d", now.Month())
		day := fmt.Sprintf("%02d", now.Day())

		dateStr := year + "-" + month + "-" + day

		fmt.Println("Sending email to", user.Email)
		body := helper.CreateDeleteStaffBody(user.Firstname, dateStr, user.Salon.Name, user.Salon.Phone)
		helper.SendConfirmationEmail(body, user.Email, "Compte supprimé")

		// Delete all staff related and the salon
		result = database.DB.Db.Unscoped().Delete(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Cannot remove user Staff",
			})
		}

	}

	// Delete the salon
	var salon models.Salon
	result = database.DB.Db.Where("id = ?", id).Delete(&salon)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot remove Salon",
		})
	}

	return c.SendString("Salon successfully deleted")
}

func DeleteStaff(c *fiber.Ctx) error {
	staffIdString := c.Params("staffID")

	// Convert the Salon Id string into UUID
	staffId, err := uuid.Parse(staffIdString)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid salon ID format",
		})
	}

	// Delete all reservations relation to this staff
	var slot models.Slot
	result := database.DB.Db.Where("hairdressing_staff_id = ?", staffId).Unscoped().Delete(&slot)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot remove user slots",
		})
	}

	// Find the staff to retrieve the mail
	var staff models.User
	result = database.DB.Db.Preload("Salon").Where("id = ? AND roles = ?", staffId, "staff").First(&staff)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot find manager",
		})
	}

	result = database.DB.Db.Where("id = ?", staffId).Unscoped().Delete(&staff)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot remove user Staff",
		})
	}

	// Send Email to the staff
	now := time.Now()
	year := fmt.Sprintf("%d", now.Year())
	month := fmt.Sprintf("%02d", now.Month())
	day := fmt.Sprintf("%02d", now.Day())

	dateStr := year + "-" + month + "-" + day

	fmt.Println("Sending email to", staff.Email)
	body := helper.CreateDeleteStaffBody(staff.Firstname, dateStr, staff.Salon.Name, staff.Salon.Phone)
	helper.SendConfirmationEmail(body, staff.Email, "Compte supprimé")

	return c.SendString("Salon successfully deleted")
}

func CreateHours(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hours := new(models.Hours)
	if err := c.BodyParser(hours); err != nil {
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
		hours.SalonID = &salonIdToken
	}

	hours.ID, _ = uuid.NewUUID()
	database.DB.Db.Create(&hours)
	return c.JSON(hours)
}
