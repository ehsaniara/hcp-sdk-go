// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudConsul20200826DeleteSnapshotResponse DeleteSnapshotResponse is a response to deleting a snapshot.
//
// swagger:model hashicorp.cloud.consul_20200826.DeleteSnapshotResponse
type HashicorpCloudConsul20200826DeleteSnapshotResponse struct {

	// operation represents the deletion of the requested snapshot.
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`
}

// Validate validates this hashicorp cloud consul 20200826 delete snapshot response
func (m *HashicorpCloudConsul20200826DeleteSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826DeleteSnapshotResponse) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud consul 20200826 delete snapshot response based on the context it is used
func (m *HashicorpCloudConsul20200826DeleteSnapshotResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsul20200826DeleteSnapshotResponse) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {
		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826DeleteSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsul20200826DeleteSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsul20200826DeleteSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
