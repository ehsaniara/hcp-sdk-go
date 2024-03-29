// Code generated by go-swagger; DO NOT EDIT.

package groups_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// GroupsServiceUpdateGroup2Reader is a Reader for the GroupsServiceUpdateGroup2 structure.
type GroupsServiceUpdateGroup2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsServiceUpdateGroup2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsServiceUpdateGroup2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupsServiceUpdateGroup2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupsServiceUpdateGroup2OK creates a GroupsServiceUpdateGroup2OK with default headers values
func NewGroupsServiceUpdateGroup2OK() *GroupsServiceUpdateGroup2OK {
	return &GroupsServiceUpdateGroup2OK{}
}

/*
GroupsServiceUpdateGroup2OK describes a response with status code 200, with default header values.

A successful response.
*/
type GroupsServiceUpdateGroup2OK struct {
	Payload interface{}
}

// IsSuccess returns true when this groups service update group2 o k response has a 2xx status code
func (o *GroupsServiceUpdateGroup2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups service update group2 o k response has a 3xx status code
func (o *GroupsServiceUpdateGroup2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups service update group2 o k response has a 4xx status code
func (o *GroupsServiceUpdateGroup2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups service update group2 o k response has a 5xx status code
func (o *GroupsServiceUpdateGroup2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups service update group2 o k response a status code equal to that given
func (o *GroupsServiceUpdateGroup2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the groups service update group2 o k response
func (o *GroupsServiceUpdateGroup2OK) Code() int {
	return 200
}

func (o *GroupsServiceUpdateGroup2OK) Error() string {
	return fmt.Sprintf("[PATCH /iam/2019-12-10/{resource_name}][%d] groupsServiceUpdateGroup2OK  %+v", 200, o.Payload)
}

func (o *GroupsServiceUpdateGroup2OK) String() string {
	return fmt.Sprintf("[PATCH /iam/2019-12-10/{resource_name}][%d] groupsServiceUpdateGroup2OK  %+v", 200, o.Payload)
}

func (o *GroupsServiceUpdateGroup2OK) GetPayload() interface{} {
	return o.Payload
}

func (o *GroupsServiceUpdateGroup2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupsServiceUpdateGroup2Default creates a GroupsServiceUpdateGroup2Default with default headers values
func NewGroupsServiceUpdateGroup2Default(code int) *GroupsServiceUpdateGroup2Default {
	return &GroupsServiceUpdateGroup2Default{
		_statusCode: code,
	}
}

/*
GroupsServiceUpdateGroup2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GroupsServiceUpdateGroup2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this groups service update group2 default response has a 2xx status code
func (o *GroupsServiceUpdateGroup2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this groups service update group2 default response has a 3xx status code
func (o *GroupsServiceUpdateGroup2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this groups service update group2 default response has a 4xx status code
func (o *GroupsServiceUpdateGroup2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this groups service update group2 default response has a 5xx status code
func (o *GroupsServiceUpdateGroup2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this groups service update group2 default response a status code equal to that given
func (o *GroupsServiceUpdateGroup2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the groups service update group2 default response
func (o *GroupsServiceUpdateGroup2Default) Code() int {
	return o._statusCode
}

func (o *GroupsServiceUpdateGroup2Default) Error() string {
	return fmt.Sprintf("[PATCH /iam/2019-12-10/{resource_name}][%d] GroupsService_UpdateGroup2 default  %+v", o._statusCode, o.Payload)
}

func (o *GroupsServiceUpdateGroup2Default) String() string {
	return fmt.Sprintf("[PATCH /iam/2019-12-10/{resource_name}][%d] GroupsService_UpdateGroup2 default  %+v", o._statusCode, o.Payload)
}

func (o *GroupsServiceUpdateGroup2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GroupsServiceUpdateGroup2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}