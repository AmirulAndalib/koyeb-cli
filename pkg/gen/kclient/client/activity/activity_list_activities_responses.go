// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/gen/kclient/models"
)

// ActivityListActivitiesReader is a Reader for the ActivityListActivities structure.
type ActivityListActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivityListActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivityListActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewActivityListActivitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActivityListActivitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActivityListActivitiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewActivityListActivitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActivityListActivitiesOK creates a ActivityListActivitiesOK with default headers values
func NewActivityListActivitiesOK() *ActivityListActivitiesOK {
	return &ActivityListActivitiesOK{}
}

/*ActivityListActivitiesOK handles this case with default header values.

A successful response.
*/
type ActivityListActivitiesOK struct {
	Payload *models.ActivityActivityList
}

func (o *ActivityListActivitiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/activities][%d] activityListActivitiesOK  %+v", 200, o.Payload)
}

func (o *ActivityListActivitiesOK) GetPayload() *models.ActivityActivityList {
	return o.Payload
}

func (o *ActivityListActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivityActivityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivityListActivitiesBadRequest creates a ActivityListActivitiesBadRequest with default headers values
func NewActivityListActivitiesBadRequest() *ActivityListActivitiesBadRequest {
	return &ActivityListActivitiesBadRequest{}
}

/*ActivityListActivitiesBadRequest handles this case with default header values.

Validation error
*/
type ActivityListActivitiesBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *ActivityListActivitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/activities][%d] activityListActivitiesBadRequest  %+v", 400, o.Payload)
}

func (o *ActivityListActivitiesBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *ActivityListActivitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivityListActivitiesForbidden creates a ActivityListActivitiesForbidden with default headers values
func NewActivityListActivitiesForbidden() *ActivityListActivitiesForbidden {
	return &ActivityListActivitiesForbidden{}
}

/*ActivityListActivitiesForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type ActivityListActivitiesForbidden struct {
	Payload *models.CommonError
}

func (o *ActivityListActivitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/activities][%d] activityListActivitiesForbidden  %+v", 403, o.Payload)
}

func (o *ActivityListActivitiesForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ActivityListActivitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivityListActivitiesNotFound creates a ActivityListActivitiesNotFound with default headers values
func NewActivityListActivitiesNotFound() *ActivityListActivitiesNotFound {
	return &ActivityListActivitiesNotFound{}
}

/*ActivityListActivitiesNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type ActivityListActivitiesNotFound struct {
	Payload *models.CommonError
}

func (o *ActivityListActivitiesNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/activities][%d] activityListActivitiesNotFound  %+v", 404, o.Payload)
}

func (o *ActivityListActivitiesNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ActivityListActivitiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivityListActivitiesDefault creates a ActivityListActivitiesDefault with default headers values
func NewActivityListActivitiesDefault(code int) *ActivityListActivitiesDefault {
	return &ActivityListActivitiesDefault{
		_statusCode: code,
	}
}

/*ActivityListActivitiesDefault handles this case with default header values.

An unexpected error response
*/
type ActivityListActivitiesDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the activity list activities default response
func (o *ActivityListActivitiesDefault) Code() int {
	return o._statusCode
}

func (o *ActivityListActivitiesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/activities][%d] activity_ListActivities default  %+v", o._statusCode, o.Payload)
}

func (o *ActivityListActivitiesDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *ActivityListActivitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}