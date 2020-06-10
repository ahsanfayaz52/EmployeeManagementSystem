package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	domainErr "github.com/ahsanfayaz52/EmployeeManagementSystem/errors"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/models"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/restapi/operations"
)

// NewListEmployees handles a request for retrieving employees.
func NewListEmployees(rt *runtime.Runtime) operations.ListEmployeesHandler {
	return &listEmployees{rt: rt}
}

type listEmployees struct {
	rt *runtime.Runtime
}

// Handle the list employees request.
func (e *listEmployees) Handle(params operations.ListEmployeesParams) middleware.Responder {
	employee, err := e.rt.Service().ListEmployees()
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetEmployeeByIDNotFound()
		default:
			return operations.NewGetEmployeeByIDInternalServerError()
		}
	}
	emList := make([]*models.Employee, 0)
	for i := range employee {
		emList = append(emList, &models.Employee{
			ID:      employee[i].ID,
			Name:    employee[i].Name,
			Address: employee[i].Address,
			Age:     int64(employee[i].Age),
			Salary:  int64(employee[i].Salary),
			Phone:   int64(employee[i].Phone),
		})
	}
	return operations.NewListEmployeesOK().WithPayload(emList)
}
