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

// DeleteAppReader is a Reader for the DeleteApp structure.
type DeleteAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAppOK creates a DeleteAppOK with default headers values
func NewDeleteAppOK() *DeleteAppOK {
	return &DeleteAppOK{}
}

/*
DeleteAppOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteAppOK struct {
	Payload models.Secrets20231128DeleteAppResponse
}

// IsSuccess returns true when this delete app o k response has a 2xx status code
func (o *DeleteAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete app o k response has a 3xx status code
func (o *DeleteAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete app o k response has a 4xx status code
func (o *DeleteAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete app o k response has a 5xx status code
func (o *DeleteAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete app o k response a status code equal to that given
func (o *DeleteAppOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete app o k response
func (o *DeleteAppOK) Code() int {
	return 200
}

func (o *DeleteAppOK) Error() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{name}][%d] deleteAppOK  %+v", 200, o.Payload)
}

func (o *DeleteAppOK) String() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{name}][%d] deleteAppOK  %+v", 200, o.Payload)
}

func (o *DeleteAppOK) GetPayload() models.Secrets20231128DeleteAppResponse {
	return o.Payload
}

func (o *DeleteAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppDefault creates a DeleteAppDefault with default headers values
func NewDeleteAppDefault(code int) *DeleteAppDefault {
	return &DeleteAppDefault{
		_statusCode: code,
	}
}

/*
DeleteAppDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteAppDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete app default response has a 2xx status code
func (o *DeleteAppDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete app default response has a 3xx status code
func (o *DeleteAppDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete app default response has a 4xx status code
func (o *DeleteAppDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete app default response has a 5xx status code
func (o *DeleteAppDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete app default response a status code equal to that given
func (o *DeleteAppDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete app default response
func (o *DeleteAppDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAppDefault) Error() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{name}][%d] DeleteApp default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAppDefault) String() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{name}][%d] DeleteApp default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAppDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
