// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointGetAddOnInputVariablesResponse hashicorp cloud waypoint get add on input variables response
//
// swagger:model hashicorp.cloud.waypoint.GetAddOnInputVariablesResponse
type HashicorpCloudWaypointGetAddOnInputVariablesResponse struct {

	// input variables
	InputVariables []*HashicorpCloudWaypointInputVariable `json:"input_variables"`
}

// Validate validates this hashicorp cloud waypoint get add on input variables response
func (m *HashicorpCloudWaypointGetAddOnInputVariablesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointGetAddOnInputVariablesResponse) validateInputVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.InputVariables) { // not required
		return nil
	}

	for i := 0; i < len(m.InputVariables); i++ {
		if swag.IsZero(m.InputVariables[i]) { // not required
			continue
		}

		if m.InputVariables[i] != nil {
			if err := m.InputVariables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input_variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input_variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint get add on input variables response based on the context it is used
func (m *HashicorpCloudWaypointGetAddOnInputVariablesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointGetAddOnInputVariablesResponse) contextValidateInputVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InputVariables); i++ {

		if m.InputVariables[i] != nil {

			if swag.IsZero(m.InputVariables[i]) { // not required
				return nil
			}

			if err := m.InputVariables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input_variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input_variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointGetAddOnInputVariablesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointGetAddOnInputVariablesResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointGetAddOnInputVariablesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
