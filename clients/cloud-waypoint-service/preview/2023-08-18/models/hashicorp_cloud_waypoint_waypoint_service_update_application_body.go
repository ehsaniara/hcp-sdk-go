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

// HashicorpCloudWaypointWaypointServiceUpdateApplicationBody hashicorp cloud waypoint waypoint service update application body
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.UpdateApplicationBody
type HashicorpCloudWaypointWaypointServiceUpdateApplicationBody struct {

	// Update with any *new* Action Configs. Old action configs will not be removed.
	ActionCfgRefs []*HashicorpCloudWaypointActionCfgRef `json:"action_cfg_refs"`

	// application
	Application *HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication `json:"application,omitempty"`

	// Updated application name
	Name string `json:"name,omitempty"`

	// Global references the entire server. This is used in some APIs
	// as a way to read/write values that are server-global.
	Namespace interface{} `json:"namespace,omitempty"`

	// readme_markdown is markdown formatted instructions on how to operate the application.
	// This may be populated from a application template.
	// Format: byte
	ReadmeMarkdown strfmt.Base64 `json:"readme_markdown,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service update application body
func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionCfgRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBody) validateActionCfgRefs(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionCfgRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionCfgRefs); i++ {
		if swag.IsZero(m.ActionCfgRefs[i]) { // not required
			continue
		}

		if m.ActionCfgRefs[i] != nil {
			if err := m.ActionCfgRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_cfg_refs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_cfg_refs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBody) validateApplication(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint waypoint service update application body based on the context it is used
func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionCfgRefs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBody) contextValidateActionCfgRefs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActionCfgRefs); i++ {

		if m.ActionCfgRefs[i] != nil {

			if swag.IsZero(m.ActionCfgRefs[i]) { // not required
				return nil
			}

			if err := m.ActionCfgRefs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("action_cfg_refs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("action_cfg_refs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBody) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceUpdateApplicationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication Application Ref Without ID
//
// # Reference to an existing Application
//
// swagger:model HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication
type HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service update application body application
func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint waypoint service update application body application based on context it is used
func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceUpdateApplicationBodyApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
