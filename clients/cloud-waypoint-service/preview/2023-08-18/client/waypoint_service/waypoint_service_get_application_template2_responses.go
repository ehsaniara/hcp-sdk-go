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

// WaypointServiceGetApplicationTemplate2Reader is a Reader for the WaypointServiceGetApplicationTemplate2 structure.
type WaypointServiceGetApplicationTemplate2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetApplicationTemplate2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetApplicationTemplate2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetApplicationTemplate2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetApplicationTemplate2OK creates a WaypointServiceGetApplicationTemplate2OK with default headers values
func NewWaypointServiceGetApplicationTemplate2OK() *WaypointServiceGetApplicationTemplate2OK {
	return &WaypointServiceGetApplicationTemplate2OK{}
}

/*
WaypointServiceGetApplicationTemplate2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetApplicationTemplate2OK struct {
	Payload *models.HashicorpCloudWaypointGetApplicationTemplateResponse
}

// IsSuccess returns true when this waypoint service get application template2 o k response has a 2xx status code
func (o *WaypointServiceGetApplicationTemplate2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get application template2 o k response has a 3xx status code
func (o *WaypointServiceGetApplicationTemplate2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get application template2 o k response has a 4xx status code
func (o *WaypointServiceGetApplicationTemplate2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get application template2 o k response has a 5xx status code
func (o *WaypointServiceGetApplicationTemplate2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get application template2 o k response a status code equal to that given
func (o *WaypointServiceGetApplicationTemplate2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get application template2 o k response
func (o *WaypointServiceGetApplicationTemplate2OK) Code() int {
	return 200
}

func (o *WaypointServiceGetApplicationTemplate2OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/application-templates/by-name/{application_template.name}][%d] waypointServiceGetApplicationTemplate2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceGetApplicationTemplate2OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/application-templates/by-name/{application_template.name}][%d] waypointServiceGetApplicationTemplate2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceGetApplicationTemplate2OK) GetPayload() *models.HashicorpCloudWaypointGetApplicationTemplateResponse {
	return o.Payload
}

func (o *WaypointServiceGetApplicationTemplate2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetApplicationTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetApplicationTemplate2Default creates a WaypointServiceGetApplicationTemplate2Default with default headers values
func NewWaypointServiceGetApplicationTemplate2Default(code int) *WaypointServiceGetApplicationTemplate2Default {
	return &WaypointServiceGetApplicationTemplate2Default{
		_statusCode: code,
	}
}

/*
WaypointServiceGetApplicationTemplate2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetApplicationTemplate2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get application template2 default response has a 2xx status code
func (o *WaypointServiceGetApplicationTemplate2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get application template2 default response has a 3xx status code
func (o *WaypointServiceGetApplicationTemplate2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get application template2 default response has a 4xx status code
func (o *WaypointServiceGetApplicationTemplate2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get application template2 default response has a 5xx status code
func (o *WaypointServiceGetApplicationTemplate2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get application template2 default response a status code equal to that given
func (o *WaypointServiceGetApplicationTemplate2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get application template2 default response
func (o *WaypointServiceGetApplicationTemplate2Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetApplicationTemplate2Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/application-templates/by-name/{application_template.name}][%d] WaypointService_GetApplicationTemplate2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceGetApplicationTemplate2Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/application-templates/by-name/{application_template.name}][%d] WaypointService_GetApplicationTemplate2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceGetApplicationTemplate2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetApplicationTemplate2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
