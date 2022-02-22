// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SentryBootstrapInfra sentry bootstrap infra
//
// swagger:model sentryBootstrapInfra
type SentryBootstrapInfra struct {

	// API Version
	//
	// API Version of the resource
	// Read Only: true
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind
	//
	// Kind of the resource
	// Read Only: true
	Kind string `json:"kind,omitempty"`

	// Metadata
	//
	// Metadata of the resource
	Metadata *V3Metadata `json:"metadata,omitempty"`

	// spec
	Spec *SentryBootstrapInfraSpec `json:"spec,omitempty"`

	// status
	Status SentryBootstrapInfraStatus `json:"status,omitempty"`
}

// Validate validates this sentry bootstrap infra
func (m *SentryBootstrapInfra) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SentryBootstrapInfra) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *SentryBootstrapInfra) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sentry bootstrap infra based on the context it is used
func (m *SentryBootstrapInfra) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SentryBootstrapInfra) contextValidateAPIVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "apiVersion", "body", string(m.APIVersion)); err != nil {
		return err
	}

	return nil
}

func (m *SentryBootstrapInfra) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "kind", "body", string(m.Kind)); err != nil {
		return err
	}

	return nil
}

func (m *SentryBootstrapInfra) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *SentryBootstrapInfra) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {
		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SentryBootstrapInfra) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SentryBootstrapInfra) UnmarshalBinary(b []byte) error {
	var res SentryBootstrapInfra
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
