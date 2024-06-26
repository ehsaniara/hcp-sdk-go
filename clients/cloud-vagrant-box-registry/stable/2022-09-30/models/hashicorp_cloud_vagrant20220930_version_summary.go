// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVagrant20220930VersionSummary hashicorp cloud vagrant 20220930 version summary
//
// swagger:model hashicorp.cloud.vagrant_20220930.VersionSummary
type HashicorpCloudVagrant20220930VersionSummary struct {

	// List of providers with versions available for this box
	ProviderNames []string `json:"provider_names"`

	// Total number of providers available on this version
	ProvidersCount string `json:"providers_count,omitempty"`
}

// Validate validates this hashicorp cloud vagrant 20220930 version summary
func (m *HashicorpCloudVagrant20220930VersionSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vagrant 20220930 version summary based on context it is used
func (m *HashicorpCloudVagrant20220930VersionSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930VersionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930VersionSummary) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930VersionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
