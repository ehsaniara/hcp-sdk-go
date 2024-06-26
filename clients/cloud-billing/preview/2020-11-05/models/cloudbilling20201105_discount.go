// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cloudbilling20201105Discount Discount represents a consumption pool discount
//
// swagger:model cloudbilling_20201105Discount
type Cloudbilling20201105Discount struct {

	// percentage is the discount percentage.
	Percentage string `json:"percentage,omitempty"`

	// product_type is the product type this discount is for.
	ProductType string `json:"product_type,omitempty"`
}

// Validate validates this cloudbilling 20201105 discount
func (m *Cloudbilling20201105Discount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloudbilling 20201105 discount based on context it is used
func (m *Cloudbilling20201105Discount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cloudbilling20201105Discount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cloudbilling20201105Discount) UnmarshalBinary(b []byte) error {
	var res Cloudbilling20201105Discount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
