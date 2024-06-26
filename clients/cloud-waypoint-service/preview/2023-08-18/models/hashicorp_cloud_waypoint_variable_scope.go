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

// HashicorpCloudWaypointVariableScope hashicorp cloud waypoint variable scope
//
// swagger:model hashicorp.cloud.waypoint.Variable.Scope
type HashicorpCloudWaypointVariableScope struct {

	// The variable is only scoped to a specific action
	Action *HashicorpCloudWaypointActionCfgRef `json:"action,omitempty"`

	// Not supported. The variable is only scoped to a specific application
	Application *HashicorpCloudWaypointRefApplication `json:"application,omitempty"`

	// Not supported. The variable is only scoped to a specific template
	ApplicationTemplate *HashicorpCloudWaypointRefApplicationTemplate `json:"application_template,omitempty"`

	// A variable that can be accessed globally per-namespace
	Global HashicorpCloudWaypointVariableScopeGlobal `json:"global,omitempty"`
}

// Validate validates this hashicorp cloud waypoint variable scope
func (m *HashicorpCloudWaypointVariableScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointVariableScope) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointVariableScope) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointVariableScope) validateApplicationTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationTemplate) { // not required
		return nil
	}

	if m.ApplicationTemplate != nil {
		if err := m.ApplicationTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application_template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint variable scope based on the context it is used
func (m *HashicorpCloudWaypointVariableScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplicationTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointVariableScope) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {

		if swag.IsZero(m.Action) { // not required
			return nil
		}

		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointVariableScope) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {

		if swag.IsZero(m.Application) { // not required
			return nil
		}

		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointVariableScope) contextValidateApplicationTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationTemplate != nil {

		if swag.IsZero(m.ApplicationTemplate) { // not required
			return nil
		}

		if err := m.ApplicationTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application_template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointVariableScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointVariableScope) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointVariableScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
