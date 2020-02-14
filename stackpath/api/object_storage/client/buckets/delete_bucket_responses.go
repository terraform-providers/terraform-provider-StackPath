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

// DeleteBucketReader is a Reader for the DeleteBucket structure.
type DeleteBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBucketNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteBucketUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteBucketInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBucketNoContent creates a DeleteBucketNoContent with default headers values
func NewDeleteBucketNoContent() *DeleteBucketNoContent {
	return &DeleteBucketNoContent{}
}

/*DeleteBucketNoContent handles this case with default header values.

No content
*/
type DeleteBucketNoContent struct {
}

func (o *DeleteBucketNoContent) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] deleteBucketNoContent ", 204)
}

func (o *DeleteBucketNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBucketUnauthorized creates a DeleteBucketUnauthorized with default headers values
func NewDeleteBucketUnauthorized() *DeleteBucketUnauthorized {
	return &DeleteBucketUnauthorized{}
}

/*DeleteBucketUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteBucketUnauthorized struct {
	Payload *models.DeleteBucketUnauthorizedBody
}

func (o *DeleteBucketUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] deleteBucketUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteBucketUnauthorized) GetPayload() *models.DeleteBucketUnauthorizedBody {
	return o.Payload
}

func (o *DeleteBucketUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteBucketUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBucketInternalServerError creates a DeleteBucketInternalServerError with default headers values
func NewDeleteBucketInternalServerError() *DeleteBucketInternalServerError {
	return &DeleteBucketInternalServerError{}
}

/*DeleteBucketInternalServerError handles this case with default header values.

Internal server error.
*/
type DeleteBucketInternalServerError struct {
	Payload *models.DeleteBucketInternalServerErrorBody
}

func (o *DeleteBucketInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] deleteBucketInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteBucketInternalServerError) GetPayload() *models.DeleteBucketInternalServerErrorBody {
	return o.Payload
}

func (o *DeleteBucketInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteBucketInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBucketDefault creates a DeleteBucketDefault with default headers values
func NewDeleteBucketDefault(code int) *DeleteBucketDefault {
	return &DeleteBucketDefault{
		_statusCode: code,
	}
}

/*DeleteBucketDefault handles this case with default header values.

Default error structure.
*/
type DeleteBucketDefault struct {
	_statusCode int

	Payload *models.DeleteBucketDefaultBody
}

// Code gets the status code for the delete bucket default response
func (o *DeleteBucketDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBucketDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] DeleteBucket default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBucketDefault) GetPayload() *models.DeleteBucketDefaultBody {
	return o.Payload
}

func (o *DeleteBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteBucketDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
