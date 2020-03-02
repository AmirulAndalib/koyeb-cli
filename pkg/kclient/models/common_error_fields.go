// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonErrorFields common error fields
// swagger:model commonErrorFields
type CommonErrorFields struct {

	// field
	Field string `json:"field,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this common error fields
func (m *CommonErrorFields) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonErrorFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonErrorFields) UnmarshalBinary(b []byte) error {
	var res CommonErrorFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}