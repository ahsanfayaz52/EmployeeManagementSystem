// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ahsanfayaz52/EmployeeManagementSystem/gen/models"
)

// SaveEmployeeCreatedCode is the HTTP code returned for type SaveEmployeeCreated
const SaveEmployeeCreatedCode int = 201

/*SaveEmployeeCreated employee added

swagger:response saveEmployeeCreated
*/
type SaveEmployeeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Employee `json:"body,omitempty"`
}

// NewSaveEmployeeCreated creates SaveEmployeeCreated with default headers values
func NewSaveEmployeeCreated() *SaveEmployeeCreated {

	return &SaveEmployeeCreated{}
}

// WithPayload adds the payload to the save employee created response
func (o *SaveEmployeeCreated) WithPayload(payload *models.Employee) *SaveEmployeeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save employee created response
func (o *SaveEmployeeCreated) SetPayload(payload *models.Employee) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveEmployeeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveEmployeeInternalServerErrorCode is the HTTP code returned for type SaveEmployeeInternalServerError
const SaveEmployeeInternalServerErrorCode int = 500

/*SaveEmployeeInternalServerError internal server error

swagger:response saveEmployeeInternalServerError
*/
type SaveEmployeeInternalServerError struct {
}

// NewSaveEmployeeInternalServerError creates SaveEmployeeInternalServerError with default headers values
func NewSaveEmployeeInternalServerError() *SaveEmployeeInternalServerError {

	return &SaveEmployeeInternalServerError{}
}

// WriteResponse to the client
func (o *SaveEmployeeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}