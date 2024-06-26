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

// ReadProviderReader is a Reader for the ReadProvider structure.
type ReadProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadProviderOK creates a ReadProviderOK with default headers values
func NewReadProviderOK() *ReadProviderOK {
	return &ReadProviderOK{}
}

/*
ReadProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadProviderOK struct {
	Payload *models.HashicorpCloudVagrant20220930ReadProviderResponse
}

// IsSuccess returns true when this read provider o k response has a 2xx status code
func (o *ReadProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read provider o k response has a 3xx status code
func (o *ReadProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read provider o k response has a 4xx status code
func (o *ReadProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read provider o k response has a 5xx status code
func (o *ReadProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read provider o k response a status code equal to that given
func (o *ReadProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read provider o k response
func (o *ReadProviderOK) Code() int {
	return 200
}

func (o *ReadProviderOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}][%d] readProviderOK  %+v", 200, o.Payload)
}

func (o *ReadProviderOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}][%d] readProviderOK  %+v", 200, o.Payload)
}

func (o *ReadProviderOK) GetPayload() *models.HashicorpCloudVagrant20220930ReadProviderResponse {
	return o.Payload
}

func (o *ReadProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930ReadProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadProviderDefault creates a ReadProviderDefault with default headers values
func NewReadProviderDefault(code int) *ReadProviderDefault {
	return &ReadProviderDefault{
		_statusCode: code,
	}
}

/*
ReadProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ReadProviderDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this read provider default response has a 2xx status code
func (o *ReadProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read provider default response has a 3xx status code
func (o *ReadProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read provider default response has a 4xx status code
func (o *ReadProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read provider default response has a 5xx status code
func (o *ReadProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read provider default response a status code equal to that given
func (o *ReadProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read provider default response
func (o *ReadProviderDefault) Code() int {
	return o._statusCode
}

func (o *ReadProviderDefault) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}][%d] ReadProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ReadProviderDefault) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}][%d] ReadProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ReadProviderDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ReadProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
