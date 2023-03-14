// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OperationState State is one of the states that an Operation can be in.
//
// The states are purposely coarse grained to make it easy to understand
// the operation state machine: pending => running => done. Or pending =>
// queued => running => done. No other state transitions are possible.
// Success/failure can be determined based on the result oneof.
//
// swagger:model OperationState
type OperationState string

func NewOperationState(value OperationState) *OperationState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OperationState.
func (m OperationState) Pointer() *OperationState {
	return &m
}

const (

	// OperationStatePENDING captures enum value "PENDING"
	OperationStatePENDING OperationState = "PENDING"

	// OperationStateRUNNING captures enum value "RUNNING"
	OperationStateRUNNING OperationState = "RUNNING"

	// OperationStateDONE captures enum value "DONE"
	OperationStateDONE OperationState = "DONE"

	// OperationStateQUEUED captures enum value "QUEUED"
	OperationStateQUEUED OperationState = "QUEUED"
)

// for schema
var operationStateEnum []interface{}

func init() {
	var res []OperationState
	if err := json.Unmarshal([]byte(`["PENDING","RUNNING","DONE","QUEUED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationStateEnum = append(operationStateEnum, v)
	}
}

func (m OperationState) validateOperationStateEnum(path, location string, value OperationState) error {
	if err := validate.EnumCase(path, location, value, operationStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this operation state
func (m OperationState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOperationStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this operation state based on context it is used
func (m OperationState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}