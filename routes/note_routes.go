package routes

import (
	"notes-app/handlers"
	"notes-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(app *fiber.App) {
	// group route dengan middleware JWT
	notes := app.Group("/api/notes", middleware.JWTProtected())

	notes.Post("/", handlers.CreateNoteHandler)
	notes.Get("/", handlers.GetNotesHandler)
	notes.Put("/:id", handlers.UpdateNoteHandler)
	notes.Delete("/:id", handlers.DeleteNoteHandler)
}
