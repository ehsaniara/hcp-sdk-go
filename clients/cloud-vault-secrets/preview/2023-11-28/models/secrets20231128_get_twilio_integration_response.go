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

// Secrets20231128GetTwilioIntegrationResponse secrets 20231128 get twilio integration response
//
// swagger:model secrets_20231128GetTwilioIntegrationResponse
type Secrets20231128GetTwilioIntegrationResponse struct {

	// integration
	Integration *Secrets20231128TwilioIntegration `json:"integration,omitempty"`
}

// Validate validates this secrets 20231128 get twilio integration response
func (m *Secrets20231128GetTwilioIntegrationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128GetTwilioIntegrationResponse) validateIntegration(formats strfmt.Registry) error {
	if swag.IsZero(m.Integration) { // not required
		return nil
	}

	if m.Integration != nil {
		if err := m.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20231128 get twilio integration response based on the context it is used
func (m *Secrets20231128GetTwilioIntegrationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128GetTwilioIntegrationResponse) contextValidateIntegration(ctx context.Context, formats strfmt.Registry) error {

	if m.Integration != nil {

		if swag.IsZero(m.Integration) { // not required
			return nil
		}

		if err := m.Integration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128GetTwilioIntegrationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128GetTwilioIntegrationResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128GetTwilioIntegrationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
