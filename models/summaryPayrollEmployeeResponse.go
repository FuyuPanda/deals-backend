package models

type SummaryPayrollEmployeeResponse struct {
	EmployeeName string `json:"employee_name"`
	Amount       int    `json:"amount"`
}

type TotalPayrollResponse struct {
	Amount int `json:"amount"`
}
