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

// HashicorpCloudWaypointWaypointServiceCreateAddOnBody CreateAddOnRequest is the request used to add an Add-on to a project
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.CreateAddOnBody
type HashicorpCloudWaypointWaypointServiceCreateAddOnBody struct {

	// The application for which the Add-on is to be created
	Application *HashicorpCloudWaypointRefApplication `json:"application,omitempty"`

	// The Add-on definition from which this Add-on is to be created
	Definition *HashicorpCloudWaypointRefAddOnDefinition `json:"definition,omitempty"`

	// name is the name of the Add-on
	Name string `json:"name,omitempty"`

	// Global references the entire server. This is used in some APIs
	// as a way to read/write values that are server-global.
	Namespace interface{} `json:"namespace,omitempty"`

	// variables is the series of input variables which have been set by the
	// application developer for the new add-on being created. This may be empty.
	Variables []*HashicorpCloudWaypointInputVariable `json:"variables"`
}

// Validate validates this hashicorp cloud waypoint waypoint service create add on body
func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) validateApplication(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) validateDefinition(formats strfmt.Registry) error {
	if swag.IsZero(m.Definition) { // not required
		return nil
	}

	if m.Definition != nil {
		if err := m.Definition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("definition")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) validateVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint waypoint service create add on body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) contextValidateDefinition(ctx context.Context, formats strfmt.Registry) error {

	if m.Definition != nil {

		if swag.IsZero(m.Definition) { // not required
			return nil
		}

		if err := m.Definition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("definition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("definition")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variables); i++ {

		if m.Variables[i] != nil {

			if swag.IsZero(m.Variables[i]) { // not required
				return nil
			}

			if err := m.Variables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceCreateAddOnBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceCreateAddOnBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}