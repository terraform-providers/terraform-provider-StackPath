// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StackpathRPCLocalizedMessageAllOf0 stackpath Rpc localized message all of0
// swagger:discriminator stackpathRpcLocalizedMessageAllOf0 @type
type StackpathRPCLocalizedMessageAllOf0 interface {
	runtime.Validatable

	// at type
	// Required: true
	AtType() string
	SetAtType(string)
}

type stackpathRpcLocalizedMessageAllOf0 struct {
	atTypeField string
}

// AtType gets the at type of this polymorphic type
func (m *stackpathRpcLocalizedMessageAllOf0) AtType() string {
	return "stackpathRpcLocalizedMessageAllOf0"
}

// SetAtType sets the at type of this polymorphic type
func (m *stackpathRpcLocalizedMessageAllOf0) SetAtType(val string) {

}

// UnmarshalStackpathRPCLocalizedMessageAllOf0Slice unmarshals polymorphic slices of StackpathRPCLocalizedMessageAllOf0
func UnmarshalStackpathRPCLocalizedMessageAllOf0Slice(reader io.Reader, consumer runtime.Consumer) ([]StackpathRPCLocalizedMessageAllOf0, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []StackpathRPCLocalizedMessageAllOf0
	for _, element := range elements {
		obj, err := unmarshalStackpathRPCLocalizedMessageAllOf0(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalStackpathRPCLocalizedMessageAllOf0 unmarshals polymorphic StackpathRPCLocalizedMessageAllOf0
func UnmarshalStackpathRPCLocalizedMessageAllOf0(reader io.Reader, consumer runtime.Consumer) (StackpathRPCLocalizedMessageAllOf0, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalStackpathRPCLocalizedMessageAllOf0(data, consumer)
}

func unmarshalStackpathRPCLocalizedMessageAllOf0(data []byte, consumer runtime.Consumer) (StackpathRPCLocalizedMessageAllOf0, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the @type property.
	var getType struct {
		AtType string `json:"@type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("@type", "body", getType.AtType); err != nil {
		return nil, err
	}

	// The value of @type is used to determine which type to create and unmarshal the data into
	switch getType.AtType {
	case "stackpath.rpc.LocalizedMessage":
		var result StackpathRPCLocalizedMessage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "stackpathRpcLocalizedMessageAllOf0":
		var result stackpathRpcLocalizedMessageAllOf0
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)

}

// Validate validates this stackpath Rpc localized message all of0
func (m *stackpathRpcLocalizedMessageAllOf0) Validate(formats strfmt.Registry) error {
	return nil
}
