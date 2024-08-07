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

// SecretServiceCreateAzureKvSyncIntegrationBody secret service create azure kv sync integration body
//
// swagger:model SecretServiceCreateAzureKvSyncIntegrationBody
type SecretServiceCreateAzureKvSyncIntegrationBody struct {

	// azure kv connection details
	AzureKvConnectionDetails *Secrets20231128AzureKvConnectionDetailsRequest `json:"azure_kv_connection_details,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this secret service create azure kv sync integration body
func (m *SecretServiceCreateAzureKvSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureKvConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateAzureKvSyncIntegrationBody) validateAzureKvConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureKvConnectionDetails) { // not required
		return nil
	}

	if m.AzureKvConnectionDetails != nil {
		if err := m.AzureKvConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_kv_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_kv_connection_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret service create azure kv sync integration body based on the context it is used
func (m *SecretServiceCreateAzureKvSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzureKvConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretServiceCreateAzureKvSyncIntegrationBody) contextValidateAzureKvConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureKvConnectionDetails != nil {

		if swag.IsZero(m.AzureKvConnectionDetails) { // not required
			return nil
		}

		if err := m.AzureKvConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_kv_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_kv_connection_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceCreateAzureKvSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceCreateAzureKvSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceCreateAzureKvSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
