// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-06-13/models"
)

// UpdateGhRepoSyncIntegrationReader is a Reader for the UpdateGhRepoSyncIntegration structure.
type UpdateGhRepoSyncIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGhRepoSyncIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGhRepoSyncIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateGhRepoSyncIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateGhRepoSyncIntegrationOK creates a UpdateGhRepoSyncIntegrationOK with default headers values
func NewUpdateGhRepoSyncIntegrationOK() *UpdateGhRepoSyncIntegrationOK {
	return &UpdateGhRepoSyncIntegrationOK{}
}

/*
UpdateGhRepoSyncIntegrationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateGhRepoSyncIntegrationOK struct {
	Payload *models.Secrets20230613UpdateSyncIntegrationResponse
}

// IsSuccess returns true when this update gh repo sync integration o k response has a 2xx status code
func (o *UpdateGhRepoSyncIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update gh repo sync integration o k response has a 3xx status code
func (o *UpdateGhRepoSyncIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gh repo sync integration o k response has a 4xx status code
func (o *UpdateGhRepoSyncIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update gh repo sync integration o k response has a 5xx status code
func (o *UpdateGhRepoSyncIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update gh repo sync integration o k response a status code equal to that given
func (o *UpdateGhRepoSyncIntegrationOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateGhRepoSyncIntegrationOK) Error() string {
	return fmt.Sprintf("[PATCH /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/gh-repo/{name}][%d] updateGhRepoSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateGhRepoSyncIntegrationOK) String() string {
	return fmt.Sprintf("[PATCH /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/gh-repo/{name}][%d] updateGhRepoSyncIntegrationOK  %+v", 200, o.Payload)
}

func (o *UpdateGhRepoSyncIntegrationOK) GetPayload() *models.Secrets20230613UpdateSyncIntegrationResponse {
	return o.Payload
}

func (o *UpdateGhRepoSyncIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613UpdateSyncIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGhRepoSyncIntegrationDefault creates a UpdateGhRepoSyncIntegrationDefault with default headers values
func NewUpdateGhRepoSyncIntegrationDefault(code int) *UpdateGhRepoSyncIntegrationDefault {
	return &UpdateGhRepoSyncIntegrationDefault{
		_statusCode: code,
	}
}

/*
UpdateGhRepoSyncIntegrationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateGhRepoSyncIntegrationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the update gh repo sync integration default response
func (o *UpdateGhRepoSyncIntegrationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update gh repo sync integration default response has a 2xx status code
func (o *UpdateGhRepoSyncIntegrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update gh repo sync integration default response has a 3xx status code
func (o *UpdateGhRepoSyncIntegrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update gh repo sync integration default response has a 4xx status code
func (o *UpdateGhRepoSyncIntegrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update gh repo sync integration default response has a 5xx status code
func (o *UpdateGhRepoSyncIntegrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update gh repo sync integration default response a status code equal to that given
func (o *UpdateGhRepoSyncIntegrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateGhRepoSyncIntegrationDefault) Error() string {
	return fmt.Sprintf("[PATCH /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/gh-repo/{name}][%d] UpdateGhRepoSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateGhRepoSyncIntegrationDefault) String() string {
	return fmt.Sprintf("[PATCH /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/gh-repo/{name}][%d] UpdateGhRepoSyncIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateGhRepoSyncIntegrationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *UpdateGhRepoSyncIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateGhRepoSyncIntegrationBody update gh repo sync integration body
swagger:model UpdateGhRepoSyncIntegrationBody
*/
type UpdateGhRepoSyncIntegrationBody struct {

	// gh repo connection details
	GhRepoConnectionDetails *models.Secrets20230613GhRepoConnectionDetailsRequest `json:"gh_repo_connection_details,omitempty"`

	// location
	Location *UpdateGhRepoSyncIntegrationParamsBodyLocation `json:"location,omitempty"`
}

// Validate validates this update gh repo sync integration body
func (o *UpdateGhRepoSyncIntegrationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGhRepoConnectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGhRepoSyncIntegrationBody) validateGhRepoConnectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.GhRepoConnectionDetails) { // not required
		return nil
	}

	if o.GhRepoConnectionDetails != nil {
		if err := o.GhRepoConnectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "gh_repo_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "gh_repo_connection_details")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateGhRepoSyncIntegrationBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this update gh repo sync integration body based on the context it is used
func (o *UpdateGhRepoSyncIntegrationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGhRepoConnectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGhRepoSyncIntegrationBody) contextValidateGhRepoConnectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if o.GhRepoConnectionDetails != nil {
		if err := o.GhRepoConnectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "gh_repo_connection_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "gh_repo_connection_details")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateGhRepoSyncIntegrationBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (o *UpdateGhRepoSyncIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateGhRepoSyncIntegrationBody) UnmarshalBinary(b []byte) error {
	var res UpdateGhRepoSyncIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateGhRepoSyncIntegrationParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model UpdateGhRepoSyncIntegrationParamsBodyLocation
*/
type UpdateGhRepoSyncIntegrationParamsBodyLocation struct {

	// region
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this update gh repo sync integration params body location
func (o *UpdateGhRepoSyncIntegrationParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGhRepoSyncIntegrationParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this update gh repo sync integration params body location based on the context it is used
func (o *UpdateGhRepoSyncIntegrationParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGhRepoSyncIntegrationParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UpdateGhRepoSyncIntegrationParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateGhRepoSyncIntegrationParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res UpdateGhRepoSyncIntegrationParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
