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

// HashicorpWaypointDecodeTokenResponse hashicorp waypoint decode token response
//
// swagger:model hashicorp.waypoint.DecodeTokenResponse
type HashicorpWaypointDecodeTokenResponse struct {

	// The decoded token.
	Token *HashicorpWaypointToken `json:"token,omitempty"`

	// Transport is the wrapper around the token. This may be useful
	// to look into the metadata around the token.
	Transport *HashicorpWaypointTokenTransport `json:"transport,omitempty"`
}

// Validate validates this hashicorp waypoint decode token response
func (m *HashicorpWaypointDecodeTokenResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointDecodeTokenResponse) validateToken(formats strfmt.Registry) error {
	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDecodeTokenResponse) validateTransport(formats strfmt.Registry) error {
	if swag.IsZero(m.Transport) { // not required
		return nil
	}

	if m.Transport != nil {
		if err := m.Transport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transport")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint decode token response based on the context it is used
func (m *HashicorpWaypointDecodeTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointDecodeTokenResponse) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {

		if swag.IsZero(m.Token) { // not required
			return nil
		}

		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointDecodeTokenResponse) contextValidateTransport(ctx context.Context, formats strfmt.Registry) error {

	if m.Transport != nil {

		if swag.IsZero(m.Transport) { // not required
			return nil
		}

		if err := m.Transport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transport")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointDecodeTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointDecodeTokenResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointDecodeTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}