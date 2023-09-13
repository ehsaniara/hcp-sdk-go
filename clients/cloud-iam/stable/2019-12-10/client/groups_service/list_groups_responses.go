// Code generated by go-swagger; DO NOT EDIT.

package groups_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListGroupsReader is a Reader for the ListGroups structure.
type ListGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGroupsOK creates a ListGroupsOK with default headers values
func NewListGroupsOK() *ListGroupsOK {
	return &ListGroupsOK{}
}

/*
ListGroupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListGroupsOK struct {
	Payload *models.HashicorpCloudIamListGroupsResponse
}

// IsSuccess returns true when this list groups o k response has a 2xx status code
func (o *ListGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list groups o k response has a 3xx status code
func (o *ListGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list groups o k response has a 4xx status code
func (o *ListGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list groups o k response has a 5xx status code
func (o *ListGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list groups o k response a status code equal to that given
func (o *ListGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGroupsOK) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/iam/{parent_resource_name}/groups][%d] listGroupsOK  %+v", 200, o.Payload)
}

func (o *ListGroupsOK) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/iam/{parent_resource_name}/groups][%d] listGroupsOK  %+v", 200, o.Payload)
}

func (o *ListGroupsOK) GetPayload() *models.HashicorpCloudIamListGroupsResponse {
	return o.Payload
}

func (o *ListGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamListGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupsDefault creates a ListGroupsDefault with default headers values
func NewListGroupsDefault(code int) *ListGroupsDefault {
	return &ListGroupsDefault{
		_statusCode: code,
	}
}

/*
ListGroupsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListGroupsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the list groups default response
func (o *ListGroupsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list groups default response has a 2xx status code
func (o *ListGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list groups default response has a 3xx status code
func (o *ListGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list groups default response has a 4xx status code
func (o *ListGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list groups default response has a 5xx status code
func (o *ListGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list groups default response a status code equal to that given
func (o *ListGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/iam/{parent_resource_name}/groups][%d] ListGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListGroupsDefault) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/iam/{parent_resource_name}/groups][%d] ListGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListGroupsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ListGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}