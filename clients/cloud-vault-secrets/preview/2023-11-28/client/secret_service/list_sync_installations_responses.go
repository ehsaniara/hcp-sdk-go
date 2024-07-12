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

// ListSyncInstallationsReader is a Reader for the ListSyncInstallations structure.
type ListSyncInstallationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSyncInstallationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSyncInstallationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSyncInstallationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSyncInstallationsOK creates a ListSyncInstallationsOK with default headers values
func NewListSyncInstallationsOK() *ListSyncInstallationsOK {
	return &ListSyncInstallationsOK{}
}

/*
ListSyncInstallationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListSyncInstallationsOK struct {
	Payload *models.Secrets20231128ListSyncInstallationsResponse
}

// IsSuccess returns true when this list sync installations o k response has a 2xx status code
func (o *ListSyncInstallationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list sync installations o k response has a 3xx status code
func (o *ListSyncInstallationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list sync installations o k response has a 4xx status code
func (o *ListSyncInstallationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list sync installations o k response has a 5xx status code
func (o *ListSyncInstallationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list sync installations o k response a status code equal to that given
func (o *ListSyncInstallationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list sync installations o k response
func (o *ListSyncInstallationsOK) Code() int {
	return 200
}

func (o *ListSyncInstallationsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/installations][%d] listSyncInstallationsOK  %+v", 200, o.Payload)
}

func (o *ListSyncInstallationsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/installations][%d] listSyncInstallationsOK  %+v", 200, o.Payload)
}

func (o *ListSyncInstallationsOK) GetPayload() *models.Secrets20231128ListSyncInstallationsResponse {
	return o.Payload
}

func (o *ListSyncInstallationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20231128ListSyncInstallationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSyncInstallationsDefault creates a ListSyncInstallationsDefault with default headers values
func NewListSyncInstallationsDefault(code int) *ListSyncInstallationsDefault {
	return &ListSyncInstallationsDefault{
		_statusCode: code,
	}
}

/*
ListSyncInstallationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListSyncInstallationsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this list sync installations default response has a 2xx status code
func (o *ListSyncInstallationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list sync installations default response has a 3xx status code
func (o *ListSyncInstallationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list sync installations default response has a 4xx status code
func (o *ListSyncInstallationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list sync installations default response has a 5xx status code
func (o *ListSyncInstallationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list sync installations default response a status code equal to that given
func (o *ListSyncInstallationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list sync installations default response
func (o *ListSyncInstallationsDefault) Code() int {
	return o._statusCode
}

func (o *ListSyncInstallationsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/installations][%d] ListSyncInstallations default  %+v", o._statusCode, o.Payload)
}

func (o *ListSyncInstallationsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-11-28/organizations/{organization_id}/projects/{project_id}/sync/installations][%d] ListSyncInstallations default  %+v", o._statusCode, o.Payload)
}

func (o *ListSyncInstallationsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ListSyncInstallationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
