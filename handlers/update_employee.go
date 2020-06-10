package handlers

import (
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/restapi/operations"
)

// NewUpdateEmployee function is for update employee.
func NewUpdateEmployee(rt *runtime.Runtime) operations.UpdateEmployeeHandler {
	return &updateEmployee{rt: rt}
}

type updateEmployee struct {
	rt *runtime.Runtime
}

// Handle the put employee request.
func (d *updateEmployee) Handle(params operations.UpdateEmployeeParams) middleware.Responder {
	employee := models.Employee{
		ID:      params.ID,
		Name:    params.Employee.Name,
		Address: params.Employee.Address,
		Age:     int(params.Employee.Age),
		Salary:  int(params.Employee.Salary),
		Phone:   int(params.Employee.Phone),
	}
	if err := d.rt.Service().UpdateEmployee(&employee); err != nil {
		return operations.NewUpdateEmployeeInternalServerError()
	}

	return operations.NewUpdateEmployeeOK()
}
