package main

import (
	"log"
	"net/http"

	"github.com/FuyuPanda/deals-backend/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/FuyuPanda/deals-backend/db"
	"github.com/FuyuPanda/deals-backend/handlers"

	"github.com/FuyuPanda/deals-backend/models"
)

func main() {
	db.InitDB()
	db.InitRedis()
	db.DB.AutoMigrate(&models.Employees{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handlers.Login)
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Post("/", handlers.CreateEmployee)
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
