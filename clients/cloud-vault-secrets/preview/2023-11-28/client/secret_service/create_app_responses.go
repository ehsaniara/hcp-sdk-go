// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// CreateAppReader is a Reader for the CreateApp structure.
type CreateAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAppOK creates a CreateAppOK with default headers values
func NewCreateAppOK() *CreateAppOK {
	return &CreateAppOK{}
}

/*
CreateAppOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAppOK struct {
	Payload *models.Secrets20231128CreateAppResponse
}

// IsSuccess returns true when this create app o k response has a 2xx status code
func (o *CreateAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app o k response has a 3xx status code
func (o *CreateAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app o k response has a 4xx status code
func (o *CreateAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app o k response has a 5xx status code
func (o *CreateAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create app o k response a status code equal to that given
func (o *CreateAppOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create app o k response
func (o *CreateAppOK) Code() int {
	return 200
}

func (o *CreateAppOK) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps][%d] createAppOK  %+v", 200, o.Payload)
}

func (o *CreateAppOK) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps][%d] createAppOK  %+v", 200, o.Payload)
}

func (o *CreateAppOK) GetPayload() *models.Secrets20231128CreateAppResponse {
	return o.Payload
}

func (o *CreateAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128CreateAppResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppDefault creates a CreateAppDefault with default headers values
func NewCreateAppDefault(code int) *CreateAppDefault {
	return &CreateAppDefault{
		_statusCode: code,
	}
}

/*
CreateAppDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateAppDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this create app default response has a 2xx status code
func (o *CreateAppDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create app default response has a 3xx status code
func (o *CreateAppDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create app default response has a 4xx status code
func (o *CreateAppDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create app default response has a 5xx status code
func (o *CreateAppDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create app default response a status code equal to that given
func (o *CreateAppDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create app default response
func (o *CreateAppDefault) Code() int {
	return o._statusCode
}

func (o *CreateAppDefault) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps][%d] CreateApp default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAppDefault) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps][%d] CreateApp default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAppDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateAppBody create app body
swagger:model CreateAppBody
*/
type CreateAppBody struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sync integrations
	SyncIntegrations []string `json:"sync_integrations"`
}

// Validate validates this create app body
func (o *CreateAppBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create app body based on context it is used
func (o *CreateAppBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateAppBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAppBody) UnmarshalBinary(b []byte) error {
	var res CreateAppBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}