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
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("PGHOST"),
        os.Getenv("PGPORT"),
        os.Getenv("PGUSER"),
        os.Getenv("PGPASSWORD"),
        os.Getenv("PGDATABASE"),
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
