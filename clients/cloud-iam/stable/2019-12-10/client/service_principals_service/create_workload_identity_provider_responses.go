// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// CreateWorkloadIdentityProviderReader is a Reader for the CreateWorkloadIdentityProvider structure.
type CreateWorkloadIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWorkloadIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateWorkloadIdentityProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateWorkloadIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateWorkloadIdentityProviderOK creates a CreateWorkloadIdentityProviderOK with default headers values
func NewCreateWorkloadIdentityProviderOK() *CreateWorkloadIdentityProviderOK {
	return &CreateWorkloadIdentityProviderOK{}
}

/*
CreateWorkloadIdentityProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateWorkloadIdentityProviderOK struct {
	Payload *models.HashicorpCloudIamCreateWorkloadIdentityProviderResponse
}

// IsSuccess returns true when this create workload identity provider o k response has a 2xx status code
func (o *CreateWorkloadIdentityProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create workload identity provider o k response has a 3xx status code
func (o *CreateWorkloadIdentityProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workload identity provider o k response has a 4xx status code
func (o *CreateWorkloadIdentityProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create workload identity provider o k response has a 5xx status code
func (o *CreateWorkloadIdentityProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create workload identity provider o k response a status code equal to that given
func (o *CreateWorkloadIdentityProviderOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateWorkloadIdentityProviderOK) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] createWorkloadIdentityProviderOK  %+v", 200, o.Payload)
}

func (o *CreateWorkloadIdentityProviderOK) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] createWorkloadIdentityProviderOK  %+v", 200, o.Payload)
}

func (o *CreateWorkloadIdentityProviderOK) GetPayload() *models.HashicorpCloudIamCreateWorkloadIdentityProviderResponse {
	return o.Payload
}

func (o *CreateWorkloadIdentityProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamCreateWorkloadIdentityProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkloadIdentityProviderDefault creates a CreateWorkloadIdentityProviderDefault with default headers values
func NewCreateWorkloadIdentityProviderDefault(code int) *CreateWorkloadIdentityProviderDefault {
	return &CreateWorkloadIdentityProviderDefault{
		_statusCode: code,
	}
}

/*
CreateWorkloadIdentityProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateWorkloadIdentityProviderDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the create workload identity provider default response
func (o *CreateWorkloadIdentityProviderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create workload identity provider default response has a 2xx status code
func (o *CreateWorkloadIdentityProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create workload identity provider default response has a 3xx status code
func (o *CreateWorkloadIdentityProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create workload identity provider default response has a 4xx status code
func (o *CreateWorkloadIdentityProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create workload identity provider default response has a 5xx status code
func (o *CreateWorkloadIdentityProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create workload identity provider default response a status code equal to that given
func (o *CreateWorkloadIdentityProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateWorkloadIdentityProviderDefault) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] CreateWorkloadIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *CreateWorkloadIdentityProviderDefault) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/workload-identity-providers][%d] CreateWorkloadIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *CreateWorkloadIdentityProviderDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *CreateWorkloadIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}