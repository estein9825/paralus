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

// ConfigDriftReconcillationAction config drift reconcillation action
//
// swagger:model configDriftReconcillationAction
type ConfigDriftReconcillationAction string

func NewConfigDriftReconcillationAction(value ConfigDriftReconcillationAction) *ConfigDriftReconcillationAction {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigDriftReconcillationAction.
func (m ConfigDriftReconcillationAction) Pointer() *ConfigDriftReconcillationAction {
	return &m
}

const (

	// ConfigDriftReconcillationActionDriftReconcillationActionNotSet captures enum value "DriftReconcillationActionNotSet"
	ConfigDriftReconcillationActionDriftReconcillationActionNotSet ConfigDriftReconcillationAction = "DriftReconcillationActionNotSet"

	// ConfigDriftReconcillationActionDriftReconcillationActionNotify captures enum value "DriftReconcillationActionNotify"
	ConfigDriftReconcillationActionDriftReconcillationActionNotify ConfigDriftReconcillationAction = "DriftReconcillationActionNotify"

	// ConfigDriftReconcillationActionDriftReconcillationActionDeny captures enum value "DriftReconcillationActionDeny"
	ConfigDriftReconcillationActionDriftReconcillationActionDeny ConfigDriftReconcillationAction = "DriftReconcillationActionDeny"
)

// for schema
var configDriftReconcillationActionEnum []interface{}

func init() {
	var res []ConfigDriftReconcillationAction
	if err := json.Unmarshal([]byte(`["DriftReconcillationActionNotSet","DriftReconcillationActionNotify","DriftReconcillationActionDeny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configDriftReconcillationActionEnum = append(configDriftReconcillationActionEnum, v)
	}
}

func (m ConfigDriftReconcillationAction) validateConfigDriftReconcillationActionEnum(path, location string, value ConfigDriftReconcillationAction) error {
	if err := validate.EnumCase(path, location, value, configDriftReconcillationActionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config drift reconcillation action
func (m ConfigDriftReconcillationAction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigDriftReconcillationActionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config drift reconcillation action based on context it is used
func (m ConfigDriftReconcillationAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
