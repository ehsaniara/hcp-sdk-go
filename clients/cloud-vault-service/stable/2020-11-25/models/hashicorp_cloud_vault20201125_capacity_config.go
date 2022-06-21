// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125CapacityConfig hashicorp cloud vault 20201125 capacity config
//
// swagger:model hashicorp.cloud.vault_20201125.CapacityConfig
type HashicorpCloudVault20201125CapacityConfig struct {

	// num_servers is the number of nodes this Vault cluster should have.
	NumServers int32 `json:"num_servers,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 capacity config
func (m *HashicorpCloudVault20201125CapacityConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125CapacityConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125CapacityConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125CapacityConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}