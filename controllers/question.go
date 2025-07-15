// controllers/question.go
package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"

	auth "go-interview-api/middleware"
	"go-interview-api/models"
	"go-interview-api/utils"
)

type CreateQuestionRequest struct {
	Topic    string `json:"topic"`
	Question string `json:"question"`
}

func CreateQuestion(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := auth.GetUserID(r.Context())
		if !ok {
			utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			return
		}

		var req CreateQuestionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid input"})
			return
		}

		q := models.Question{
			Topic:     req.Topic,
			Question:  req.Question,
			AuthorID:  userID,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&q).Error; err != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create question"})
			return
		}

		utils.RespondJSON(w, http.StatusCreated, q)
	}
}

func ListQuestions(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var questions []models.Question
		if err := db.Find(&questions).Error; err != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not fetch questions"})
			return
		}
		utils.RespondJSON(w, http.StatusOK, questions)
	}
}
