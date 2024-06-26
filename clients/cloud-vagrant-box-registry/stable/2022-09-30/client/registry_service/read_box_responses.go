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

// ReadBoxReader is a Reader for the ReadBox structure.
type ReadBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadBoxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadBoxDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadBoxOK creates a ReadBoxOK with default headers values
func NewReadBoxOK() *ReadBoxOK {
	return &ReadBoxOK{}
}

/*
ReadBoxOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadBoxOK struct {
	Payload *models.HashicorpCloudVagrant20220930ReadBoxResponse
}

// IsSuccess returns true when this read box o k response has a 2xx status code
func (o *ReadBoxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read box o k response has a 3xx status code
func (o *ReadBoxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read box o k response has a 4xx status code
func (o *ReadBoxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read box o k response has a 5xx status code
func (o *ReadBoxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read box o k response a status code equal to that given
func (o *ReadBoxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read box o k response
func (o *ReadBoxOK) Code() int {
	return 200
}

func (o *ReadBoxOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}][%d] readBoxOK  %+v", 200, o.Payload)
}

func (o *ReadBoxOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}][%d] readBoxOK  %+v", 200, o.Payload)
}

func (o *ReadBoxOK) GetPayload() *models.HashicorpCloudVagrant20220930ReadBoxResponse {
	return o.Payload
}

func (o *ReadBoxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930ReadBoxResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadBoxDefault creates a ReadBoxDefault with default headers values
func NewReadBoxDefault(code int) *ReadBoxDefault {
	return &ReadBoxDefault{
		_statusCode: code,
	}
}

/*
ReadBoxDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ReadBoxDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this read box default response has a 2xx status code
func (o *ReadBoxDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read box default response has a 3xx status code
func (o *ReadBoxDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read box default response has a 4xx status code
func (o *ReadBoxDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read box default response has a 5xx status code
func (o *ReadBoxDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read box default response a status code equal to that given
func (o *ReadBoxDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read box default response
func (o *ReadBoxDefault) Code() int {
	return o._statusCode
}

func (o *ReadBoxDefault) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}][%d] ReadBox default  %+v", o._statusCode, o.Payload)
}

func (o *ReadBoxDefault) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/box/{box}][%d] ReadBox default  %+v", o._statusCode, o.Payload)
}

func (o *ReadBoxDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ReadBoxDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
