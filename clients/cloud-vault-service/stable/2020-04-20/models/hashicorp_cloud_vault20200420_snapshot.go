// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudVault20200420Snapshot Snapshot is our representation needed to back-up a Vault cluster.
//
// swagger:model hashicorp.cloud.vault_20200420.Snapshot
type HashicorpCloudVault20200420Snapshot struct {

	// cluster_id is the cluster id that this snapshot backs.
	ClusterID string `json:"cluster_id,omitempty"`

	// finished_at notes the time that this snapshot was finished.
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finished_at,omitempty"`

	// location is the location of the Snapshot.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// Name of the snapshot
	Name string `json:"name,omitempty"`

	// requested_at notes the time that this snapshot was requested.
	// Format: date-time
	RequestedAt strfmt.DateTime `json:"requested_at,omitempty"`

	// snapshot_id is the snapshots UUID.
	SnapshotID string `json:"snapshot_id,omitempty"`

	// state is represents the current status for this snapshot.
	State *HashicorpCloudVault20200420SnapshotState `json:"state,omitempty"`

	// type is the type of snapshot.
	Type *HashicorpCloudVault20200420SnapshotType `json:"type,omitempty"`

	// vault_version is the version of the Vault cluster this snapshot was taken from.
	VaultVersion string `json:"vault_version,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 snapshot
func (m *HashicorpCloudVault20200420Snapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420Snapshot) validateFinishedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20200420Snapshot) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420Snapshot) validateRequestedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("requested_at", "body", "date-time", m.RequestedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20200420Snapshot) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420Snapshot) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20200420 snapshot based on the context it is used
func (m *HashicorpCloudVault20200420Snapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420Snapshot) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420Snapshot) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20200420Snapshot) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420Snapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420Snapshot) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420Snapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
