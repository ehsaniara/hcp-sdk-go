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

// HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus PeeringConnectionStatus is the overall, summarized status of the peering connection, as determined by the states
// of the individual peers as well as the state of the PeeringConnection
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.PeeringConnectionStatus
type HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus string

func NewHashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus(value HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus) *HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus.
func (m HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus) Pointer() *HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus {
	return &m
}

const (

	// HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGUNDEFINED captures enum value "PEERING_UNDEFINED"
	HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGUNDEFINED HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus = "PEERING_UNDEFINED"

	// HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGCONNECTING captures enum value "PEERING_CONNECTING"
	HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGCONNECTING HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus = "PEERING_CONNECTING"

	// HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGACTIVE captures enum value "PEERING_ACTIVE"
	HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGACTIVE HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus = "PEERING_ACTIVE"

	// HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGDELETING captures enum value "PEERING_DELETING"
	HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGDELETING HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus = "PEERING_DELETING"

	// HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGFAILING captures enum value "PEERING_FAILING"
	HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusPEERINGFAILING HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus = "PEERING_FAILING"
)

// for schema
var hashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusEnum []interface{}

func init() {
	var res []HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus
	if err := json.Unmarshal([]byte(`["PEERING_UNDEFINED","PEERING_CONNECTING","PEERING_ACTIVE","PEERING_DELETING","PEERING_FAILING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusEnum = append(hashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusEnum, v)
	}
}

func (m HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus) validateHashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusEnum(path, location string, value HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud global network manager 20220215 peering connection status
func (m HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud global network manager 20220215 peering connection status based on context it is used
func (m HashicorpCloudGlobalNetworkManager20220215PeeringConnectionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}