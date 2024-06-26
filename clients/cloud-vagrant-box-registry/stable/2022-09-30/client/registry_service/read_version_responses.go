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

// ReadVersionReader is a Reader for the ReadVersion structure.
type ReadVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadVersionOK creates a ReadVersionOK with default headers values
func NewReadVersionOK() *ReadVersionOK {
	return &ReadVersionOK{}
}

/*
ReadVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadVersionOK struct {
	Payload *models.HashicorpCloudVagrant20220930ReadVersionResponse
}

// IsSuccess returns true when this read version o k response has a 2xx status code
func (o *ReadVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read version o k response has a 3xx status code
func (o *ReadVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read version o k response has a 4xx status code
func (o *ReadVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read version o k response has a 5xx status code
func (o *ReadVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read version o k response a status code equal to that given
func (o *ReadVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read version o k response
func (o *ReadVersionOK) Code() int {
	return 200
}

func (o *ReadVersionOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}][%d] readVersionOK  %+v", 200, o.Payload)
}

func (o *ReadVersionOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}][%d] readVersionOK  %+v", 200, o.Payload)
}

func (o *ReadVersionOK) GetPayload() *models.HashicorpCloudVagrant20220930ReadVersionResponse {
	return o.Payload
}

func (o *ReadVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930ReadVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadVersionDefault creates a ReadVersionDefault with default headers values
func NewReadVersionDefault(code int) *ReadVersionDefault {
	return &ReadVersionDefault{
		_statusCode: code,
	}
}

/*
ReadVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ReadVersionDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this read version default response has a 2xx status code
func (o *ReadVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read version default response has a 3xx status code
func (o *ReadVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read version default response has a 4xx status code
func (o *ReadVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read version default response has a 5xx status code
func (o *ReadVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read version default response a status code equal to that given
func (o *ReadVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read version default response
func (o *ReadVersionDefault) Code() int {
	return o._statusCode
}

func (o *ReadVersionDefault) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}][%d] ReadVersion default  %+v", o._statusCode, o.Payload)
}

func (o *ReadVersionDefault) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}][%d] ReadVersion default  %+v", o._statusCode, o.Payload)
}

func (o *ReadVersionDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ReadVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
