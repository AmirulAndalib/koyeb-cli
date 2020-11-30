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
	"github.com/go-openapi/strfmt"
)

// NewStackListStacksParams creates a new StackListStacksParams object
// with the default values initialized.
func NewStackListStacksParams() *StackListStacksParams {
	var (
		repositoryTypeDefault = string("GITHUB")
	)
	return &StackListStacksParams{
		RepositoryType: &repositoryTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewStackListStacksParamsWithTimeout creates a new StackListStacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStackListStacksParamsWithTimeout(timeout time.Duration) *StackListStacksParams {
	var (
		repositoryTypeDefault = string("GITHUB")
	)
	return &StackListStacksParams{
		RepositoryType: &repositoryTypeDefault,

		timeout: timeout,
	}
}

// NewStackListStacksParamsWithContext creates a new StackListStacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewStackListStacksParamsWithContext(ctx context.Context) *StackListStacksParams {
	var (
		repositoryTypeDefault = string("GITHUB")
	)
	return &StackListStacksParams{
		RepositoryType: &repositoryTypeDefault,

		Context: ctx,
	}
}

// NewStackListStacksParamsWithHTTPClient creates a new StackListStacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStackListStacksParamsWithHTTPClient(client *http.Client) *StackListStacksParams {
	var (
		repositoryTypeDefault = string("GITHUB")
	)
	return &StackListStacksParams{
		RepositoryType: &repositoryTypeDefault,
		HTTPClient:     client,
	}
}

/*StackListStacksParams contains all the parameters to send to the API endpoint
for the stack list stacks operation typically these are written to a http.Request
*/
type StackListStacksParams struct {

	/*Limit*/
	Limit *string
	/*Name*/
	Name *string
	/*Offset*/
	Offset *string
	/*RepositoryBranch
	  The branch to track changes on.

	*/
	RepositoryBranch *string
	/*RepositoryName
	  The url to find the repo (.e.g: koyeb/gateway).

	*/
	RepositoryName *string
	/*RepositoryType
	  Where the repo lives.

	*/
	RepositoryType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stack list stacks params
func (o *StackListStacksParams) WithTimeout(timeout time.Duration) *StackListStacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stack list stacks params
func (o *StackListStacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stack list stacks params
func (o *StackListStacksParams) WithContext(ctx context.Context) *StackListStacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stack list stacks params
func (o *StackListStacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stack list stacks params
func (o *StackListStacksParams) WithHTTPClient(client *http.Client) *StackListStacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stack list stacks params
func (o *StackListStacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the stack list stacks params
func (o *StackListStacksParams) WithLimit(limit *string) *StackListStacksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the stack list stacks params
func (o *StackListStacksParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the stack list stacks params
func (o *StackListStacksParams) WithName(name *string) *StackListStacksParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the stack list stacks params
func (o *StackListStacksParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the stack list stacks params
func (o *StackListStacksParams) WithOffset(offset *string) *StackListStacksParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the stack list stacks params
func (o *StackListStacksParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithRepositoryBranch adds the repositoryBranch to the stack list stacks params
func (o *StackListStacksParams) WithRepositoryBranch(repositoryBranch *string) *StackListStacksParams {
	o.SetRepositoryBranch(repositoryBranch)
	return o
}

// SetRepositoryBranch adds the repositoryBranch to the stack list stacks params
func (o *StackListStacksParams) SetRepositoryBranch(repositoryBranch *string) {
	o.RepositoryBranch = repositoryBranch
}

// WithRepositoryName adds the repositoryName to the stack list stacks params
func (o *StackListStacksParams) WithRepositoryName(repositoryName *string) *StackListStacksParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the stack list stacks params
func (o *StackListStacksParams) SetRepositoryName(repositoryName *string) {
	o.RepositoryName = repositoryName
}

// WithRepositoryType adds the repositoryType to the stack list stacks params
func (o *StackListStacksParams) WithRepositoryType(repositoryType *string) *StackListStacksParams {
	o.SetRepositoryType(repositoryType)
	return o
}

// SetRepositoryType adds the repositoryType to the stack list stacks params
func (o *StackListStacksParams) SetRepositoryType(repositoryType *string) {
	o.RepositoryType = repositoryType
}

// WriteToRequest writes these params to a swagger request
func (o *StackListStacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.RepositoryBranch != nil {

		// query param repository.branch
		var qrRepositoryBranch string
		if o.RepositoryBranch != nil {
			qrRepositoryBranch = *o.RepositoryBranch
		}
		qRepositoryBranch := qrRepositoryBranch
		if qRepositoryBranch != "" {
			if err := r.SetQueryParam("repository.branch", qRepositoryBranch); err != nil {
				return err
			}
		}

	}

	if o.RepositoryName != nil {

		// query param repository.name
		var qrRepositoryName string
		if o.RepositoryName != nil {
			qrRepositoryName = *o.RepositoryName
		}
		qRepositoryName := qrRepositoryName
		if qRepositoryName != "" {
			if err := r.SetQueryParam("repository.name", qRepositoryName); err != nil {
				return err
			}
		}

	}

	if o.RepositoryType != nil {

		// query param repository.type
		var qrRepositoryType string
		if o.RepositoryType != nil {
			qrRepositoryType = *o.RepositoryType
		}
		qRepositoryType := qrRepositoryType
		if qRepositoryType != "" {
			if err := r.SetQueryParam("repository.type", qRepositoryType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}