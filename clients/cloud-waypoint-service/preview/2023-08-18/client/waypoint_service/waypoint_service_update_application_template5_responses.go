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

// WaypointServiceUpdateApplicationTemplate5Reader is a Reader for the WaypointServiceUpdateApplicationTemplate5 structure.
type WaypointServiceUpdateApplicationTemplate5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUpdateApplicationTemplate5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUpdateApplicationTemplate5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUpdateApplicationTemplate5Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUpdateApplicationTemplate5OK creates a WaypointServiceUpdateApplicationTemplate5OK with default headers values
func NewWaypointServiceUpdateApplicationTemplate5OK() *WaypointServiceUpdateApplicationTemplate5OK {
	return &WaypointServiceUpdateApplicationTemplate5OK{}
}

/*
WaypointServiceUpdateApplicationTemplate5OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUpdateApplicationTemplate5OK struct {
	Payload *models.HashicorpCloudWaypointUpdateApplicationTemplateResponse
}

// IsSuccess returns true when this waypoint service update application template5 o k response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplate5OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service update application template5 o k response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplate5OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service update application template5 o k response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplate5OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service update application template5 o k response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplate5OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service update application template5 o k response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplate5OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service update application template5 o k response
func (o *WaypointServiceUpdateApplicationTemplate5OK) Code() int {
	return 200
}

func (o *WaypointServiceUpdateApplicationTemplate5OK) Error() string {
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/templates/{existing_application_template.id}][%d] waypointServiceUpdateApplicationTemplate5OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceUpdateApplicationTemplate5OK) String() string {
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/templates/{existing_application_template.id}][%d] waypointServiceUpdateApplicationTemplate5OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceUpdateApplicationTemplate5OK) GetPayload() *models.HashicorpCloudWaypointUpdateApplicationTemplateResponse {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplate5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUpdateApplicationTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUpdateApplicationTemplate5Default creates a WaypointServiceUpdateApplicationTemplate5Default with default headers values
func NewWaypointServiceUpdateApplicationTemplate5Default(code int) *WaypointServiceUpdateApplicationTemplate5Default {
	return &WaypointServiceUpdateApplicationTemplate5Default{
		_statusCode: code,
	}
}

/*
WaypointServiceUpdateApplicationTemplate5Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUpdateApplicationTemplate5Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service update application template5 default response has a 2xx status code
func (o *WaypointServiceUpdateApplicationTemplate5Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service update application template5 default response has a 3xx status code
func (o *WaypointServiceUpdateApplicationTemplate5Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service update application template5 default response has a 4xx status code
func (o *WaypointServiceUpdateApplicationTemplate5Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service update application template5 default response has a 5xx status code
func (o *WaypointServiceUpdateApplicationTemplate5Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service update application template5 default response a status code equal to that given
func (o *WaypointServiceUpdateApplicationTemplate5Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service update application template5 default response
func (o *WaypointServiceUpdateApplicationTemplate5Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUpdateApplicationTemplate5Default) Error() string {
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/templates/{existing_application_template.id}][%d] WaypointService_UpdateApplicationTemplate5 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceUpdateApplicationTemplate5Default) String() string {
	return fmt.Sprintf("[PATCH /waypoint/2023-08-18/namespace/{namespace.id}/templates/{existing_application_template.id}][%d] WaypointService_UpdateApplicationTemplate5 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceUpdateApplicationTemplate5Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUpdateApplicationTemplate5Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}