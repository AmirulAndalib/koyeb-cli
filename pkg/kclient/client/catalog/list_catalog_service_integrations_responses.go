// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// ListCatalogServiceIntegrationsReader is a Reader for the ListCatalogServiceIntegrations structure.
type ListCatalogServiceIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCatalogServiceIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCatalogServiceIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCatalogServiceIntegrationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCatalogServiceIntegrationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCatalogServiceIntegrationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCatalogServiceIntegrationsOK creates a ListCatalogServiceIntegrationsOK with default headers values
func NewListCatalogServiceIntegrationsOK() *ListCatalogServiceIntegrationsOK {
	return &ListCatalogServiceIntegrationsOK{}
}

/*ListCatalogServiceIntegrationsOK handles this case with default header values.

A successful response.
*/
type ListCatalogServiceIntegrationsOK struct {
	Payload *models.StorageListCatalogIntegrationsReply
}

func (o *ListCatalogServiceIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListCatalogServiceIntegrationsOK) GetPayload() *models.StorageListCatalogIntegrationsReply {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageListCatalogIntegrationsReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServiceIntegrationsBadRequest creates a ListCatalogServiceIntegrationsBadRequest with default headers values
func NewListCatalogServiceIntegrationsBadRequest() *ListCatalogServiceIntegrationsBadRequest {
	return &ListCatalogServiceIntegrationsBadRequest{}
}

/*ListCatalogServiceIntegrationsBadRequest handles this case with default header values.

Validation error
*/
type ListCatalogServiceIntegrationsBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *ListCatalogServiceIntegrationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsBadRequest  %+v", 400, o.Payload)
}

func (o *ListCatalogServiceIntegrationsBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServiceIntegrationsForbidden creates a ListCatalogServiceIntegrationsForbidden with default headers values
func NewListCatalogServiceIntegrationsForbidden() *ListCatalogServiceIntegrationsForbidden {
	return &ListCatalogServiceIntegrationsForbidden{}
}

/*ListCatalogServiceIntegrationsForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type ListCatalogServiceIntegrationsForbidden struct {
	Payload *models.CommonError
}

func (o *ListCatalogServiceIntegrationsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsForbidden  %+v", 403, o.Payload)
}

func (o *ListCatalogServiceIntegrationsForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServiceIntegrationsNotFound creates a ListCatalogServiceIntegrationsNotFound with default headers values
func NewListCatalogServiceIntegrationsNotFound() *ListCatalogServiceIntegrationsNotFound {
	return &ListCatalogServiceIntegrationsNotFound{}
}

/*ListCatalogServiceIntegrationsNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type ListCatalogServiceIntegrationsNotFound struct {
	Payload *models.CommonError
}

func (o *ListCatalogServiceIntegrationsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsNotFound  %+v", 404, o.Payload)
}

func (o *ListCatalogServiceIntegrationsNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}