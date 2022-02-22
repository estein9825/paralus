// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RPCGetKubectlClusterSettingsResponse rpc get kubectl cluster settings response
//
// swagger:model rpcGetKubectlClusterSettingsResponse
type RPCGetKubectlClusterSettingsResponse struct {

	// disable c l i kubectl
	DisableCLIKubectl bool `json:"disableCLIKubectl,omitempty"`

	// disable web kubectl
	DisableWebKubectl bool `json:"disableWebKubectl,omitempty"`
}

// Validate validates this rpc get kubectl cluster settings response
func (m *RPCGetKubectlClusterSettingsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rpc get kubectl cluster settings response based on context it is used
func (m *RPCGetKubectlClusterSettingsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RPCGetKubectlClusterSettingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RPCGetKubectlClusterSettingsResponse) UnmarshalBinary(b []byte) error {
	var res RPCGetKubectlClusterSettingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
