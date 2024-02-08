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

// WaypointServiceDestroyApplicationReader is a Reader for the WaypointServiceDestroyApplication structure.
type WaypointServiceDestroyApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceDestroyApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceDestroyApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceDestroyApplicationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceDestroyApplicationOK creates a WaypointServiceDestroyApplicationOK with default headers values
func NewWaypointServiceDestroyApplicationOK() *WaypointServiceDestroyApplicationOK {
	return &WaypointServiceDestroyApplicationOK{}
}

/*
WaypointServiceDestroyApplicationOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceDestroyApplicationOK struct {
	Payload models.HashicorpCloudWaypointDestroyApplicationResponse
}

// IsSuccess returns true when this waypoint service destroy application o k response has a 2xx status code
func (o *WaypointServiceDestroyApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service destroy application o k response has a 3xx status code
func (o *WaypointServiceDestroyApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service destroy application o k response has a 4xx status code
func (o *WaypointServiceDestroyApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service destroy application o k response has a 5xx status code
func (o *WaypointServiceDestroyApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service destroy application o k response a status code equal to that given
func (o *WaypointServiceDestroyApplicationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service destroy application o k response
func (o *WaypointServiceDestroyApplicationOK) Code() int {
	return 200
}

func (o *WaypointServiceDestroyApplicationOK) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/applications/{application.id}][%d] waypointServiceDestroyApplicationOK  %+v", 200, o.Payload)
}

func (o *WaypointServiceDestroyApplicationOK) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/applications/{application.id}][%d] waypointServiceDestroyApplicationOK  %+v", 200, o.Payload)
}

func (o *WaypointServiceDestroyApplicationOK) GetPayload() models.HashicorpCloudWaypointDestroyApplicationResponse {
	return o.Payload
}

func (o *WaypointServiceDestroyApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceDestroyApplicationDefault creates a WaypointServiceDestroyApplicationDefault with default headers values
func NewWaypointServiceDestroyApplicationDefault(code int) *WaypointServiceDestroyApplicationDefault {
	return &WaypointServiceDestroyApplicationDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceDestroyApplicationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceDestroyApplicationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service destroy application default response has a 2xx status code
func (o *WaypointServiceDestroyApplicationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service destroy application default response has a 3xx status code
func (o *WaypointServiceDestroyApplicationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service destroy application default response has a 4xx status code
func (o *WaypointServiceDestroyApplicationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service destroy application default response has a 5xx status code
func (o *WaypointServiceDestroyApplicationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service destroy application default response a status code equal to that given
func (o *WaypointServiceDestroyApplicationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service destroy application default response
func (o *WaypointServiceDestroyApplicationDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceDestroyApplicationDefault) Error() string {
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/applications/{application.id}][%d] WaypointService_DestroyApplication default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceDestroyApplicationDefault) String() string {
	return fmt.Sprintf("[DELETE /waypoint/2023-08-18/namespace/{namespace.id}/applications/{application.id}][%d] WaypointService_DestroyApplication default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceDestroyApplicationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceDestroyApplicationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
