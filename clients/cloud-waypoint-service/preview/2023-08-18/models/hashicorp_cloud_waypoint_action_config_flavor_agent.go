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

// HashicorpCloudWaypointActionConfigFlavorAgent hashicorp cloud waypoint action config flavor agent
//
// swagger:model hashicorp.cloud.waypoint.ActionConfig.Flavor.Agent
type HashicorpCloudWaypointActionConfigFlavorAgent struct {

	// The operation information. The action_run_id isn't used.
	Op *HashicorpCloudWaypointAgentOperation `json:"op,omitempty"`
}

// Validate validates this hashicorp cloud waypoint action config flavor agent
func (m *HashicorpCloudWaypointActionConfigFlavorAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointActionConfigFlavorAgent) validateOp(formats strfmt.Registry) error {
	if swag.IsZero(m.Op) { // not required
		return nil
	}

	if m.Op != nil {
		if err := m.Op.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint action config flavor agent based on the context it is used
func (m *HashicorpCloudWaypointActionConfigFlavorAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointActionConfigFlavorAgent) contextValidateOp(ctx context.Context, formats strfmt.Registry) error {

	if m.Op != nil {

		if swag.IsZero(m.Op) { // not required
			return nil
		}

		if err := m.Op.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionConfigFlavorAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionConfigFlavorAgent) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointActionConfigFlavorAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
