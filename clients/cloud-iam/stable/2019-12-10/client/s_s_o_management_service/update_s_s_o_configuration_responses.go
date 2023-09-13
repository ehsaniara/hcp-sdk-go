// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

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

// UpdateSSOConfigurationReader is a Reader for the UpdateSSOConfiguration structure.
type UpdateSSOConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSSOConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSSOConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSSOConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSSOConfigurationOK creates a UpdateSSOConfigurationOK with default headers values
func NewUpdateSSOConfigurationOK() *UpdateSSOConfigurationOK {
	return &UpdateSSOConfigurationOK{}
}

/*
UpdateSSOConfigurationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateSSOConfigurationOK struct {
	Payload models.HashicorpCloudIamUpdateSSOConfigurationResponse
}

// IsSuccess returns true when this update s s o configuration o k response has a 2xx status code
func (o *UpdateSSOConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update s s o configuration o k response has a 3xx status code
func (o *UpdateSSOConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update s s o configuration o k response has a 4xx status code
func (o *UpdateSSOConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update s s o configuration o k response has a 5xx status code
func (o *UpdateSSOConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update s s o configuration o k response a status code equal to that given
func (o *UpdateSSOConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateSSOConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] updateSSOConfigurationOK  %+v", 200, o.Payload)
}

func (o *UpdateSSOConfigurationOK) String() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] updateSSOConfigurationOK  %+v", 200, o.Payload)
}

func (o *UpdateSSOConfigurationOK) GetPayload() models.HashicorpCloudIamUpdateSSOConfigurationResponse {
	return o.Payload
}

func (o *UpdateSSOConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSSOConfigurationDefault creates a UpdateSSOConfigurationDefault with default headers values
func NewUpdateSSOConfigurationDefault(code int) *UpdateSSOConfigurationDefault {
	return &UpdateSSOConfigurationDefault{
		_statusCode: code,
	}
}

/*
UpdateSSOConfigurationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateSSOConfigurationDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the update s s o configuration default response
func (o *UpdateSSOConfigurationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update s s o configuration default response has a 2xx status code
func (o *UpdateSSOConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update s s o configuration default response has a 3xx status code
func (o *UpdateSSOConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update s s o configuration default response has a 4xx status code
func (o *UpdateSSOConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update s s o configuration default response has a 5xx status code
func (o *UpdateSSOConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update s s o configuration default response a status code equal to that given
func (o *UpdateSSOConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateSSOConfigurationDefault) Error() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] UpdateSSOConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSSOConfigurationDefault) String() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/organizations/{organization_id}/sso-configurations/{type}][%d] UpdateSSOConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSSOConfigurationDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *UpdateSSOConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}