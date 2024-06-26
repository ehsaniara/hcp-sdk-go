// Code generated by go-swagger; DO NOT EDIT.

package registry_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/stable/2022-09-30/models"
)

// CreateProviderReader is a Reader for the CreateProvider structure.
type CreateProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProviderOK creates a CreateProviderOK with default headers values
func NewCreateProviderOK() *CreateProviderOK {
	return &CreateProviderOK{}
}

/*
CreateProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateProviderOK struct {
	Payload *models.HashicorpCloudVagrant20220930CreateProviderResponse
}

// IsSuccess returns true when this create provider o k response has a 2xx status code
func (o *CreateProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create provider o k response has a 3xx status code
func (o *CreateProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create provider o k response has a 4xx status code
func (o *CreateProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create provider o k response has a 5xx status code
func (o *CreateProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create provider o k response a status code equal to that given
func (o *CreateProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create provider o k response
func (o *CreateProviderOK) Code() int {
	return 200
}

func (o *CreateProviderOK) Error() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/providers][%d] createProviderOK  %+v", 200, o.Payload)
}

func (o *CreateProviderOK) String() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/providers][%d] createProviderOK  %+v", 200, o.Payload)
}

func (o *CreateProviderOK) GetPayload() *models.HashicorpCloudVagrant20220930CreateProviderResponse {
	return o.Payload
}

func (o *CreateProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930CreateProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProviderDefault creates a CreateProviderDefault with default headers values
func NewCreateProviderDefault(code int) *CreateProviderDefault {
	return &CreateProviderDefault{
		_statusCode: code,
	}
}

/*
CreateProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateProviderDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this create provider default response has a 2xx status code
func (o *CreateProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create provider default response has a 3xx status code
func (o *CreateProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create provider default response has a 4xx status code
func (o *CreateProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create provider default response has a 5xx status code
func (o *CreateProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create provider default response a status code equal to that given
func (o *CreateProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create provider default response
func (o *CreateProviderDefault) Code() int {
	return o._statusCode
}

func (o *CreateProviderDefault) Error() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/providers][%d] CreateProvider default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProviderDefault) String() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/providers][%d] CreateProvider default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProviderDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *CreateProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
