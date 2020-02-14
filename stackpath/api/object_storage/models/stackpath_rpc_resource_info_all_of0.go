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

// StackpathRPCResourceInfoAllOf0 stackpath Rpc resource info all of0
// swagger:discriminator stackpathRpcResourceInfoAllOf0 @type
type StackpathRPCResourceInfoAllOf0 interface {
	runtime.Validatable

	// at type
	// Required: true
	AtType() string
	SetAtType(string)
}

type stackpathRpcResourceInfoAllOf0 struct {
	atTypeField string
}

// AtType gets the at type of this polymorphic type
func (m *stackpathRpcResourceInfoAllOf0) AtType() string {
	return "stackpathRpcResourceInfoAllOf0"
}

// SetAtType sets the at type of this polymorphic type
func (m *stackpathRpcResourceInfoAllOf0) SetAtType(val string) {

}

// UnmarshalStackpathRPCResourceInfoAllOf0Slice unmarshals polymorphic slices of StackpathRPCResourceInfoAllOf0
func UnmarshalStackpathRPCResourceInfoAllOf0Slice(reader io.Reader, consumer runtime.Consumer) ([]StackpathRPCResourceInfoAllOf0, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []StackpathRPCResourceInfoAllOf0
	for _, element := range elements {
		obj, err := unmarshalStackpathRPCResourceInfoAllOf0(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalStackpathRPCResourceInfoAllOf0 unmarshals polymorphic StackpathRPCResourceInfoAllOf0
func UnmarshalStackpathRPCResourceInfoAllOf0(reader io.Reader, consumer runtime.Consumer) (StackpathRPCResourceInfoAllOf0, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalStackpathRPCResourceInfoAllOf0(data, consumer)
}

func unmarshalStackpathRPCResourceInfoAllOf0(data []byte, consumer runtime.Consumer) (StackpathRPCResourceInfoAllOf0, error) {
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
	case "stackpath.rpc.ResourceInfo":
		var result StackpathRPCResourceInfo
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "stackpathRpcResourceInfoAllOf0":
		var result stackpathRpcResourceInfoAllOf0
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)

}

// Validate validates this stackpath Rpc resource info all of0
func (m *stackpathRpcResourceInfoAllOf0) Validate(formats strfmt.Registry) error {
	return nil
}
