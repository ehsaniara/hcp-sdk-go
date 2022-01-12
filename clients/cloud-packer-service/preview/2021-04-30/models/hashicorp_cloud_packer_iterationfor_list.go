// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudPackerIterationforList The list endpoint does not return build information.
//
// swagger:model hashicorp.cloud.packer.IterationforList
type HashicorpCloudPackerIterationforList struct {

	// Who created the iteration.
	AuthorID string `json:"author_id,omitempty"`

	// Human-readable name for the bucket.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// Maps the build component type to its status enum, for displaying build
	// status in the iterations view.
	BuildStatuses map[string]string `json:"build_statuses,omitempty"`

	// If true, all builds associated with this iteration have successfully
	// completed and uploaded metadata to the registry. When "complete" is true,
	// This iteration is considered ready to use, and can have channels assigned
	// to it.
	Complete bool `json:"complete,omitempty"`

	// When the iteration was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Fingerprint of the iteration. The fingerprint is set by Packer when you
	// call `packer build`. It will most often correspond to a git commit sha,
	// but can be manually overridden by setting the environment variable
	// `HCP_PACKER_BUILD_FINGERPRINT`
	Fingerprint string `json:"fingerprint,omitempty"`

	// Unique identifier of the iteration; created and set by the HCP Packer
	// registry when the iteration is created.
	ID string `json:"id,omitempty"`

	// The human-readable version number assigned to this iteration. This
	// field will only be set if the iteration is "complete".
	IncrementalVersion int32 `json:"incremental_version,omitempty"`

	// The unique identifier of the iteration that was used as a source
	// for this iteration, if this iteration was built on a base layer.
	IterationAncestorID string `json:"iteration_ancestor_id,omitempty"`

	// A short explanation of why this iteration was revoked.
	RevocationMessage string `json:"revocation_message,omitempty"`

	// Timestamp from when the iteration is revoked an no longer trusted to be secure.
	// Format: date-time
	RevokeAt strfmt.DateTime `json:"revoke_at,omitempty"`

	// When the iteration was most recently updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud packer iterationfor list
func (m *HashicorpCloudPackerIterationforList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevokeAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerIterationforList) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPackerIterationforList) validateRevokeAt(formats strfmt.Registry) error {

	if swag.IsZero(m.RevokeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("revoke_at", "body", "date-time", m.RevokeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPackerIterationforList) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerIterationforList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerIterationforList) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerIterationforList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
