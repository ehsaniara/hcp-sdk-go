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
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudNetwork20200907Network Network represents a single operation.
//
// swagger:model hashicorp.cloud.network_20200907.Network
type HashicorpCloudNetwork20200907Network struct {

	// cidr_block is the IP range of the HVN.
	CidrBlock string `json:"cidr_block,omitempty"`

	// created_at is a timestamp when the network was originally created
	//
	// Output only.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id is the unique ID for this operation used in other RPC calls.
	// This ID is only guaranteed to be unique within the region that
	// the operation is running in.
	ID string `json:"id,omitempty"`

	// location is the location of the HVN
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// provider_network_data contains information about the underlying cloud provider network.
	//
	// Output only.
	// Read Only: true
	ProviderNetworkData *HashicorpCloudNetwork20200907NetworkProviderNetworkData `json:"provider_network_data,omitempty"`

	// state is the current state of the network.
	//
	// Output only.
	// Read Only: true
	State *HashicorpCloudNetwork20200907NetworkState `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 network
func (m *HashicorpCloudNetwork20200907Network) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderNetworkData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907Network) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Network) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Network) validateProviderNetworkData(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderNetworkData) { // not required
		return nil
	}

	if m.ProviderNetworkData != nil {
		if err := m.ProviderNetworkData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider_network_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider_network_data")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Network) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 network based on the context it is used
func (m *HashicorpCloudNetwork20200907Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProviderNetworkData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907Network) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Network) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Network) contextValidateProviderNetworkData(ctx context.Context, formats strfmt.Registry) error {

	if m.ProviderNetworkData != nil {
		if err := m.ProviderNetworkData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider_network_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider_network_data")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907Network) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907Network) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907Network) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}