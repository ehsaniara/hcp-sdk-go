// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Secrets20231128UpsertSyncInstallationResponse secrets 20231128 upsert sync installation response
//
// swagger:model secrets_20231128UpsertSyncInstallationResponse
type Secrets20231128UpsertSyncInstallationResponse struct {

	// sync installation
	SyncInstallation *Secrets20231128SyncInstallation `json:"sync_installation,omitempty"`
}

// Validate validates this secrets 20231128 upsert sync installation response
func (m *Secrets20231128UpsertSyncInstallationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncInstallation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128UpsertSyncInstallationResponse) validateSyncInstallation(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncInstallation) { // not required
		return nil
	}

	if m.SyncInstallation != nil {
		if err := m.SyncInstallation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync_installation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync_installation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20231128 upsert sync installation response based on the context it is used
func (m *Secrets20231128UpsertSyncInstallationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSyncInstallation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20231128UpsertSyncInstallationResponse) contextValidateSyncInstallation(ctx context.Context, formats strfmt.Registry) error {

	if m.SyncInstallation != nil {

		if swag.IsZero(m.SyncInstallation) { // not required
			return nil
		}

		if err := m.SyncInstallation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync_installation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync_installation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20231128UpsertSyncInstallationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20231128UpsertSyncInstallationResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20231128UpsertSyncInstallationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
