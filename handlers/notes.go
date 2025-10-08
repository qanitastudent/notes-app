package handlers

import (
	"notes-app/db"
	"notes-app/middleware"
	"notes-app/models"

	"github.com/gofiber/fiber/v2"
)

// Buat note baru
func CreateNoteHandler(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request invalid"})
	}

	_, err := db.DB.Exec("INSERT INTO notes(user_id, title, content) VALUES($1,$2,$3)", userID, note.Title, note.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat note"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Note berhasil dibuat"})
}

// Ambil semua note milik user
func GetNotesHandler(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	rows, err := db.DB.Query("SELECT id, title, content, created_at, updated_at FROM notes WHERE user_id=$1", userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil notes"})
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		notes = append(notes, n)
	}

	return c.JSON(notes)
}

// Update note
func UpdateNoteHandler(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	noteID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID note invalid"})
	}

	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request invalid"})
	}

	_, err = db.DB.Exec("UPDATE notes SET title=$1, content=$2, updated_at=NOW() WHERE id=$3 AND user_id=$4",
		note.Title, note.Content, noteID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal update note"})
	}

	return c.JSON(fiber.Map{"message": "Note berhasil diupdate"})
}

// Hapus note
func DeleteNoteHandler(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	noteID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID note invalid"})
	}

	_, err = db.DB.Exec("DELETE FROM notes WHERE id=$1 AND user_id=$2", noteID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal hapus note"})
	}

	return c.JSON(fiber.Map{"message": "Note berhasil dihapus"})
}
