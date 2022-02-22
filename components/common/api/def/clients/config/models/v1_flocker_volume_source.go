// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1FlockerVolumeSource Represents a Flocker volume mounted by the Flocker agent.
// One and only one of datasetName and datasetUUID should be set.
// Flocker volumes do not support ownership management or SELinux relabeling.
//
// swagger:model v1FlockerVolumeSource
type V1FlockerVolumeSource struct {

	// Name of the dataset stored as metadata -> name on the dataset for Flocker
	// should be considered as deprecated
	// +optional
	DatasetName string `json:"datasetName,omitempty"`

	// UUID of the dataset. This is unique identifier of a Flocker dataset
	// +optional
	DatasetUUID string `json:"datasetUUID,omitempty"`
}

// Validate validates this v1 flocker volume source
func (m *V1FlockerVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 flocker volume source based on context it is used
func (m *V1FlockerVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FlockerVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FlockerVolumeSource) UnmarshalBinary(b []byte) error {
	var res V1FlockerVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
