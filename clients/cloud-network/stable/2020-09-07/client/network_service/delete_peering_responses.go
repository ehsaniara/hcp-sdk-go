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

// DeletePeeringReader is a Reader for the DeletePeering structure.
type DeletePeeringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePeeringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePeeringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePeeringDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePeeringOK creates a DeletePeeringOK with default headers values
func NewDeletePeeringOK() *DeletePeeringOK {
	return &DeletePeeringOK{}
}

/*
DeletePeeringOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeletePeeringOK struct {
	Payload *models.HashicorpCloudNetwork20200907DeletePeeringResponse
}

// IsSuccess returns true when this delete peering o k response has a 2xx status code
func (o *DeletePeeringOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete peering o k response has a 3xx status code
func (o *DeletePeeringOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete peering o k response has a 4xx status code
func (o *DeletePeeringOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete peering o k response has a 5xx status code
func (o *DeletePeeringOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete peering o k response a status code equal to that given
func (o *DeletePeeringOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete peering o k response
func (o *DeletePeeringOK) Code() int {
	return 200
}

func (o *DeletePeeringOK) Error() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] deletePeeringOK  %+v", 200, o.Payload)
}

func (o *DeletePeeringOK) String() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] deletePeeringOK  %+v", 200, o.Payload)
}

func (o *DeletePeeringOK) GetPayload() *models.HashicorpCloudNetwork20200907DeletePeeringResponse {
	return o.Payload
}

func (o *DeletePeeringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907DeletePeeringResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePeeringDefault creates a DeletePeeringDefault with default headers values
func NewDeletePeeringDefault(code int) *DeletePeeringDefault {
	return &DeletePeeringDefault{
		_statusCode: code,
	}
}

/*
DeletePeeringDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeletePeeringDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this delete peering default response has a 2xx status code
func (o *DeletePeeringDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete peering default response has a 3xx status code
func (o *DeletePeeringDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete peering default response has a 4xx status code
func (o *DeletePeeringDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete peering default response has a 5xx status code
func (o *DeletePeeringDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete peering default response a status code equal to that given
func (o *DeletePeeringDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete peering default response
func (o *DeletePeeringDefault) Code() int {
	return o._statusCode
}

func (o *DeletePeeringDefault) Error() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] DeletePeering default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePeeringDefault) String() string {
	return fmt.Sprintf("[DELETE /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{hvn_id}/peerings/{id}][%d] DeletePeering default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePeeringDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DeletePeeringDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
