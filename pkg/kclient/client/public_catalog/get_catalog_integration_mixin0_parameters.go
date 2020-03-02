// Code generated by go-swagger; DO NOT EDIT.

package public_catalog

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

// NewGetCatalogIntegrationMixin0Params creates a new GetCatalogIntegrationMixin0Params object
// with the default values initialized.
func NewGetCatalogIntegrationMixin0Params() *GetCatalogIntegrationMixin0Params {
	var ()
	return &GetCatalogIntegrationMixin0Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogIntegrationMixin0ParamsWithTimeout creates a new GetCatalogIntegrationMixin0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCatalogIntegrationMixin0ParamsWithTimeout(timeout time.Duration) *GetCatalogIntegrationMixin0Params {
	var ()
	return &GetCatalogIntegrationMixin0Params{

		timeout: timeout,
	}
}

// NewGetCatalogIntegrationMixin0ParamsWithContext creates a new GetCatalogIntegrationMixin0Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetCatalogIntegrationMixin0ParamsWithContext(ctx context.Context) *GetCatalogIntegrationMixin0Params {
	var ()
	return &GetCatalogIntegrationMixin0Params{

		Context: ctx,
	}
}

// NewGetCatalogIntegrationMixin0ParamsWithHTTPClient creates a new GetCatalogIntegrationMixin0Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCatalogIntegrationMixin0ParamsWithHTTPClient(client *http.Client) *GetCatalogIntegrationMixin0Params {
	var ()
	return &GetCatalogIntegrationMixin0Params{
		HTTPClient: client,
	}
}

/*GetCatalogIntegrationMixin0Params contains all the parameters to send to the API endpoint
for the get catalog integration mixin0 operation typically these are written to a http.Request
*/
type GetCatalogIntegrationMixin0Params struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get catalog integration mixin0 params
func (o *GetCatalogIntegrationMixin0Params) WithTimeout(timeout time.Duration) *GetCatalogIntegrationMixin0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog integration mixin0 params
func (o *GetCatalogIntegrationMixin0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog integration mixin0 params
func (o *GetCatalogIntegrationMixin0Params) WithContext(ctx context.Context) *GetCatalogIntegrationMixin0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog integration mixin0 params
func (o *GetCatalogIntegrationMixin0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog integration mixin0 params
func (o *GetCatalogIntegrationMixin0Params) WithHTTPClient(client *http.Client) *GetCatalogIntegrationMixin0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog integration mixin0 params
func (o *GetCatalogIntegrationMixin0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get catalog integration mixin0 params
func (o *GetCatalogIntegrationMixin0Params) WithID(id string) *GetCatalogIntegrationMixin0Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get catalog integration mixin0 params
func (o *GetCatalogIntegrationMixin0Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogIntegrationMixin0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}