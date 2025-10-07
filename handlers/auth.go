package handlers

import (
	"encoding/json"
	"net/http"
	"notes-app/config"
	"notes-app/db"
	"notes-app/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Register user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Gagal hash password", http.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec("INSERT INTO users(username,password) VALUES($1,$2)", user.Username, string(hashedPassword))
	if err != nil {
		http.Error(w, "Gagal menyimpan user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User berhasil dibuat!"})
}

// Login user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Request invalid", http.StatusBadRequest)
		return
	}

	var storedUser models.User
	var hashedPassword string

	err = db.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", creds.Username).
		Scan(&storedUser.ID, &storedUser.Username, &hashedPassword)
	if err != nil {
		http.Error(w, "User tidak ditemukan", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Password salah", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		http.Error(w, "Gagal generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
