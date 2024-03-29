// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128AwsSmConnectionDetailsRequest secrets 20231128 aws sm connection details request
//
// swagger:model secrets_20231128AwsSmConnectionDetailsRequest
type Secrets20231128AwsSmConnectionDetailsRequest struct {

	// access key id
	AccessKeyID string `json:"access_key_id,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// role arn
	RoleArn string `json:"role_arn,omitempty"`

	// secret access key
	SecretAccessKey string `json:"secret_access_key,omitempty"`
}

// Validate validates this secrets 20231128 aws sm connection details request
func (m *Secrets20231128AwsSmConnectionDetailsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 aws sm connection details request based on context it is used
func (m *Secrets20231128AwsSmConnectionDetailsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128AwsSmConnectionDetailsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128AwsSmConnectionDetailsRequest) UnmarshalBinary(b []byte) error {
	var res Secrets20231128AwsSmConnectionDetailsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}