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

// HashicorpWaypointEntrypointExecResponse hashicorp waypoint entrypoint exec response
//
// swagger:model hashicorp.waypoint.EntrypointExecResponse
type HashicorpWaypointEntrypointExecResponse struct {

	// input is raw stdin input from the client
	// Format: byte
	Input strfmt.Base64 `json:"input,omitempty"`

	// input_eof means that stdin is now closed
	InputEOF interface{} `json:"input_eof,omitempty"`

	// opened is sent when the entrypoint session is successfully opened.
	// The value of this message is meaningless. The existence of the message
	// itself is a signal that the stream was opened properly.
	Opened bool `json:"opened,omitempty"`

	// winch is SIGWNCH information for window sizing
	Winch *HashicorpWaypointExecStreamRequestWindowSize `json:"winch,omitempty"`
}

// Validate validates this hashicorp waypoint entrypoint exec response
func (m *HashicorpWaypointEntrypointExecResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWinch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointEntrypointExecResponse) validateWinch(formats strfmt.Registry) error {
	if swag.IsZero(m.Winch) { // not required
		return nil
	}

	if m.Winch != nil {
		if err := m.Winch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("winch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("winch")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint entrypoint exec response based on the context it is used
func (m *HashicorpWaypointEntrypointExecResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWinch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointEntrypointExecResponse) contextValidateWinch(ctx context.Context, formats strfmt.Registry) error {

	if m.Winch != nil {

		if swag.IsZero(m.Winch) { // not required
			return nil
		}

		if err := m.Winch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("winch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("winch")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointEntrypointExecResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointEntrypointExecResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointEntrypointExecResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}