// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2022-02-15/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// GetServiceSummariesReader is a Reader for the GetServiceSummaries structure.
type GetServiceSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServiceSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceSummariesOK creates a GetServiceSummariesOK with default headers values
func NewGetServiceSummariesOK() *GetServiceSummariesOK {
	return &GetServiceSummariesOK{}
}

/*
GetServiceSummariesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetServiceSummariesOK struct {
	Payload *models.HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesResponse
}

// IsSuccess returns true when this get service summaries o k response has a 2xx status code
func (o *GetServiceSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get service summaries o k response has a 3xx status code
func (o *GetServiceSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service summaries o k response has a 4xx status code
func (o *GetServiceSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get service summaries o k response has a 5xx status code
func (o *GetServiceSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get service summaries o k response a status code equal to that given
func (o *GetServiceSummariesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetServiceSummariesOK) Error() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/services][%d] getServiceSummariesOK  %+v", 200, o.Payload)
}

func (o *GetServiceSummariesOK) String() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/services][%d] getServiceSummariesOK  %+v", 200, o.Payload)
}

func (o *GetServiceSummariesOK) GetPayload() *models.HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesResponse {
	return o.Payload
}

func (o *GetServiceSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceSummariesDefault creates a GetServiceSummariesDefault with default headers values
func NewGetServiceSummariesDefault(code int) *GetServiceSummariesDefault {
	return &GetServiceSummariesDefault{
		_statusCode: code,
	}
}

/*
GetServiceSummariesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetServiceSummariesDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get service summaries default response
func (o *GetServiceSummariesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get service summaries default response has a 2xx status code
func (o *GetServiceSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get service summaries default response has a 3xx status code
func (o *GetServiceSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get service summaries default response has a 4xx status code
func (o *GetServiceSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get service summaries default response has a 5xx status code
func (o *GetServiceSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get service summaries default response a status code equal to that given
func (o *GetServiceSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetServiceSummariesDefault) Error() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/services][%d] GetServiceSummaries default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceSummariesDefault) String() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/services][%d] GetServiceSummaries default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceSummariesDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetServiceSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}