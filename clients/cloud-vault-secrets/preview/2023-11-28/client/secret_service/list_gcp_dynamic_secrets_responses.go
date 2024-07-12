// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// ListGcpDynamicSecretsReader is a Reader for the ListGcpDynamicSecrets structure.
type ListGcpDynamicSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGcpDynamicSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGcpDynamicSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGcpDynamicSecretsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGcpDynamicSecretsOK creates a ListGcpDynamicSecretsOK with default headers values
func NewListGcpDynamicSecretsOK() *ListGcpDynamicSecretsOK {
	return &ListGcpDynamicSecretsOK{}
}

/*
ListGcpDynamicSecretsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListGcpDynamicSecretsOK struct {
	Payload *models.Secrets20231128ListGcpDynamicSecretsResponse
}

// IsSuccess returns true when this list gcp dynamic secrets o k response has a 2xx status code
func (o *ListGcpDynamicSecretsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list gcp dynamic secrets o k response has a 3xx status code
func (o *ListGcpDynamicSecretsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list gcp dynamic secrets o k response has a 4xx status code
func (o *ListGcpDynamicSecretsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list gcp dynamic secrets o k response has a 5xx status code
func (o *ListGcpDynamicSecretsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list gcp dynamic secrets o k response a status code equal to that given
func (o *ListGcpDynamicSecretsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list gcp dynamic secrets o k response
func (o *ListGcpDynamicSecretsOK) Code() int {
	return 200
}

func (o *ListGcpDynamicSecretsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret][%d] listGcpDynamicSecretsOK  %+v", 200, o.Payload)
}

func (o *ListGcpDynamicSecretsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret][%d] listGcpDynamicSecretsOK  %+v", 200, o.Payload)
}

func (o *ListGcpDynamicSecretsOK) GetPayload() *models.Secrets20231128ListGcpDynamicSecretsResponse {
	return o.Payload
}

func (o *ListGcpDynamicSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128ListGcpDynamicSecretsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGcpDynamicSecretsDefault creates a ListGcpDynamicSecretsDefault with default headers values
func NewListGcpDynamicSecretsDefault(code int) *ListGcpDynamicSecretsDefault {
	return &ListGcpDynamicSecretsDefault{
		_statusCode: code,
	}
}

/*
ListGcpDynamicSecretsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListGcpDynamicSecretsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this list gcp dynamic secrets default response has a 2xx status code
func (o *ListGcpDynamicSecretsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list gcp dynamic secrets default response has a 3xx status code
func (o *ListGcpDynamicSecretsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list gcp dynamic secrets default response has a 4xx status code
func (o *ListGcpDynamicSecretsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list gcp dynamic secrets default response has a 5xx status code
func (o *ListGcpDynamicSecretsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list gcp dynamic secrets default response a status code equal to that given
func (o *ListGcpDynamicSecretsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list gcp dynamic secrets default response
func (o *ListGcpDynamicSecretsDefault) Code() int {
	return o._statusCode
}

func (o *ListGcpDynamicSecretsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret][%d] ListGcpDynamicSecrets default  %+v", o._statusCode, o.Payload)
}

func (o *ListGcpDynamicSecretsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/dynamic/gcp/secret][%d] ListGcpDynamicSecrets default  %+v", o._statusCode, o.Payload)
}

func (o *ListGcpDynamicSecretsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ListGcpDynamicSecretsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
