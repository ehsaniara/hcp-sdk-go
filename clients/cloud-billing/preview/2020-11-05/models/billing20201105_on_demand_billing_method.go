// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Billing20201105OnDemandBillingMethod OnDemandBillingMethod contains the information used to register a Billing
// Account for periodic invoices and charges through our payments processor.
//
// swagger:model billing_20201105OnDemandBillingMethod
type Billing20201105OnDemandBillingMethod struct {

	// billing_address is the address that will show on invoices.
	BillingAddress *Billing20201105Address `json:"billing_address,omitempty"`

	// email is the email address to which invoices will be sent.
	Email string `json:"email,omitempty"`

	// name is the customer's full name or business name that will appear on
	// invoices.
	Name string `json:"name,omitempty"`

	// stripe_setup_intent_id is the ID of the Stripe Setup Intent used to collect
	// and tokenize payment details.
	//
	// https://stripe.com/docs/payments/setup-intents
	//
	// Call the CreateSetupIntent endpoint first to obtain a client secret for
	// use with Stripe.js (e.g. given as an argument to confirmCardSetup) and
	// then provide the resulting Setup Intent's ID in this call.
	//
	// Note: this field should not be set if you're not updating the payment
	// method.
	StripeSetupIntentID string `json:"stripe_setup_intent_id,omitempty"`

	// tax_settings determine how tax will be calculated.
	TaxSettings *Billing20201105TaxSettings `json:"tax_settings,omitempty"`
}

// Validate validates this billing 20201105 on demand billing method
func (m *Billing20201105OnDemandBillingMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105OnDemandBillingMethod) validateBillingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingAddress) { // not required
		return nil
	}

	if m.BillingAddress != nil {
		if err := m.BillingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_address")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105OnDemandBillingMethod) validateTaxSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxSettings) { // not required
		return nil
	}

	if m.TaxSettings != nil {
		if err := m.TaxSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tax_settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this billing 20201105 on demand billing method based on the context it is used
func (m *Billing20201105OnDemandBillingMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105OnDemandBillingMethod) contextValidateBillingAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingAddress != nil {

		if swag.IsZero(m.BillingAddress) { // not required
			return nil
		}

		if err := m.BillingAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_address")
			}
			return err
		}
	}

	return nil
}

func (m *Billing20201105OnDemandBillingMethod) contextValidateTaxSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxSettings != nil {

		if swag.IsZero(m.TaxSettings) { // not required
			return nil
		}

		if err := m.TaxSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tax_settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105OnDemandBillingMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105OnDemandBillingMethod) UnmarshalBinary(b []byte) error {
	var res Billing20201105OnDemandBillingMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
