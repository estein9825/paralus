// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OwnerReference OwnerReference contains enough information to let you identify an owning
// object. An owning object must be in the same namespace as the dependent, or
// be cluster-scoped, so there is no namespace field.
//
// swagger:model v1OwnerReference
type V1OwnerReference struct {

	// API version of the referent.
	APIVersion string `json:"apiVersion,omitempty"`

	// If true, AND if the owner has the "foregroundDeletion" finalizer, then
	// the owner cannot be deleted from the key-value store until this
	// reference is removed.
	// Defaults to false.
	// To set this field, a user needs "delete" permission of the owner,
	// otherwise 422 (Unprocessable Entity) will be returned.
	// +optional
	BlockOwnerDeletion bool `json:"blockOwnerDeletion,omitempty"`

	// If true, this reference points to the managing controller.
	// +optional
	Controller bool `json:"controller,omitempty"`

	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Name of the referent.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name,omitempty"`

	// UID of the referent.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 owner reference
func (m *V1OwnerReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 owner reference based on context it is used
func (m *V1OwnerReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1OwnerReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OwnerReference) UnmarshalBinary(b []byte) error {
	var res V1OwnerReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
