// Code generated by go-swagger; DO NOT EDIT.

package box_service

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

// DeleteBoxReader is a Reader for the DeleteBox structure.
type DeleteBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBoxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteBoxDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBoxOK creates a DeleteBoxOK with default headers values
func NewDeleteBoxOK() *DeleteBoxOK {
	return &DeleteBoxOK{}
}

/*
DeleteBoxOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteBoxOK struct {
	Payload *models.HashicorpCloudVagrant20220930DeleteBoxResponse
}

// IsSuccess returns true when this delete box o k response has a 2xx status code
func (o *DeleteBoxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete box o k response has a 3xx status code
func (o *DeleteBoxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete box o k response has a 4xx status code
func (o *DeleteBoxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete box o k response has a 5xx status code
func (o *DeleteBoxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete box o k response a status code equal to that given
func (o *DeleteBoxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete box o k response
func (o *DeleteBoxOK) Code() int {
	return 200
}

func (o *DeleteBoxOK) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] deleteBoxOK  %+v", 200, o.Payload)
}

func (o *DeleteBoxOK) String() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] deleteBoxOK  %+v", 200, o.Payload)
}

func (o *DeleteBoxOK) GetPayload() *models.HashicorpCloudVagrant20220930DeleteBoxResponse {
	return o.Payload
}

func (o *DeleteBoxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930DeleteBoxResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBoxDefault creates a DeleteBoxDefault with default headers values
func NewDeleteBoxDefault(code int) *DeleteBoxDefault {
	return &DeleteBoxDefault{
		_statusCode: code,
	}
}

/*
DeleteBoxDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteBoxDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this delete box default response has a 2xx status code
func (o *DeleteBoxDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete box default response has a 3xx status code
func (o *DeleteBoxDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete box default response has a 4xx status code
func (o *DeleteBoxDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete box default response has a 5xx status code
func (o *DeleteBoxDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete box default response a status code equal to that given
func (o *DeleteBoxDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete box default response
func (o *DeleteBoxDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBoxDefault) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] DeleteBox default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBoxDefault) String() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/boxes/{box}][%d] DeleteBox default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBoxDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeleteBoxDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
