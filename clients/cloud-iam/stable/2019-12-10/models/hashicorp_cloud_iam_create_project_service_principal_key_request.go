// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamCreateProjectServicePrincipalKeyRequest CreateProjectServicePrincipalKeyRequest is the request message used when creating a
// service principal key for a service principal on project level.
//
// swagger:model hashicorp.cloud.iam.CreateProjectServicePrincipalKeyRequest
type HashicorpCloudIamCreateProjectServicePrincipalKeyRequest struct {

	// organization_id is the unique identifier (UUID) of the organization under
	// which the service principal key should be created.
	// It must be the organization of the provided service principal.
	OrganizationID string `json:"organization_id,omitempty"`

	// principal_id is the ID of the service principal for which the new service
	// principal key should be created.
	PrincipalID string `json:"principal_id,omitempty"`

	// project_id is the unique identifier (UUID) of the project under
	// which the service principal key should be created.
	// It must be the project_id of the provided service principal.
	ProjectID string `json:"project_id,omitempty"`
}

// Validate validates this hashicorp cloud iam create project service principal key request
func (m *HashicorpCloudIamCreateProjectServicePrincipalKeyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam create project service principal key request based on context it is used
func (m *HashicorpCloudIamCreateProjectServicePrincipalKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamCreateProjectServicePrincipalKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamCreateProjectServicePrincipalKeyRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamCreateProjectServicePrincipalKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}