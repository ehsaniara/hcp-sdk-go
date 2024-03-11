// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128Principal Information about the principal who created or updated the resource
//
// swagger:model secrets_20231128Principal
type Secrets20231128Principal struct {

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this secrets 20231128 principal
func (m *Secrets20231128Principal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 principal based on context it is used
func (m *Secrets20231128Principal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128Principal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128Principal) UnmarshalBinary(b []byte) error {
	var res Secrets20231128Principal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
