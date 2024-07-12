// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128TwilioStaticCredentialsResponse secrets 20231128 twilio static credentials response
//
// swagger:model secrets_20231128TwilioStaticCredentialsResponse
type Secrets20231128TwilioStaticCredentialsResponse struct {

	// account sid
	AccountSid string `json:"account_sid,omitempty"`

	// api key sid
	APIKeySid string `json:"api_key_sid,omitempty"`
}

// Validate validates this secrets 20231128 twilio static credentials response
func (m *Secrets20231128TwilioStaticCredentialsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 twilio static credentials response based on context it is used
func (m *Secrets20231128TwilioStaticCredentialsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128TwilioStaticCredentialsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128TwilioStaticCredentialsResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128TwilioStaticCredentialsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
