// Code generated by go-swagger; DO NOT EDIT.

package version_service

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

// DeleteVersionReader is a Reader for the DeleteVersion structure.
type DeleteVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVersionOK creates a DeleteVersionOK with default headers values
func NewDeleteVersionOK() *DeleteVersionOK {
	return &DeleteVersionOK{}
}

/*
DeleteVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteVersionOK struct {
	Payload models.HashicorpCloudVagrant20220930DeleteVersionResponse
}

// IsSuccess returns true when this delete version o k response has a 2xx status code
func (o *DeleteVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete version o k response has a 3xx status code
func (o *DeleteVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete version o k response has a 4xx status code
func (o *DeleteVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete version o k response has a 5xx status code
func (o *DeleteVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete version o k response a status code equal to that given
func (o *DeleteVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete version o k response
func (o *DeleteVersionOK) Code() int {
	return 200
}

func (o *DeleteVersionOK) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}][%d] deleteVersionOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionOK) String() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}][%d] deleteVersionOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionOK) GetPayload() models.HashicorpCloudVagrant20220930DeleteVersionResponse {
	return o.Payload
}

func (o *DeleteVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionDefault creates a DeleteVersionDefault with default headers values
func NewDeleteVersionDefault(code int) *DeleteVersionDefault {
	return &DeleteVersionDefault{
		_statusCode: code,
	}
}

/*
DeleteVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteVersionDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this delete version default response has a 2xx status code
func (o *DeleteVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete version default response has a 3xx status code
func (o *DeleteVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete version default response has a 4xx status code
func (o *DeleteVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete version default response has a 5xx status code
func (o *DeleteVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete version default response a status code equal to that given
func (o *DeleteVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete version default response
func (o *DeleteVersionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVersionDefault) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}][%d] DeleteVersion default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVersionDefault) String() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}][%d] DeleteVersion default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVersionDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeleteVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
