// Code generated by go-swagger; DO NOT EDIT.

package architecture_service

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

// ListArchitecturesReader is a Reader for the ListArchitectures structure.
type ListArchitecturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArchitecturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArchitecturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListArchitecturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListArchitecturesOK creates a ListArchitecturesOK with default headers values
func NewListArchitecturesOK() *ListArchitecturesOK {
	return &ListArchitecturesOK{}
}

/*
ListArchitecturesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListArchitecturesOK struct {
	Payload *models.HashicorpCloudVagrant20220930ListArchitecturesResponse
}

// IsSuccess returns true when this list architectures o k response has a 2xx status code
func (o *ListArchitecturesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list architectures o k response has a 3xx status code
func (o *ListArchitecturesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list architectures o k response has a 4xx status code
func (o *ListArchitecturesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list architectures o k response has a 5xx status code
func (o *ListArchitecturesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list architectures o k response a status code equal to that given
func (o *ListArchitecturesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list architectures o k response
func (o *ListArchitecturesOK) Code() int {
	return 200
}

func (o *ListArchitecturesOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}/architectures][%d] listArchitecturesOK  %+v", 200, o.Payload)
}

func (o *ListArchitecturesOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}/architectures][%d] listArchitecturesOK  %+v", 200, o.Payload)
}

func (o *ListArchitecturesOK) GetPayload() *models.HashicorpCloudVagrant20220930ListArchitecturesResponse {
	return o.Payload
}

func (o *ListArchitecturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930ListArchitecturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchitecturesDefault creates a ListArchitecturesDefault with default headers values
func NewListArchitecturesDefault(code int) *ListArchitecturesDefault {
	return &ListArchitecturesDefault{
		_statusCode: code,
	}
}

/*
ListArchitecturesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListArchitecturesDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this list architectures default response has a 2xx status code
func (o *ListArchitecturesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list architectures default response has a 3xx status code
func (o *ListArchitecturesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list architectures default response has a 4xx status code
func (o *ListArchitecturesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list architectures default response has a 5xx status code
func (o *ListArchitecturesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list architectures default response a status code equal to that given
func (o *ListArchitecturesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list architectures default response
func (o *ListArchitecturesDefault) Code() int {
	return o._statusCode
}

func (o *ListArchitecturesDefault) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}/architectures][%d] ListArchitectures default  %+v", o._statusCode, o.Payload)
}

func (o *ListArchitecturesDefault) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers/{provider}/architectures][%d] ListArchitectures default  %+v", o._statusCode, o.Payload)
}

func (o *ListArchitecturesDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ListArchitecturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
