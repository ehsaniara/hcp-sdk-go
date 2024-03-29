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

// WaypointServiceUIGetApplicationTemplateBundle2Reader is a Reader for the WaypointServiceUIGetApplicationTemplateBundle2 structure.
type WaypointServiceUIGetApplicationTemplateBundle2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointServiceUIGetApplicationTemplateBundle2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointServiceUIGetApplicationTemplateBundle2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointServiceUIGetApplicationTemplateBundle2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointServiceUIGetApplicationTemplateBundle2OK creates a WaypointServiceUIGetApplicationTemplateBundle2OK with default headers values
func NewWaypointServiceUIGetApplicationTemplateBundle2OK() *WaypointServiceUIGetApplicationTemplateBundle2OK {
	return &WaypointServiceUIGetApplicationTemplateBundle2OK{}
}

/*
WaypointServiceUIGetApplicationTemplateBundle2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointServiceUIGetApplicationTemplateBundle2OK struct {
	Payload *models.HashicorpCloudWaypointUIGetApplicationTemplateBundleResponse
}

// IsSuccess returns true when this waypoint service Ui get application template bundle2 o k response has a 2xx status code
func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint service Ui get application template bundle2 o k response has a 3xx status code
func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint service Ui get application template bundle2 o k response has a 4xx status code
func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint service Ui get application template bundle2 o k response has a 5xx status code
func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint service Ui get application template bundle2 o k response a status code equal to that given
func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the waypoint service Ui get application template bundle2 o k response
func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) Code() int {
	return 200
}

func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/application-templates/by-name/{application_template.name}][%d] waypointServiceUiGetApplicationTemplateBundle2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/application-templates/by-name/{application_template.name}][%d] waypointServiceUiGetApplicationTemplateBundle2OK  %+v", 200, o.Payload)
}

func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) GetPayload() *models.HashicorpCloudWaypointUIGetApplicationTemplateBundleResponse {
	return o.Payload
}

func (o *WaypointServiceUIGetApplicationTemplateBundle2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointUIGetApplicationTemplateBundleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointServiceUIGetApplicationTemplateBundle2Default creates a WaypointServiceUIGetApplicationTemplateBundle2Default with default headers values
func NewWaypointServiceUIGetApplicationTemplateBundle2Default(code int) *WaypointServiceUIGetApplicationTemplateBundle2Default {
	return &WaypointServiceUIGetApplicationTemplateBundle2Default{
		_statusCode: code,
	}
}

/*
WaypointServiceUIGetApplicationTemplateBundle2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointServiceUIGetApplicationTemplateBundle2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this waypoint service UI get application template bundle2 default response has a 2xx status code
func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint service UI get application template bundle2 default response has a 3xx status code
func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint service UI get application template bundle2 default response has a 4xx status code
func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint service UI get application template bundle2 default response has a 5xx status code
func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint service UI get application template bundle2 default response a status code equal to that given
func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the waypoint service UI get application template bundle2 default response
func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) Code() int {
	return o._statusCode
}

func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/application-templates/by-name/{application_template.name}][%d] WaypointService_UI_GetApplicationTemplateBundle2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) String() string {
	return fmt.Sprintf("[GET /waypoint/2023-08-18/namespace/{namespace.id}/ui/application-templates/by-name/{application_template.name}][%d] WaypointService_UI_GetApplicationTemplateBundle2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *WaypointServiceUIGetApplicationTemplateBundle2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
