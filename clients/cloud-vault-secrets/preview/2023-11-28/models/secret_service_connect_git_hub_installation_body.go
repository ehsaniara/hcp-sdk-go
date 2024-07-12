// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretServiceConnectGitHubInstallationBody secret service connect git hub installation body
//
// swagger:model SecretServiceConnectGitHubInstallationBody
type SecretServiceConnectGitHubInstallationBody struct {

	// installation id
	InstallationID string `json:"installation_id,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this secret service connect git hub installation body
func (m *SecretServiceConnectGitHubInstallationBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret service connect git hub installation body based on context it is used
func (m *SecretServiceConnectGitHubInstallationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretServiceConnectGitHubInstallationBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretServiceConnectGitHubInstallationBody) UnmarshalBinary(b []byte) error {
	var res SecretServiceConnectGitHubInstallationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}