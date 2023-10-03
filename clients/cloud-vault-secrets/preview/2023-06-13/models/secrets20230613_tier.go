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

// Secrets20230613Tier secrets 20230613 tier
//
// swagger:model secrets_20230613Tier
type Secrets20230613Tier string

func NewSecrets20230613Tier(value Secrets20230613Tier) *Secrets20230613Tier {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Secrets20230613Tier.
func (m Secrets20230613Tier) Pointer() *Secrets20230613Tier {
	return &m
}

const (

	// Secrets20230613TierUNKNOWN captures enum value "UNKNOWN"
	Secrets20230613TierUNKNOWN Secrets20230613Tier = "UNKNOWN"

	// Secrets20230613TierFREE captures enum value "FREE"
	Secrets20230613TierFREE Secrets20230613Tier = "FREE"

	// Secrets20230613TierSTANDARD captures enum value "STANDARD"
	Secrets20230613TierSTANDARD Secrets20230613Tier = "STANDARD"
)

// for schema
var secrets20230613TierEnum []interface{}

func init() {
	var res []Secrets20230613Tier
	if err := json.Unmarshal([]byte(`["UNKNOWN","FREE","STANDARD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		secrets20230613TierEnum = append(secrets20230613TierEnum, v)
	}
}

func (m Secrets20230613Tier) validateSecrets20230613TierEnum(path, location string, value Secrets20230613Tier) error {
	if err := validate.EnumCase(path, location, value, secrets20230613TierEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this secrets 20230613 tier
func (m Secrets20230613Tier) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSecrets20230613TierEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this secrets 20230613 tier based on context it is used
func (m Secrets20230613Tier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
