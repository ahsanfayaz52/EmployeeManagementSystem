package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/restapi/operations"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
)

// NewSaveEmployee handles request for saving employee.
func NewSaveEmployee(rt *runtime.Runtime) operations.SaveEmployeeHandler {
	return &saveEmployee{rt: rt}
}

type saveEmployee struct {
	rt *runtime.Runtime
}

// Handle the save employee request.
func (e *saveEmployee) Handle(params operations.SaveEmployeeParams) middleware.Responder {
	employee := models.Employee{
		ID:      params.Employee.ID,
		Name:    params.Employee.Name,
		Address: params.Employee.Address,
		Age:     int(params.Employee.Age),
		Salary:  int(params.Employee.Salary),
		Phone:   int(params.Employee.Phone),
	}

	id, err := e.rt.Service().SaveEmployee(&employee)
	if err != nil {
		log().Errorf("failed to save/update employee: %s", err)
		return operations.NewSaveEmployeeInternalServerError()
	}

	params.Employee.ID = id
	return operations.NewSaveEmployeeCreated().WithPayload(params.Employee)
}
