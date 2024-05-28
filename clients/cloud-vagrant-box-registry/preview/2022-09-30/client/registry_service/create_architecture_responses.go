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
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// CreateArchitectureReader is a Reader for the CreateArchitecture structure.
type CreateArchitectureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateArchitectureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateArchitectureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateArchitectureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateArchitectureOK creates a CreateArchitectureOK with default headers values
func NewCreateArchitectureOK() *CreateArchitectureOK {
	return &CreateArchitectureOK{}
}

/*
CreateArchitectureOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateArchitectureOK struct {
	Payload *models.HashicorpCloudVagrant20220930CreateArchitectureResponse
}

// IsSuccess returns true when this create architecture o k response has a 2xx status code
func (o *CreateArchitectureOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create architecture o k response has a 3xx status code
func (o *CreateArchitectureOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create architecture o k response has a 4xx status code
func (o *CreateArchitectureOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create architecture o k response has a 5xx status code
func (o *CreateArchitectureOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create architecture o k response a status code equal to that given
func (o *CreateArchitectureOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create architecture o k response
func (o *CreateArchitectureOK) Code() int {
	return 200
}

func (o *CreateArchitectureOK) Error() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}/architectures][%d] createArchitectureOK  %+v", 200, o.Payload)
}

func (o *CreateArchitectureOK) String() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}/architectures][%d] createArchitectureOK  %+v", 200, o.Payload)
}

func (o *CreateArchitectureOK) GetPayload() *models.HashicorpCloudVagrant20220930CreateArchitectureResponse {
	return o.Payload
}

func (o *CreateArchitectureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930CreateArchitectureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateArchitectureDefault creates a CreateArchitectureDefault with default headers values
func NewCreateArchitectureDefault(code int) *CreateArchitectureDefault {
	return &CreateArchitectureDefault{
		_statusCode: code,
	}
}

/*
CreateArchitectureDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateArchitectureDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this create architecture default response has a 2xx status code
func (o *CreateArchitectureDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create architecture default response has a 3xx status code
func (o *CreateArchitectureDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create architecture default response has a 4xx status code
func (o *CreateArchitectureDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create architecture default response has a 5xx status code
func (o *CreateArchitectureDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create architecture default response a status code equal to that given
func (o *CreateArchitectureDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create architecture default response
func (o *CreateArchitectureDefault) Code() int {
	return o._statusCode
}

func (o *CreateArchitectureDefault) Error() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}/architectures][%d] CreateArchitecture default  %+v", o._statusCode, o.Payload)
}

func (o *CreateArchitectureDefault) String() string {
	return fmt.Sprintf("[POST /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}/architectures][%d] CreateArchitecture default  %+v", o._statusCode, o.Payload)
}

func (o *CreateArchitectureDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *CreateArchitectureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}