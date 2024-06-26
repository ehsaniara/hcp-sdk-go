// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointActionRunVariable NOTE(briancain): In action vars phase 1, this might turn into its own proto
// message. For now, we'll define it inline here but will likely break our API
// to change ActionRuns to reference those instead.
//
// swagger:model hashicorp.cloud.waypoint.ActionRun.Variable
type HashicorpCloudWaypointActionRunVariable struct {

	// The name of the variable
	Name string `json:"name,omitempty"`

	// Whether the variable is sensitive or not
	Sensitive bool `json:"sensitive,omitempty"`

	// The value of the variable
	Value string `json:"value,omitempty"`
}

// Validate validates this hashicorp cloud waypoint action run variable
func (m *HashicorpCloudWaypointActionRunVariable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint action run variable based on context it is used
func (m *HashicorpCloudWaypointActionRunVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionRunVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointActionRunVariable) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointActionRunVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
