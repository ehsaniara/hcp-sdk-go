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

// DeleteProviderReader is a Reader for the DeleteProvider structure.
type DeleteProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProviderOK creates a DeleteProviderOK with default headers values
func NewDeleteProviderOK() *DeleteProviderOK {
	return &DeleteProviderOK{}
}

/*
DeleteProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteProviderOK struct {
	Payload models.HashicorpCloudVagrant20220930DeleteProviderResponse
}

// IsSuccess returns true when this delete provider o k response has a 2xx status code
func (o *DeleteProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete provider o k response has a 3xx status code
func (o *DeleteProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete provider o k response has a 4xx status code
func (o *DeleteProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete provider o k response has a 5xx status code
func (o *DeleteProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete provider o k response a status code equal to that given
func (o *DeleteProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete provider o k response
func (o *DeleteProviderOK) Code() int {
	return 200
}

func (o *DeleteProviderOK) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}][%d] deleteProviderOK  %+v", 200, o.Payload)
}

func (o *DeleteProviderOK) String() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}][%d] deleteProviderOK  %+v", 200, o.Payload)
}

func (o *DeleteProviderOK) GetPayload() models.HashicorpCloudVagrant20220930DeleteProviderResponse {
	return o.Payload
}

func (o *DeleteProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProviderDefault creates a DeleteProviderDefault with default headers values
func NewDeleteProviderDefault(code int) *DeleteProviderDefault {
	return &DeleteProviderDefault{
		_statusCode: code,
	}
}

/*
DeleteProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteProviderDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this delete provider default response has a 2xx status code
func (o *DeleteProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete provider default response has a 3xx status code
func (o *DeleteProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete provider default response has a 4xx status code
func (o *DeleteProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete provider default response has a 5xx status code
func (o *DeleteProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete provider default response a status code equal to that given
func (o *DeleteProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete provider default response
func (o *DeleteProviderDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProviderDefault) Error() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}][%d] DeleteProvider default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProviderDefault) String() string {
	return fmt.Sprintf("[DELETE /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}][%d] DeleteProvider default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProviderDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeleteProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
