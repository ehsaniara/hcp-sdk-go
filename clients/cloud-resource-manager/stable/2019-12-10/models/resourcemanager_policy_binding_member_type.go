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

// ResourcemanagerPolicyBindingMemberType PolicyBindingMemberType is the type of the policy binding member.
//
// swagger:model resourcemanagerPolicyBindingMemberType
type ResourcemanagerPolicyBindingMemberType string

func NewResourcemanagerPolicyBindingMemberType(value ResourcemanagerPolicyBindingMemberType) *ResourcemanagerPolicyBindingMemberType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ResourcemanagerPolicyBindingMemberType.
func (m ResourcemanagerPolicyBindingMemberType) Pointer() *ResourcemanagerPolicyBindingMemberType {
	return &m
}

const (

	// ResourcemanagerPolicyBindingMemberTypeUNSET captures enum value "UNSET"
	ResourcemanagerPolicyBindingMemberTypeUNSET ResourcemanagerPolicyBindingMemberType = "UNSET"

	// ResourcemanagerPolicyBindingMemberTypeUSER captures enum value "USER"
	ResourcemanagerPolicyBindingMemberTypeUSER ResourcemanagerPolicyBindingMemberType = "USER"

	// ResourcemanagerPolicyBindingMemberTypeGROUP captures enum value "GROUP"
	ResourcemanagerPolicyBindingMemberTypeGROUP ResourcemanagerPolicyBindingMemberType = "GROUP"

	// ResourcemanagerPolicyBindingMemberTypeSERVICEPRINCIPAL captures enum value "SERVICE_PRINCIPAL"
	ResourcemanagerPolicyBindingMemberTypeSERVICEPRINCIPAL ResourcemanagerPolicyBindingMemberType = "SERVICE_PRINCIPAL"
)

// for schema
var resourcemanagerPolicyBindingMemberTypeEnum []interface{}

func init() {
	var res []ResourcemanagerPolicyBindingMemberType
	if err := json.Unmarshal([]byte(`["UNSET","USER","GROUP","SERVICE_PRINCIPAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourcemanagerPolicyBindingMemberTypeEnum = append(resourcemanagerPolicyBindingMemberTypeEnum, v)
	}
}

func (m ResourcemanagerPolicyBindingMemberType) validateResourcemanagerPolicyBindingMemberTypeEnum(path, location string, value ResourcemanagerPolicyBindingMemberType) error {
	if err := validate.EnumCase(path, location, value, resourcemanagerPolicyBindingMemberTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this resourcemanager policy binding member type
func (m ResourcemanagerPolicyBindingMemberType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResourcemanagerPolicyBindingMemberTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this resourcemanager policy binding member type based on context it is used
func (m ResourcemanagerPolicyBindingMemberType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
