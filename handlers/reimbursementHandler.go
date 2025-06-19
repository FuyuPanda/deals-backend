package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FuyuPanda/deals-backend/db"
	"github.com/FuyuPanda/deals-backend/middleware"
	"github.com/FuyuPanda/deals-backend/models"
)

func CreateReimbursement(w http.ResponseWriter, r *http.Request) {
	var reimbursements models.EmployeeReimbursement

	if err := json.NewDecoder(r.Body).Decode(&reimbursements); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	reimbursements.EmployeeID = int(userID)
	reimbursements.CreatedBy = int(userID)
	reimbursements.UpdatedBy = int(userID)

	if err := db.DB.Create(&reimbursements).Error; err != nil {
		http.Error(w, "Failed to submit reimbursement", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reimbursements)

}

func GetReimbursementHistory(w http.ResponseWriter, r *http.Request) {
	var reimbursements []models.EmployeeReimbursement
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	if err := db.DB.Where("employee_id = ?", userID).Find(&reimbursements).Error; err != nil {
		http.Error(w, "Reimbursement record not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(reimbursements)
}
