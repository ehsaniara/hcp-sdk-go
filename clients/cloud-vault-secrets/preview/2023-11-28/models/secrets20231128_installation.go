// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128Installation secrets 20231128 installation
//
// swagger:model secrets_20231128Installation
type Secrets20231128Installation struct {

	// account login
	AccountLogin string `json:"account_login,omitempty"`

	// account type
	AccountType *GhAppMetadataAccountType `json:"account_type,omitempty"`

	// app slug
	AppSlug string `json:"app_slug,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this secrets 20231128 installation
func (m *Secrets20231128Installation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128Installation) validateAccountType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if m.AccountType != nil {
		if err := m.AccountType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account_type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20231128 installation based on the context it is used
func (m *Secrets20231128Installation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128Installation) contextValidateAccountType(ctx context.Context, formats strfmt.Registry) error {

	if m.AccountType != nil {

		if swag.IsZero(m.AccountType) { // not required
			return nil
		}

		if err := m.AccountType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128Installation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128Installation) UnmarshalBinary(b []byte) error {
	var res Secrets20231128Installation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
