// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudConsul20200826CreateSnapshotResponse CreateSnapshotResponse is a response from creating a snapshot.
//
// swagger:model hashicorp.cloud.consul_20200826.CreateSnapshotResponse
type HashicorpCloudConsul20200826CreateSnapshotResponse struct {

	// operation represents the newly created snapshot.
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`

	// snapshot_id is the identifer of the newly created snapshot.
	SnapshotID string `json:"snapshot_id,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200826 create snapshot response
func (m *HashicorpCloudConsul20200826CreateSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826CreateSnapshotResponse) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826CreateSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826CreateSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200826CreateSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
