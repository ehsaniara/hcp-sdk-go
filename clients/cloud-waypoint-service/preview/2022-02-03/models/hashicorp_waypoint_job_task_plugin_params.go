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

// HashicorpWaypointJobTaskPluginParams TaskPluginParams contains the information about a specific task plugin
// that is used by both StartTask and StopTask
//
// swagger:model hashicorp.waypoint.Job.TaskPluginParams
type HashicorpWaypointJobTaskPluginParams struct {

	// The configuration information for the task. This is HCL that is
	// decoded to figure out the task plugin and then provide that
	// task plugin with configuration
	// Format: byte
	HclConfig strfmt.Base64 `json:"hcl_config,omitempty"`

	// hcl format
	HclFormat *HashicorpWaypointHclFormat `json:"hcl_format,omitempty"`

	// The plugin type to invoke for the task plugin.
	PluginType string `json:"plugin_type,omitempty"`
}

// Validate validates this hashicorp waypoint job task plugin params
func (m *HashicorpWaypointJobTaskPluginParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHclFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobTaskPluginParams) validateHclFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.HclFormat) { // not required
		return nil
	}

	if m.HclFormat != nil {
		if err := m.HclFormat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hcl_format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hcl_format")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint job task plugin params based on the context it is used
func (m *HashicorpWaypointJobTaskPluginParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHclFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobTaskPluginParams) contextValidateHclFormat(ctx context.Context, formats strfmt.Registry) error {

	if m.HclFormat != nil {

		if swag.IsZero(m.HclFormat) { // not required
			return nil
		}

		if err := m.HclFormat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hcl_format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hcl_format")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointJobTaskPluginParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointJobTaskPluginParams) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointJobTaskPluginParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}