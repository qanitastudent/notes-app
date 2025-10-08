package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )


    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Gagal konek ke database:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Tidak bisa ping ke database:", err)
    }

    log.Println("✅ Koneksi database berhasil!")
}
