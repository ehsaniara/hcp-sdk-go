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

// HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility hashicorp cloud global network manager 20220215 admin partition with eligibility
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.AdminPartitionWithEligibility
type HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility struct {

	// description is the description of the consul admin partition
	Description string `json:"description,omitempty"`

	// ineligibility_reasons is a list of reasons why this particular partition is ineligible for peering.
	// Empty list means the partition is eligible.
	IneligibilityReasons []*HashicorpCloudGlobalNetworkManager20220215PartitionPeeringIneligibilityReason `json:"ineligibility_reasons"`

	// name is the name of the consul admin partition
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 admin partition with eligibility
func (m *HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIneligibilityReasons(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility) validateIneligibilityReasons(formats strfmt.Registry) error {
	if swag.IsZero(m.IneligibilityReasons) { // not required
		return nil
	}

	for i := 0; i < len(m.IneligibilityReasons); i++ {
		if swag.IsZero(m.IneligibilityReasons[i]) { // not required
			continue
		}

		if m.IneligibilityReasons[i] != nil {
			if err := m.IneligibilityReasons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ineligibility_reasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ineligibility_reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 admin partition with eligibility based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIneligibilityReasons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility) contextValidateIneligibilityReasons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IneligibilityReasons); i++ {

		if m.IneligibilityReasons[i] != nil {
			if err := m.IneligibilityReasons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ineligibility_reasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ineligibility_reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}