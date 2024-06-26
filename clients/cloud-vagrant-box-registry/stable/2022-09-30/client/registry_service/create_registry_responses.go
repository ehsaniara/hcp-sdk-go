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

// CreateRegistryReader is a Reader for the CreateRegistry structure.
type CreateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRegistryOK creates a CreateRegistryOK with default headers values
func NewCreateRegistryOK() *CreateRegistryOK {
	return &CreateRegistryOK{}
}

/*
CreateRegistryOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateRegistryOK struct {
	Payload *models.HashicorpCloudVagrant20220930CreateRegistryResponse
}

// IsSuccess returns true when this create registry o k response has a 2xx status code
func (o *CreateRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create registry o k response has a 3xx status code
func (o *CreateRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create registry o k response has a 4xx status code
func (o *CreateRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create registry o k response has a 5xx status code
func (o *CreateRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create registry o k response a status code equal to that given
func (o *CreateRegistryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create registry o k response
func (o *CreateRegistryOK) Code() int {
	return 200
}

func (o *CreateRegistryOK) Error() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registries][%d] createRegistryOK  %+v", 200, o.Payload)
}

func (o *CreateRegistryOK) String() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registries][%d] createRegistryOK  %+v", 200, o.Payload)
}

func (o *CreateRegistryOK) GetPayload() *models.HashicorpCloudVagrant20220930CreateRegistryResponse {
	return o.Payload
}

func (o *CreateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930CreateRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegistryDefault creates a CreateRegistryDefault with default headers values
func NewCreateRegistryDefault(code int) *CreateRegistryDefault {
	return &CreateRegistryDefault{
		_statusCode: code,
	}
}

/*
CreateRegistryDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateRegistryDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this create registry default response has a 2xx status code
func (o *CreateRegistryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create registry default response has a 3xx status code
func (o *CreateRegistryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create registry default response has a 4xx status code
func (o *CreateRegistryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create registry default response has a 5xx status code
func (o *CreateRegistryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create registry default response a status code equal to that given
func (o *CreateRegistryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create registry default response
func (o *CreateRegistryDefault) Code() int {
	return o._statusCode
}

func (o *CreateRegistryDefault) Error() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registries][%d] CreateRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRegistryDefault) String() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registries][%d] CreateRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRegistryDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *CreateRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
