// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20220411BucketAncestryChild hashicorp cloud packer 20220411 bucket ancestry child
//
// swagger:model hashicorp.cloud.packer_20220411.BucketAncestry.Child
type HashicorpCloudPacker20220411BucketAncestryChild struct {

	// The child bucket slug.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// The child iteration fingerprint.
	IterationFingerprint string `json:"iteration_fingerprint,omitempty"`

	// The child iteration ULID.
	IterationID string `json:"iteration_id,omitempty"`

	// The child iteration incremental version.
	IterationIncrementalVersion int32 `json:"iteration_incremental_version,omitempty"`
}

// Validate validates this hashicorp cloud packer 20220411 bucket ancestry child
func (m *HashicorpCloudPacker20220411BucketAncestryChild) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411BucketAncestryChild) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20220411BucketAncestryChild) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20220411BucketAncestryChild
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}