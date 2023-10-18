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

// Billing20201105ListConsumptionPoolsResponse ListConsumptionPoolsResponse is a response for ListConsumptionPools endpoint
//
// swagger:model billing_20201105ListConsumptionPoolsResponse
type Billing20201105ListConsumptionPoolsResponse struct {

	// consumption_pools is the list of fetched consumption pools.
	ConsumptionPools []*Billing20201105ConsumptionPool `json:"consumption_pools"`

	// summary is the aggregated summary for all active consumption pools.
	Summary *Billing20201105ConsumptionPoolsSummary `json:"summary,omitempty"`
}

// Validate validates this billing 20201105 list consumption pools response
func (m *Billing20201105ListConsumptionPoolsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsumptionPools(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ListConsumptionPoolsResponse) validateConsumptionPools(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumptionPools) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsumptionPools); i++ {
		if swag.IsZero(m.ConsumptionPools[i]) { // not required
			continue
		}

		if m.ConsumptionPools[i] != nil {
			if err := m.ConsumptionPools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumption_pools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consumption_pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Billing20201105ListConsumptionPoolsResponse) validateSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing 20201105 list consumption pools response based on the context it is used
func (m *Billing20201105ListConsumptionPoolsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsumptionPools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ListConsumptionPoolsResponse) contextValidateConsumptionPools(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsumptionPools); i++ {

		if m.ConsumptionPools[i] != nil {
			if err := m.ConsumptionPools[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumption_pools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consumption_pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Billing20201105ListConsumptionPoolsResponse) contextValidateSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.Summary != nil {
		if err := m.Summary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105ListConsumptionPoolsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105ListConsumptionPoolsResponse) UnmarshalBinary(b []byte) error {
	var res Billing20201105ListConsumptionPoolsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}