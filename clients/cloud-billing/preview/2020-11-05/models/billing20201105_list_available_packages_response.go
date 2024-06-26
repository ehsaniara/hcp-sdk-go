// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Billing20201105ListAvailablePackagesResponse ListAvailablePackagesResponse is the response to ListAvailablePackages & contains a list of available packages.
//
// swagger:model billing_20201105ListAvailablePackagesResponse
type Billing20201105ListAvailablePackagesResponse struct {

	// packages is the list of available Packages in response to the ListAvailablePackagesRequest.
	Packages []*Billing20201105AvailablePackage `json:"packages"`
}

// Validate validates this billing 20201105 list available packages response
func (m *Billing20201105ListAvailablePackagesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ListAvailablePackagesResponse) validatePackages(formats strfmt.Registry) error {
	if swag.IsZero(m.Packages) { // not required
		return nil
	}

	for i := 0; i < len(m.Packages); i++ {
		if swag.IsZero(m.Packages[i]) { // not required
			continue
		}

		if m.Packages[i] != nil {
			if err := m.Packages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this billing 20201105 list available packages response based on the context it is used
func (m *Billing20201105ListAvailablePackagesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePackages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ListAvailablePackagesResponse) contextValidatePackages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Packages); i++ {

		if m.Packages[i] != nil {

			if swag.IsZero(m.Packages[i]) { // not required
				return nil
			}

			if err := m.Packages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105ListAvailablePackagesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105ListAvailablePackagesResponse) UnmarshalBinary(b []byte) error {
	var res Billing20201105ListAvailablePackagesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
