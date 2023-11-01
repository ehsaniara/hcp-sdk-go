// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudWaypointTerraformTFRunState hashicorp cloud waypoint terraform t f run state
//
// swagger:model hashicorp.cloud.waypoint.Terraform.TFRunState
type HashicorpCloudWaypointTerraformTFRunState string

func NewHashicorpCloudWaypointTerraformTFRunState(value HashicorpCloudWaypointTerraformTFRunState) *HashicorpCloudWaypointTerraformTFRunState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudWaypointTerraformTFRunState.
func (m HashicorpCloudWaypointTerraformTFRunState) Pointer() *HashicorpCloudWaypointTerraformTFRunState {
	return &m
}

const (

	// HashicorpCloudWaypointTerraformTFRunStateUNKNOWN captures enum value "UNKNOWN"
	HashicorpCloudWaypointTerraformTFRunStateUNKNOWN HashicorpCloudWaypointTerraformTFRunState = "UNKNOWN"

	// HashicorpCloudWaypointTerraformTFRunStateRUNNING captures enum value "RUNNING"
	HashicorpCloudWaypointTerraformTFRunStateRUNNING HashicorpCloudWaypointTerraformTFRunState = "RUNNING"

	// HashicorpCloudWaypointTerraformTFRunStateSUCCESS captures enum value "SUCCESS"
	HashicorpCloudWaypointTerraformTFRunStateSUCCESS HashicorpCloudWaypointTerraformTFRunState = "SUCCESS"

	// HashicorpCloudWaypointTerraformTFRunStateERROR captures enum value "ERROR"
	HashicorpCloudWaypointTerraformTFRunStateERROR HashicorpCloudWaypointTerraformTFRunState = "ERROR"
)

// for schema
var hashicorpCloudWaypointTerraformTFRunStateEnum []interface{}

func init() {
	var res []HashicorpCloudWaypointTerraformTFRunState
	if err := json.Unmarshal([]byte(`["UNKNOWN","RUNNING","SUCCESS","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudWaypointTerraformTFRunStateEnum = append(hashicorpCloudWaypointTerraformTFRunStateEnum, v)
	}
}

func (m HashicorpCloudWaypointTerraformTFRunState) validateHashicorpCloudWaypointTerraformTFRunStateEnum(path, location string, value HashicorpCloudWaypointTerraformTFRunState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudWaypointTerraformTFRunStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud waypoint terraform t f run state
func (m HashicorpCloudWaypointTerraformTFRunState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudWaypointTerraformTFRunStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud waypoint terraform t f run state based on context it is used
func (m HashicorpCloudWaypointTerraformTFRunState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}