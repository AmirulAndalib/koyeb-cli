// Code generated by go-swagger; DO NOT EDIT.

package stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStackGetStackActivitiesParams creates a new StackGetStackActivitiesParams object
// with the default values initialized.
func NewStackGetStackActivitiesParams() *StackGetStackActivitiesParams {
	var ()
	return &StackGetStackActivitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStackGetStackActivitiesParamsWithTimeout creates a new StackGetStackActivitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStackGetStackActivitiesParamsWithTimeout(timeout time.Duration) *StackGetStackActivitiesParams {
	var ()
	return &StackGetStackActivitiesParams{

		timeout: timeout,
	}
}

// NewStackGetStackActivitiesParamsWithContext creates a new StackGetStackActivitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewStackGetStackActivitiesParamsWithContext(ctx context.Context) *StackGetStackActivitiesParams {
	var ()
	return &StackGetStackActivitiesParams{

		Context: ctx,
	}
}

// NewStackGetStackActivitiesParamsWithHTTPClient creates a new StackGetStackActivitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStackGetStackActivitiesParamsWithHTTPClient(client *http.Client) *StackGetStackActivitiesParams {
	var ()
	return &StackGetStackActivitiesParams{
		HTTPClient: client,
	}
}

/*StackGetStackActivitiesParams contains all the parameters to send to the API endpoint
for the stack get stack activities operation typically these are written to a http.Request
*/
type StackGetStackActivitiesParams struct {

	/*ID*/
	ID string
	/*Limit*/
	Limit *string
	/*Offset*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stack get stack activities params
func (o *StackGetStackActivitiesParams) WithTimeout(timeout time.Duration) *StackGetStackActivitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack get stack activities params
func (o *StackGetStackActivitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack get stack activities params
func (o *StackGetStackActivitiesParams) WithContext(ctx context.Context) *StackGetStackActivitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack get stack activities params
func (o *StackGetStackActivitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack get stack activities params
func (o *StackGetStackActivitiesParams) WithHTTPClient(client *http.Client) *StackGetStackActivitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack get stack activities params
func (o *StackGetStackActivitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the stack get stack activities params
func (o *StackGetStackActivitiesParams) WithID(id string) *StackGetStackActivitiesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stack get stack activities params
func (o *StackGetStackActivitiesParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the stack get stack activities params
func (o *StackGetStackActivitiesParams) WithLimit(limit *string) *StackGetStackActivitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the stack get stack activities params
func (o *StackGetStackActivitiesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the stack get stack activities params
func (o *StackGetStackActivitiesParams) WithOffset(offset *string) *StackGetStackActivitiesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the stack get stack activities params
func (o *StackGetStackActivitiesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *StackGetStackActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}