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

// Secrets20231128OpenSecretDynamicInstance secrets 20231128 open secret dynamic instance
//
// swagger:model secrets_20231128OpenSecretDynamicInstance
type Secrets20231128OpenSecretDynamicInstance struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// ttl
	TTL string `json:"ttl,omitempty"`

	// values
	Values map[string]string `json:"values,omitempty"`
}

// Validate validates this secrets 20231128 open secret dynamic instance
func (m *Secrets20231128OpenSecretDynamicInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128OpenSecretDynamicInstance) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20231128OpenSecretDynamicInstance) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this secrets 20231128 open secret dynamic instance based on context it is used
func (m *Secrets20231128OpenSecretDynamicInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128OpenSecretDynamicInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128OpenSecretDynamicInstance) UnmarshalBinary(b []byte) error {
	var res Secrets20231128OpenSecretDynamicInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}