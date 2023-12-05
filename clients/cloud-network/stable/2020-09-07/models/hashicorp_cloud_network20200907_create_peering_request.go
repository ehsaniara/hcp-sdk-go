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

// HashicorpCloudNetwork20200907CreatePeeringRequest CreatePeeringRequest is a request type for CreatePeering endpoint
//
// swagger:model hashicorp.cloud.network_20200907.CreatePeeringRequest
type HashicorpCloudNetwork20200907CreatePeeringRequest struct {

	// peering is the peering to be created.
	Peering *HashicorpCloudNetwork20200907Peering `json:"peering,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 create peering request
func (m *HashicorpCloudNetwork20200907CreatePeeringRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeering(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreatePeeringRequest) validatePeering(formats strfmt.Registry) error {
	if swag.IsZero(m.Peering) { // not required
		return nil
	}

	if m.Peering != nil {
		if err := m.Peering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peering")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 create peering request based on the context it is used
func (m *HashicorpCloudNetwork20200907CreatePeeringRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeering(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreatePeeringRequest) contextValidatePeering(ctx context.Context, formats strfmt.Registry) error {

	if m.Peering != nil {

		if swag.IsZero(m.Peering) { // not required
			return nil
		}

		if err := m.Peering.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peering")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peering")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreatePeeringRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreatePeeringRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907CreatePeeringRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
