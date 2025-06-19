package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/FuyuPanda/deals-backend/db"
	"github.com/FuyuPanda/deals-backend/middleware"
	"github.com/FuyuPanda/deals-backend/models"
)

func CreateOvertime(w http.ResponseWriter, r *http.Request) {
	var overtimes models.EmployeeOvertime
	if err := json.NewDecoder(r.Body).Decode(&overtimes); err != nil {
		fmt.Println("JSON decode error:", err)
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	totalHours := time.Since(overtimes.StartTime).Round(time.Hour).Hours()
	if totalHours > 3 {
		totalHours = 3
	}
	overtimes.EmployeeID = int(userID)
	overtimes.CreatedBy = int(userID)
	overtimes.UpdatedBy = int(userID)
	overtimes.TotalHours = int(totalHours)

	if err := db.DB.Create(&overtimes).Error; err != nil {
		http.Error(w, "Failed to submit overtime", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(overtimes)

}

func GetOvertimeHistory(w http.ResponseWriter, r *http.Request) {
	var overtimes []models.EmployeeOvertime
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	if err := db.DB.Where("employee_id = ?", userID).Find(&overtimes).Error; err != nil {
		http.Error(w, "Overtime record not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(overtimes)
}
