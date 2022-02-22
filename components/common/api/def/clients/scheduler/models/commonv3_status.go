// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Commonv3Status Status
//
// status of a resource
//
// swagger:model commonv3Status
type Commonv3Status struct {

	// Condition Status
	//
	// status of the condition
	// Read Only: true
	// Enum: [StatusNotSet StatusSubmitted StatusOK StatusFailed]
	ConditionStatus *V3ConditionStatus `json:"conditionStatus,omitempty"`

	// Condition Type
	//
	// type of the status condition
	// Read Only: true
	ConditionType string `json:"conditionType,omitempty"`

	// Last Updated
	//
	// when the condition status is last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// Reason
	//
	// reason of the last condition status
	// Read Only: true
	Reason string `json:"reason,omitempty"`
}

// Validate validates this commonv3 status
func (m *Commonv3Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var commonv3StatusTypeConditionStatusPropEnum []interface{}

func init() {
	var res []V3ConditionStatus
	if err := json.Unmarshal([]byte(`["StatusNotSet","StatusSubmitted","StatusOK","StatusFailed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonv3StatusTypeConditionStatusPropEnum = append(commonv3StatusTypeConditionStatusPropEnum, v)
	}
}

const (

	// Commonv3StatusConditionStatusStatusNotSet captures enum value "StatusNotSet"
	Commonv3StatusConditionStatusStatusNotSet V3ConditionStatus = "StatusNotSet"

	// Commonv3StatusConditionStatusStatusSubmitted captures enum value "StatusSubmitted"
	Commonv3StatusConditionStatusStatusSubmitted V3ConditionStatus = "StatusSubmitted"

	// Commonv3StatusConditionStatusStatusOK captures enum value "StatusOK"
	Commonv3StatusConditionStatusStatusOK V3ConditionStatus = "StatusOK"

	// Commonv3StatusConditionStatusStatusFailed captures enum value "StatusFailed"
	Commonv3StatusConditionStatusStatusFailed V3ConditionStatus = "StatusFailed"
)

// prop value enum
func (m *Commonv3Status) validateConditionStatusEnum(path, location string, value V3ConditionStatus) error {
	if err := validate.EnumCase(path, location, value, commonv3StatusTypeConditionStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Commonv3Status) validateConditionStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ConditionStatus) { // not required
		return nil
	}

	if m.ConditionStatus != nil {
		if err := m.ConditionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditionStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditionStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Commonv3Status) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this commonv3 status based on the context it is used
func (m *Commonv3Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditionStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Commonv3Status) contextValidateConditionStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ConditionStatus != nil {
		if err := m.ConditionStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditionStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditionStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Commonv3Status) contextValidateConditionType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "conditionType", "body", string(m.ConditionType)); err != nil {
		return err
	}

	return nil
}

func (m *Commonv3Status) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *Commonv3Status) contextValidateReason(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reason", "body", string(m.Reason)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Commonv3Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Commonv3Status) UnmarshalBinary(b []byte) error {
	var res Commonv3Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
