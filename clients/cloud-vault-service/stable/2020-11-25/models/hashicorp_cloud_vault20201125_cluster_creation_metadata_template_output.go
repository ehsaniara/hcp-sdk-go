// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125ClusterCreationMetadataTemplateOutput TemplateOutput represents the output from a template
//
// swagger:model hashicorp.cloud.vault_20201125.Cluster.CreationMetadata.TemplateOutput
type HashicorpCloudVault20201125ClusterCreationMetadataTemplateOutput struct {

	// accessor is the token accessor of the token used to apply the template
	Accessor string `json:"accessor,omitempty"`

	// fields is the map of json encoded values for each template output
	Fields map[string]strfmt.Base64 `json:"fields,omitempty"`

	// id is the identifier of the template applied
	ID string `json:"id,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 cluster creation metadata template output
func (m *HashicorpCloudVault20201125ClusterCreationMetadataTemplateOutput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125ClusterCreationMetadataTemplateOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125ClusterCreationMetadataTemplateOutput) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125ClusterCreationMetadataTemplateOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
