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

// GetMongoDBAtlasRotatingSecretConfigReader is a Reader for the GetMongoDBAtlasRotatingSecretConfig structure.
type GetMongoDBAtlasRotatingSecretConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMongoDBAtlasRotatingSecretConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMongoDBAtlasRotatingSecretConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMongoDBAtlasRotatingSecretConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMongoDBAtlasRotatingSecretConfigOK creates a GetMongoDBAtlasRotatingSecretConfigOK with default headers values
func NewGetMongoDBAtlasRotatingSecretConfigOK() *GetMongoDBAtlasRotatingSecretConfigOK {
	return &GetMongoDBAtlasRotatingSecretConfigOK{}
}

/*
GetMongoDBAtlasRotatingSecretConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetMongoDBAtlasRotatingSecretConfigOK struct {
	Payload *models.Secrets20231128GetMongoDBAtlasRotatingSecretConfigResponse
}

// IsSuccess returns true when this get mongo d b atlas rotating secret config o k response has a 2xx status code
func (o *GetMongoDBAtlasRotatingSecretConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mongo d b atlas rotating secret config o k response has a 3xx status code
func (o *GetMongoDBAtlasRotatingSecretConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mongo d b atlas rotating secret config o k response has a 4xx status code
func (o *GetMongoDBAtlasRotatingSecretConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mongo d b atlas rotating secret config o k response has a 5xx status code
func (o *GetMongoDBAtlasRotatingSecretConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mongo d b atlas rotating secret config o k response a status code equal to that given
func (o *GetMongoDBAtlasRotatingSecretConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mongo d b atlas rotating secret config o k response
func (o *GetMongoDBAtlasRotatingSecretConfigOK) Code() int {
	return 200
}

func (o *GetMongoDBAtlasRotatingSecretConfigOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secret/rotating/mongodb-atlas/{secret_name}][%d] getMongoDBAtlasRotatingSecretConfigOK  %+v", 200, o.Payload)
}

func (o *GetMongoDBAtlasRotatingSecretConfigOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secret/rotating/mongodb-atlas/{secret_name}][%d] getMongoDBAtlasRotatingSecretConfigOK  %+v", 200, o.Payload)
}

func (o *GetMongoDBAtlasRotatingSecretConfigOK) GetPayload() *models.Secrets20231128GetMongoDBAtlasRotatingSecretConfigResponse {
	return o.Payload
}

func (o *GetMongoDBAtlasRotatingSecretConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128GetMongoDBAtlasRotatingSecretConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMongoDBAtlasRotatingSecretConfigDefault creates a GetMongoDBAtlasRotatingSecretConfigDefault with default headers values
func NewGetMongoDBAtlasRotatingSecretConfigDefault(code int) *GetMongoDBAtlasRotatingSecretConfigDefault {
	return &GetMongoDBAtlasRotatingSecretConfigDefault{
		_statusCode: code,
	}
}

/*
GetMongoDBAtlasRotatingSecretConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetMongoDBAtlasRotatingSecretConfigDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get mongo d b atlas rotating secret config default response has a 2xx status code
func (o *GetMongoDBAtlasRotatingSecretConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get mongo d b atlas rotating secret config default response has a 3xx status code
func (o *GetMongoDBAtlasRotatingSecretConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get mongo d b atlas rotating secret config default response has a 4xx status code
func (o *GetMongoDBAtlasRotatingSecretConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get mongo d b atlas rotating secret config default response has a 5xx status code
func (o *GetMongoDBAtlasRotatingSecretConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get mongo d b atlas rotating secret config default response a status code equal to that given
func (o *GetMongoDBAtlasRotatingSecretConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get mongo d b atlas rotating secret config default response
func (o *GetMongoDBAtlasRotatingSecretConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetMongoDBAtlasRotatingSecretConfigDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secret/rotating/mongodb-atlas/{secret_name}][%d] GetMongoDBAtlasRotatingSecretConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetMongoDBAtlasRotatingSecretConfigDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/apps/{app_name}/secret/rotating/mongodb-atlas/{secret_name}][%d] GetMongoDBAtlasRotatingSecretConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetMongoDBAtlasRotatingSecretConfigDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetMongoDBAtlasRotatingSecretConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}