package routes

import (
	"notes-app/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/register", handlers.RegisterHandler) // sesuaikan nama handler
	app.Post("/api/login", handlers.LoginHandler)
}
