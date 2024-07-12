// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128GcpSmConnectionDetailsRequest secrets 20231128 gcp sm connection details request
//
// swagger:model secrets_20231128GcpSmConnectionDetailsRequest
type Secrets20231128GcpSmConnectionDetailsRequest struct {

	// credentials
	Credentials string `json:"credentials,omitempty"`

	// locations
	Locations []string `json:"locations"`
}

// Validate validates this secrets 20231128 gcp sm connection details request
func (m *Secrets20231128GcpSmConnectionDetailsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 gcp sm connection details request based on context it is used
func (m *Secrets20231128GcpSmConnectionDetailsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128GcpSmConnectionDetailsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128GcpSmConnectionDetailsRequest) UnmarshalBinary(b []byte) error {
	var res Secrets20231128GcpSmConnectionDetailsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
