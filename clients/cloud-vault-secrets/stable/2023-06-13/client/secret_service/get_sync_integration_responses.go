// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/models"
)

// GetSyncIntegrationReader is a Reader for the GetSyncIntegration structure.
type GetSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSyncIntegrationOK creates a GetSyncIntegrationOK with default headers values
func NewGetSyncIntegrationOK() *GetSyncIntegrationOK {
	return &GetSyncIntegrationOK{}
}

/*
GetSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetSyncIntegrationOK struct {
	Payload *models.Secrets20230613GetSyncIntegrationResponse
}

// IsSuccess returns true when this get sync integration o k response has a 2xx status code
func (o *GetSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sync integration o k response has a 3xx status code
func (o *GetSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sync integration o k response has a 4xx status code
func (o *GetSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sync integration o k response has a 5xx status code
func (o *GetSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sync integration o k response a status code equal to that given
func (o *GetSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sync integration o k response
func (o *GetSyncIntegrationOK) Code() int {
	return 200
}

func (o *GetSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/integrations/{name}][%d] getSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetSyncIntegrationOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/integrations/{name}][%d] getSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetSyncIntegrationOK) GetPayload() *models.Secrets20230613GetSyncIntegrationResponse {
	return o.Payload
}

func (o *GetSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613GetSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyncIntegrationDefault creates a GetSyncIntegrationDefault with default headers values
func NewGetSyncIntegrationDefault(code int) *GetSyncIntegrationDefault {
	return &GetSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
GetSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get sync integration default response has a 2xx status code
func (o *GetSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get sync integration default response has a 3xx status code
func (o *GetSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get sync integration default response has a 4xx status code
func (o *GetSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get sync integration default response has a 5xx status code
func (o *GetSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get sync integration default response a status code equal to that given
func (o *GetSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get sync integration default response
func (o *GetSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *GetSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/integrations/{name}][%d] GetSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *GetSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/integrations/{name}][%d] GetSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *GetSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
