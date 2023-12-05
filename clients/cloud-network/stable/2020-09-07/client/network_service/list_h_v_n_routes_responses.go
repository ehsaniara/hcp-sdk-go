// Code generated by go-swagger; DO NOT EDIT.

package network_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-network/stable/2020-09-07/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListHVNRoutesReader is a Reader for the ListHVNRoutes structure.
type ListHVNRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHVNRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListHVNRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListHVNRoutesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListHVNRoutesOK creates a ListHVNRoutesOK with default headers values
func NewListHVNRoutesOK() *ListHVNRoutesOK {
	return &ListHVNRoutesOK{}
}

/*
ListHVNRoutesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListHVNRoutesOK struct {
	Payload *models.HashicorpCloudNetwork20200907ListHVNRoutesResponse
}

// IsSuccess returns true when this list h v n routes o k response has a 2xx status code
func (o *ListHVNRoutesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list h v n routes o k response has a 3xx status code
func (o *ListHVNRoutesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list h v n routes o k response has a 4xx status code
func (o *ListHVNRoutesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list h v n routes o k response has a 5xx status code
func (o *ListHVNRoutesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list h v n routes o k response a status code equal to that given
func (o *ListHVNRoutesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list h v n routes o k response
func (o *ListHVNRoutesOK) Code() int {
	return 200
}

func (o *ListHVNRoutesOK) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes][%d] listHVNRoutesOK  %+v", 200, o.Payload)
}

func (o *ListHVNRoutesOK) String() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes][%d] listHVNRoutesOK  %+v", 200, o.Payload)
}

func (o *ListHVNRoutesOK) GetPayload() *models.HashicorpCloudNetwork20200907ListHVNRoutesResponse {
	return o.Payload
}

func (o *ListHVNRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907ListHVNRoutesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHVNRoutesDefault creates a ListHVNRoutesDefault with default headers values
func NewListHVNRoutesDefault(code int) *ListHVNRoutesDefault {
	return &ListHVNRoutesDefault{
		_statusCode: code,
	}
}

/*
ListHVNRoutesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListHVNRoutesDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this list h v n routes default response has a 2xx status code
func (o *ListHVNRoutesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list h v n routes default response has a 3xx status code
func (o *ListHVNRoutesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list h v n routes default response has a 4xx status code
func (o *ListHVNRoutesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list h v n routes default response has a 5xx status code
func (o *ListHVNRoutesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list h v n routes default response a status code equal to that given
func (o *ListHVNRoutesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list h v n routes default response
func (o *ListHVNRoutesDefault) Code() int {
	return o._statusCode
}

func (o *ListHVNRoutesDefault) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes][%d] ListHVNRoutes default  %+v", o._statusCode, o.Payload)
}

func (o *ListHVNRoutesDefault) String() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{hvn.location.organization_id}/projects/{hvn.location.project_id}/networks/{hvn.id}/routes][%d] ListHVNRoutes default  %+v", o._statusCode, o.Payload)
}

func (o *ListHVNRoutesDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListHVNRoutesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
