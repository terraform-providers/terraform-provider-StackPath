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

// StackpathRPCBadRequestAllOf0 stackpath Rpc bad request all of0
// swagger:discriminator stackpathRpcBadRequestAllOf0 @type
type StackpathRPCBadRequestAllOf0 interface {
	runtime.Validatable

	// at type
	// Required: true
	AtType() string
	SetAtType(string)
}

type stackpathRpcBadRequestAllOf0 struct {
	atTypeField string
}

// AtType gets the at type of this polymorphic type
func (m *stackpathRpcBadRequestAllOf0) AtType() string {
	return "stackpathRpcBadRequestAllOf0"
}

// SetAtType sets the at type of this polymorphic type
func (m *stackpathRpcBadRequestAllOf0) SetAtType(val string) {

}

// UnmarshalStackpathRPCBadRequestAllOf0Slice unmarshals polymorphic slices of StackpathRPCBadRequestAllOf0
func UnmarshalStackpathRPCBadRequestAllOf0Slice(reader io.Reader, consumer runtime.Consumer) ([]StackpathRPCBadRequestAllOf0, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []StackpathRPCBadRequestAllOf0
	for _, element := range elements {
		obj, err := unmarshalStackpathRPCBadRequestAllOf0(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalStackpathRPCBadRequestAllOf0 unmarshals polymorphic StackpathRPCBadRequestAllOf0
func UnmarshalStackpathRPCBadRequestAllOf0(reader io.Reader, consumer runtime.Consumer) (StackpathRPCBadRequestAllOf0, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalStackpathRPCBadRequestAllOf0(data, consumer)
}

func unmarshalStackpathRPCBadRequestAllOf0(data []byte, consumer runtime.Consumer) (StackpathRPCBadRequestAllOf0, error) {
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
	case "stackpath.rpc.BadRequest":
		var result StackpathRPCBadRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "stackpathRpcBadRequestAllOf0":
		var result stackpathRpcBadRequestAllOf0
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)

}

// Validate validates this stackpath Rpc bad request all of0
func (m *stackpathRpcBadRequestAllOf0) Validate(formats strfmt.Registry) error {
	return nil
}
