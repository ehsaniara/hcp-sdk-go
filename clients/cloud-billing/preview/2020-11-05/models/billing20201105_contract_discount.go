// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Billing20201105ContractDiscount Discount represents discount under a contract.
//
// swagger:model billing_20201105ContractDiscount
type Billing20201105ContractDiscount struct {

	// percentage is the discount percentage.
	Percentage string `json:"percentage,omitempty"`

	// product_type is the product type this discount is for.
	ProductType string `json:"product_type,omitempty"`
}

// Validate validates this billing 20201105 contract discount
func (m *Billing20201105ContractDiscount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this billing 20201105 contract discount based on context it is used
func (m *Billing20201105ContractDiscount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105ContractDiscount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105ContractDiscount) UnmarshalBinary(b []byte) error {
	var res Billing20201105ContractDiscount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
