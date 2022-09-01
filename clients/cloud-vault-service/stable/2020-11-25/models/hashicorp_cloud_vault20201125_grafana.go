// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125Grafana hashicorp cloud vault 20201125 grafana
//
// swagger:model hashicorp.cloud.vault_20201125.Grafana
type HashicorpCloudVault20201125Grafana struct {

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 grafana
func (m *HashicorpCloudVault20201125Grafana) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 grafana based on context it is used
func (m *HashicorpCloudVault20201125Grafana) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125Grafana) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125Grafana) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125Grafana
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
