// Code generated by go-swagger; DO NOT EDIT.

package waypoint_control_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2022-02-03/models"
)

// WaypointControlServiceGetNamespaceReader is a Reader for the WaypointControlServiceGetNamespace structure.
type WaypointControlServiceGetNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointControlServiceGetNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointControlServiceGetNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointControlServiceGetNamespaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointControlServiceGetNamespaceOK creates a WaypointControlServiceGetNamespaceOK with default headers values
func NewWaypointControlServiceGetNamespaceOK() *WaypointControlServiceGetNamespaceOK {
	return &WaypointControlServiceGetNamespaceOK{}
}

/*
WaypointControlServiceGetNamespaceOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointControlServiceGetNamespaceOK struct {
	Payload *models.HashicorpCloudWaypointGetNamespaceResponse
}

// IsSuccess returns true when this waypoint control service get namespace o k response has a 2xx status code
func (o *WaypointControlServiceGetNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint control service get namespace o k response has a 3xx status code
func (o *WaypointControlServiceGetNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint control service get namespace o k response has a 4xx status code
func (o *WaypointControlServiceGetNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint control service get namespace o k response has a 5xx status code
func (o *WaypointControlServiceGetNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint control service get namespace o k response a status code equal to that given
func (o *WaypointControlServiceGetNamespaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint control service get namespace o k response
func (o *WaypointControlServiceGetNamespaceOK) Code() int {
	return 200
}

func (o *WaypointControlServiceGetNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/organizations/{location.organization_id}/projects/{location.project_id}/namespace][%d] waypointControlServiceGetNamespaceOK  %+v", 200, o.Payload)
}

func (o *WaypointControlServiceGetNamespaceOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/organizations/{location.organization_id}/projects/{location.project_id}/namespace][%d] waypointControlServiceGetNamespaceOK  %+v", 200, o.Payload)
}

func (o *WaypointControlServiceGetNamespaceOK) GetPayload() *models.HashicorpCloudWaypointGetNamespaceResponse {
	return o.Payload
}

func (o *WaypointControlServiceGetNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetNamespaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointControlServiceGetNamespaceDefault creates a WaypointControlServiceGetNamespaceDefault with default headers values
func NewWaypointControlServiceGetNamespaceDefault(code int) *WaypointControlServiceGetNamespaceDefault {
	return &WaypointControlServiceGetNamespaceDefault{
		_statusCode: code,
	}
}

/*
WaypointControlServiceGetNamespaceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointControlServiceGetNamespaceDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint control service get namespace default response has a 2xx status code
func (o *WaypointControlServiceGetNamespaceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint control service get namespace default response has a 3xx status code
func (o *WaypointControlServiceGetNamespaceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint control service get namespace default response has a 4xx status code
func (o *WaypointControlServiceGetNamespaceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint control service get namespace default response has a 5xx status code
func (o *WaypointControlServiceGetNamespaceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint control service get namespace default response a status code equal to that given
func (o *WaypointControlServiceGetNamespaceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint control service get namespace default response
func (o *WaypointControlServiceGetNamespaceDefault) Code() int {
	return o._statusCode
}

func (o *WaypointControlServiceGetNamespaceDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/organizations/{location.organization_id}/projects/{location.project_id}/namespace][%d] WaypointControlService_GetNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointControlServiceGetNamespaceDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/organizations/{location.organization_id}/projects/{location.project_id}/namespace][%d] WaypointControlService_GetNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointControlServiceGetNamespaceDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointControlServiceGetNamespaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}