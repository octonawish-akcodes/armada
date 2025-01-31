// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetJobRunErrorHandlerFunc turns a function with the right signature into a get job run error handler
type GetJobRunErrorHandlerFunc func(GetJobRunErrorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetJobRunErrorHandlerFunc) Handle(params GetJobRunErrorParams) middleware.Responder {
	return fn(params)
}

// GetJobRunErrorHandler interface for that can handle valid get job run error params
type GetJobRunErrorHandler interface {
	Handle(GetJobRunErrorParams) middleware.Responder
}

// NewGetJobRunError creates a new http.Handler for the get job run error operation
func NewGetJobRunError(ctx *middleware.Context, handler GetJobRunErrorHandler) *GetJobRunError {
	return &GetJobRunError{Context: ctx, Handler: handler}
}

/* GetJobRunError swagger:route POST /api/v1/jobRunError getJobRunError

GetJobRunError get job run error API

*/
type GetJobRunError struct {
	Context *middleware.Context
	Handler GetJobRunErrorHandler
}

func (o *GetJobRunError) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetJobRunErrorParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetJobRunErrorBody get job run error body
//
// swagger:model GetJobRunErrorBody
type GetJobRunErrorBody struct {

	// run Id
	// Required: true
	RunID string `json:"runId"`
}

// Validate validates this get job run error body
func (o *GetJobRunErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRunID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetJobRunErrorBody) validateRunID(formats strfmt.Registry) error {

	if err := validate.RequiredString("getJobRunErrorRequest"+"."+"runId", "body", o.RunID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get job run error body based on context it is used
func (o *GetJobRunErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetJobRunErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetJobRunErrorBody) UnmarshalBinary(b []byte) error {
	var res GetJobRunErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetJobRunErrorOKBody get job run error o k body
//
// swagger:model GetJobRunErrorOKBody
type GetJobRunErrorOKBody struct {

	// Error for individual job run
	ErrorString string `json:"errorString,omitempty"`
}

// Validate validates this get job run error o k body
func (o *GetJobRunErrorOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get job run error o k body based on context it is used
func (o *GetJobRunErrorOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetJobRunErrorOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetJobRunErrorOKBody) UnmarshalBinary(b []byte) error {
	var res GetJobRunErrorOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
