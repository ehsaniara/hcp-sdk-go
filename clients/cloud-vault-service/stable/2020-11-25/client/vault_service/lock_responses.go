// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// LockReader is a Reader for the Lock structure.
type LockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLockOK creates a LockOK with default headers values
func NewLockOK() *LockOK {
	return &LockOK{}
}

/*
LockOK describes a response with status code 200, with default header values.

A successful response.
*/
type LockOK struct {
	Payload *models.HashicorpCloudVault20201125LockResponse
}

// IsSuccess returns true when this lock o k response has a 2xx status code
func (o *LockOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lock o k response has a 3xx status code
func (o *LockOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lock o k response has a 4xx status code
func (o *LockOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lock o k response has a 5xx status code
func (o *LockOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lock o k response a status code equal to that given
func (o *LockOK) IsCode(code int) bool {
	return code == 200
}

func (o *LockOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/lock][%d] lockOK  %+v", 200, o.Payload)
}

func (o *LockOK) String() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/lock][%d] lockOK  %+v", 200, o.Payload)
}

func (o *LockOK) GetPayload() *models.HashicorpCloudVault20201125LockResponse {
	return o.Payload
}

func (o *LockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125LockResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLockDefault creates a LockDefault with default headers values
func NewLockDefault(code int) *LockDefault {
	return &LockDefault{
		_statusCode: code,
	}
}

/*
LockDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LockDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the lock default response
func (o *LockDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this lock default response has a 2xx status code
func (o *LockDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lock default response has a 3xx status code
func (o *LockDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lock default response has a 4xx status code
func (o *LockDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lock default response has a 5xx status code
func (o *LockDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lock default response a status code equal to that given
func (o *LockDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *LockDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/lock][%d] Lock default  %+v", o._statusCode, o.Payload)
}

func (o *LockDefault) String() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/lock][%d] Lock default  %+v", o._statusCode, o.Payload)
}

func (o *LockDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *LockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
