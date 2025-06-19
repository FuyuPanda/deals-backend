package models

import (
	"time"
)

type PayslipResponse struct {
	ID             uint                                  `json:"Id"`
	EmployeeID     int                                   `json:"employee_id"`
	CreatedBy      int                                   `json:"CreatedBy"`
	UpdatedBy      int                                   `json:"UpdatedBy"`
	IsDownloaded   bool                                  `json:"IsDownloaded"`
	Status         string                                `json:"status"`
	TotalAmount    int                                   `json:"TotalAmount"`
	DownloadedOn   time.Time                             `json:"DownloadedOn"`
	Period         time.Time                             `json:"period"`
	CreatedAt      time.Time                             `json:"CreatedAt"`
	UpdatedAt      time.Time                             `json:"UpdatedAt"`
	Overtimes      []PayslipOvertimeDetailsResponse      `json:"overtimes"`
	Reimbursements []PayslipReimbursementDetailsResponse `json:"reimbursements"`
	Attendances    []AttendanceResponse                  `json:"attendances"`
}
