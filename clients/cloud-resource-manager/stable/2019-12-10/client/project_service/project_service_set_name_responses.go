// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
)

// ProjectServiceSetNameReader is a Reader for the ProjectServiceSetName structure.
type ProjectServiceSetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceSetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceSetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceSetNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceSetNameOK creates a ProjectServiceSetNameOK with default headers values
func NewProjectServiceSetNameOK() *ProjectServiceSetNameOK {
	return &ProjectServiceSetNameOK{}
}

/*
ProjectServiceSetNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectServiceSetNameOK struct {
	Payload interface{}
}

// IsSuccess returns true when this project service set name o k response has a 2xx status code
func (o *ProjectServiceSetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project service set name o k response has a 3xx status code
func (o *ProjectServiceSetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project service set name o k response has a 4xx status code
func (o *ProjectServiceSetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project service set name o k response has a 5xx status code
func (o *ProjectServiceSetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project service set name o k response a status code equal to that given
func (o *ProjectServiceSetNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectServiceSetNameOK) Error() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/name][%d] projectServiceSetNameOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceSetNameOK) String() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/name][%d] projectServiceSetNameOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceSetNameOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectServiceSetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceSetNameDefault creates a ProjectServiceSetNameDefault with default headers values
func NewProjectServiceSetNameDefault(code int) *ProjectServiceSetNameDefault {
	return &ProjectServiceSetNameDefault{
		_statusCode: code,
	}
}

/*
ProjectServiceSetNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectServiceSetNameDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the project service set name default response
func (o *ProjectServiceSetNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this project service set name default response has a 2xx status code
func (o *ProjectServiceSetNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this project service set name default response has a 3xx status code
func (o *ProjectServiceSetNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this project service set name default response has a 4xx status code
func (o *ProjectServiceSetNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this project service set name default response has a 5xx status code
func (o *ProjectServiceSetNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this project service set name default response a status code equal to that given
func (o *ProjectServiceSetNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProjectServiceSetNameDefault) Error() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/name][%d] ProjectService_SetName default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceSetNameDefault) String() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/name][%d] ProjectService_SetName default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceSetNameDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ProjectServiceSetNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ProjectServiceSetNameBody ProjectSetNameRequest see ProjectService.SetName
swagger:model ProjectServiceSetNameBody
*/
type ProjectServiceSetNameBody struct {

	// name is the value the project's name should be updated to.
	Name string `json:"name,omitempty"`
}

// Validate validates this project service set name body
func (o *ProjectServiceSetNameBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project service set name body based on context it is used
func (o *ProjectServiceSetNameBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectServiceSetNameBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectServiceSetNameBody) UnmarshalBinary(b []byte) error {
	var res ProjectServiceSetNameBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
