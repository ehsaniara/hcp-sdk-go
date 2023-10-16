// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility hashicorp cloud global network manager 20220215 list clusters with dialer eligibility response cluster with dialer eligibility
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ListClustersWithDialerEligibilityResponse.ClusterWithDialerEligibility
type HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility struct {

	// connectivity_options is a list of network connectivity options for the two clusters
	// involved in the prospective peering connection
	ConnectivityOptions []*HashicorpCloudGlobalNetworkManager20220215PeeringConnectivityOption `json:"connectivity_options"`

	// id is the user settable GNM cluster name.
	ID string `json:"id,omitempty"`

	// ineligibility_reasons is the list of reasons why this cluster is ineligible. Empty
	// list means this cluster is eligible.
	IneligibilityReasons []*HashicorpCloudGlobalNetworkManager20220215PeeringIneligibilityReason `json:"ineligibility_reasons"`

	// partitions are the Consul admin partitions on the cluster
	Partitions []*HashicorpCloudGlobalNetworkManager20220215AdminPartitionWithEligibility `json:"partitions"`
}

// Validate validates this hashicorp cloud global network manager 20220215 list clusters with dialer eligibility response cluster with dialer eligibility
func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectivityOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIneligibilityReasons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) validateConnectivityOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectivityOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectivityOptions); i++ {
		if swag.IsZero(m.ConnectivityOptions[i]) { // not required
			continue
		}

		if m.ConnectivityOptions[i] != nil {
			if err := m.ConnectivityOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectivity_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectivity_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) validateIneligibilityReasons(formats strfmt.Registry) error {
	if swag.IsZero(m.IneligibilityReasons) { // not required
		return nil
	}

	for i := 0; i < len(m.IneligibilityReasons); i++ {
		if swag.IsZero(m.IneligibilityReasons[i]) { // not required
			continue
		}

		if m.IneligibilityReasons[i] != nil {
			if err := m.IneligibilityReasons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ineligibility_reasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ineligibility_reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) validatePartitions(formats strfmt.Registry) error {
	if swag.IsZero(m.Partitions) { // not required
		return nil
	}

	for i := 0; i < len(m.Partitions); i++ {
		if swag.IsZero(m.Partitions[i]) { // not required
			continue
		}

		if m.Partitions[i] != nil {
			if err := m.Partitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("partitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 list clusters with dialer eligibility response cluster with dialer eligibility based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectivityOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIneligibilityReasons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) contextValidateConnectivityOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConnectivityOptions); i++ {

		if m.ConnectivityOptions[i] != nil {

			if swag.IsZero(m.ConnectivityOptions[i]) { // not required
				return nil
			}

			if err := m.ConnectivityOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectivity_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectivity_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) contextValidateIneligibilityReasons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IneligibilityReasons); i++ {

		if m.IneligibilityReasons[i] != nil {

			if swag.IsZero(m.IneligibilityReasons[i]) { // not required
				return nil
			}

			if err := m.IneligibilityReasons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ineligibility_reasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ineligibility_reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) contextValidatePartitions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Partitions); i++ {

		if m.Partitions[i] != nil {

			if swag.IsZero(m.Partitions[i]) { // not required
				return nil
			}

			if err := m.Partitions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("partitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ListClustersWithDialerEligibilityResponseClusterWithDialerEligibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
