// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128AzureKvConnectionDetailsResponse secrets 20231128 azure kv connection details response
//
// swagger:model secrets_20231128AzureKvConnectionDetailsResponse
type Secrets20231128AzureKvConnectionDetailsResponse struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// key vault uri
	KeyVaultURI string `json:"key_vault_uri,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`
}

// Validate validates this secrets 20231128 azure kv connection details response
func (m *Secrets20231128AzureKvConnectionDetailsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 azure kv connection details response based on context it is used
func (m *Secrets20231128AzureKvConnectionDetailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128AzureKvConnectionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128AzureKvConnectionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128AzureKvConnectionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
