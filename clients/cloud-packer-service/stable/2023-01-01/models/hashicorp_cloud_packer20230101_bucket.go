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

// HashicorpCloudPacker20230101Bucket hashicorp cloud packer 20230101 bucket
//
// swagger:model hashicorp.cloud.packer_20230101.Bucket
type HashicorpCloudPacker20230101Bucket struct {

	// Information about this bucket's children. Children are buckets that used
	// any version of this bucket as the base artifact for their latest complete
	// version.
	Children *HashicorpCloudPacker20230101BucketChildren `json:"children,omitempty"`

	// Created datetime.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Short description of what this bucket is for.
	Description string `json:"description,omitempty"`

	// Unique identifier (ULID).
	ID string `json:"id,omitempty"`

	// A key:value map for custom, user-settable metadata about your bucket.
	Labels map[string]string `json:"labels,omitempty"`

	// The bucket's most recent valid version, same assigned to the latest channel.
	LatestVersion *HashicorpCloudPacker20230101Version `json:"latest_version,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// Human-readable name for the bucket.
	Name string `json:"name,omitempty"`

	// Information about this bucket's parents. Parents are the base artifacts
	// OSS Packer used for building the latest complete version in this bucket.
	Parents *HashicorpCloudPacker20230101BucketParents `json:"parents,omitempty"`

	// List of the cloud providers or other platforms that are included in the
	// latest complete version. e.g aws, gcp, or azure.
	Platforms []string `json:"platforms"`

	// The human-readable identifier in the platform. For example, `packer/project/<project-id>/bucket/<bucket-name>`.
	ResourceName string `json:"resource_name,omitempty"`

	// Last updated datetime.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Total number of versions in this bucket.
	VersionCount string `json:"version_count,omitempty"`
}

// Validate validates this hashicorp cloud packer 20230101 bucket
func (m *HashicorpCloudPacker20230101Bucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParents(formats); err != nil {
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

func (m *HashicorpCloudPacker20230101Bucket) validateChildren(formats strfmt.Registry) error {
	if swag.IsZero(m.Children) { // not required
		return nil
	}

	if m.Children != nil {
		if err := m.Children.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("children")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("children")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPacker20230101Bucket) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudPacker20230101Bucket) validateLatestVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestVersion) { // not required
		return nil
	}

	if m.LatestVersion != nil {
		if err := m.LatestVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest_version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latest_version")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPacker20230101Bucket) validateLocation(formats strfmt.Registry) error {
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

func (m *HashicorpCloudPacker20230101Bucket) validateParents(formats strfmt.Registry) error {
	if swag.IsZero(m.Parents) { // not required
		return nil
	}

	if m.Parents != nil {
		if err := m.Parents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parents")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPacker20230101Bucket) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer 20230101 bucket based on the context it is used
func (m *HashicorpCloudPacker20230101Bucket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChildren(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPacker20230101Bucket) contextValidateChildren(ctx context.Context, formats strfmt.Registry) error {

	if m.Children != nil {

		if swag.IsZero(m.Children) { // not required
			return nil
		}

		if err := m.Children.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("children")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("children")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPacker20230101Bucket) contextValidateLatestVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.LatestVersion != nil {

		if swag.IsZero(m.LatestVersion) { // not required
			return nil
		}

		if err := m.LatestVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest_version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latest_version")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudPacker20230101Bucket) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if swag.IsZero(m.Location) { // not required
			return nil
		}

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

func (m *HashicorpCloudPacker20230101Bucket) contextValidateParents(ctx context.Context, formats strfmt.Registry) error {

	if m.Parents != nil {

		if swag.IsZero(m.Parents) { // not required
			return nil
		}

		if err := m.Parents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parents")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101Bucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPacker20230101Bucket) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPacker20230101Bucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
