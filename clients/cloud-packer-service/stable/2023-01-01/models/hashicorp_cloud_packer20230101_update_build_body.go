// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPacker20230101UpdateBuildBody hashicorp cloud packer 20230101 update build body
//
// swagger:model hashicorp.cloud.packer_20230101.UpdateBuildBody
type HashicorpCloudPacker20230101UpdateBuildBody struct {

	// A list of artifacts to create and associate with this build.
	Artifacts []*HashicorpCloudPacker20230101ArtifactCreateBody `json:"artifacts"`

	// A key:value map for custom, user-settable metadata about your build.
	Labels map[string]string `json:"labels,omitempty"`

	// Additional information set by Packer about a build, such as plugins used.
	Metadata *HashicorpCloudPacker20230101BuildMetadata `json:"metadata,omitempty"`

	// The UUID specific to this call to Packer build. If you use the manifest
	// post-processor, this UUID will match the UUID present there.
	PackerRunUUID string `json:"packer_run_uuid,omitempty"`

	// The ID of the channel that was used to fetch the parent_version_id.
	// When the parent_channel_id is set, parent_version_id should also be set.
	ParentChannelID string `json:"parent_channel_id,omitempty"`

	// The ID of the parent version associated with the `source_external_identifier`.
	ParentVersionID string `json:"parent_version_id,omitempty"`

	// The platform that this build produced artifacts for.
	// For example, AWS, GCP, or Azure.
	Platform string `json:"platform,omitempty"`

	// The ID or URL of the remote cloud source artifact. Used for tracking artifact
	// dependencies for build pipelines.
	SourceExternalIdentifier string `json:"source_external_identifier,omitempty"`

	// Status of the build. The status can be RUNNING, DONE, CANCELLED, FAILED,
	// or UNSET.
	Status *HashicorpCloudPacker20230101BuildStatus `json:"status,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 update build body
func (m *HashicorpCloudPacker20230101UpdateBuildBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101UpdateBuildBody) validateArtifacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Artifacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Artifacts); i++ {
		if swag.IsZero(m.Artifacts[i]) { // not required
			continue
		}

		if m.Artifacts[i] != nil {
			if err := m.Artifacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artifacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPacker20230101UpdateBuildBody) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPacker20230101UpdateBuildBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer 20230101 update build body based on the context it is used
func (m *HashicorpCloudPacker20230101UpdateBuildBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArtifacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101UpdateBuildBody) contextValidateArtifacts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Artifacts); i++ {

		if m.Artifacts[i] != nil {

			if swag.IsZero(m.Artifacts[i]) { // not required
				return nil
			}

			if err := m.Artifacts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artifacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPacker20230101UpdateBuildBody) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPacker20230101UpdateBuildBody) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101UpdateBuildBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101UpdateBuildBody) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101UpdateBuildBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
