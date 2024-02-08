// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudWaypointWaypointServiceValidateAgentGroupsBody hashicorp cloud waypoint waypoint service validate agent groups body
//
// swagger:model hashicorp.cloud.waypoint.WaypointService.ValidateAgentGroupsBody
type HashicorpCloudWaypointWaypointServiceValidateAgentGroupsBody struct {

	// The list of groups that the caller is interested in
	Groups []string `json:"groups"`

	// Global references the entire server. This is used in some APIs
	// as a way to read/write values that are server-global.
	Namespace interface{} `json:"namespace,omitempty"`
}

// Validate validates this hashicorp cloud waypoint waypoint service validate agent groups body
func (m *HashicorpCloudWaypointWaypointServiceValidateAgentGroupsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint waypoint service validate agent groups body based on context it is used
func (m *HashicorpCloudWaypointWaypointServiceValidateAgentGroupsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceValidateAgentGroupsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointWaypointServiceValidateAgentGroupsBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointWaypointServiceValidateAgentGroupsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
