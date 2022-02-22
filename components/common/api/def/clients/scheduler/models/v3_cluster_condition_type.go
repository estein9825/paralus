// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V3ClusterConditionType v3 cluster condition type
//
// swagger:model v3ClusterConditionType
type V3ClusterConditionType string

func NewV3ClusterConditionType(value V3ClusterConditionType) *V3ClusterConditionType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V3ClusterConditionType.
func (m V3ClusterConditionType) Pointer() *V3ClusterConditionType {
	return &m
}

const (

	// V3ClusterConditionTypeClusterRegister captures enum value "ClusterRegister"
	V3ClusterConditionTypeClusterRegister V3ClusterConditionType = "ClusterRegister"

	// V3ClusterConditionTypeClusterApprove captures enum value "ClusterApprove"
	V3ClusterConditionTypeClusterApprove V3ClusterConditionType = "ClusterApprove"

	// V3ClusterConditionTypeClusterCheckIn captures enum value "ClusterCheckIn"
	V3ClusterConditionTypeClusterCheckIn V3ClusterConditionType = "ClusterCheckIn"

	// V3ClusterConditionTypeClusterNodeSync captures enum value "ClusterNodeSync"
	V3ClusterConditionTypeClusterNodeSync V3ClusterConditionType = "ClusterNodeSync"

	// V3ClusterConditionTypeClusterBlueprintSync captures enum value "ClusterBlueprintSync"
	V3ClusterConditionTypeClusterBlueprintSync V3ClusterConditionType = "ClusterBlueprintSync"

	// V3ClusterConditionTypeClusterNamespaceSync captures enum value "ClusterNamespaceSync"
	V3ClusterConditionTypeClusterNamespaceSync V3ClusterConditionType = "ClusterNamespaceSync"

	// V3ClusterConditionTypeClusterReady captures enum value "ClusterReady"
	V3ClusterConditionTypeClusterReady V3ClusterConditionType = "ClusterReady"

	// V3ClusterConditionTypeClusterAuxiliaryTaskSync captures enum value "ClusterAuxiliaryTaskSync"
	V3ClusterConditionTypeClusterAuxiliaryTaskSync V3ClusterConditionType = "ClusterAuxiliaryTaskSync"

	// V3ClusterConditionTypeClusterBootstrapAgent captures enum value "ClusterBootstrapAgent"
	V3ClusterConditionTypeClusterBootstrapAgent V3ClusterConditionType = "ClusterBootstrapAgent"

	// V3ClusterConditionTypeClusterDelete captures enum value "ClusterDelete"
	V3ClusterConditionTypeClusterDelete V3ClusterConditionType = "ClusterDelete"
)

// for schema
var v3ClusterConditionTypeEnum []interface{}

func init() {
	var res []V3ClusterConditionType
	if err := json.Unmarshal([]byte(`["ClusterRegister","ClusterApprove","ClusterCheckIn","ClusterNodeSync","ClusterBlueprintSync","ClusterNamespaceSync","ClusterReady","ClusterAuxiliaryTaskSync","ClusterBootstrapAgent","ClusterDelete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3ClusterConditionTypeEnum = append(v3ClusterConditionTypeEnum, v)
	}
}

func (m V3ClusterConditionType) validateV3ClusterConditionTypeEnum(path, location string, value V3ClusterConditionType) error {
	if err := validate.EnumCase(path, location, value, v3ClusterConditionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 cluster condition type
func (m V3ClusterConditionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3ClusterConditionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v3 cluster condition type based on context it is used
func (m V3ClusterConditionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
