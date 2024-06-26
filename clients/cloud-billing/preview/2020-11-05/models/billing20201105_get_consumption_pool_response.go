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

// Billing20201105GetConsumptionPoolResponse GetConsumptionPoolResponse is a response for GetConsumptionPoolById endpoint
//
// swagger:model billing_20201105GetConsumptionPoolResponse
type Billing20201105GetConsumptionPoolResponse struct {

	// consumption_pool is the consumption pool with the corresponding id
	ConsumptionPool *Billing20201105ConsumptionPool `json:"consumption_pool,omitempty"`

	// consumption_pool_updates is the information about any updates that have been made to the consumption pool
	ConsumptionPoolUpdates []*Billing20201105ConsumptionPoolUpdateInfo `json:"consumption_pool_updates"`
}

// Validate validates this billing 20201105 get consumption pool response
func (m *Billing20201105GetConsumptionPoolResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsumptionPool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumptionPoolUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105GetConsumptionPoolResponse) validateConsumptionPool(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumptionPool) { // not required
		return nil
	}

	if m.ConsumptionPool != nil {
		if err := m.ConsumptionPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consumption_pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consumption_pool")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105GetConsumptionPoolResponse) validateConsumptionPoolUpdates(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumptionPoolUpdates) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsumptionPoolUpdates); i++ {
		if swag.IsZero(m.ConsumptionPoolUpdates[i]) { // not required
			continue
		}

		if m.ConsumptionPoolUpdates[i] != nil {
			if err := m.ConsumptionPoolUpdates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumption_pool_updates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consumption_pool_updates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this billing 20201105 get consumption pool response based on the context it is used
func (m *Billing20201105GetConsumptionPoolResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsumptionPool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsumptionPoolUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105GetConsumptionPoolResponse) contextValidateConsumptionPool(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsumptionPool != nil {

		if swag.IsZero(m.ConsumptionPool) { // not required
			return nil
		}

		if err := m.ConsumptionPool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consumption_pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consumption_pool")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105GetConsumptionPoolResponse) contextValidateConsumptionPoolUpdates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsumptionPoolUpdates); i++ {

		if m.ConsumptionPoolUpdates[i] != nil {

			if swag.IsZero(m.ConsumptionPoolUpdates[i]) { // not required
				return nil
			}

			if err := m.ConsumptionPoolUpdates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumption_pool_updates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consumption_pool_updates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105GetConsumptionPoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105GetConsumptionPoolResponse) UnmarshalBinary(b []byte) error {
	var res Billing20201105GetConsumptionPoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
