// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudPacker20220411RegistryBillingDeprovision hashicorp cloud packer 20220411 registry billing deprovision
//
// swagger:model hashicorp.cloud.packer_20220411.RegistryBillingDeprovision
type HashicorpCloudPacker20220411RegistryBillingDeprovision struct {

	// The time the registry was deactivated of billing.
	// Format: date-time
	At strfmt.DateTime `json:"at,omitempty"`

	// Reason of why the registry was deactivated.
	Reason HashicorpCloudPacker20220411RegistryBillingDeprovisionReason `json:"reason,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 registry billing deprovision
func (m *HashicorpCloudPacker20220411RegistryBillingDeprovision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20220411RegistryBillingDeprovision) validateAt(formats strfmt.Registry) error {

	if swag.IsZero(m.At) { // not required
		return nil
	}

	if err := validate.FormatOf("at", "body", "date-time", m.At.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPacker20220411RegistryBillingDeprovision) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := m.Reason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411RegistryBillingDeprovision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411RegistryBillingDeprovision) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411RegistryBillingDeprovision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}