// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/models"
)

// NewSaveEmployeeParams creates a new SaveEmployeeParams object
// no default values defined in spec.
func NewSaveEmployeeParams() SaveEmployeeParams {

	return SaveEmployeeParams{}
}

// SaveEmployeeParams contains all the bound params for the save employee operation
// typically these are obtained from a http.Request
//
// swagger:parameters saveEmployee
type SaveEmployeeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*contains employee information
	  Required: true
	  In: body
	*/
	Employee *models.Employee
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSaveEmployeeParams() beforehand.
func (o *SaveEmployeeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Employee
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("employee", "body"))
			} else {
				res = append(res, errors.NewParseError("employee", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Employee = &body
			}
		}
	} else {
		res = append(res, errors.Required("employee", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
