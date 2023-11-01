// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointActionsdriverActionConfigFlavorCustomHeader hashicorp cloud waypoint actionsdriver action config flavor custom header
//
// swagger:model hashicorp.cloud.waypoint.actionsdriver.ActionConfig.Flavor.Custom.Header
type HashicorpCloudWaypointActionsdriverActionConfigFlavorCustomHeader struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this hashicorp cloud waypoint actionsdriver action config flavor custom header
func (m *HashicorpCloudWaypointActionsdriverActionConfigFlavorCustomHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint actionsdriver action config flavor custom header based on context it is used
func (m *HashicorpCloudWaypointActionsdriverActionConfigFlavorCustomHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionsdriverActionConfigFlavorCustomHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionsdriverActionConfigFlavorCustomHeader) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointActionsdriverActionConfigFlavorCustomHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}