// Code generated by go-swagger; DO NOT EDIT.

package consul_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2020-08-26/models"
)

// ConsulServiceDeleteSnapshotReader is a Reader for the ConsulServiceDeleteSnapshot structure.
type ConsulServiceDeleteSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsulServiceDeleteSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsulServiceDeleteSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsulServiceDeleteSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsulServiceDeleteSnapshotOK creates a ConsulServiceDeleteSnapshotOK with default headers values
func NewConsulServiceDeleteSnapshotOK() *ConsulServiceDeleteSnapshotOK {
	return &ConsulServiceDeleteSnapshotOK{}
}

/*ConsulServiceDeleteSnapshotOK handles this case with default header values.

A successful response.
*/
type ConsulServiceDeleteSnapshotOK struct {
	Payload *models.HashicorpCloudConsul20200826DeleteSnapshotResponse
}

func (o *ConsulServiceDeleteSnapshotOK) Error() string {
	return fmt.Sprintf("[DELETE /consul/2020-08-26/organizations/{location.organization_id}/projects/{location.project_id}/snapshots/{snapshot_id}][%d] consulServiceDeleteSnapshotOK  %+v", 200, o.Payload)
}

func (o *ConsulServiceDeleteSnapshotOK) GetPayload() *models.HashicorpCloudConsul20200826DeleteSnapshotResponse {
	return o.Payload
}

func (o *ConsulServiceDeleteSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudConsul20200826DeleteSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsulServiceDeleteSnapshotDefault creates a ConsulServiceDeleteSnapshotDefault with default headers values
func NewConsulServiceDeleteSnapshotDefault(code int) *ConsulServiceDeleteSnapshotDefault {
	return &ConsulServiceDeleteSnapshotDefault{
		_statusCode: code,
	}
}

/*ConsulServiceDeleteSnapshotDefault handles this case with default header values.

An unexpected error response.
*/
type ConsulServiceDeleteSnapshotDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the consul service delete snapshot default response
func (o *ConsulServiceDeleteSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *ConsulServiceDeleteSnapshotDefault) Error() string {
	return fmt.Sprintf("[DELETE /consul/2020-08-26/organizations/{location.organization_id}/projects/{location.project_id}/snapshots/{snapshot_id}][%d] ConsulService_DeleteSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *ConsulServiceDeleteSnapshotDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ConsulServiceDeleteSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}