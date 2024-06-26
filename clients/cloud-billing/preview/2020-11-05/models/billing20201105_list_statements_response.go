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

// Billing20201105ListStatementsResponse ListStatementsResponse is the response message returned by the ListStatements RPC.
//
// swagger:model billing_20201105ListStatementsResponse
type Billing20201105ListStatementsResponse struct {

	// statement_summaries is the array of statements for an org and Billing Account.
	StatementSummaries []*Billing20201105StatementSummary `json:"statement_summaries"`
}

// Validate validates this billing 20201105 list statements response
func (m *Billing20201105ListStatementsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatementSummaries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ListStatementsResponse) validateStatementSummaries(formats strfmt.Registry) error {
	if swag.IsZero(m.StatementSummaries) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementSummaries); i++ {
		if swag.IsZero(m.StatementSummaries[i]) { // not required
			continue
		}

		if m.StatementSummaries[i] != nil {
			if err := m.StatementSummaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statement_summaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statement_summaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this billing 20201105 list statements response based on the context it is used
func (m *Billing20201105ListStatementsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatementSummaries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Billing20201105ListStatementsResponse) contextValidateStatementSummaries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatementSummaries); i++ {

		if m.StatementSummaries[i] != nil {

			if swag.IsZero(m.StatementSummaries[i]) { // not required
				return nil
			}

			if err := m.StatementSummaries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statement_summaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statement_summaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Billing20201105ListStatementsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Billing20201105ListStatementsResponse) UnmarshalBinary(b []byte) error {
	var res Billing20201105ListStatementsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
