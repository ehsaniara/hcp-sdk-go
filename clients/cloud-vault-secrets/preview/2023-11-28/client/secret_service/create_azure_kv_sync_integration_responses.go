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

// CreateAzureKvSyncIntegrationReader is a Reader for the CreateAzureKvSyncIntegration structure.
type CreateAzureKvSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAzureKvSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAzureKvSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAzureKvSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAzureKvSyncIntegrationOK creates a CreateAzureKvSyncIntegrationOK with default headers values
func NewCreateAzureKvSyncIntegrationOK() *CreateAzureKvSyncIntegrationOK {
	return &CreateAzureKvSyncIntegrationOK{}
}

/*
CreateAzureKvSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAzureKvSyncIntegrationOK struct {
	Payload *models.Secrets20231128CreateSyncIntegrationResponse
}

// IsSuccess returns true when this create azure kv sync integration o k response has a 2xx status code
func (o *CreateAzureKvSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create azure kv sync integration o k response has a 3xx status code
func (o *CreateAzureKvSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure kv sync integration o k response has a 4xx status code
func (o *CreateAzureKvSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create azure kv sync integration o k response has a 5xx status code
func (o *CreateAzureKvSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create azure kv sync integration o k response a status code equal to that given
func (o *CreateAzureKvSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create azure kv sync integration o k response
func (o *CreateAzureKvSyncIntegrationOK) Code() int {
	return 200
}

func (o *CreateAzureKvSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/azure-kv][%d] createAzureKvSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *CreateAzureKvSyncIntegrationOK) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/azure-kv][%d] createAzureKvSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *CreateAzureKvSyncIntegrationOK) GetPayload() *models.Secrets20231128CreateSyncIntegrationResponse {
	return o.Payload
}

func (o *CreateAzureKvSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureKvSyncIntegrationDefault creates a CreateAzureKvSyncIntegrationDefault with default headers values
func NewCreateAzureKvSyncIntegrationDefault(code int) *CreateAzureKvSyncIntegrationDefault {
	return &CreateAzureKvSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
CreateAzureKvSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateAzureKvSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create azure kv sync integration default response has a 2xx status code
func (o *CreateAzureKvSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create azure kv sync integration default response has a 3xx status code
func (o *CreateAzureKvSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create azure kv sync integration default response has a 4xx status code
func (o *CreateAzureKvSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create azure kv sync integration default response has a 5xx status code
func (o *CreateAzureKvSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create azure kv sync integration default response a status code equal to that given
func (o *CreateAzureKvSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create azure kv sync integration default response
func (o *CreateAzureKvSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateAzureKvSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/azure-kv][%d] CreateAzureKvSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAzureKvSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/azure-kv][%d] CreateAzureKvSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAzureKvSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateAzureKvSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateAzureKvSyncIntegrationBody create azure kv sync integration body
swagger:model CreateAzureKvSyncIntegrationBody
*/
type CreateAzureKvSyncIntegrationBody struct {

	// azure kv connection details
	AzureKvConnectionDetails *models.Secrets20231128AzureKvConnectionDetailsRequest `json:"azure_kv_connection_details,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create azure kv sync integration body
func (o *CreateAzureKvSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAzureKvConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAzureKvSyncIntegrationBody) validateAzureKvConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.AzureKvConnectionDetails) { // not required
		return nil
	}

	if o.AzureKvConnectionDetails != nil {
		if err := o.AzureKvConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "azure_kv_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "azure_kv_connection_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create azure kv sync integration body based on the context it is used
func (o *CreateAzureKvSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAzureKvConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAzureKvSyncIntegrationBody) contextValidateAzureKvConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if o.AzureKvConnectionDetails != nil {

		if swag.IsZero(o.AzureKvConnectionDetails) { // not required
			return nil
		}

		if err := o.AzureKvConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "azure_kv_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "azure_kv_connection_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateAzureKvSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAzureKvSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res CreateAzureKvSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
