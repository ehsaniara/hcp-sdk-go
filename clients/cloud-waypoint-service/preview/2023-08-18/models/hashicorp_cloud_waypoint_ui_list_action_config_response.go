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

// HashicorpCloudWaypointUIListActionConfigResponse hashicorp cloud waypoint UI list action config response
//
// swagger:model hashicorp.cloud.waypoint.UI.ListActionConfigResponse
type HashicorpCloudWaypointUIListActionConfigResponse struct {

	// action config bundles
	ActionConfigBundles []*HashicorpCloudWaypointUIListActionConfigBundle `json:"action_config_bundles"`
}

// Validate validates this hashicorp cloud waypoint UI list action config response
func (m *HashicorpCloudWaypointUIListActionConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionConfigBundles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUIListActionConfigResponse) validateActionConfigBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionConfigBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionConfigBundles); i++ {
		if swag.IsZero(m.ActionConfigBundles[i]) { // not required
			continue
		}

		if m.ActionConfigBundles[i] != nil {
			if err := m.ActionConfigBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_config_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_config_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint UI list action config response based on the context it is used
func (m *HashicorpCloudWaypointUIListActionConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionConfigBundles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointUIListActionConfigResponse) contextValidateActionConfigBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActionConfigBundles); i++ {

		if m.ActionConfigBundles[i] != nil {

			if swag.IsZero(m.ActionConfigBundles[i]) { // not required
				return nil
			}

			if err := m.ActionConfigBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_config_bundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_config_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIListActionConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointUIListActionConfigResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointUIListActionConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
