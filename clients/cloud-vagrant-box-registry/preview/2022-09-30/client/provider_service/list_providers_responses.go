// Code generated by go-swagger; DO NOT EDIT.

package provider_service

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

// ListProvidersReader is a Reader for the ListProviders structure.
type ListProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProvidersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProvidersOK creates a ListProvidersOK with default headers values
func NewListProvidersOK() *ListProvidersOK {
	return &ListProvidersOK{}
}

/*
ListProvidersOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProvidersOK struct {
	Payload *models.HashicorpCloudVagrant20220930ListProvidersResponse
}

// IsSuccess returns true when this list providers o k response has a 2xx status code
func (o *ListProvidersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list providers o k response has a 3xx status code
func (o *ListProvidersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list providers o k response has a 4xx status code
func (o *ListProvidersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list providers o k response has a 5xx status code
func (o *ListProvidersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list providers o k response a status code equal to that given
func (o *ListProvidersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list providers o k response
func (o *ListProvidersOK) Code() int {
	return 200
}

func (o *ListProvidersOK) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers][%d] listProvidersOK  %+v", 200, o.Payload)
}

func (o *ListProvidersOK) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers][%d] listProvidersOK  %+v", 200, o.Payload)
}

func (o *ListProvidersOK) GetPayload() *models.HashicorpCloudVagrant20220930ListProvidersResponse {
	return o.Payload
}

func (o *ListProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrant20220930ListProvidersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersDefault creates a ListProvidersDefault with default headers values
func NewListProvidersDefault(code int) *ListProvidersDefault {
	return &ListProvidersDefault{
		_statusCode: code,
	}
}

/*
ListProvidersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListProvidersDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this list providers default response has a 2xx status code
func (o *ListProvidersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list providers default response has a 3xx status code
func (o *ListProvidersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list providers default response has a 4xx status code
func (o *ListProvidersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list providers default response has a 5xx status code
func (o *ListProvidersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list providers default response a status code equal to that given
func (o *ListProvidersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list providers default response
func (o *ListProvidersDefault) Code() int {
	return o._statusCode
}

func (o *ListProvidersDefault) Error() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers][%d] ListProviders default  %+v", o._statusCode, o.Payload)
}

func (o *ListProvidersDefault) String() string {
	return fmt.Sprintf("[GET /vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/providers][%d] ListProviders default  %+v", o._statusCode, o.Payload)
}

func (o *ListProvidersDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ListProvidersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
