// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetBucketOKBodyBucket get bucket o k body bucket
// swagger:model getBucketOKBodyBucket
type GetBucketOKBodyBucket struct {

	// The date when the bucket was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// The URL used to access the bucket
	EndpointURL string `json:"endpointUrl,omitempty"`

	// The ID for the bucket
	ID string `json:"id,omitempty"`

	// The name of the bucket
	Label string `json:"label,omitempty"`

	// The region in which the bucket is created. Available regions are: us-east-1, us-west-1, eu-central-1
	Region string `json:"region,omitempty"`

	// The date when the bucket was last updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// - PRIVATE: The bucket is private and only accessibly with credentials
	//  - PUBLIC: The bucket is public and accessible over the internet
	// Enum: [PRIVATE PUBLIC]
	Visibility *string `json:"visibility,omitempty"`
}

// Validate validates this get bucket o k body bucket
func (m *GetBucketOKBodyBucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBucketOKBodyBucket) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetBucketOKBodyBucket) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var getBucketOKBodyBucketTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRIVATE","PUBLIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getBucketOKBodyBucketTypeVisibilityPropEnum = append(getBucketOKBodyBucketTypeVisibilityPropEnum, v)
	}
}

const (

	// GetBucketOKBodyBucketVisibilityPRIVATE captures enum value "PRIVATE"
	GetBucketOKBodyBucketVisibilityPRIVATE string = "PRIVATE"

	// GetBucketOKBodyBucketVisibilityPUBLIC captures enum value "PUBLIC"
	GetBucketOKBodyBucketVisibilityPUBLIC string = "PUBLIC"
)

// prop value enum
func (m *GetBucketOKBodyBucket) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getBucketOKBodyBucketTypeVisibilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetBucketOKBodyBucket) validateVisibility(formats strfmt.Registry) error {

	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", *m.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetBucketOKBodyBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBucketOKBodyBucket) UnmarshalBinary(b []byte) error {
	var res GetBucketOKBodyBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
