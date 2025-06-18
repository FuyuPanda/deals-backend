package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FuyuPanda/deals-backend/db"
	"github.com/FuyuPanda/deals-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var employees models.Employees
	if err := json.NewDecoder(r.Body).Decode(&employees); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employees.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error encrypting password", http.StatusBadRequest)
	}
	employees.Password = string(hashedPassword)
	if err := db.DB.Create(&employees).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
}
