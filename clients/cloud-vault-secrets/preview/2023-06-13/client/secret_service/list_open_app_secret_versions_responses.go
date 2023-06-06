// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-06-13/models"
)

// ListOpenAppSecretVersionsReader is a Reader for the ListOpenAppSecretVersions structure.
type ListOpenAppSecretVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenAppSecretVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenAppSecretVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenAppSecretVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenAppSecretVersionsOK creates a ListOpenAppSecretVersionsOK with default headers values
func NewListOpenAppSecretVersionsOK() *ListOpenAppSecretVersionsOK {
	return &ListOpenAppSecretVersionsOK{}
}

/*
ListOpenAppSecretVersionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListOpenAppSecretVersionsOK struct {
	Payload *models.Secrets20230613ListOpenAppSecretVersionsResponse
}

// IsSuccess returns true when this list open app secret versions o k response has a 2xx status code
func (o *ListOpenAppSecretVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list open app secret versions o k response has a 3xx status code
func (o *ListOpenAppSecretVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list open app secret versions o k response has a 4xx status code
func (o *ListOpenAppSecretVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list open app secret versions o k response has a 5xx status code
func (o *ListOpenAppSecretVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list open app secret versions o k response a status code equal to that given
func (o *ListOpenAppSecretVersionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListOpenAppSecretVersionsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/open/{secret_name}/versions][%d] listOpenAppSecretVersionsOK  %+v", 200, o.Payload)
}

func (o *ListOpenAppSecretVersionsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/open/{secret_name}/versions][%d] listOpenAppSecretVersionsOK  %+v", 200, o.Payload)
}

func (o *ListOpenAppSecretVersionsOK) GetPayload() *models.Secrets20230613ListOpenAppSecretVersionsResponse {
	return o.Payload
}

func (o *ListOpenAppSecretVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613ListOpenAppSecretVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenAppSecretVersionsDefault creates a ListOpenAppSecretVersionsDefault with default headers values
func NewListOpenAppSecretVersionsDefault(code int) *ListOpenAppSecretVersionsDefault {
	return &ListOpenAppSecretVersionsDefault{
		_statusCode: code,
	}
}

/*
ListOpenAppSecretVersionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListOpenAppSecretVersionsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the list open app secret versions default response
func (o *ListOpenAppSecretVersionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list open app secret versions default response has a 2xx status code
func (o *ListOpenAppSecretVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list open app secret versions default response has a 3xx status code
func (o *ListOpenAppSecretVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list open app secret versions default response has a 4xx status code
func (o *ListOpenAppSecretVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list open app secret versions default response has a 5xx status code
func (o *ListOpenAppSecretVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list open app secret versions default response a status code equal to that given
func (o *ListOpenAppSecretVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListOpenAppSecretVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/open/{secret_name}/versions][%d] ListOpenAppSecretVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenAppSecretVersionsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/open/{secret_name}/versions][%d] ListOpenAppSecretVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenAppSecretVersionsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ListOpenAppSecretVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
