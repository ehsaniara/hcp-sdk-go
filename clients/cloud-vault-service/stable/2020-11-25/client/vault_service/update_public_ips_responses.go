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
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// UpdatePublicIpsReader is a Reader for the UpdatePublicIps structure.
type UpdatePublicIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePublicIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePublicIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdatePublicIpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePublicIpsOK creates a UpdatePublicIpsOK with default headers values
func NewUpdatePublicIpsOK() *UpdatePublicIpsOK {
	return &UpdatePublicIpsOK{}
}

/* UpdatePublicIpsOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdatePublicIpsOK struct {
	Payload *models.HashicorpCloudVault20201125UpdatePublicIpsResponse
}

func (o *UpdatePublicIpsOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/public-ips][%d] updatePublicIpsOK  %+v", 200, o.Payload)
}
func (o *UpdatePublicIpsOK) GetPayload() *models.HashicorpCloudVault20201125UpdatePublicIpsResponse {
	return o.Payload
}

func (o *UpdatePublicIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125UpdatePublicIpsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePublicIpsDefault creates a UpdatePublicIpsDefault with default headers values
func NewUpdatePublicIpsDefault(code int) *UpdatePublicIpsDefault {
	return &UpdatePublicIpsDefault{
		_statusCode: code,
	}
}

/* UpdatePublicIpsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdatePublicIpsDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the update public ips default response
func (o *UpdatePublicIpsDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePublicIpsDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/public-ips][%d] UpdatePublicIps default  %+v", o._statusCode, o.Payload)
}
func (o *UpdatePublicIpsDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *UpdatePublicIpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
