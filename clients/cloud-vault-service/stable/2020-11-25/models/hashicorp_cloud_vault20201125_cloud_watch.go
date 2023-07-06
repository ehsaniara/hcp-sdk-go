// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125CloudWatch hashicorp cloud vault 20201125 cloud watch
//
// swagger:model hashicorp.cloud.vault_20201125.CloudWatch
type HashicorpCloudVault20201125CloudWatch struct {

	// access key id
	AccessKeyID string `json:"access_key_id,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// secret access key
	SecretAccessKey string `json:"secret_access_key,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 cloud watch
func (m *HashicorpCloudVault20201125CloudWatch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 cloud watch based on context it is used
func (m *HashicorpCloudVault20201125CloudWatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125CloudWatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125CloudWatch) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125CloudWatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
