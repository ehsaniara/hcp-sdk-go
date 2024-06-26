// Code generated by go-swagger; DO NOT EDIT.

package registry_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/stable/2022-09-30/models"
)

// CompleteDirectUploadBoxReader is a Reader for the CompleteDirectUploadBox structure.
type CompleteDirectUploadBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteDirectUploadBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompleteDirectUploadBoxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCompleteDirectUploadBoxDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompleteDirectUploadBoxOK creates a CompleteDirectUploadBoxOK with default headers values
func NewCompleteDirectUploadBoxOK() *CompleteDirectUploadBoxOK {
	return &CompleteDirectUploadBoxOK{}
}

/*
CompleteDirectUploadBoxOK describes a response with status code 200, with default header values.

A successful response.
*/
type CompleteDirectUploadBoxOK struct {
	Payload models.HashicorpCloudVagrant20220930CompleteDirectUploadBoxResponse
}

// IsSuccess returns true when this complete direct upload box o k response has a 2xx status code
func (o *CompleteDirectUploadBoxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this complete direct upload box o k response has a 3xx status code
func (o *CompleteDirectUploadBoxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this complete direct upload box o k response has a 4xx status code
func (o *CompleteDirectUploadBoxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this complete direct upload box o k response has a 5xx status code
func (o *CompleteDirectUploadBoxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this complete direct upload box o k response a status code equal to that given
func (o *CompleteDirectUploadBoxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the complete direct upload box o k response
func (o *CompleteDirectUploadBoxOK) Code() int {
	return 200
}

func (o *CompleteDirectUploadBoxOK) Error() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}/architecture/{architecture}/direct/complete/{object}][%d] completeDirectUploadBoxOK  %+v", 200, o.Payload)
}

func (o *CompleteDirectUploadBoxOK) String() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}/architecture/{architecture}/direct/complete/{object}][%d] completeDirectUploadBoxOK  %+v", 200, o.Payload)
}

func (o *CompleteDirectUploadBoxOK) GetPayload() models.HashicorpCloudVagrant20220930CompleteDirectUploadBoxResponse {
	return o.Payload
}

func (o *CompleteDirectUploadBoxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteDirectUploadBoxDefault creates a CompleteDirectUploadBoxDefault with default headers values
func NewCompleteDirectUploadBoxDefault(code int) *CompleteDirectUploadBoxDefault {
	return &CompleteDirectUploadBoxDefault{
		_statusCode: code,
	}
}

/*
CompleteDirectUploadBoxDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CompleteDirectUploadBoxDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this complete direct upload box default response has a 2xx status code
func (o *CompleteDirectUploadBoxDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this complete direct upload box default response has a 3xx status code
func (o *CompleteDirectUploadBoxDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this complete direct upload box default response has a 4xx status code
func (o *CompleteDirectUploadBoxDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this complete direct upload box default response has a 5xx status code
func (o *CompleteDirectUploadBoxDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this complete direct upload box default response a status code equal to that given
func (o *CompleteDirectUploadBoxDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the complete direct upload box default response
func (o *CompleteDirectUploadBoxDefault) Code() int {
	return o._statusCode
}

func (o *CompleteDirectUploadBoxDefault) Error() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}/architecture/{architecture}/direct/complete/{object}][%d] CompleteDirectUploadBox default  %+v", o._statusCode, o.Payload)
}

func (o *CompleteDirectUploadBoxDefault) String() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry/{registry}/box/{box}/version/{version}/provider/{provider}/architecture/{architecture}/direct/complete/{object}][%d] CompleteDirectUploadBox default  %+v", o._statusCode, o.Payload)
}

func (o *CompleteDirectUploadBoxDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *CompleteDirectUploadBoxDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
