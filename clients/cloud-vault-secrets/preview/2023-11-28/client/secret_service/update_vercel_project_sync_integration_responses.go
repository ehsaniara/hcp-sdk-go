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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// UpdateVercelProjectSyncIntegrationReader is a Reader for the UpdateVercelProjectSyncIntegration structure.
type UpdateVercelProjectSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVercelProjectSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVercelProjectSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateVercelProjectSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateVercelProjectSyncIntegrationOK creates a UpdateVercelProjectSyncIntegrationOK with default headers values
func NewUpdateVercelProjectSyncIntegrationOK() *UpdateVercelProjectSyncIntegrationOK {
	return &UpdateVercelProjectSyncIntegrationOK{}
}

/*
UpdateVercelProjectSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateVercelProjectSyncIntegrationOK struct {
	Payload *models.Secrets20231128UpdateSyncIntegrationResponse
}

// IsSuccess returns true when this update vercel project sync integration o k response has a 2xx status code
func (o *UpdateVercelProjectSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update vercel project sync integration o k response has a 3xx status code
func (o *UpdateVercelProjectSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update vercel project sync integration o k response has a 4xx status code
func (o *UpdateVercelProjectSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update vercel project sync integration o k response has a 5xx status code
func (o *UpdateVercelProjectSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update vercel project sync integration o k response a status code equal to that given
func (o *UpdateVercelProjectSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update vercel project sync integration o k response
func (o *UpdateVercelProjectSyncIntegrationOK) Code() int {
	return 200
}

func (o *UpdateVercelProjectSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[PATCH /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/vercel-project/{name}][%d] updateVercelProjectSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateVercelProjectSyncIntegrationOK) String() string {
	return fmt.Sprintf("[PATCH /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/vercel-project/{name}][%d] updateVercelProjectSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateVercelProjectSyncIntegrationOK) GetPayload() *models.Secrets20231128UpdateSyncIntegrationResponse {
	return o.Payload
}

func (o *UpdateVercelProjectSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128UpdateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVercelProjectSyncIntegrationDefault creates a UpdateVercelProjectSyncIntegrationDefault with default headers values
func NewUpdateVercelProjectSyncIntegrationDefault(code int) *UpdateVercelProjectSyncIntegrationDefault {
	return &UpdateVercelProjectSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
UpdateVercelProjectSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateVercelProjectSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this update vercel project sync integration default response has a 2xx status code
func (o *UpdateVercelProjectSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update vercel project sync integration default response has a 3xx status code
func (o *UpdateVercelProjectSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update vercel project sync integration default response has a 4xx status code
func (o *UpdateVercelProjectSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update vercel project sync integration default response has a 5xx status code
func (o *UpdateVercelProjectSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update vercel project sync integration default response a status code equal to that given
func (o *UpdateVercelProjectSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update vercel project sync integration default response
func (o *UpdateVercelProjectSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVercelProjectSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[PATCH /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/vercel-project/{name}][%d] UpdateVercelProjectSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateVercelProjectSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[PATCH /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/vercel-project/{name}][%d] UpdateVercelProjectSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateVercelProjectSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *UpdateVercelProjectSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateVercelProjectSyncIntegrationBody update vercel project sync integration body
swagger:model UpdateVercelProjectSyncIntegrationBody
*/
type UpdateVercelProjectSyncIntegrationBody struct {

	// vercel project connection details
	VercelProjectConnectionDetails *models.Secrets20231128VercelProjectConnectionDetailsRequest `json:"vercel_project_connection_details,omitempty"`
}

// Validate validates this update vercel project sync integration body
func (o *UpdateVercelProjectSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVercelProjectConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVercelProjectSyncIntegrationBody) validateVercelProjectConnectionDetails(formats strfmt.Registry) error {
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

// ContextValidate validate this update vercel project sync integration body based on the context it is used
func (o *UpdateVercelProjectSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVercelProjectConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVercelProjectSyncIntegrationBody) contextValidateVercelProjectConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UpdateVercelProjectSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVercelProjectSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res UpdateVercelProjectSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
