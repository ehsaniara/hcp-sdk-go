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

// UpsertVercelProjectSyncIntegrationReader is a Reader for the UpsertVercelProjectSyncIntegration structure.
type UpsertVercelProjectSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpsertVercelProjectSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpsertVercelProjectSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpsertVercelProjectSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpsertVercelProjectSyncIntegrationOK creates a UpsertVercelProjectSyncIntegrationOK with default headers values
func NewUpsertVercelProjectSyncIntegrationOK() *UpsertVercelProjectSyncIntegrationOK {
	return &UpsertVercelProjectSyncIntegrationOK{}
}

/*
UpsertVercelProjectSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpsertVercelProjectSyncIntegrationOK struct {
	Payload *models.Secrets20230613UpsertSyncIntegrationResponse
}

// IsSuccess returns true when this upsert vercel project sync integration o k response has a 2xx status code
func (o *UpsertVercelProjectSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upsert vercel project sync integration o k response has a 3xx status code
func (o *UpsertVercelProjectSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upsert vercel project sync integration o k response has a 4xx status code
func (o *UpsertVercelProjectSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upsert vercel project sync integration o k response has a 5xx status code
func (o *UpsertVercelProjectSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upsert vercel project sync integration o k response a status code equal to that given
func (o *UpsertVercelProjectSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upsert vercel project sync integration o k response
func (o *UpsertVercelProjectSyncIntegrationOK) Code() int {
	return 200
}

func (o *UpsertVercelProjectSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel-project][%d] upsertVercelProjectSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpsertVercelProjectSyncIntegrationOK) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel-project][%d] upsertVercelProjectSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpsertVercelProjectSyncIntegrationOK) GetPayload() *models.Secrets20230613UpsertSyncIntegrationResponse {
	return o.Payload
}

func (o *UpsertVercelProjectSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613UpsertSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpsertVercelProjectSyncIntegrationDefault creates a UpsertVercelProjectSyncIntegrationDefault with default headers values
func NewUpsertVercelProjectSyncIntegrationDefault(code int) *UpsertVercelProjectSyncIntegrationDefault {
	return &UpsertVercelProjectSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
UpsertVercelProjectSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpsertVercelProjectSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this upsert vercel project sync integration default response has a 2xx status code
func (o *UpsertVercelProjectSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upsert vercel project sync integration default response has a 3xx status code
func (o *UpsertVercelProjectSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upsert vercel project sync integration default response has a 4xx status code
func (o *UpsertVercelProjectSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upsert vercel project sync integration default response has a 5xx status code
func (o *UpsertVercelProjectSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upsert vercel project sync integration default response a status code equal to that given
func (o *UpsertVercelProjectSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the upsert vercel project sync integration default response
func (o *UpsertVercelProjectSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *UpsertVercelProjectSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[PUT /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel-project][%d] UpsertVercelProjectSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertVercelProjectSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[PUT /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/vercel-project][%d] UpsertVercelProjectSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertVercelProjectSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *UpsertVercelProjectSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpsertVercelProjectSyncIntegrationBody upsert vercel project sync integration body
swagger:model UpsertVercelProjectSyncIntegrationBody
*/
type UpsertVercelProjectSyncIntegrationBody struct {

	// location
	Location *UpsertVercelProjectSyncIntegrationParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// vercel project connection details
	VercelProjectConnectionDetails *models.Secrets20230613VercelProjectConnectionDetailsRequest `json:"vercel_project_connection_details,omitempty"`
}

// Validate validates this upsert vercel project sync integration body
func (o *UpsertVercelProjectSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVercelProjectConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpsertVercelProjectSyncIntegrationBody) validateLocation(formats strfmt.Registry) error {
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

func (o *UpsertVercelProjectSyncIntegrationBody) validateVercelProjectConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.VercelProjectConnectionDetails) { // not required
		return nil
	}

	if o.VercelProjectConnectionDetails != nil {
		if err := o.VercelProjectConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "vercel_project_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "vercel_project_connection_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upsert vercel project sync integration body based on the context it is used
func (o *UpsertVercelProjectSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVercelProjectConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpsertVercelProjectSyncIntegrationBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

func (o *UpsertVercelProjectSyncIntegrationBody) contextValidateVercelProjectConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if o.VercelProjectConnectionDetails != nil {

		if swag.IsZero(o.VercelProjectConnectionDetails) { // not required
			return nil
		}

		if err := o.VercelProjectConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "vercel_project_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "vercel_project_connection_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpsertVercelProjectSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpsertVercelProjectSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res UpsertVercelProjectSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpsertVercelProjectSyncIntegrationParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model UpsertVercelProjectSyncIntegrationParamsBodyLocation
*/
type UpsertVercelProjectSyncIntegrationParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this upsert vercel project sync integration params body location
func (o *UpsertVercelProjectSyncIntegrationParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpsertVercelProjectSyncIntegrationParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this upsert vercel project sync integration params body location based on the context it is used
func (o *UpsertVercelProjectSyncIntegrationParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpsertVercelProjectSyncIntegrationParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UpsertVercelProjectSyncIntegrationParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpsertVercelProjectSyncIntegrationParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res UpsertVercelProjectSyncIntegrationParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
