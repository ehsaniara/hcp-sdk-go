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

// LogService20210330ControlPlaneEvent ControlPlaneEvent captures an action that happened on the HCP control-plane.
//
// swagger:model log_service_20210330ControlPlaneEvent
type LogService20210330ControlPlaneEvent struct {

	// Action classifies the performed action.
	Action *LogService20210330Action `json:"action,omitempty"`

	// authentication_info holds information about the authenticated principal.
	AuthenticationInfo *LogService20210330AuthenticationInfo `json:"authentication_info,omitempty"`

	// authorization_info describes what permissions were required to perform
	// the action and whether the authenticated principal was granted these.
	AuthorizationInfo []*LogService20210330AuthorizationInfo `json:"authorization_info"`

	// description describes the event in a human comprehensible manner.
	Description string `json:"description,omitempty"`

	// metadata is action-specific data about the request, response, or other
	// information associated with the current audited event.
	Metadata interface{} `json:"metadata,omitempty"`

	// operation_info describes the operation this event is part of.
	OperationInfo *LogService20210330OperationInfo `json:"operation_info,omitempty"`

	// principal holds information about the principal that performed the action.
	Principal *CloudlogService20210330Principal `json:"principal,omitempty"`

	// request_info contains information about the request triggering this event.
	RequestInfo *LogService20210330RequestInfo `json:"request_info,omitempty"`

	// service_principal_delegation_chain store the chain of assumed principals
	// leading up to the current principal. The order of principals will be in the
	// order the identities were assumed.
	ServicePrincipalDelegationChain []*LogService20210330ServicePrincipalDelegationInfo `json:"service_principal_delegation_chain"`

	// status_code contains the status code the audited action resulted in. This
	// is the string representation of the event's status.code.
	StatusCode string `json:"status_code,omitempty"`
}

// Validate validates this log service 20210330 control plane event
func (m *LogService20210330ControlPlaneEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrincipal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePrincipalDelegationChain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330ControlPlaneEvent) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) validateAuthenticationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationInfo) { // not required
		return nil
	}

	if m.AuthenticationInfo != nil {
		if err := m.AuthenticationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication_info")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) validateAuthorizationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthorizationInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.AuthorizationInfo); i++ {
		if swag.IsZero(m.AuthorizationInfo[i]) { // not required
			continue
		}

		if m.AuthorizationInfo[i] != nil {
			if err := m.AuthorizationInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authorization_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("authorization_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) validateOperationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.OperationInfo) { // not required
		return nil
	}

	if m.OperationInfo != nil {
		if err := m.OperationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation_info")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) validatePrincipal(formats strfmt.Registry) error {
	if swag.IsZero(m.Principal) { // not required
		return nil
	}

	if m.Principal != nil {
		if err := m.Principal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("principal")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) validateRequestInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestInfo) { // not required
		return nil
	}

	if m.RequestInfo != nil {
		if err := m.RequestInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request_info")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) validateServicePrincipalDelegationChain(formats strfmt.Registry) error {
	if swag.IsZero(m.ServicePrincipalDelegationChain) { // not required
		return nil
	}

	for i := 0; i < len(m.ServicePrincipalDelegationChain); i++ {
		if swag.IsZero(m.ServicePrincipalDelegationChain[i]) { // not required
			continue
		}

		if m.ServicePrincipalDelegationChain[i] != nil {
			if err := m.ServicePrincipalDelegationChain[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_principal_delegation_chain" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_principal_delegation_chain" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this log service 20210330 control plane event based on the context it is used
func (m *LogService20210330ControlPlaneEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthenticationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthorizationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrincipal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServicePrincipalDelegationChain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService20210330ControlPlaneEvent) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {

		if swag.IsZero(m.Action) { // not required
			return nil
		}

		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) contextValidateAuthenticationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthenticationInfo != nil {

		if swag.IsZero(m.AuthenticationInfo) { // not required
			return nil
		}

		if err := m.AuthenticationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication_info")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) contextValidateAuthorizationInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuthorizationInfo); i++ {

		if m.AuthorizationInfo[i] != nil {

			if swag.IsZero(m.AuthorizationInfo[i]) { // not required
				return nil
			}

			if err := m.AuthorizationInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authorization_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("authorization_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) contextValidateOperationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.OperationInfo != nil {

		if swag.IsZero(m.OperationInfo) { // not required
			return nil
		}

		if err := m.OperationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation_info")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) contextValidatePrincipal(ctx context.Context, formats strfmt.Registry) error {

	if m.Principal != nil {

		if swag.IsZero(m.Principal) { // not required
			return nil
		}

		if err := m.Principal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("principal")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) contextValidateRequestInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestInfo != nil {

		if swag.IsZero(m.RequestInfo) { // not required
			return nil
		}

		if err := m.RequestInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request_info")
			}
			return err
		}
	}

	return nil
}

func (m *LogService20210330ControlPlaneEvent) contextValidateServicePrincipalDelegationChain(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServicePrincipalDelegationChain); i++ {

		if m.ServicePrincipalDelegationChain[i] != nil {

			if swag.IsZero(m.ServicePrincipalDelegationChain[i]) { // not required
				return nil
			}

			if err := m.ServicePrincipalDelegationChain[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_principal_delegation_chain" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_principal_delegation_chain" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogService20210330ControlPlaneEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogService20210330ControlPlaneEvent) UnmarshalBinary(b []byte) error {
	var res LogService20210330ControlPlaneEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}