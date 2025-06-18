package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/FuyuPanda/deals-backend/db"
	"github.com/FuyuPanda/deals-backend/middleware"
	"github.com/FuyuPanda/deals-backend/models"
	"github.com/FuyuPanda/deals-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employees
	username := r.Context().Value(middleware.UsernameKey).(string)

	if username != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Access",
		})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&employees); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	for _, employee := range employees {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error encrypting password", http.StatusBadRequest)
			return
		}

		employee.Password = string(hashedPassword)

		intID, ok := r.Context().Value(middleware.UserIDKey).(float64)
		if !ok {
			http.Error(w, "Invalid user ID in context", http.StatusInternalServerError)
			return
		}

		employee.CreatedBy = int(intID)
		employee.UpdatedBy = int(intID)

		if err := db.DB.Create(&employees).Error; err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		employee.Password = ""

	}

	json.NewEncoder(w).Encode(employees)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var employees models.Employees
	if err := json.NewDecoder(r.Body).Decode(&employees); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	var data models.Employees
	if err := db.DB.Where("Username = ?", employees.Username).First(&data).Error; err != nil {
		http.Error(w, "Invalid Username", http.StatusUnauthorized)
		return
	}

	if !CheckPassword(employees.Password, data.Password) {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(data.Id, data.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	redisKey := fmt.Sprintf("user:%d:token", data.Id)
	expiryTime := 24 * time.Hour
	err = db.RedisClient.Set(db.Ctx, redisKey, token, 24*time.Hour).Err()
	if err != nil {
		http.Error(w, "Could not store token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message":      "OK",
		"token":        token,
		"expiry_time:": expiryTime.String(),
	})

}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
