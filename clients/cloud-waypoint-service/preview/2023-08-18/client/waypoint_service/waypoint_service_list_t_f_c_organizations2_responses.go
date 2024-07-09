// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
)

// WaypointServiceListTFCOrganizations2Reader is a Reader for the WaypointServiceListTFCOrganizations2 structure.
type WaypointServiceListTFCOrganizations2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceListTFCOrganizations2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceListTFCOrganizations2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceListTFCOrganizations2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceListTFCOrganizations2OK creates a WaypointServiceListTFCOrganizations2OK with default headers values
func NewWaypointServiceListTFCOrganizations2OK() *WaypointServiceListTFCOrganizations2OK {
	return &WaypointServiceListTFCOrganizations2OK{}
}

/*
WaypointServiceListTFCOrganizations2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceListTFCOrganizations2OK struct {
	Payload *models.HashicorpCloudWaypointListTFCOrganizationsResponse
}

// IsSuccess returns true when this waypoint service list t f c organizations2 o k response has a 2xx status code
func (o *WaypointServiceListTFCOrganizations2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service list t f c organizations2 o k response has a 3xx status code
func (o *WaypointServiceListTFCOrganizations2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service list t f c organizations2 o k response has a 4xx status code
func (o *WaypointServiceListTFCOrganizations2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service list t f c organizations2 o k response has a 5xx status code
func (o *WaypointServiceListTFCOrganizations2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service list t f c organizations2 o k response a status code equal to that given
func (o *WaypointServiceListTFCOrganizations2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service list t f c organizations2 o k response
func (o *WaypointServiceListTFCOrganizations2OK) Code() int {
	return 200
}

func (o *WaypointServiceListTFCOrganizations2OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfc-organizations][%d] waypointServiceListTFCOrganizations2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceListTFCOrganizations2OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfc-organizations][%d] waypointServiceListTFCOrganizations2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceListTFCOrganizations2OK) GetPayload() *models.HashicorpCloudWaypointListTFCOrganizationsResponse {
	return o.Payload
}

func (o *WaypointServiceListTFCOrganizations2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointListTFCOrganizationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceListTFCOrganizations2Default creates a WaypointServiceListTFCOrganizations2Default with default headers values
func NewWaypointServiceListTFCOrganizations2Default(code int) *WaypointServiceListTFCOrganizations2Default {
	return &WaypointServiceListTFCOrganizations2Default{
		_statusCode: code,
	}
}

/*
WaypointServiceListTFCOrganizations2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceListTFCOrganizations2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service list t f c organizations2 default response has a 2xx status code
func (o *WaypointServiceListTFCOrganizations2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service list t f c organizations2 default response has a 3xx status code
func (o *WaypointServiceListTFCOrganizations2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service list t f c organizations2 default response has a 4xx status code
func (o *WaypointServiceListTFCOrganizations2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service list t f c organizations2 default response has a 5xx status code
func (o *WaypointServiceListTFCOrganizations2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service list t f c organizations2 default response a status code equal to that given
func (o *WaypointServiceListTFCOrganizations2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service list t f c organizations2 default response
func (o *WaypointServiceListTFCOrganizations2Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceListTFCOrganizations2Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfc-organizations][%d] WaypointService_ListTFCOrganizations2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceListTFCOrganizations2Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/tfc-organizations][%d] WaypointService_ListTFCOrganizations2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceListTFCOrganizations2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceListTFCOrganizations2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}