// Code generated by go-swagger; DO NOT EDIT.

package iam_service

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

// PingReader is a Reader for the Ping structure.
type PingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPingOK creates a PingOK with default headers values
func NewPingOK() *PingOK {
	return &PingOK{}
}

/*
PingOK describes a response with status code 200, with default header values.

A successful response.
*/
type PingOK struct {
	Payload models.HashicorpCloudIamPingResponse
}

// IsSuccess returns true when this ping o k response has a 2xx status code
func (o *PingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ping o k response has a 3xx status code
func (o *PingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping o k response has a 4xx status code
func (o *PingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping o k response has a 5xx status code
func (o *PingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ping o k response a status code equal to that given
func (o *PingOK) IsCode(code int) bool {
	return code == 200
}

func (o *PingOK) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/ping][%d] pingOK  %+v", 200, o.Payload)
}

func (o *PingOK) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/ping][%d] pingOK  %+v", 200, o.Payload)
}

func (o *PingOK) GetPayload() models.HashicorpCloudIamPingResponse {
	return o.Payload
}

func (o *PingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingDefault creates a PingDefault with default headers values
func NewPingDefault(code int) *PingDefault {
	return &PingDefault{
		_statusCode: code,
	}
}

/*
PingDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PingDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the ping default response
func (o *PingDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ping default response has a 2xx status code
func (o *PingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ping default response has a 3xx status code
func (o *PingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ping default response has a 4xx status code
func (o *PingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ping default response has a 5xx status code
func (o *PingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ping default response a status code equal to that given
func (o *PingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PingDefault) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/ping][%d] Ping default  %+v", o._statusCode, o.Payload)
}

func (o *PingDefault) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/ping][%d] Ping default  %+v", o._statusCode, o.Payload)
}

func (o *PingDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}