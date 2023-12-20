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

// IamGroup iam group
//
// swagger:model iamGroup
type IamGroup struct {

	// created_at is when the group was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// description is a description of the group.
	Description string `json:"description,omitempty"`

	// display_name is the user-specified display name of the group.
	DisplayName string `json:"display_name,omitempty"`

	// member_count is the count of members in the group.
	// We use Int32Value to make sure we return nil instead of 0 if we don't have a member count
	MemberCount int32 `json:"member_count,omitempty"`

	// resource_id is the principal ID of the group.
	ResourceID string `json:"resource_id,omitempty"`

	// resource_name is the name of the group.
	ResourceName string `json:"resource_name,omitempty"`

	// scim_synchronized denotes the group was synchronized from an upstream IdP using SCIM.
	ScimSynchronized bool `json:"scim_synchronized,omitempty"`

	// updated_at is when the group was last updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this iam group
func (m *IamGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IamGroup) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IamGroup) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iam group based on context it is used
func (m *IamGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IamGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IamGroup) UnmarshalBinary(b []byte) error {
	var res IamGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}