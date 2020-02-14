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

// UpdateBucketReader is a Reader for the UpdateBucket structure.
type UpdateBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateBucketUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateBucketInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBucketOK creates a UpdateBucketOK with default headers values
func NewUpdateBucketOK() *UpdateBucketOK {
	return &UpdateBucketOK{}
}

/*UpdateBucketOK handles this case with default header values.

UpdateBucketOK update bucket o k
*/
type UpdateBucketOK struct {
	Payload *models.UpdateBucketOKBody
}

func (o *UpdateBucketOK) Error() string {
	return fmt.Sprintf("[PUT /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] updateBucketOK  %+v", 200, o.Payload)
}

func (o *UpdateBucketOK) GetPayload() *models.UpdateBucketOKBody {
	return o.Payload
}

func (o *UpdateBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateBucketOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBucketUnauthorized creates a UpdateBucketUnauthorized with default headers values
func NewUpdateBucketUnauthorized() *UpdateBucketUnauthorized {
	return &UpdateBucketUnauthorized{}
}

/*UpdateBucketUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type UpdateBucketUnauthorized struct {
	Payload *models.UpdateBucketUnauthorizedBody
}

func (o *UpdateBucketUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] updateBucketUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateBucketUnauthorized) GetPayload() *models.UpdateBucketUnauthorizedBody {
	return o.Payload
}

func (o *UpdateBucketUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateBucketUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBucketInternalServerError creates a UpdateBucketInternalServerError with default headers values
func NewUpdateBucketInternalServerError() *UpdateBucketInternalServerError {
	return &UpdateBucketInternalServerError{}
}

/*UpdateBucketInternalServerError handles this case with default header values.

Internal server error.
*/
type UpdateBucketInternalServerError struct {
	Payload *models.UpdateBucketInternalServerErrorBody
}

func (o *UpdateBucketInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] updateBucketInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateBucketInternalServerError) GetPayload() *models.UpdateBucketInternalServerErrorBody {
	return o.Payload
}

func (o *UpdateBucketInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateBucketInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBucketDefault creates a UpdateBucketDefault with default headers values
func NewUpdateBucketDefault(code int) *UpdateBucketDefault {
	return &UpdateBucketDefault{
		_statusCode: code,
	}
}

/*UpdateBucketDefault handles this case with default header values.

Default error structure.
*/
type UpdateBucketDefault struct {
	_statusCode int

	Payload *models.UpdateBucketDefaultBody
}

// Code gets the status code for the update bucket default response
func (o *UpdateBucketDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBucketDefault) Error() string {
	return fmt.Sprintf("[PUT /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] UpdateBucket default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateBucketDefault) GetPayload() *models.UpdateBucketDefaultBody {
	return o.Payload
}

func (o *UpdateBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateBucketDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
