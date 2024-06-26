// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128MongoDBRole secrets 20231128 mongo d b role
//
// swagger:model secrets_20231128MongoDBRole
type Secrets20231128MongoDBRole struct {

	// collection name
	CollectionName string `json:"collection_name,omitempty"`

	// database name
	DatabaseName string `json:"database_name,omitempty"`

	// role name
	RoleName string `json:"role_name,omitempty"`
}

// Validate validates this secrets 20231128 mongo d b role
func (m *Secrets20231128MongoDBRole) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets 20231128 mongo d b role based on context it is used
func (m *Secrets20231128MongoDBRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128MongoDBRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128MongoDBRole) UnmarshalBinary(b []byte) error {
	var res Secrets20231128MongoDBRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
