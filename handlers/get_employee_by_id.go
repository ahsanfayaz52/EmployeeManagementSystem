package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	domainErr "github.com/ahsanfayaz52/EmployeeManagementSystem/errors"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/models"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/restapi/operations"
)

// NewGetEmployeeByID handles a request for retrieving employee.
func NewGetEmployeeByID(rt *runtime.Runtime) operations.GetEmployeeByIDHandler {
	return &getEmployeeByID{rt: rt}
}

type getEmployeeByID struct {
	rt *runtime.Runtime
}

// Handle the get employee request.
func (e *getEmployeeByID) Handle(params operations.GetEmployeeByIDParams) middleware.Responder {
	employee, err := e.rt.Service().GetEmployeeByID(params.ID)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetEmployeeByIDNotFound()
		default:
			return operations.NewGetEmployeeByIDInternalServerError()
		}
	}

	return operations.NewGetEmployeeByIDOK().WithPayload(&models.Employee{
		ID:      employee.ID,
		Name:    employee.Name,
		Address: employee.Address,
		Age:     int64(employee.Age),
		Salary:  int64(employee.Salary),
		Phone:   int64(employee.Phone),
	})
}
