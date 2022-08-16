// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20220411ParentBuildStatus The status of a parent image build.
//
// swagger:model hashicorp.cloud.packer_20220411.ParentBuildStatus
type HashicorpCloudPacker20220411ParentBuildStatus struct {

	// The parent bucket slug.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// The channel associated with this relationship.
	Channel string `json:"channel,omitempty"`

	// The parent ID/URL.
	ImageID string `json:"image_id,omitempty"`

	// The parent iteration fingerprint.
	IterationFingerprint string `json:"iteration_fingerprint,omitempty"`

	// The parent iteration ULID.
	IterationID string `json:"iteration_id,omitempty"`

	// The parent iteration incremental version.
	IterationIncrementalVersion int32 `json:"iteration_incremental_version,omitempty"`

	// The parent status.
	Status HashicorpCloudPacker20220411ParentBuildStatusStatus `json:"status,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 parent build status
func (m *HashicorpCloudPacker20220411ParentBuildStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20220411ParentBuildStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411ParentBuildStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411ParentBuildStatus) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411ParentBuildStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}