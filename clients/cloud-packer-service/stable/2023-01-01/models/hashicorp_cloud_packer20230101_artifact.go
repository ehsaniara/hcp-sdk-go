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

// HashicorpCloudPacker20230101Artifact hashicorp cloud packer 20230101 artifact
//
// swagger:model hashicorp.cloud.packer_20230101.Artifact
type HashicorpCloudPacker20230101Artifact struct {

	// Created datetime.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// ID or URL of the remote artifact as given by a build.
	// For example, ami-12345.
	ExternalIdentifier string `json:"external_identifier,omitempty"`

	// Unique identifier (ULID).
	ID string `json:"id,omitempty"`

	// External region as provided by `packer build`. For example, "ap-east-1".
	Region string `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 artifact
func (m *HashicorpCloudPacker20230101Artifact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101Artifact) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud packer 20230101 artifact based on context it is used
func (m *HashicorpCloudPacker20230101Artifact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101Artifact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101Artifact) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101Artifact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}