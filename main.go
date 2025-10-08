package main

import (
    "log"
    "net/http"
    "notes-app/config"
    "notes-app/db"
    "notes-app/handlers"
    "notes-app/middleware"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    // Load .env (khusus lokal)
    godotenv.Load()

    config.InitConfig()
    db.Connect()

    r := mux.NewRouter()

    // Auth routes
    r.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")

    // Protected routes (pakai JWT)
    notes := r.PathPrefix("/api/notes").Subrouter()
    notes.Use(middleware.JWTMiddleware)
    notes.HandleFunc("", handlers.CreateNoteHandler).Methods("POST")
    notes.HandleFunc("", handlers.GetNotesHandler).Methods("GET")
    notes.HandleFunc("/{id}", handlers.UpdateNoteHandler).Methods("PUT")
    notes.HandleFunc("/{id}", handlers.DeleteNoteHandler).Methods("DELETE")

    log.Println("🚀 Server berjalan di port:", os.Getenv("PORT"))
    http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
