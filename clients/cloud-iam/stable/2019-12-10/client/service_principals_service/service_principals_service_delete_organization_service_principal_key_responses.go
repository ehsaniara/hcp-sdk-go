// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

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

// ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyReader is a Reader for the ServicePrincipalsServiceDeleteOrganizationServicePrincipalKey structure.
type ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK creates a ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK with default headers values
func NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK() *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK {
	return &ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK{}
}

/*
ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK struct {
	Payload models.HashicorpCloudIamDeleteOrganizationServicePrincipalKeyResponse
}

// IsSuccess returns true when this service principals service delete organization service principal key o k response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service delete organization service principal key o k response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service delete organization service principal key o k response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service delete organization service principal key o k response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service delete organization service principal key o k response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/service-principal-keys/{client_id}][%d] servicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/service-principal-keys/{client_id}][%d] servicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) GetPayload() models.HashicorpCloudIamDeleteOrganizationServicePrincipalKeyResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault creates a ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault with default headers values
func NewServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault(code int) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault {
	return &ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service delete organization service principal key default response
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service delete organization service principal key default response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service delete organization service principal key default response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service delete organization service principal key default response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service delete organization service principal key default response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service delete organization service principal key default response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/service-principal-keys/{client_id}][%d] ServicePrincipalsService_DeleteOrganizationServicePrincipalKey default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/service-principal-keys/{client_id}][%d] ServicePrincipalsService_DeleteOrganizationServicePrincipalKey default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}