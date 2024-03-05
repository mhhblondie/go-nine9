package main

import (
	"github.com/Go-nine9/go-nine9/controllers"
	"github.com/Go-nine9/go-nine9/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	// PUBLIC ROUTES //
	app.Get("/", controllers.Home)
	app.Get("/salons", controllers.GetSalons)

	app.Get("/salons/:id", controllers.GetSalonById)

	// AUTH ROUTES //
	app.Post("/auth/register", controllers.CreateNewUser)
	app.Post("/auth/login", controllers.LoginUser)

	// PRIVATE ROUTES /
	api := app.Group("/api", middleware.AuthRequired())
	api.Get("/me", controllers.GetMe)
	api.Patch("/me", controllers.UpdateMe)
	api.Patch("/me/password", controllers.UpdateMePassword)
	api.Post("/reservations", controllers.CreateReservation)
	api.Post("/salons", controllers.CreateSalon)

	// GESTION ROUTES (ADMIN AND MANAGER) //
	management := api.Group("/management", middleware.RoleMiddleware("manager", "admin"))

	// USERS
	management.Get("/users", controllers.GetAllUsers)
	management.Post("/users", controllers.CreateUser)
	management.Get("/users/:id", controllers.GetUserById)
	management.Patch("/users/:id", controllers.UpdateUser)
	management.Delete("/users/:id", controllers.DeleteUser)

	// SALONS
	management.Get("/salons", controllers.GetMySalons)
	management.Post("/salons/staff", controllers.AddStaff)
	management.Patch("/salons/:id", controllers.UpdateSalon)
	management.Delete("/salons/:id", controllers.DeleteSalon)
	management.Delete("/salons/staff/:staffID", controllers.DeleteStaff)
	management.Post("/hours", controllers.CreateHours)

	// SLOTS
	management.Get("/slots", controllers.GetAllSlots)
	management.Post("/slots", controllers.CreateSlot)
	management.Patch("/slots/:id", controllers.UpdateSlot)
	management.Delete("/slots/:id", controllers.DeleteSlot)

	// RESERVATIONS
	management.Post("/reservations", controllers.CreateReservation)
	management.Delete("/reservations/:id", controllers.DeleteReservation)

	//SERVICE
	management.Post("/service", controllers.CreateService)
	management.Patch("/service/:id", controllers.UpdateService)
	management.Delete("/service/:id", controllers.DeleteService)

	//Prestations
	management.Post("/prestations", controllers.CreatePrestation)
	management.Patch("/prestations/:id", controllers.UpdatePrestation)
	management.Delete("/prestations/:id", controllers.DeletePrestation)

	// ADMIN ONLY //

	// SALONS
	admin := api.Group("/admin", middleware.RoleMiddleware("admin"))
	admin.Get("/salons", controllers.GetSalons)
	admin.Post("/salons", controllers.CreateSalonAdmin)
}
