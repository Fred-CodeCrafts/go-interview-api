package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"go-interview-api/controllers"
	auth "go-interview-api/middleware"
	"go-interview-api/models"
)

func main() {
	// Init DB
	db, err := gorm.Open(sqlite.Open("interview.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	// Auto migrate schema
	if err := db.AutoMigrate(&models.User{}, &models.Question{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Public routes
	r.Post("/users/register", controllers.RegisterUser(db))
	r.Post("/users/login", controllers.LoginUser(db))
	r.Get("/questions", controllers.ListQuestions(db))

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware) // JWT auth middleware
		r.Post("/questions", controllers.CreateQuestion(db))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Listening on port", port)
	log.Fatal(srv.ListenAndServe())
}
