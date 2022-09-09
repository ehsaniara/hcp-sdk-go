// Code generated by go-swagger; DO NOT EDIT.

package registry_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// CreateRegistryReader is a Reader for the CreateRegistry structure.
type CreateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRegistryOK creates a CreateRegistryOK with default headers values
func NewCreateRegistryOK() *CreateRegistryOK {
	return &CreateRegistryOK{}
}

/* CreateRegistryOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateRegistryOK struct {
	Payload *models.HashicorpCloudVagrantCreateRegistryResponse
}

func (o *CreateRegistryOK) Error() string {
	return fmt.Sprintf("[PUT /vagrant/2022-09-30/registry][%d] createRegistryOK  %+v", 200, o.Payload)
}
func (o *CreateRegistryOK) GetPayload() *models.HashicorpCloudVagrantCreateRegistryResponse {
	return o.Payload
}

func (o *CreateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVagrantCreateRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}