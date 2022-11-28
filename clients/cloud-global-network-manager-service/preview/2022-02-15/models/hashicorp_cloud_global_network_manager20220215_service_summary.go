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

// HashicorpCloudGlobalNetworkManager20220215ServiceSummary ServiceSummary is an aggregate summary about the health of a service across one or more clusters
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ServiceSummary
type HashicorpCloudGlobalNetworkManager20220215ServiceSummary struct {

	// critical
	Critical int32 `json:"critical,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// passing
	Passing int32 `json:"passing,omitempty"`

	// summaries
	Summaries []*HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry `json:"summaries"`

	// total service instances for the service after applying any requested filters
	Total int32 `json:"total,omitempty"`

	// warning
	Warning int32 `json:"warning,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 service summary
func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummaries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummary) validateSummaries(formats strfmt.Registry) error {
	if swag.IsZero(m.Summaries) { // not required
		return nil
	}

	for i := 0; i < len(m.Summaries); i++ {
		if swag.IsZero(m.Summaries[i]) { // not required
			continue
		}

		if m.Summaries[i] != nil {
			if err := m.Summaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("summaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 service summary based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSummaries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummary) contextValidateSummaries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Summaries); i++ {

		if m.Summaries[i] != nil {
			if err := m.Summaries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("summaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummary) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ServiceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}