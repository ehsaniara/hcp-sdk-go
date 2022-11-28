// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudGlobalNetworkManager20220215RaftInfo hashicorp cloud global network manager 20220215 raft info
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.RaftInfo
type HashicorpCloudGlobalNetworkManager20220215RaftInfo struct {

	// applied index
	AppliedIndex string `json:"applied_index,omitempty"`

	// is leader
	IsLeader bool `json:"is_leader,omitempty"`

	// known leader
	KnownLeader bool `json:"known_leader,omitempty"`

	// time since last contact
	TimeSinceLastContact string `json:"time_since_last_contact,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 raft info
func (m *HashicorpCloudGlobalNetworkManager20220215RaftInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud global network manager 20220215 raft info based on context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215RaftInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215RaftInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215RaftInfo) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215RaftInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}