// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulTelemetry20230414GetLabelValuesResponse hashicorp cloud consul telemetry 20230414 get label values response
//
// swagger:model hashicorp.cloud.consul_telemetry_20230414.GetLabelValuesResponse
type HashicorpCloudConsulTelemetry20230414GetLabelValuesResponse struct {

	// data
	Data []string `json:"data"`
}

// Validate validates this hashicorp cloud consul telemetry 20230414 get label values response
func (m *HashicorpCloudConsulTelemetry20230414GetLabelValuesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud consul telemetry 20230414 get label values response based on context it is used
func (m *HashicorpCloudConsulTelemetry20230414GetLabelValuesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414GetLabelValuesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulTelemetry20230414GetLabelValuesResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulTelemetry20230414GetLabelValuesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
