// Code generated by go-swagger; DO NOT EDIT.

package iam_service

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

// DeleteOrganizationMembershipReader is a Reader for the DeleteOrganizationMembership structure.
type DeleteOrganizationMembershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationMembershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationMembershipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteOrganizationMembershipDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrganizationMembershipOK creates a DeleteOrganizationMembershipOK with default headers values
func NewDeleteOrganizationMembershipOK() *DeleteOrganizationMembershipOK {
	return &DeleteOrganizationMembershipOK{}
}

/*
DeleteOrganizationMembershipOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteOrganizationMembershipOK struct {
	Payload models.HashicorpCloudIamDeleteOrganizationMembershipResponse
}

// IsSuccess returns true when this delete organization membership o k response has a 2xx status code
func (o *DeleteOrganizationMembershipOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization membership o k response has a 3xx status code
func (o *DeleteOrganizationMembershipOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization membership o k response has a 4xx status code
func (o *DeleteOrganizationMembershipOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization membership o k response has a 5xx status code
func (o *DeleteOrganizationMembershipOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization membership o k response a status code equal to that given
func (o *DeleteOrganizationMembershipOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteOrganizationMembershipOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}][%d] deleteOrganizationMembershipOK  %+v", 200, o.Payload)
}

func (o *DeleteOrganizationMembershipOK) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}][%d] deleteOrganizationMembershipOK  %+v", 200, o.Payload)
}

func (o *DeleteOrganizationMembershipOK) GetPayload() models.HashicorpCloudIamDeleteOrganizationMembershipResponse {
	return o.Payload
}

func (o *DeleteOrganizationMembershipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationMembershipDefault creates a DeleteOrganizationMembershipDefault with default headers values
func NewDeleteOrganizationMembershipDefault(code int) *DeleteOrganizationMembershipDefault {
	return &DeleteOrganizationMembershipDefault{
		_statusCode: code,
	}
}

/*
DeleteOrganizationMembershipDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteOrganizationMembershipDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the delete organization membership default response
func (o *DeleteOrganizationMembershipDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete organization membership default response has a 2xx status code
func (o *DeleteOrganizationMembershipDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete organization membership default response has a 3xx status code
func (o *DeleteOrganizationMembershipDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete organization membership default response has a 4xx status code
func (o *DeleteOrganizationMembershipDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete organization membership default response has a 5xx status code
func (o *DeleteOrganizationMembershipDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete organization membership default response a status code equal to that given
func (o *DeleteOrganizationMembershipDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteOrganizationMembershipDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}][%d] DeleteOrganizationMembership default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationMembershipDefault) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}][%d] DeleteOrganizationMembership default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationMembershipDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeleteOrganizationMembershipDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}