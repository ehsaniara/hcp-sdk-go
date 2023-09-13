// Code generated by go-swagger; DO NOT EDIT.

package invitations_service

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

// DeleteOrganizationInvitationReader is a Reader for the DeleteOrganizationInvitation structure.
type DeleteOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteOrganizationInvitationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrganizationInvitationOK creates a DeleteOrganizationInvitationOK with default headers values
func NewDeleteOrganizationInvitationOK() *DeleteOrganizationInvitationOK {
	return &DeleteOrganizationInvitationOK{}
}

/*
DeleteOrganizationInvitationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteOrganizationInvitationOK struct {
	Payload models.HashicorpCloudIamDeleteOrganizationInvitationResponse
}

// IsSuccess returns true when this delete organization invitation o k response has a 2xx status code
func (o *DeleteOrganizationInvitationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization invitation o k response has a 3xx status code
func (o *DeleteOrganizationInvitationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization invitation o k response has a 4xx status code
func (o *DeleteOrganizationInvitationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization invitation o k response has a 5xx status code
func (o *DeleteOrganizationInvitationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization invitation o k response a status code equal to that given
func (o *DeleteOrganizationInvitationOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/invitations/{invitation_id}][%d] deleteOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *DeleteOrganizationInvitationOK) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/invitations/{invitation_id}][%d] deleteOrganizationInvitationOK  %+v", 200, o.Payload)
}

func (o *DeleteOrganizationInvitationOK) GetPayload() models.HashicorpCloudIamDeleteOrganizationInvitationResponse {
	return o.Payload
}

func (o *DeleteOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationInvitationDefault creates a DeleteOrganizationInvitationDefault with default headers values
func NewDeleteOrganizationInvitationDefault(code int) *DeleteOrganizationInvitationDefault {
	return &DeleteOrganizationInvitationDefault{
		_statusCode: code,
	}
}

/*
DeleteOrganizationInvitationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteOrganizationInvitationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the delete organization invitation default response
func (o *DeleteOrganizationInvitationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete organization invitation default response has a 2xx status code
func (o *DeleteOrganizationInvitationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete organization invitation default response has a 3xx status code
func (o *DeleteOrganizationInvitationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete organization invitation default response has a 4xx status code
func (o *DeleteOrganizationInvitationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete organization invitation default response has a 5xx status code
func (o *DeleteOrganizationInvitationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete organization invitation default response a status code equal to that given
func (o *DeleteOrganizationInvitationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteOrganizationInvitationDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/invitations/{invitation_id}][%d] DeleteOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationInvitationDefault) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/invitations/{invitation_id}][%d] DeleteOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationInvitationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *DeleteOrganizationInvitationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}