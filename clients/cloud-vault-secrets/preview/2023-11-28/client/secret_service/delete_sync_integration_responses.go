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

// DeleteSyncIntegrationReader is a Reader for the DeleteSyncIntegration structure.
type DeleteSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSyncIntegrationOK creates a DeleteSyncIntegrationOK with default headers values
func NewDeleteSyncIntegrationOK() *DeleteSyncIntegrationOK {
	return &DeleteSyncIntegrationOK{}
}

/*
DeleteSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteSyncIntegrationOK struct {
	Payload models.Secrets20231128DeleteSyncIntegrationResponse
}

// IsSuccess returns true when this delete sync integration o k response has a 2xx status code
func (o *DeleteSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete sync integration o k response has a 3xx status code
func (o *DeleteSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sync integration o k response has a 4xx status code
func (o *DeleteSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete sync integration o k response has a 5xx status code
func (o *DeleteSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sync integration o k response a status code equal to that given
func (o *DeleteSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete sync integration o k response
func (o *DeleteSyncIntegrationOK) Code() int {
	return 200
}

func (o *DeleteSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/integrations/{name}][%d] deleteSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *DeleteSyncIntegrationOK) String() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/integrations/{name}][%d] deleteSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *DeleteSyncIntegrationOK) GetPayload() models.Secrets20231128DeleteSyncIntegrationResponse {
	return o.Payload
}

func (o *DeleteSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSyncIntegrationDefault creates a DeleteSyncIntegrationDefault with default headers values
func NewDeleteSyncIntegrationDefault(code int) *DeleteSyncIntegrationDefault {
	return &DeleteSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
DeleteSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete sync integration default response has a 2xx status code
func (o *DeleteSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete sync integration default response has a 3xx status code
func (o *DeleteSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete sync integration default response has a 4xx status code
func (o *DeleteSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete sync integration default response has a 5xx status code
func (o *DeleteSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete sync integration default response a status code equal to that given
func (o *DeleteSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete sync integration default response
func (o *DeleteSyncIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/integrations/{name}][%d] DeleteSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[DELETE /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/integrations/{name}][%d] DeleteSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSyncIntegrationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
