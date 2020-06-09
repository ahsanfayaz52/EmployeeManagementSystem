package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/restapi/operations"
)

// Handler replaces swagger handler.
type Handler *operations.EmployeeManagementSystemAPI

// NewHandler overrides swagger api handlers.
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewEmployeeManagementSystemAPI(spec)

	// employee handlers.
	handler.AddEmployeeHandler = NewAddEmployee(rt)
	handler.GetEmployeeByIDHandler = NewGetEmployeeByID(rt)
	handler.DeleteEmployeeHandler = NewDeleteEmployee(rt)
	handler.ListEmployeesHandler = NewListEmployees(rt)
	handler.UpdateEmployeeHandler = NewUpdateEmployee(rt)

	return handler
}
