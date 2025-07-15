// controllers/auth.go
package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	auth "go-interview-api/middleware"
	"go-interview-api/models"
	"go-interview-api/utils"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid input"})
			return
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not hash password"})
			return
		}

		user := models.User{
			Username:  req.Username,
			Email:     req.Email,
			Password:  string(hashed),
			CreatedAt: time.Now(),
		}

		if err := db.Create(&user).Error; err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "username or email already exists"})
			return
		}

		utils.RespondJSON(w, http.StatusCreated, map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		})
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid input"})
			return
		}

		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
			return
		}

		token, err := auth.GenerateJWT(user.ID, user.Username)
		if err != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not generate token"})
			return
		}

		utils.RespondJSON(w, http.StatusOK, map[string]string{"token": token})
	}
}
