package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	domainErr "github.com/ahsanfayaz52/EmployeeManagementSystem/errors"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/restapi/operations"
)

// NewDeleteEmployee function will delete the employee.
func NewDeleteEmployee(rt *runtime.Runtime) operations.DeleteEmployeeHandler {
	return &deleteEmployee{rt: rt}
}

type deleteEmployee struct {
	rt *runtime.Runtime
}

// Handle the delete entry request.
func (e *deleteEmployee) Handle(params operations.DeleteEmployeeParams) middleware.Responder {
	if err := e.rt.Service().DeleteEmployee(params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewDeleteEmployeeNotFound()
		default:
			return operations.NewDeleteEmployeeInternalServerError()
		}
	}
	return operations.NewDeleteEmployeeNoContent()
}
