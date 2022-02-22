// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EnvVar EnvVar represents an environment variable present in a Container.
//
// swagger:model v1EnvVar
type V1EnvVar struct {

	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name,omitempty"`

	// Variable references $(VAR_NAME) are expanded
	// using the previous defined environment variables in the container and
	// any service environment variables. If a variable cannot be resolved,
	// the reference in the input string will be unchanged. The $(VAR_NAME)
	// syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	// references will never be expanded, regardless of whether the variable
	// exists or not.
	// Defaults to "".
	// +optional
	Value string `json:"value,omitempty"`

	// Source for the environment variable's value. Cannot be used if value is not empty.
	// +optional
	ValueFrom *V1EnvVarSource `json:"valueFrom,omitempty"`
}

// Validate validates this v1 env var
func (m *V1EnvVar) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValueFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EnvVar) validateValueFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.ValueFrom) { // not required
		return nil
	}

	if m.ValueFrom != nil {
		if err := m.ValueFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueFrom")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 env var based on the context it is used
func (m *V1EnvVar) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValueFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EnvVar) contextValidateValueFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.ValueFrom != nil {
		if err := m.ValueFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueFrom")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EnvVar) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EnvVar) UnmarshalBinary(b []byte) error {
	var res V1EnvVar
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
