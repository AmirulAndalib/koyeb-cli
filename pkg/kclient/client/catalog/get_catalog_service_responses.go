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

// GetCatalogServiceReader is a Reader for the GetCatalogService structure.
type GetCatalogServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCatalogServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCatalogServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCatalogServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCatalogServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCatalogServiceOK creates a GetCatalogServiceOK with default headers values
func NewGetCatalogServiceOK() *GetCatalogServiceOK {
	return &GetCatalogServiceOK{}
}

/*GetCatalogServiceOK handles this case with default header values.

A successful response.
*/
type GetCatalogServiceOK struct {
	Payload *models.StorageGetCatalogServiceReply
}

func (o *GetCatalogServiceOK) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}][%d] getCatalogServiceOK  %+v", 200, o.Payload)
}

func (o *GetCatalogServiceOK) GetPayload() *models.StorageGetCatalogServiceReply {
	return o.Payload
}

func (o *GetCatalogServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageGetCatalogServiceReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogServiceBadRequest creates a GetCatalogServiceBadRequest with default headers values
func NewGetCatalogServiceBadRequest() *GetCatalogServiceBadRequest {
	return &GetCatalogServiceBadRequest{}
}

/*GetCatalogServiceBadRequest handles this case with default header values.

Validation error
*/
type GetCatalogServiceBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *GetCatalogServiceBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}][%d] getCatalogServiceBadRequest  %+v", 400, o.Payload)
}

func (o *GetCatalogServiceBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *GetCatalogServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogServiceForbidden creates a GetCatalogServiceForbidden with default headers values
func NewGetCatalogServiceForbidden() *GetCatalogServiceForbidden {
	return &GetCatalogServiceForbidden{}
}

/*GetCatalogServiceForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type GetCatalogServiceForbidden struct {
	Payload *models.CommonError
}

func (o *GetCatalogServiceForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}][%d] getCatalogServiceForbidden  %+v", 403, o.Payload)
}

func (o *GetCatalogServiceForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *GetCatalogServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogServiceNotFound creates a GetCatalogServiceNotFound with default headers values
func NewGetCatalogServiceNotFound() *GetCatalogServiceNotFound {
	return &GetCatalogServiceNotFound{}
}

/*GetCatalogServiceNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type GetCatalogServiceNotFound struct {
	Payload *models.CommonError
}

func (o *GetCatalogServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}][%d] getCatalogServiceNotFound  %+v", 404, o.Payload)
}

func (o *GetCatalogServiceNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *GetCatalogServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}