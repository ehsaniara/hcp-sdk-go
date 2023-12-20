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

// UserPrincipalIdentityType IdentityType contains the identity types for users.
//
//   - UNSET: UNSET is the default value, should not be used.
//   - SOCIAL_GITHUB: SOCIAL_GITHUB is the GitHub Social login provider.
//   - EMAIL_PASSWORD: EMAIL_PASSWORD is the auth0 email/password database provider.
//   - SAMLP: SAMLP is the auth0 SAML provider.
//   - HASHICORP_SSO: HASHICORP_SSO is the HashiCorp SSO provider.
//
// swagger:model UserPrincipalIdentityType
type UserPrincipalIdentityType string

func NewUserPrincipalIdentityType(value UserPrincipalIdentityType) *UserPrincipalIdentityType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UserPrincipalIdentityType.
func (m UserPrincipalIdentityType) Pointer() *UserPrincipalIdentityType {
	return &m
}

const (

	// UserPrincipalIdentityTypeUNSET captures enum value "UNSET"
	UserPrincipalIdentityTypeUNSET UserPrincipalIdentityType = "UNSET"

	// UserPrincipalIdentityTypeSOCIALGITHUB captures enum value "SOCIAL_GITHUB"
	UserPrincipalIdentityTypeSOCIALGITHUB UserPrincipalIdentityType = "SOCIAL_GITHUB"

	// UserPrincipalIdentityTypeEMAILPASSWORD captures enum value "EMAIL_PASSWORD"
	UserPrincipalIdentityTypeEMAILPASSWORD UserPrincipalIdentityType = "EMAIL_PASSWORD"

	// UserPrincipalIdentityTypeSAMLP captures enum value "SAMLP"
	UserPrincipalIdentityTypeSAMLP UserPrincipalIdentityType = "SAMLP"

	// UserPrincipalIdentityTypeHASHICORPSSO captures enum value "HASHICORP_SSO"
	UserPrincipalIdentityTypeHASHICORPSSO UserPrincipalIdentityType = "HASHICORP_SSO"
)

// for schema
var userPrincipalIdentityTypeEnum []interface{}

func init() {
	var res []UserPrincipalIdentityType
	if err := json.Unmarshal([]byte(`["UNSET","SOCIAL_GITHUB","EMAIL_PASSWORD","SAMLP","HASHICORP_SSO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userPrincipalIdentityTypeEnum = append(userPrincipalIdentityTypeEnum, v)
	}
}

func (m UserPrincipalIdentityType) validateUserPrincipalIdentityTypeEnum(path, location string, value UserPrincipalIdentityType) error {
	if err := validate.EnumCase(path, location, value, userPrincipalIdentityTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user principal identity type
func (m UserPrincipalIdentityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserPrincipalIdentityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user principal identity type based on context it is used
func (m UserPrincipalIdentityType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}