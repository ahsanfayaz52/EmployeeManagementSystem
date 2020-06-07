// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/models"
)

// ListEmployeeOKCode is the HTTP code returned for type ListEmployeeOK
const ListEmployeeOKCode int = 200

/*ListEmployeeOK employee response

swagger:response listEmployeeOK
*/
type ListEmployeeOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Employee `json:"body,omitempty"`
}

// NewListEmployeeOK creates ListEmployeeOK with default headers values
func NewListEmployeeOK() *ListEmployeeOK {

	return &ListEmployeeOK{}
}

// WithPayload adds the payload to the list employee o k response
func (o *ListEmployeeOK) WithPayload(payload []*models.Employee) *ListEmployeeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list employee o k response
func (o *ListEmployeeOK) SetPayload(payload []*models.Employee) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEmployeeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Employee, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListEmployeeInternalServerErrorCode is the HTTP code returned for type ListEmployeeInternalServerError
const ListEmployeeInternalServerErrorCode int = 500

/*ListEmployeeInternalServerError internal server error

swagger:response listEmployeeInternalServerError
*/
type ListEmployeeInternalServerError struct {
}

// NewListEmployeeInternalServerError creates ListEmployeeInternalServerError with default headers values
func NewListEmployeeInternalServerError() *ListEmployeeInternalServerError {

	return &ListEmployeeInternalServerError{}
}

// WriteResponse to the client
func (o *ListEmployeeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
