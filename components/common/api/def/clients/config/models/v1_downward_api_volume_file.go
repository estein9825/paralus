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

// V1DownwardAPIVolumeFile DownwardAPIVolumeFile represents information to create the file containing the pod field
//
// swagger:model v1DownwardAPIVolumeFile
type V1DownwardAPIVolumeFile struct {

	// Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.
	// +optional
	FieldRef *V1ObjectFieldSelector `json:"fieldRef,omitempty"`

	// Optional: mode bits used to set permissions on this file, must be an octal value
	// between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// If not specified, the volume defaultMode will be used.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	Mode int32 `json:"mode,omitempty"`

	// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	Path string `json:"path,omitempty"`

	// Selects a resource of the container: only resources limits and requests
	// (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.
	// +optional
	ResourceFieldRef *V1ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}

// Validate validates this v1 downward API volume file
func (m *V1DownwardAPIVolumeFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceFieldRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DownwardAPIVolumeFile) validateFieldRef(formats strfmt.Registry) error {
	if swag.IsZero(m.FieldRef) { // not required
		return nil
	}

	if m.FieldRef != nil {
		if err := m.FieldRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1DownwardAPIVolumeFile) validateResourceFieldRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceFieldRef) { // not required
		return nil
	}

	if m.ResourceFieldRef != nil {
		if err := m.ResourceFieldRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceFieldRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 downward API volume file based on the context it is used
func (m *V1DownwardAPIVolumeFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFieldRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceFieldRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DownwardAPIVolumeFile) contextValidateFieldRef(ctx context.Context, formats strfmt.Registry) error {

	if m.FieldRef != nil {
		if err := m.FieldRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1DownwardAPIVolumeFile) contextValidateResourceFieldRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceFieldRef != nil {
		if err := m.ResourceFieldRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceFieldRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1DownwardAPIVolumeFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DownwardAPIVolumeFile) UnmarshalBinary(b []byte) error {
	var res V1DownwardAPIVolumeFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
