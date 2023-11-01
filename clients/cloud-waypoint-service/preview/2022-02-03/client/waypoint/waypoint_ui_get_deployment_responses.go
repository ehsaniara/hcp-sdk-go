// Code generated by go-swagger; DO NOT EDIT.

package waypoint

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

// WaypointUIGetDeploymentReader is a Reader for the WaypointUIGetDeployment structure.
type WaypointUIGetDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUIGetDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUIGetDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUIGetDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUIGetDeploymentOK creates a WaypointUIGetDeploymentOK with default headers values
func NewWaypointUIGetDeploymentOK() *WaypointUIGetDeploymentOK {
	return &WaypointUIGetDeploymentOK{}
}

/*
WaypointUIGetDeploymentOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointUIGetDeploymentOK struct {
	Payload *models.HashicorpWaypointUIGetDeploymentResponse
}

// IsSuccess returns true when this waypoint Ui get deployment o k response has a 2xx status code
func (o *WaypointUIGetDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint Ui get deployment o k response has a 3xx status code
func (o *WaypointUIGetDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint Ui get deployment o k response has a 4xx status code
func (o *WaypointUIGetDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint Ui get deployment o k response has a 5xx status code
func (o *WaypointUIGetDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint Ui get deployment o k response a status code equal to that given
func (o *WaypointUIGetDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint Ui get deployment o k response
func (o *WaypointUIGetDeploymentOK) Code() int {
	return 200
}

func (o *WaypointUIGetDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/ui/deployment/{ref.id}][%d] waypointUiGetDeploymentOK  %+v", 200, o.Payload)
}

func (o *WaypointUIGetDeploymentOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/ui/deployment/{ref.id}][%d] waypointUiGetDeploymentOK  %+v", 200, o.Payload)
}

func (o *WaypointUIGetDeploymentOK) GetPayload() *models.HashicorpWaypointUIGetDeploymentResponse {
	return o.Payload
}

func (o *WaypointUIGetDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUIGetDeploymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUIGetDeploymentDefault creates a WaypointUIGetDeploymentDefault with default headers values
func NewWaypointUIGetDeploymentDefault(code int) *WaypointUIGetDeploymentDefault {
	return &WaypointUIGetDeploymentDefault{
		_statusCode: code,
	}
}

/*
WaypointUIGetDeploymentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointUIGetDeploymentDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this waypoint UI get deployment default response has a 2xx status code
func (o *WaypointUIGetDeploymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint UI get deployment default response has a 3xx status code
func (o *WaypointUIGetDeploymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint UI get deployment default response has a 4xx status code
func (o *WaypointUIGetDeploymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint UI get deployment default response has a 5xx status code
func (o *WaypointUIGetDeploymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint UI get deployment default response a status code equal to that given
func (o *WaypointUIGetDeploymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint UI get deployment default response
func (o *WaypointUIGetDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *WaypointUIGetDeploymentDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/ui/deployment/{ref.id}][%d] Waypoint_UI_GetDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUIGetDeploymentDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/ui/deployment/{ref.id}][%d] Waypoint_UI_GetDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUIGetDeploymentDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUIGetDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}