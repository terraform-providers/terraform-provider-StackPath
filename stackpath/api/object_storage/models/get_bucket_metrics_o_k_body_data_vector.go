// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetBucketMetricsOKBodyDataVector A set of time series containing a single sample for each time series, all sharing the same timestamp
// swagger:model getBucketMetricsOKBodyDataVector
type GetBucketMetricsOKBodyDataVector struct {

	// A data point's value
	Results []*GetBucketMetricsOKBodyDataVectorResultsItems `json:"results"`
}

// Validate validates this get bucket metrics o k body data vector
func (m *GetBucketMetricsOKBodyDataVector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBucketMetricsOKBodyDataVector) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetBucketMetricsOKBodyDataVector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBucketMetricsOKBodyDataVector) UnmarshalBinary(b []byte) error {
	var res GetBucketMetricsOKBodyDataVector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
