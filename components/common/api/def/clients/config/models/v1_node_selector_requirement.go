// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NodeSelectorRequirement A node selector requirement is a selector that contains values, a key, and an operator
// that relates the key and values.
//
// swagger:model v1NodeSelectorRequirement
type V1NodeSelectorRequirement struct {

	// The label key that the selector applies to.
	Key string `json:"key,omitempty"`

	// Represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Operator string `json:"operator,omitempty"`

	// An array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. If the operator is Gt or Lt, the values
	// array must have a single element, which will be interpreted as an integer.
	// This array is replaced during a strategic merge patch.
	// +optional
	Values []string `json:"values"`
}

// Validate validates this v1 node selector requirement
func (m *V1NodeSelectorRequirement) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 node selector requirement based on context it is used
func (m *V1NodeSelectorRequirement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NodeSelectorRequirement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NodeSelectorRequirement) UnmarshalBinary(b []byte) error {
	var res V1NodeSelectorRequirement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
