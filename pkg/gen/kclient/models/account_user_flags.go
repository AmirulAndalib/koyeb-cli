// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AccountUserFlags - ADMIN: A user is an admin user
//  - TEST: A user is a test user
//  - RESTRICTED: Whether this account as restricted access
//
// swagger:model accountUserFlags
type AccountUserFlags string

const (

	// AccountUserFlagsADMIN captures enum value "ADMIN"
	AccountUserFlagsADMIN AccountUserFlags = "ADMIN"

	// AccountUserFlagsTEST captures enum value "TEST"
	AccountUserFlagsTEST AccountUserFlags = "TEST"

	// AccountUserFlagsRESTRICTED captures enum value "RESTRICTED"
	AccountUserFlagsRESTRICTED AccountUserFlags = "RESTRICTED"
)

// for schema
var accountUserFlagsEnum []interface{}

func init() {
	var res []AccountUserFlags
	if err := json.Unmarshal([]byte(`["ADMIN","TEST","RESTRICTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountUserFlagsEnum = append(accountUserFlagsEnum, v)
	}
}

func (m AccountUserFlags) validateAccountUserFlagsEnum(path, location string, value AccountUserFlags) error {
	if err := validate.Enum(path, location, value, accountUserFlagsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account user flags
func (m AccountUserFlags) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountUserFlagsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}