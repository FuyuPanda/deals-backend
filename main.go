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
	db.DB.AutoMigrate(&models.Attendance{})
	db.DB.AutoMigrate(&models.EmployeeReimbursement{})
	db.DB.AutoMigrate(&models.EmployeeOvertime{})
	db.DB.AutoMigrate(&models.Payslip{})
	db.DB.AutoMigrate(&models.PayslipOvertimeDetails{})
	db.DB.AutoMigrate(&models.PayslipReimbursementDetails{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handlers.Login)
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Post("/", handlers.CreateEmployee)
	})

	r.Route("/attendance", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Get("/", handlers.GetAttendanceHistory)
		r.Post("/clockin", handlers.CheckInAttendance)
		r.Put("/clockout", handlers.CheckOutAttendance)
	})

	r.Route("/overtime", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Get("/", handlers.GetOvertimeHistory)
		r.Post("/", handlers.CreateOvertime)
	})

	r.Route("/reimbursement", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Get("/", handlers.GetReimbursementHistory)
		r.Post("/", handlers.CreateReimbursement)
	})

	r.Route("/payroll", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Post("/", handlers.CreatePayroll)
		r.Get("/", handlers.GetPayroll)
		r.Get("/employee/summary", handlers.GetSummaryPayrollEmployee)
		r.Get("/total", handlers.GetTotalPayrollEmployee)
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
