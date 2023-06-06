// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-06-13/models"
)

// UpsertSyncInstallationReader is a Reader for the UpsertSyncInstallation structure.
type UpsertSyncInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpsertSyncInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpsertSyncInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpsertSyncInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpsertSyncInstallationOK creates a UpsertSyncInstallationOK with default headers values
func NewUpsertSyncInstallationOK() *UpsertSyncInstallationOK {
	return &UpsertSyncInstallationOK{}
}

/*
UpsertSyncInstallationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpsertSyncInstallationOK struct {
	Payload *models.Secrets20230613UpsertSyncInstallationResponse
}

// IsSuccess returns true when this upsert sync installation o k response has a 2xx status code
func (o *UpsertSyncInstallationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upsert sync installation o k response has a 3xx status code
func (o *UpsertSyncInstallationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upsert sync installation o k response has a 4xx status code
func (o *UpsertSyncInstallationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upsert sync installation o k response has a 5xx status code
func (o *UpsertSyncInstallationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upsert sync installation o k response a status code equal to that given
func (o *UpsertSyncInstallationOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpsertSyncInstallationOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/installations][%d] upsertSyncInstallationOK  %+v", 200, o.Payload)
}

func (o *UpsertSyncInstallationOK) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/installations][%d] upsertSyncInstallationOK  %+v", 200, o.Payload)
}

func (o *UpsertSyncInstallationOK) GetPayload() *models.Secrets20230613UpsertSyncInstallationResponse {
	return o.Payload
}

func (o *UpsertSyncInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613UpsertSyncInstallationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpsertSyncInstallationDefault creates a UpsertSyncInstallationDefault with default headers values
func NewUpsertSyncInstallationDefault(code int) *UpsertSyncInstallationDefault {
	return &UpsertSyncInstallationDefault{
		_statusCode: code,
	}
}

/*
UpsertSyncInstallationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpsertSyncInstallationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the upsert sync installation default response
func (o *UpsertSyncInstallationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this upsert sync installation default response has a 2xx status code
func (o *UpsertSyncInstallationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upsert sync installation default response has a 3xx status code
func (o *UpsertSyncInstallationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upsert sync installation default response has a 4xx status code
func (o *UpsertSyncInstallationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upsert sync installation default response has a 5xx status code
func (o *UpsertSyncInstallationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upsert sync installation default response a status code equal to that given
func (o *UpsertSyncInstallationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpsertSyncInstallationDefault) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/installations][%d] UpsertSyncInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertSyncInstallationDefault) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/installations][%d] UpsertSyncInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertSyncInstallationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *UpsertSyncInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpsertSyncInstallationBody upsert sync installation body
swagger:model UpsertSyncInstallationBody
*/
type UpsertSyncInstallationBody struct {

	// aws external id
	AwsExternalID string `json:"aws_external_id,omitempty"`

	// gh installation id
	GhInstallationID string `json:"gh_installation_id,omitempty"`

	// location
	Location *UpsertSyncInstallationParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this upsert sync installation body
func (o *UpsertSyncInstallationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpsertSyncInstallationBody) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.Location) { // not required
		return nil
	}

	if o.Location != nil {
		if err := o.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upsert sync installation body based on the context it is used
func (o *UpsertSyncInstallationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpsertSyncInstallationBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.Location != nil {
		if err := o.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpsertSyncInstallationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpsertSyncInstallationBody) UnmarshalBinary(b []byte) error {
	var res UpsertSyncInstallationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpsertSyncInstallationParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model UpsertSyncInstallationParamsBodyLocation
*/
type UpsertSyncInstallationParamsBodyLocation struct {

	// region
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this upsert sync installation params body location
func (o *UpsertSyncInstallationParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpsertSyncInstallationParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upsert sync installation params body location based on the context it is used
func (o *UpsertSyncInstallationParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpsertSyncInstallationParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {
		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpsertSyncInstallationParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpsertSyncInstallationParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res UpsertSyncInstallationParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
