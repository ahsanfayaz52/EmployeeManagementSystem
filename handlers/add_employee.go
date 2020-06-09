package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/restapi/operations"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
)

// NewAddEmployee handles request for saving employee.
func NewAddEmployee(rt *runtime.Runtime) operations.AddEmployeeHandler {
	return &addEmployee{rt: rt}
}

type addEmployee struct {
	rt *runtime.Runtime
}

// Handle the save employee request.
func (e *addEmployee) Handle(params operations.AddEmployeeParams) middleware.Responder {
	employee := models.Employee{
		ID:      params.Employee.ID,
		Name:    params.Employee.Name,
		Address: params.Employee.Address,
		Age:     int(params.Employee.Age),
		Salary:  int(params.Employee.Salary),
		Phone:   int(params.Employee.Phone),
	}

	id, err := e.rt.Service().AddEmployee(&employee)
	if err != nil {
		log().Errorf("failed to add employee: %s", err)
		return operations.NewAddEmployeeInternalServerError()
	}

	params.Employee.ID = id
	return operations.NewAddEmployeeCreated().WithPayload(params.Employee)
}
