package main

import (
    "log"
    "notes-app/config"
    "notes-app/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    // Load .env (opsional, khusus lokal)
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️ .env tidak ditemukan, pastikan environment sudah di-set")
    }

    // Connect ke PostgreSQL
    config.ConnectDB()

    // Inisialisasi Fiber app
    app := fiber.New()

    // Setup routes modular
    routes.SetupRoutes(app)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Println("🚀 Server berjalan di port:", port)
    log.Fatal(app.Listen(":" + port))
}
