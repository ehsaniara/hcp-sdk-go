// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128GcpFederatedWorkloadIdentityRequest secrets 20231128 gcp federated workload identity request
//
// swagger:model secrets_20231128GcpFederatedWorkloadIdentityRequest
type Secrets20231128GcpFederatedWorkloadIdentityRequest struct {

	// audience
	Audience string `json:"audience,omitempty"`

	// service account email
	ServiceAccountEmail string `json:"service_account_email,omitempty"`
}

// Validate validates this secrets 20231128 gcp federated workload identity request
func (m *Secrets20231128GcpFederatedWorkloadIdentityRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 gcp federated workload identity request based on context it is used
func (m *Secrets20231128GcpFederatedWorkloadIdentityRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128GcpFederatedWorkloadIdentityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128GcpFederatedWorkloadIdentityRequest) UnmarshalBinary(b []byte) error {
	var res Secrets20231128GcpFederatedWorkloadIdentityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
