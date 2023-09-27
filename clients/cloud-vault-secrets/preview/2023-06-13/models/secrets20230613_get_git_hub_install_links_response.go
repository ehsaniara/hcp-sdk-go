// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20230613GetGitHubInstallLinksResponse secrets 20230613 get git hub install links response
//
// swagger:model secrets_20230613GetGitHubInstallLinksResponse
type Secrets20230613GetGitHubInstallLinksResponse struct {

	// installation url
	InstallationURL string `json:"installation_url,omitempty"`
}

// Validate validates this secrets 20230613 get git hub install links response
func (m *Secrets20230613GetGitHubInstallLinksResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20230613 get git hub install links response based on context it is used
func (m *Secrets20230613GetGitHubInstallLinksResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613GetGitHubInstallLinksResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613GetGitHubInstallLinksResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20230613GetGitHubInstallLinksResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
