// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/object_storage/models"
)

// CreateBucketReader is a Reader for the CreateBucket structure.
type CreateBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateBucketUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateBucketInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBucketOK creates a CreateBucketOK with default headers values
func NewCreateBucketOK() *CreateBucketOK {
	return &CreateBucketOK{}
}

/*CreateBucketOK handles this case with default header values.

CreateBucketOK create bucket o k
*/
type CreateBucketOK struct {
	Payload *models.CreateBucketOKBody
}

func (o *CreateBucketOK) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/buckets][%d] createBucketOK  %+v", 200, o.Payload)
}

func (o *CreateBucketOK) GetPayload() *models.CreateBucketOKBody {
	return o.Payload
}

func (o *CreateBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateBucketOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBucketUnauthorized creates a CreateBucketUnauthorized with default headers values
func NewCreateBucketUnauthorized() *CreateBucketUnauthorized {
	return &CreateBucketUnauthorized{}
}

/*CreateBucketUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type CreateBucketUnauthorized struct {
	Payload *models.CreateBucketUnauthorizedBody
}

func (o *CreateBucketUnauthorized) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/buckets][%d] createBucketUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateBucketUnauthorized) GetPayload() *models.CreateBucketUnauthorizedBody {
	return o.Payload
}

func (o *CreateBucketUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateBucketUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBucketInternalServerError creates a CreateBucketInternalServerError with default headers values
func NewCreateBucketInternalServerError() *CreateBucketInternalServerError {
	return &CreateBucketInternalServerError{}
}

/*CreateBucketInternalServerError handles this case with default header values.

Internal server error.
*/
type CreateBucketInternalServerError struct {
	Payload *models.CreateBucketInternalServerErrorBody
}

func (o *CreateBucketInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/buckets][%d] createBucketInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateBucketInternalServerError) GetPayload() *models.CreateBucketInternalServerErrorBody {
	return o.Payload
}

func (o *CreateBucketInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateBucketInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBucketDefault creates a CreateBucketDefault with default headers values
func NewCreateBucketDefault(code int) *CreateBucketDefault {
	return &CreateBucketDefault{
		_statusCode: code,
	}
}

/*CreateBucketDefault handles this case with default header values.

Default error structure.
*/
type CreateBucketDefault struct {
	_statusCode int

	Payload *models.CreateBucketDefaultBody
}

// Code gets the status code for the create bucket default response
func (o *CreateBucketDefault) Code() int {
	return o._statusCode
}

func (o *CreateBucketDefault) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/buckets][%d] CreateBucket default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBucketDefault) GetPayload() *models.CreateBucketDefaultBody {
	return o.Payload
}

func (o *CreateBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateBucketDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
