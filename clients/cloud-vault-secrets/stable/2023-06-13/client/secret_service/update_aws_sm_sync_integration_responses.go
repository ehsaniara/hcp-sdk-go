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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/models"
)

// UpdateAwsSmSyncIntegrationReader is a Reader for the UpdateAwsSmSyncIntegration structure.
type UpdateAwsSmSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAwsSmSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAwsSmSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAwsSmSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAwsSmSyncIntegrationOK creates a UpdateAwsSmSyncIntegrationOK with default headers values
func NewUpdateAwsSmSyncIntegrationOK() *UpdateAwsSmSyncIntegrationOK {
	return &UpdateAwsSmSyncIntegrationOK{}
}

/*
UpdateAwsSmSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateAwsSmSyncIntegrationOK struct {
	Payload *models.Secrets20230613UpdateSyncIntegrationResponse
}

// IsSuccess returns true when this update aws sm sync integration o k response has a 2xx status code
func (o *UpdateAwsSmSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update aws sm sync integration o k response has a 3xx status code
func (o *UpdateAwsSmSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update aws sm sync integration o k response has a 4xx status code
func (o *UpdateAwsSmSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update aws sm sync integration o k response has a 5xx status code
func (o *UpdateAwsSmSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update aws sm sync integration o k response a status code equal to that given
func (o *UpdateAwsSmSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update aws sm sync integration o k response
func (o *UpdateAwsSmSyncIntegrationOK) Code() int {
	return 200
}

func (o *UpdateAwsSmSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[PATCH /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/aws-sm/{name}][%d] updateAwsSmSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateAwsSmSyncIntegrationOK) String() string {
	return fmt.Sprintf("[PATCH /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/aws-sm/{name}][%d] updateAwsSmSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateAwsSmSyncIntegrationOK) GetPayload() *models.Secrets20230613UpdateSyncIntegrationResponse {
	return o.Payload
}

func (o *UpdateAwsSmSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613UpdateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAwsSmSyncIntegrationDefault creates a UpdateAwsSmSyncIntegrationDefault with default headers values
func NewUpdateAwsSmSyncIntegrationDefault(code int) *UpdateAwsSmSyncIntegrationDefault {
	return &UpdateAwsSmSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
UpdateAwsSmSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateAwsSmSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this update aws sm sync integration default response has a 2xx status code
func (o *UpdateAwsSmSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update aws sm sync integration default response has a 3xx status code
func (o *UpdateAwsSmSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update aws sm sync integration default response has a 4xx status code
func (o *UpdateAwsSmSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update aws sm sync integration default response has a 5xx status code
func (o *UpdateAwsSmSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update aws sm sync integration default response a status code equal to that given
func (o *UpdateAwsSmSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update aws sm sync integration default response
func (o *UpdateAwsSmSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAwsSmSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[PATCH /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/aws-sm/{name}][%d] UpdateAwsSmSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAwsSmSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[PATCH /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/aws-sm/{name}][%d] UpdateAwsSmSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAwsSmSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *UpdateAwsSmSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateAwsSmSyncIntegrationBody update aws sm sync integration body
swagger:model UpdateAwsSmSyncIntegrationBody
*/
type UpdateAwsSmSyncIntegrationBody struct {

	// aws sm connection details
	AwsSmConnectionDetails *models.Secrets20230613AwsSmConnectionDetailsRequest `json:"aws_sm_connection_details,omitempty"`

	// location
	Location *UpdateAwsSmSyncIntegrationParamsBodyLocation `json:"location,omitempty"`
}

// Validate validates this update aws sm sync integration body
func (o *UpdateAwsSmSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAwsSmConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAwsSmSyncIntegrationBody) validateAwsSmConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.AwsSmConnectionDetails) { // not required
		return nil
	}

	if o.AwsSmConnectionDetails != nil {
		if err := o.AwsSmConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "aws_sm_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "aws_sm_connection_details")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateAwsSmSyncIntegrationBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this update aws sm sync integration body based on the context it is used
func (o *UpdateAwsSmSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAwsSmConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAwsSmSyncIntegrationBody) contextValidateAwsSmConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if o.AwsSmConnectionDetails != nil {

		if swag.IsZero(o.AwsSmConnectionDetails) { // not required
			return nil
		}

		if err := o.AwsSmConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "aws_sm_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "aws_sm_connection_details")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateAwsSmSyncIntegrationBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.Location != nil {

		if swag.IsZero(o.Location) { // not required
			return nil
		}

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
func (o *UpdateAwsSmSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAwsSmSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res UpdateAwsSmSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateAwsSmSyncIntegrationParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model UpdateAwsSmSyncIntegrationParamsBodyLocation
*/
type UpdateAwsSmSyncIntegrationParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this update aws sm sync integration params body location
func (o *UpdateAwsSmSyncIntegrationParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAwsSmSyncIntegrationParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this update aws sm sync integration params body location based on the context it is used
func (o *UpdateAwsSmSyncIntegrationParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAwsSmSyncIntegrationParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {

		if swag.IsZero(o.Region) { // not required
			return nil
		}

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
func (o *UpdateAwsSmSyncIntegrationParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAwsSmSyncIntegrationParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res UpdateAwsSmSyncIntegrationParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
