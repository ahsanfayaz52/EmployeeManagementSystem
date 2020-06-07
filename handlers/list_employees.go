package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	domainErr "github.com/ahsanfayaz52/EmployeeManagementSystem/errors"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/models"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/restapi/operations"
)

// NewListEmployee handles a request for retrieving employees.
func NewListEmployee(rt *runtime.Runtime) operations.ListEmployeeHandler {
	return &listEmployees{rt: rt}
}

type listEmployees struct {
	rt *runtime.Runtime
}

// Handle the list employees request.
func (e *listEmployees) Handle(params operations.ListEmployeeParams) middleware.Responder {
	employee, err := e.rt.Service().ListEmployee()
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetEmployeeByIDNotFound()
		default:
			return operations.NewGetEmployeeByIDInternalServerError()
		}
	}
	var emList []*models.Employee
	i := 0
	for i < len(employee) {
		emList = append(emList, &models.Employee{
			ID:      employee[i].ID,
			Name:    employee[i].Name,
			Address: employee[i].Address,
			Age:     int64(employee[i].Age),
			Salary:  int64(employee[i].Salary),
			Phone:   int64(employee[i].Phone),
		})
		i++
	}
	return operations.NewListEmployeeOK().WithPayload(emList)
}
