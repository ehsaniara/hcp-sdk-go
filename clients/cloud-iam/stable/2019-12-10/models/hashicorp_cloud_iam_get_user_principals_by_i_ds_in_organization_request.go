// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest GetUserPrincipalsByIDsInOrganizationRequest is a request for users by ID in a given organization
//
// swagger:model hashicorp.cloud.iam.GetUserPrincipalsByIDsInOrganizationRequest
type HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest struct {

	// ids is a list of user IDs to look up
	Ids []string `json:"ids"`

	// organization_id is the ID of organization for users to look up
	OrganizationID string `json:"organization_id,omitempty"`
}

// Validate validates this hashicorp cloud iam get user principals by i ds in organization request
func (m *HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam get user principals by i ds in organization request based on context it is used
func (m *HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}