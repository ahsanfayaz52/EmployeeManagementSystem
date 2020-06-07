// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListEmployeesParams creates a new ListEmployeesParams object
// with the default values initialized.
func NewListEmployeesParams() *ListEmployeesParams {

	return &ListEmployeesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListEmployeesParamsWithTimeout creates a new ListEmployeesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEmployeesParamsWithTimeout(timeout time.Duration) *ListEmployeesParams {

	return &ListEmployeesParams{

		timeout: timeout,
	}
}

// NewListEmployeesParamsWithContext creates a new ListEmployeesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListEmployeesParamsWithContext(ctx context.Context) *ListEmployeesParams {

	return &ListEmployeesParams{

		Context: ctx,
	}
}

// NewListEmployeesParamsWithHTTPClient creates a new ListEmployeesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEmployeesParamsWithHTTPClient(client *http.Client) *ListEmployeesParams {

	return &ListEmployeesParams{
		HTTPClient: client,
	}
}

/*ListEmployeesParams contains all the parameters to send to the API endpoint
for the list employees operation typically these are written to a http.Request
*/
type ListEmployeesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list employees params
func (o *ListEmployeesParams) WithTimeout(timeout time.Duration) *ListEmployeesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list employees params
func (o *ListEmployeesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list employees params
func (o *ListEmployeesParams) WithContext(ctx context.Context) *ListEmployeesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list employees params
func (o *ListEmployeesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list employees params
func (o *ListEmployeesParams) WithHTTPClient(client *http.Client) *ListEmployeesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list employees params
func (o *ListEmployeesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListEmployeesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}