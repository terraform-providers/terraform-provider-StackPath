// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StackpathRPCQuotaFailureAllOf1ViolationsItems stackpath Rpc quota failure all of1 violations items
// swagger:model stackpathRpcQuotaFailureAllOf1ViolationsItems
type StackpathRPCQuotaFailureAllOf1ViolationsItems struct {

	// description
	Description string `json:"description,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this stackpath Rpc quota failure all of1 violations items
func (m *StackpathRPCQuotaFailureAllOf1ViolationsItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCQuotaFailureAllOf1ViolationsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCQuotaFailureAllOf1ViolationsItems) UnmarshalBinary(b []byte) error {
	var res StackpathRPCQuotaFailureAllOf1ViolationsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
