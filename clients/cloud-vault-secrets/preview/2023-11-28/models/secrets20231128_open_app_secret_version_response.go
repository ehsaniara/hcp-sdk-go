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

// Secrets20231128OpenAppSecretVersionResponse secrets 20231128 open app secret version response
//
// swagger:model secrets_20231128OpenAppSecretVersionResponse
type Secrets20231128OpenAppSecretVersionResponse struct {

	// rotating version
	RotatingVersion *Secrets20231128OpenSecretRotatingVersion `json:"rotating_version,omitempty"`

	// static version
	StaticVersion *Secrets20231128OpenSecretStaticVersion `json:"static_version,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this secrets 20231128 open app secret version response
func (m *Secrets20231128OpenAppSecretVersionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRotatingVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128OpenAppSecretVersionResponse) validateRotatingVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.RotatingVersion) { // not required
		return nil
	}

	if m.RotatingVersion != nil {
		if err := m.RotatingVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rotating_version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rotating_version")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128OpenAppSecretVersionResponse) validateStaticVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticVersion) { // not required
		return nil
	}

	if m.StaticVersion != nil {
		if err := m.StaticVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static_version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static_version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20231128 open app secret version response based on the context it is used
func (m *Secrets20231128OpenAppSecretVersionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRotatingVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128OpenAppSecretVersionResponse) contextValidateRotatingVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.RotatingVersion != nil {

		if swag.IsZero(m.RotatingVersion) { // not required
			return nil
		}

		if err := m.RotatingVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rotating_version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rotating_version")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20231128OpenAppSecretVersionResponse) contextValidateStaticVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.StaticVersion != nil {

		if swag.IsZero(m.StaticVersion) { // not required
			return nil
		}

		if err := m.StaticVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static_version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static_version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128OpenAppSecretVersionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128OpenAppSecretVersionResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128OpenAppSecretVersionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
