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
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudVagrant20220930SearchIndexResponse hashicorp cloud vagrant 20220930 search index response
//
// swagger:model hashicorp.cloud.vagrant_20220930.SearchIndexResponse
type HashicorpCloudVagrant20220930SearchIndexResponse struct {

	// Boxes found from search
	Boxes []*HashicorpCloudVagrant20220930Index `json:"boxes"`

	// Pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud vagrant 20220930 search index response
func (m *HashicorpCloudVagrant20220930SearchIndexResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoxes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrant20220930SearchIndexResponse) validateBoxes(formats strfmt.Registry) error {
	if swag.IsZero(m.Boxes) { // not required
		return nil
	}

	for i := 0; i < len(m.Boxes); i++ {
		if swag.IsZero(m.Boxes[i]) { // not required
			continue
		}

		if m.Boxes[i] != nil {
			if err := m.Boxes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boxes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boxes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudVagrant20220930SearchIndexResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vagrant 20220930 search index response based on the context it is used
func (m *HashicorpCloudVagrant20220930SearchIndexResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoxes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrant20220930SearchIndexResponse) contextValidateBoxes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Boxes); i++ {

		if m.Boxes[i] != nil {

			if swag.IsZero(m.Boxes[i]) { // not required
				return nil
			}

			if err := m.Boxes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boxes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boxes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudVagrant20220930SearchIndexResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930SearchIndexResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrant20220930SearchIndexResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrant20220930SearchIndexResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
