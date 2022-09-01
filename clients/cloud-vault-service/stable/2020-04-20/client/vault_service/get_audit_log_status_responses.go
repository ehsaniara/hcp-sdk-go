// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-04-20/models"
)

// GetAuditLogStatusReader is a Reader for the GetAuditLogStatus structure.
type GetAuditLogStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditLogStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditLogStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAuditLogStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditLogStatusOK creates a GetAuditLogStatusOK with default headers values
func NewGetAuditLogStatusOK() *GetAuditLogStatusOK {
	return &GetAuditLogStatusOK{}
}

/* GetAuditLogStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAuditLogStatusOK struct {
	Payload *models.HashicorpCloudVault20200420GetAuditLogStatusResponse
}

func (o *GetAuditLogStatusOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog/{log_id}][%d] getAuditLogStatusOK  %+v", 200, o.Payload)
}
func (o *GetAuditLogStatusOK) GetPayload() *models.HashicorpCloudVault20200420GetAuditLogStatusResponse {
	return o.Payload
}

func (o *GetAuditLogStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20200420GetAuditLogStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditLogStatusDefault creates a GetAuditLogStatusDefault with default headers values
func NewGetAuditLogStatusDefault(code int) *GetAuditLogStatusDefault {
	return &GetAuditLogStatusDefault{
		_statusCode: code,
	}
}

/* GetAuditLogStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAuditLogStatusDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get audit log status default response
func (o *GetAuditLogStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditLogStatusDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-04-20/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog/{log_id}][%d] GetAuditLogStatus default  %+v", o._statusCode, o.Payload)
}
func (o *GetAuditLogStatusDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetAuditLogStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
