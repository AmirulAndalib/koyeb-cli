// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/gen/kclient/models"
)

// HooksGithubReader is a Reader for the HooksGithub structure.
type HooksGithubReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HooksGithubReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHooksGithubOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHooksGithubBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHooksGithubForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHooksGithubNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHooksGithubDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHooksGithubOK creates a HooksGithubOK with default headers values
func NewHooksGithubOK() *HooksGithubOK {
	return &HooksGithubOK{}
}

/*HooksGithubOK handles this case with default header values.

A successful response.
*/
type HooksGithubOK struct {
	Payload models.ActivityGithubHookReply
}

func (o *HooksGithubOK) Error() string {
	return fmt.Sprintf("[POST /v1/hooks/github/payload][%d] hooksGithubOK  %+v", 200, o.Payload)
}

func (o *HooksGithubOK) GetPayload() models.ActivityGithubHookReply {
	return o.Payload
}

func (o *HooksGithubOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHooksGithubBadRequest creates a HooksGithubBadRequest with default headers values
func NewHooksGithubBadRequest() *HooksGithubBadRequest {
	return &HooksGithubBadRequest{}
}

/*HooksGithubBadRequest handles this case with default header values.

Validation error
*/
type HooksGithubBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *HooksGithubBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/hooks/github/payload][%d] hooksGithubBadRequest  %+v", 400, o.Payload)
}

func (o *HooksGithubBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *HooksGithubBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHooksGithubForbidden creates a HooksGithubForbidden with default headers values
func NewHooksGithubForbidden() *HooksGithubForbidden {
	return &HooksGithubForbidden{}
}

/*HooksGithubForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type HooksGithubForbidden struct {
	Payload *models.CommonError
}

func (o *HooksGithubForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/hooks/github/payload][%d] hooksGithubForbidden  %+v", 403, o.Payload)
}

func (o *HooksGithubForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *HooksGithubForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHooksGithubNotFound creates a HooksGithubNotFound with default headers values
func NewHooksGithubNotFound() *HooksGithubNotFound {
	return &HooksGithubNotFound{}
}

/*HooksGithubNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type HooksGithubNotFound struct {
	Payload *models.CommonError
}

func (o *HooksGithubNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/hooks/github/payload][%d] hooksGithubNotFound  %+v", 404, o.Payload)
}

func (o *HooksGithubNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *HooksGithubNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHooksGithubDefault creates a HooksGithubDefault with default headers values
func NewHooksGithubDefault(code int) *HooksGithubDefault {
	return &HooksGithubDefault{
		_statusCode: code,
	}
}

/*HooksGithubDefault handles this case with default header values.

An unexpected error response
*/
type HooksGithubDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the hooks github default response
func (o *HooksGithubDefault) Code() int {
	return o._statusCode
}

func (o *HooksGithubDefault) Error() string {
	return fmt.Sprintf("[POST /v1/hooks/github/payload][%d] hooks_Github default  %+v", o._statusCode, o.Payload)
}

func (o *HooksGithubDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *HooksGithubDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}