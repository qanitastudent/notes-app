package main

import (
    "log"
    "net/http"
    "notes-app/config"
    "notes-app/db"
    "notes-app/handlers"
    "github.com/gorilla/mux"
    "os"
)

func main() {
    config.InitConfig()
    db.Connect()

    r := mux.NewRouter()
    r.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")
    r.HandleFunc("/api/notes", handlers.CreateNoteHandler).Methods("POST")
    r.HandleFunc("/api/notes", handlers.GetNotesHandler).Methods("GET")
    r.HandleFunc("/api/notes/{id}", handlers.UpdateNoteHandler).Methods("PUT")
    r.HandleFunc("/api/notes/{id}", handlers.DeleteNoteHandler).Methods("DELETE")

    log.Println("🚀 Server berjalan di port:", os.Getenv("PORT"))
    http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
