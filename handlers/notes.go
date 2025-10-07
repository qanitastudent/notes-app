package handlers

import (
	"encoding/json"
	"net/http"
	"notes-app/db"
	"notes-app/middleware"
	"notes-app/models"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("INSERT INTO notes(user_id, title, content) VALUES($1,$2,$3)", userID, note.Title, note.Content)
	if err != nil {
		http.Error(w, "Gagal membuat note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Note berhasil dibuat"})
}

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	rows, err := db.DB.Query("SELECT id, title, content, created_at, updated_at FROM notes WHERE user_id=$1", userID)
	if err != nil {
		http.Error(w, "Gagal mengambil notes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	notes := []models.Note{}
	for rows.Next() {
		var n models.Note
		rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt)
		notes = append(notes, n)
	}

	json.NewEncoder(w).Encode(notes)
}

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	noteID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)

	_, err := db.DB.Exec("UPDATE notes SET title=$1, content=$2, updated_at=NOW() WHERE id=$3 AND user_id=$4",
		note.Title, note.Content, noteID, userID)
	if err != nil {
		http.Error(w, "Gagal update note", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Note berhasil diupdate"})
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	noteID, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := db.DB.Exec("DELETE FROM notes WHERE id=$1 AND user_id=$2", noteID, userID)
	if err != nil {
		http.Error(w, "Gagal hapus note", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Note berhasil dihapus"})
}
