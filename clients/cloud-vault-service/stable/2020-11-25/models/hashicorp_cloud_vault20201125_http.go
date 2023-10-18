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

// HashicorpCloudVault20201125HTTP hashicorp cloud vault 20201125 HTTP
//
// swagger:model hashicorp.cloud.vault_20201125.HTTP
type HashicorpCloudVault20201125HTTP struct {

	// basic
	Basic *HashicorpCloudVault20201125HTTPBasicAuth `json:"basic,omitempty"`

	// bearer
	Bearer *HashicorpCloudVault20201125HTTPBearerAuth `json:"bearer,omitempty"`

	// codec
	Codec *HashicorpCloudVault20201125HTTPEncodingCodec `json:"codec,omitempty"`

	// compression
	Compression bool `json:"compression,omitempty"`

	// headers
	Headers interface{} `json:"headers,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// payload prefix
	PayloadPrefix string `json:"payload_prefix,omitempty"`

	// payload suffix
	PayloadSuffix string `json:"payload_suffix,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 HTTP
func (m *HashicorpCloudVault20201125HTTP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBearer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCodec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125HTTP) validateBasic(formats strfmt.Registry) error {
	if swag.IsZero(m.Basic) { // not required
		return nil
	}

	if m.Basic != nil {
		if err := m.Basic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125HTTP) validateBearer(formats strfmt.Registry) error {
	if swag.IsZero(m.Bearer) { // not required
		return nil
	}

	if m.Bearer != nil {
		if err := m.Bearer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bearer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bearer")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125HTTP) validateCodec(formats strfmt.Registry) error {
	if swag.IsZero(m.Codec) { // not required
		return nil
	}

	if m.Codec != nil {
		if err := m.Codec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20201125 HTTP based on the context it is used
func (m *HashicorpCloudVault20201125HTTP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBasic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBearer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCodec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125HTTP) contextValidateBasic(ctx context.Context, formats strfmt.Registry) error {

	if m.Basic != nil {
		if err := m.Basic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125HTTP) contextValidateBearer(ctx context.Context, formats strfmt.Registry) error {

	if m.Bearer != nil {
		if err := m.Bearer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bearer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bearer")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125HTTP) contextValidateCodec(ctx context.Context, formats strfmt.Registry) error {

	if m.Codec != nil {
		if err := m.Codec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125HTTP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125HTTP) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125HTTP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}