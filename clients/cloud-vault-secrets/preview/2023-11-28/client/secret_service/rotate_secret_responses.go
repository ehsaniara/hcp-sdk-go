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

// RotateSecretReader is a Reader for the RotateSecret structure.
type RotateSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRotateSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRotateSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRotateSecretOK creates a RotateSecretOK with default headers values
func NewRotateSecretOK() *RotateSecretOK {
	return &RotateSecretOK{}
}

/*
RotateSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type RotateSecretOK struct {
	Payload models.Secrets20231128RotateSecretResponse
}

// IsSuccess returns true when this rotate secret o k response has a 2xx status code
func (o *RotateSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rotate secret o k response has a 3xx status code
func (o *RotateSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rotate secret o k response has a 4xx status code
func (o *RotateSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rotate secret o k response has a 5xx status code
func (o *RotateSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rotate secret o k response a status code equal to that given
func (o *RotateSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rotate secret o k response
func (o *RotateSecretOK) Code() int {
	return 200
}

func (o *RotateSecretOK) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}:rotate][%d] rotateSecretOK  %+v", 200, o.Payload)
}

func (o *RotateSecretOK) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}:rotate][%d] rotateSecretOK  %+v", 200, o.Payload)
}

func (o *RotateSecretOK) GetPayload() models.Secrets20231128RotateSecretResponse {
	return o.Payload
}

func (o *RotateSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateSecretDefault creates a RotateSecretDefault with default headers values
func NewRotateSecretDefault(code int) *RotateSecretDefault {
	return &RotateSecretDefault{
		_statusCode: code,
	}
}

/*
RotateSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RotateSecretDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this rotate secret default response has a 2xx status code
func (o *RotateSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rotate secret default response has a 3xx status code
func (o *RotateSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rotate secret default response has a 4xx status code
func (o *RotateSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rotate secret default response has a 5xx status code
func (o *RotateSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rotate secret default response a status code equal to that given
func (o *RotateSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rotate secret default response
func (o *RotateSecretDefault) Code() int {
	return o._statusCode
}

func (o *RotateSecretDefault) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}:rotate][%d] RotateSecret default  %+v", o._statusCode, o.Payload)
}

func (o *RotateSecretDefault) String() string {
	return fmt.Sprintf("[POST /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secrets/{secret_name}:rotate][%d] RotateSecret default  %+v", o._statusCode, o.Payload)
}

func (o *RotateSecretDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *RotateSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
