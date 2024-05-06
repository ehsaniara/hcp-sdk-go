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

// HashicorpCloudVagrant20220930Index hashicorp cloud vagrant 20220930 index
//
// swagger:model hashicorp.cloud.vagrant_20220930.Index
type HashicorpCloudVagrant20220930Index struct {

	// Architectures supported in this version.
	Architectures []string `json:"architectures"`

	// Avatar is the url to use as the avatar for the index
	Avatar string `json:"avatar,omitempty"`

	// The name segment of the box.
	BoxID string `json:"box_id,omitempty"`

	// The date the box was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The current version of the box
	CurrentVersionID string `json:"current_version_id,omitempty"`

	// The long-form description of the box.
	Description string `json:"description,omitempty"`

	// The number of times this box has been downloaded.
	Downloads string `json:"downloads,omitempty"`

	// Whether or not the box is private.
	IsPrivate bool `json:"is_private,omitempty"`

	// Providers supported in this version.
	Providers []string `json:"providers"`

	// The registry the box belongs to.
	RegistryID string `json:"registry_id,omitempty"`

	// A short-form description of the box. Limited to 100 characters.
	ShortDescription string `json:"short_description,omitempty"`

	// The date that the record was last updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// A vagrant cloud url dynamically generated on each search
	VcURL string `json:"vc_url,omitempty"`

	// A long-form description of the version.
	VersionDescription string `json:"version_description,omitempty"`

	// The date that the version was updated
	// Format: date-time
	VersionUpdatedAt strfmt.DateTime `json:"version_updated_at,omitempty"`
}

// Validate validates this hashicorp cloud vagrant 20220930 index
func (m *HashicorpCloudVagrant20220930Index) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrant20220930Index) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVagrant20220930Index) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVagrant20220930Index) validateVersionUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("version_updated_at", "body", "date-time", m.VersionUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud vagrant 20220930 index based on context it is used
func (m *HashicorpCloudVagrant20220930Index) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930Index) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930Index) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930Index
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
