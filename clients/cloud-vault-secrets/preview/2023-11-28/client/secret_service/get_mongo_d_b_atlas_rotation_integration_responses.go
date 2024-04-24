// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// GetMongoDBAtlasRotationIntegrationReader is a Reader for the GetMongoDBAtlasRotationIntegration structure.
type GetMongoDBAtlasRotationIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMongoDBAtlasRotationIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMongoDBAtlasRotationIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMongoDBAtlasRotationIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMongoDBAtlasRotationIntegrationOK creates a GetMongoDBAtlasRotationIntegrationOK with default headers values
func NewGetMongoDBAtlasRotationIntegrationOK() *GetMongoDBAtlasRotationIntegrationOK {
	return &GetMongoDBAtlasRotationIntegrationOK{}
}

/*
GetMongoDBAtlasRotationIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetMongoDBAtlasRotationIntegrationOK struct {
	Payload *models.Secrets20231128GetMongoDBAtlasRotationIntegrationResponse
}

// IsSuccess returns true when this get mongo d b atlas rotation integration o k response has a 2xx status code
func (o *GetMongoDBAtlasRotationIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mongo d b atlas rotation integration o k response has a 3xx status code
func (o *GetMongoDBAtlasRotationIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mongo d b atlas rotation integration o k response has a 4xx status code
func (o *GetMongoDBAtlasRotationIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mongo d b atlas rotation integration o k response has a 5xx status code
func (o *GetMongoDBAtlasRotationIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mongo d b atlas rotation integration o k response a status code equal to that given
func (o *GetMongoDBAtlasRotationIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mongo d b atlas rotation integration o k response
func (o *GetMongoDBAtlasRotationIntegrationOK) Code() int {
	return 200
}

func (o *GetMongoDBAtlasRotationIntegrationOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/rotation/mongodb-atlas/{integration_name}][%d] getMongoDBAtlasRotationIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetMongoDBAtlasRotationIntegrationOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/rotation/mongodb-atlas/{integration_name}][%d] getMongoDBAtlasRotationIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetMongoDBAtlasRotationIntegrationOK) GetPayload() *models.Secrets20231128GetMongoDBAtlasRotationIntegrationResponse {
	return o.Payload
}

func (o *GetMongoDBAtlasRotationIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetMongoDBAtlasRotationIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMongoDBAtlasRotationIntegrationDefault creates a GetMongoDBAtlasRotationIntegrationDefault with default headers values
func NewGetMongoDBAtlasRotationIntegrationDefault(code int) *GetMongoDBAtlasRotationIntegrationDefault {
	return &GetMongoDBAtlasRotationIntegrationDefault{
		_statusCode: code,
	}
}

/*
GetMongoDBAtlasRotationIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetMongoDBAtlasRotationIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get mongo d b atlas rotation integration default response has a 2xx status code
func (o *GetMongoDBAtlasRotationIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get mongo d b atlas rotation integration default response has a 3xx status code
func (o *GetMongoDBAtlasRotationIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get mongo d b atlas rotation integration default response has a 4xx status code
func (o *GetMongoDBAtlasRotationIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get mongo d b atlas rotation integration default response has a 5xx status code
func (o *GetMongoDBAtlasRotationIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get mongo d b atlas rotation integration default response a status code equal to that given
func (o *GetMongoDBAtlasRotationIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get mongo d b atlas rotation integration default response
func (o *GetMongoDBAtlasRotationIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *GetMongoDBAtlasRotationIntegrationDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/rotation/mongodb-atlas/{integration_name}][%d] GetMongoDBAtlasRotationIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *GetMongoDBAtlasRotationIntegrationDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/integrations/rotation/mongodb-atlas/{integration_name}][%d] GetMongoDBAtlasRotationIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *GetMongoDBAtlasRotationIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetMongoDBAtlasRotationIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
