// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DeploymentSpec A deployment's specification
// swagger:model v1DeploymentSpec
type V1DeploymentSpec struct {

	// The maximum number of instances in a deployment
	MaxReplicas int32 `json:"maxReplicas,omitempty"`

	// The minimum number of instances in a deployment
	MinReplicas int32 `json:"minReplicas,omitempty"`

	// scale settings
	ScaleSettings *V1ScaleSettings `json:"scaleSettings,omitempty"`

	// A collection of filters that match the deployment's scope
	Selectors []*V1MatchExpression `json:"selectors"`
}

// Validate validates this v1 deployment spec
func (m *V1DeploymentSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScaleSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DeploymentSpec) validateScaleSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.ScaleSettings) { // not required
		return nil
	}

	if m.ScaleSettings != nil {
		if err := m.ScaleSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scaleSettings")
			}
			return err
		}
	}

	return nil
}

func (m *V1DeploymentSpec) validateSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.Selectors) { // not required
		return nil
	}

	for i := 0; i < len(m.Selectors); i++ {
		if swag.IsZero(m.Selectors[i]) { // not required
			continue
		}

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1DeploymentSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DeploymentSpec) UnmarshalBinary(b []byte) error {
	var res V1DeploymentSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}