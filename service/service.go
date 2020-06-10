package service

import (
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
)

// AddEmployee adds or update employee into database.
func (s *Service) AddEmployee(employee *models.Employee) (string, error) {
	return s.db.AddEmployee(employee)
}

// UpdateEmployee updates employee record in database.
func (s *Service) UpdateEmployee(employee *models.Employee) error {
	return s.db.UpdateEmployee(employee)
}

// GetEmployeeByID retrieves employee from database with id.
func (s *Service) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.db.GetEmployeeByID(id)
}

// DeleteEmployee deletes employee from database.
func (s *Service) DeleteEmployee(id string) error {
	return s.db.DeleteEmployee(id)
}

// ListEmployees retrieve all the employees from database.
func (s *Service) ListEmployees() ([]*models.Employee, error) {
	return s.db.ListEmployees()
}
