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

// SecretServiceCreateGcpDynamicSecretBody secret service create gcp dynamic secret body
//
// swagger:model SecretServiceCreateGcpDynamicSecretBody
type SecretServiceCreateGcpDynamicSecretBody struct {

	// default ttl
	DefaultTTL string `json:"default_ttl,omitempty"`

	// integration name
	IntegrationName string `json:"integration_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// service account impersonation
	ServiceAccountImpersonation *Secrets20231128ServiceAccountImpersonationRequest `json:"service_account_impersonation,omitempty"`
}

// Validate validates this secret service create gcp dynamic secret body
func (m *SecretServiceCreateGcpDynamicSecretBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceAccountImpersonation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateGcpDynamicSecretBody) validateServiceAccountImpersonation(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceAccountImpersonation) { // not required
		return nil
	}

	if m.ServiceAccountImpersonation != nil {
		if err := m.ServiceAccountImpersonation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_account_impersonation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_account_impersonation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret service create gcp dynamic secret body based on the context it is used
func (m *SecretServiceCreateGcpDynamicSecretBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceAccountImpersonation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateGcpDynamicSecretBody) contextValidateServiceAccountImpersonation(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceAccountImpersonation != nil {

		if swag.IsZero(m.ServiceAccountImpersonation) { // not required
			return nil
		}

		if err := m.ServiceAccountImpersonation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_account_impersonation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_account_impersonation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateGcpDynamicSecretBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateGcpDynamicSecretBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateGcpDynamicSecretBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
