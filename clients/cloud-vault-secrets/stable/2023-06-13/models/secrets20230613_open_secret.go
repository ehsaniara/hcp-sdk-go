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
)

// Secrets20230613OpenSecret secrets 20230613 open secret
//
// swagger:model secrets_20230613OpenSecret
type Secrets20230613OpenSecret struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by
	CreatedBy *Secrets20230613Principal `json:"created_by,omitempty"`

	// latest version
	LatestVersion string `json:"latest_version,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sync status
	SyncStatus map[string]Secrets20230613SyncStatus `json:"sync_status,omitempty"`

	// version
	Version *Secrets20230613OpenSecretVersion `json:"version,omitempty"`
}

// Validate validates this secrets 20230613 open secret
func (m *Secrets20230613OpenSecret) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20230613OpenSecret) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Secrets20230613OpenSecret) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20230613OpenSecret) validateSyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncStatus) { // not required
		return nil
	}

	for k := range m.SyncStatus {

		if err := validate.Required("sync_status"+"."+k, "body", m.SyncStatus[k]); err != nil {
			return err
		}
		if val, ok := m.SyncStatus[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sync_status" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sync_status" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *Secrets20230613OpenSecret) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secrets 20230613 open secret based on the context it is used
func (m *Secrets20230613OpenSecret) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20230613OpenSecret) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *Secrets20230613OpenSecret) contextValidateSyncStatus(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.SyncStatus {

		if val, ok := m.SyncStatus[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Secrets20230613OpenSecret) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {

		if swag.IsZero(m.Version) { // not required
			return nil
		}

		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613OpenSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613OpenSecret) UnmarshalBinary(b []byte) error {
	var res Secrets20230613OpenSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
