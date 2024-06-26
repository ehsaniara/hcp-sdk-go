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

// Secrets20231128MongoDBAtlasIntegration secrets 20231128 mongo d b atlas integration
//
// swagger:model secrets_20231128MongoDBAtlasIntegration
type Secrets20231128MongoDBAtlasIntegration struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// integration name
	IntegrationName string `json:"integration_name,omitempty"`

	// mongodb api public key
	MongodbAPIPublicKey string `json:"mongodb_api_public_key,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// updated by id
	UpdatedByID string `json:"updated_by_id,omitempty"`
}

// Validate validates this secrets 20231128 mongo d b atlas integration
func (m *Secrets20231128MongoDBAtlasIntegration) Validate(formats strfmt.Registry) error {
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

func (m *Secrets20231128MongoDBAtlasIntegration) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128MongoDBAtlasIntegration) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this secrets 20231128 mongo d b atlas integration based on context it is used
func (m *Secrets20231128MongoDBAtlasIntegration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128MongoDBAtlasIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128MongoDBAtlasIntegration) UnmarshalBinary(b []byte) error {
	var res Secrets20231128MongoDBAtlasIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
