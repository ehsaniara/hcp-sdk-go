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

// WaypointServiceGetNamespaceReader is a Reader for the WaypointServiceGetNamespace structure.
type WaypointServiceGetNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceGetNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceGetNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceGetNamespaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceGetNamespaceOK creates a WaypointServiceGetNamespaceOK with default headers values
func NewWaypointServiceGetNamespaceOK() *WaypointServiceGetNamespaceOK {
	return &WaypointServiceGetNamespaceOK{}
}

/*
WaypointServiceGetNamespaceOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceGetNamespaceOK struct {
	Payload *models.HashicorpCloudWaypointGetNamespaceResponse
}

// IsSuccess returns true when this waypoint service get namespace o k response has a 2xx status code
func (o *WaypointServiceGetNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service get namespace o k response has a 3xx status code
func (o *WaypointServiceGetNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service get namespace o k response has a 4xx status code
func (o *WaypointServiceGetNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service get namespace o k response has a 5xx status code
func (o *WaypointServiceGetNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service get namespace o k response a status code equal to that given
func (o *WaypointServiceGetNamespaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service get namespace o k response
func (o *WaypointServiceGetNamespaceOK) Code() int {
	return 200
}

func (o *WaypointServiceGetNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/organizations/{location.organization_id}/projects/{location.project_id}/namespaces][%d] waypointServiceGetNamespaceOK  %+v", 200, o.Payload)
}

func (o *WaypointServiceGetNamespaceOK) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/organizations/{location.organization_id}/projects/{location.project_id}/namespaces][%d] waypointServiceGetNamespaceOK  %+v", 200, o.Payload)
}

func (o *WaypointServiceGetNamespaceOK) GetPayload() *models.HashicorpCloudWaypointGetNamespaceResponse {
	return o.Payload
}

func (o *WaypointServiceGetNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointGetNamespaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceGetNamespaceDefault creates a WaypointServiceGetNamespaceDefault with default headers values
func NewWaypointServiceGetNamespaceDefault(code int) *WaypointServiceGetNamespaceDefault {
	return &WaypointServiceGetNamespaceDefault{
		_statusCode: code,
	}
}

/*
WaypointServiceGetNamespaceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceGetNamespaceDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service get namespace default response has a 2xx status code
func (o *WaypointServiceGetNamespaceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service get namespace default response has a 3xx status code
func (o *WaypointServiceGetNamespaceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service get namespace default response has a 4xx status code
func (o *WaypointServiceGetNamespaceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service get namespace default response has a 5xx status code
func (o *WaypointServiceGetNamespaceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service get namespace default response a status code equal to that given
func (o *WaypointServiceGetNamespaceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service get namespace default response
func (o *WaypointServiceGetNamespaceDefault) Code() int {
	return o._statusCode
}

func (o *WaypointServiceGetNamespaceDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/organizations/{location.organization_id}/projects/{location.project_id}/namespaces][%d] WaypointService_GetNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceGetNamespaceDefault) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/organizations/{location.organization_id}/projects/{location.project_id}/namespaces][%d] WaypointService_GetNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceGetNamespaceDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceGetNamespaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
