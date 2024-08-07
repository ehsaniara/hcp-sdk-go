// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128HcpTerraformMetadata secrets 20231128 hcp terraform metadata
//
// swagger:model secrets_20231128HcpTerraformMetadata
type Secrets20231128HcpTerraformMetadata struct {

	// org
	Org string `json:"org,omitempty"`

	// team
	Team string `json:"team,omitempty"`
}

// Validate validates this secrets 20231128 hcp terraform metadata
func (m *Secrets20231128HcpTerraformMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 hcp terraform metadata based on context it is used
func (m *Secrets20231128HcpTerraformMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128HcpTerraformMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128HcpTerraformMetadata) UnmarshalBinary(b []byte) error {
	var res Secrets20231128HcpTerraformMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
