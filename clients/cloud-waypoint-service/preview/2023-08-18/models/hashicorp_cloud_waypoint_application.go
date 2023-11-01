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

// HashicorpCloudWaypointApplication hashicorp cloud waypoint application
//
// swagger:model hashicorp.cloud.waypoint.Application
type HashicorpCloudWaypointApplication struct {

	// application template
	ApplicationTemplate *HashicorpCloudWaypointRefApplicationTemplate `json:"application_template,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// readme_markdown is markdown formatted instructions on how to
	// operate the application. This may be populated from a application template.
	// Format: byte
	ReadmeMarkdown strfmt.Base64 `json:"readme_markdown,omitempty"`
}

// Validate validates this hashicorp cloud waypoint application
func (m *HashicorpCloudWaypointApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointApplication) validateApplicationTemplate(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud waypoint application based on the context it is used
func (m *HashicorpCloudWaypointApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointApplication) contextValidateApplicationTemplate(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudWaypointApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointApplication) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}