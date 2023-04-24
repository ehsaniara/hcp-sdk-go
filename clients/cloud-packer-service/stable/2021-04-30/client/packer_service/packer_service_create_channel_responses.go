// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceCreateChannelReader is a Reader for the PackerServiceCreateChannel structure.
type PackerServiceCreateChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceCreateChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceCreateChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceCreateChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceCreateChannelOK creates a PackerServiceCreateChannelOK with default headers values
func NewPackerServiceCreateChannelOK() *PackerServiceCreateChannelOK {
	return &PackerServiceCreateChannelOK{}
}

/*
PackerServiceCreateChannelOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceCreateChannelOK struct {
	Payload *models.HashicorpCloudPackerCreateChannelResponse
}

// IsSuccess returns true when this packer service create channel o k response has a 2xx status code
func (o *PackerServiceCreateChannelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service create channel o k response has a 3xx status code
func (o *PackerServiceCreateChannelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service create channel o k response has a 4xx status code
func (o *PackerServiceCreateChannelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service create channel o k response has a 5xx status code
func (o *PackerServiceCreateChannelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service create channel o k response a status code equal to that given
func (o *PackerServiceCreateChannelOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceCreateChannelOK) Error() string {
	return fmt.Sprintf("[POST /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels][%d] packerServiceCreateChannelOK  %+v", 200, o.Payload)
}

func (o *PackerServiceCreateChannelOK) String() string {
	return fmt.Sprintf("[POST /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels][%d] packerServiceCreateChannelOK  %+v", 200, o.Payload)
}

func (o *PackerServiceCreateChannelOK) GetPayload() *models.HashicorpCloudPackerCreateChannelResponse {
	return o.Payload
}

func (o *PackerServiceCreateChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerCreateChannelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceCreateChannelDefault creates a PackerServiceCreateChannelDefault with default headers values
func NewPackerServiceCreateChannelDefault(code int) *PackerServiceCreateChannelDefault {
	return &PackerServiceCreateChannelDefault{
		_statusCode: code,
	}
}

/*
PackerServiceCreateChannelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceCreateChannelDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the packer service create channel default response
func (o *PackerServiceCreateChannelDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service create channel default response has a 2xx status code
func (o *PackerServiceCreateChannelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service create channel default response has a 3xx status code
func (o *PackerServiceCreateChannelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service create channel default response has a 4xx status code
func (o *PackerServiceCreateChannelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service create channel default response has a 5xx status code
func (o *PackerServiceCreateChannelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service create channel default response a status code equal to that given
func (o *PackerServiceCreateChannelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceCreateChannelDefault) Error() string {
	return fmt.Sprintf("[POST /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels][%d] PackerService_CreateChannel default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceCreateChannelDefault) String() string {
	return fmt.Sprintf("[POST /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/channels][%d] PackerService_CreateChannel default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceCreateChannelDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceCreateChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PackerServiceCreateChannelBody packer service create channel body
swagger:model PackerServiceCreateChannelBody
*/
type PackerServiceCreateChannelBody struct {

	// Fingerprint of the iteration. The fingerprint is set by Packer when you
	// call `packer build`. It will most often correspond to a git commit sha,
	// but can be manually overridden by setting the environment variable
	// `HCP_PACKER_BUILD_FINGERPRINT`.
	Fingerprint string `json:"fingerprint,omitempty"`

	// The human-readable version number assigned to this iteration.
	IncrementalVersion int32 `json:"incremental_version,omitempty"`

	// ULID of the iteration.
	IterationID string `json:"iteration_id,omitempty"`

	// location
	Location *PackerServiceCreateChannelParamsBodyLocation `json:"location,omitempty"`

	// When set, will set the channel access in HCP Packer registry. The channel is unrestricted by default;
	Restriction *models.HashicorpCloudPackerCreateChannelRequestRestriction `json:"restriction,omitempty"`

	// Human-readable name for the channel.
	Slug string `json:"slug,omitempty"`
}

// Validate validates this packer service create channel body
func (o *PackerServiceCreateChannelBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRestriction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceCreateChannelBody) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.Location) { // not required
		return nil
	}

	if o.Location != nil {
		if err := o.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

func (o *PackerServiceCreateChannelBody) validateRestriction(formats strfmt.Registry) error {
	if swag.IsZero(o.Restriction) { // not required
		return nil
	}

	if o.Restriction != nil {
		if err := o.Restriction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "restriction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "restriction")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this packer service create channel body based on the context it is used
func (o *PackerServiceCreateChannelBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRestriction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceCreateChannelBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.Location != nil {
		if err := o.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

func (o *PackerServiceCreateChannelBody) contextValidateRestriction(ctx context.Context, formats strfmt.Registry) error {

	if o.Restriction != nil {
		if err := o.Restriction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "restriction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "restriction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PackerServiceCreateChannelBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceCreateChannelBody) UnmarshalBinary(b []byte) error {
	var res PackerServiceCreateChannelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PackerServiceCreateChannelParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model PackerServiceCreateChannelParamsBodyLocation
*/
type PackerServiceCreateChannelParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this packer service create channel params body location
func (o *PackerServiceCreateChannelParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceCreateChannelParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this packer service create channel params body location based on the context it is used
func (o *PackerServiceCreateChannelParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceCreateChannelParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {
		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PackerServiceCreateChannelParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceCreateChannelParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res PackerServiceCreateChannelParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
