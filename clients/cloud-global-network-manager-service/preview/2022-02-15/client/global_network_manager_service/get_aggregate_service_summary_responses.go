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

// GetAggregateServiceSummaryReader is a Reader for the GetAggregateServiceSummary structure.
type GetAggregateServiceSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAggregateServiceSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAggregateServiceSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAggregateServiceSummaryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAggregateServiceSummaryOK creates a GetAggregateServiceSummaryOK with default headers values
func NewGetAggregateServiceSummaryOK() *GetAggregateServiceSummaryOK {
	return &GetAggregateServiceSummaryOK{}
}

/*
GetAggregateServiceSummaryOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAggregateServiceSummaryOK struct {
	Payload *models.HashicorpCloudGlobalNetworkManager20220215GetAggregateServiceSummaryResponse
}

// IsSuccess returns true when this get aggregate service summary o k response has a 2xx status code
func (o *GetAggregateServiceSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get aggregate service summary o k response has a 3xx status code
func (o *GetAggregateServiceSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get aggregate service summary o k response has a 4xx status code
func (o *GetAggregateServiceSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get aggregate service summary o k response has a 5xx status code
func (o *GetAggregateServiceSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get aggregate service summary o k response a status code equal to that given
func (o *GetAggregateServiceSummaryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAggregateServiceSummaryOK) Error() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/aggregate_service_summary][%d] getAggregateServiceSummaryOK  %+v", 200, o.Payload)
}

func (o *GetAggregateServiceSummaryOK) String() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/aggregate_service_summary][%d] getAggregateServiceSummaryOK  %+v", 200, o.Payload)
}

func (o *GetAggregateServiceSummaryOK) GetPayload() *models.HashicorpCloudGlobalNetworkManager20220215GetAggregateServiceSummaryResponse {
	return o.Payload
}

func (o *GetAggregateServiceSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudGlobalNetworkManager20220215GetAggregateServiceSummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAggregateServiceSummaryDefault creates a GetAggregateServiceSummaryDefault with default headers values
func NewGetAggregateServiceSummaryDefault(code int) *GetAggregateServiceSummaryDefault {
	return &GetAggregateServiceSummaryDefault{
		_statusCode: code,
	}
}

/*
GetAggregateServiceSummaryDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAggregateServiceSummaryDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get aggregate service summary default response
func (o *GetAggregateServiceSummaryDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get aggregate service summary default response has a 2xx status code
func (o *GetAggregateServiceSummaryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get aggregate service summary default response has a 3xx status code
func (o *GetAggregateServiceSummaryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get aggregate service summary default response has a 4xx status code
func (o *GetAggregateServiceSummaryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get aggregate service summary default response has a 5xx status code
func (o *GetAggregateServiceSummaryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get aggregate service summary default response a status code equal to that given
func (o *GetAggregateServiceSummaryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAggregateServiceSummaryDefault) Error() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/aggregate_service_summary][%d] GetAggregateServiceSummary default  %+v", o._statusCode, o.Payload)
}

func (o *GetAggregateServiceSummaryDefault) String() string {
	return fmt.Sprintf("[GET /global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/aggregate_service_summary][%d] GetAggregateServiceSummary default  %+v", o._statusCode, o.Payload)
}

func (o *GetAggregateServiceSummaryDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetAggregateServiceSummaryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}