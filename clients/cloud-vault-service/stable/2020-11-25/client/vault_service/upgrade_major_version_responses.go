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

// UpgradeMajorVersionReader is a Reader for the UpgradeMajorVersion structure.
type UpgradeMajorVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeMajorVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpgradeMajorVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpgradeMajorVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpgradeMajorVersionOK creates a UpgradeMajorVersionOK with default headers values
func NewUpgradeMajorVersionOK() *UpgradeMajorVersionOK {
	return &UpgradeMajorVersionOK{}
}

/*UpgradeMajorVersionOK handles this case with default header values.

A successful response.
*/
type UpgradeMajorVersionOK struct {
	Payload *models.HashicorpCloudVault20201125UpgradeMajorVersionResponse
}

func (o *UpgradeMajorVersionOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/major-version/upgrade][%d] upgradeMajorVersionOK  %+v", 200, o.Payload)
}

func (o *UpgradeMajorVersionOK) GetPayload() *models.HashicorpCloudVault20201125UpgradeMajorVersionResponse {
	return o.Payload
}

func (o *UpgradeMajorVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125UpgradeMajorVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpgradeMajorVersionDefault creates a UpgradeMajorVersionDefault with default headers values
func NewUpgradeMajorVersionDefault(code int) *UpgradeMajorVersionDefault {
	return &UpgradeMajorVersionDefault{
		_statusCode: code,
	}
}

/*UpgradeMajorVersionDefault handles this case with default header values.

An unexpected error response.
*/
type UpgradeMajorVersionDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the upgrade major version default response
func (o *UpgradeMajorVersionDefault) Code() int {
	return o._statusCode
}

func (o *UpgradeMajorVersionDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/major-version/upgrade][%d] UpgradeMajorVersion default  %+v", o._statusCode, o.Payload)
}

func (o *UpgradeMajorVersionDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *UpgradeMajorVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
