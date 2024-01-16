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

// HashicorpCloudWaypointListTFCOrganizationsResponse hashicorp cloud waypoint list t f c organizations response
//
// swagger:model hashicorp.cloud.waypoint.ListTFCOrganizationsResponse
type HashicorpCloudWaypointListTFCOrganizationsResponse struct {

	// tfc organizations
	TfcOrganizations []*HashicorpCloudWaypointTFCOrganization `json:"tfc_organizations"`
}

// Validate validates this hashicorp cloud waypoint list t f c organizations response
func (m *HashicorpCloudWaypointListTFCOrganizationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTfcOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointListTFCOrganizationsResponse) validateTfcOrganizations(formats strfmt.Registry) error {
	if swag.IsZero(m.TfcOrganizations) { // not required
		return nil
	}

	for i := 0; i < len(m.TfcOrganizations); i++ {
		if swag.IsZero(m.TfcOrganizations[i]) { // not required
			continue
		}

		if m.TfcOrganizations[i] != nil {
			if err := m.TfcOrganizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tfc_organizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tfc_organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud waypoint list t f c organizations response based on the context it is used
func (m *HashicorpCloudWaypointListTFCOrganizationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTfcOrganizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudWaypointListTFCOrganizationsResponse) contextValidateTfcOrganizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TfcOrganizations); i++ {

		if m.TfcOrganizations[i] != nil {

			if swag.IsZero(m.TfcOrganizations[i]) { // not required
				return nil
			}

			if err := m.TfcOrganizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tfc_organizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tfc_organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudWaypointListTFCOrganizationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudWaypointListTFCOrganizationsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudWaypointListTFCOrganizationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}