// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudWaypointAddOn AddOn is an add-on definition rendered for a project
//
// swagger:model hashicorp.cloud.waypoint.AddOn
type HashicorpCloudWaypointAddOn struct {

	// The application for which the Add-on was created
	Application *HashicorpCloudWaypointRefApplication `json:"application,omitempty"`

	// The count of other instances of the same add-on definition for the same
	// application.
	Count string `json:"count,omitempty"`

	// The time at which the add-on was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created_by indicates what created the Add-on
	CreatedBy string `json:"created_by,omitempty"`

	// The Add-on definition from which this Add-on was created
	Definition *HashicorpCloudWaypointRefAddOnDefinition `json:"definition,omitempty"`

	// Long-form description of the Add-on
	Description string `json:"description,omitempty"`

	// Unique identifier of the Add-on
	ID string `json:"id,omitempty"`

	// Name of the Add-on
	Name string `json:"name,omitempty"`

	// Rendered README markdown template for Add-on
	// Format: byte
	ReadmeMarkdown strfmt.Base64 `json:"readme_markdown,omitempty"`

	// The status of the Terraform run for the add-on. Client should not set this
	// field.
	Status *HashicorpCloudWaypointTerraformTFRunState `json:"status,omitempty"`

	// Short description of the Add-on
	Summary string `json:"summary,omitempty"`

	// A list of descriptive tags for an Add-on
	Tags []string `json:"tags"`

	// Terraform No Code module used for provisioning the Add-on
	TerraformNocodeModule *HashicorpCloudWaypointTerraformNocodeModule `json:"terraform_nocode_module,omitempty"`
}

// Validate validates this hashicorp cloud waypoint add on
func (m *HashicorpCloudWaypointAddOn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerraformNocodeModule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointAddOn) validateApplication(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointAddOn) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudWaypointAddOn) validateDefinition(formats strfmt.Registry) error {
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

func (m *HashicorpCloudWaypointAddOn) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointAddOn) validateTerraformNocodeModule(formats strfmt.Registry) error {
	if swag.IsZero(m.TerraformNocodeModule) { // not required
		return nil
	}

	if m.TerraformNocodeModule != nil {
		if err := m.TerraformNocodeModule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_nocode_module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_nocode_module")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint add on based on the context it is used
func (m *HashicorpCloudWaypointAddOn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerraformNocodeModule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointAddOn) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointAddOn) contextValidateDefinition(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudWaypointAddOn) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudWaypointAddOn) contextValidateTerraformNocodeModule(ctx context.Context, formats strfmt.Registry) error {

	if m.TerraformNocodeModule != nil {

		if swag.IsZero(m.TerraformNocodeModule) { // not required
			return nil
		}

		if err := m.TerraformNocodeModule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_nocode_module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_nocode_module")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointAddOn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointAddOn) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointAddOn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}