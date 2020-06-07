package service

import (
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
)

// SaveEmployee adds or update employee into database.
func (s *Service) SaveEmployee(employee *models.Employee) (string, error) {
	result, err := s.db.SaveEmployee(employee)
	if err != nil {
		return "", err
	}
	return result, nil
}

// GetEmployeeByID retrieves employee from database with id.
func (s *Service) GetEmployeeByID(id string) (*models.Employee, error) {
	employee, err := s.db.GetEmployeeByID(id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// DeleteEmployee deletes employee from database.
func (s *Service) DeleteEmployee(id string) error {
	err := s.db.DeleteEmployee(id)
	if err != nil {
		return err
	}
	return nil
}

// ListEmployees retrieve all the employees from database.
func (s *Service) ListEmployees() ([]*models.Employee, error) {
	employees, err := s.db.ListEmployees()
	if err != nil {
		return nil, err
	}
	return employees, nil
}
