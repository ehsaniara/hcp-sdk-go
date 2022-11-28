// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudGlobalNetworkManager20220215TLSInfo hashicorp cloud global network manager 20220215 TLS info
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.TLSInfo
type HashicorpCloudGlobalNetworkManager20220215TLSInfo struct {

	// cert expiry
	// Format: date-time
	CertExpiry strfmt.DateTime `json:"cert_expiry,omitempty"`

	// cert name
	CertName string `json:"cert_name,omitempty"`

	// cert serial
	CertSerial string `json:"cert_serial,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// verify incoming
	VerifyIncoming bool `json:"verify_incoming,omitempty"`

	// verify outgoing
	VerifyOutgoing bool `json:"verify_outgoing,omitempty"`

	// verify server hostname
	VerifyServerHostname bool `json:"verify_server_hostname,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 TLS info
func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertExpiry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) validateCertExpiry(formats strfmt.Registry) error {
	if swag.IsZero(m.CertExpiry) { // not required
		return nil
	}

	if err := validate.FormatOf("cert_expiry", "body", "date-time", m.CertExpiry.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud global network manager 20220215 TLS info based on context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215TLSInfo) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215TLSInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}