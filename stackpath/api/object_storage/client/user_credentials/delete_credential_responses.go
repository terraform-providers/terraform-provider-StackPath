// Code generated by go-swagger; DO NOT EDIT.

package user_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/object_storage/models"
)

// DeleteCredentialReader is a Reader for the DeleteCredential structure.
type DeleteCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCredentialNoContent creates a DeleteCredentialNoContent with default headers values
func NewDeleteCredentialNoContent() *DeleteCredentialNoContent {
	return &DeleteCredentialNoContent{}
}

/*DeleteCredentialNoContent handles this case with default header values.

No content
*/
type DeleteCredentialNoContent struct {
}

func (o *DeleteCredentialNoContent) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key}][%d] deleteCredentialNoContent ", 204)
}

func (o *DeleteCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCredentialUnauthorized creates a DeleteCredentialUnauthorized with default headers values
func NewDeleteCredentialUnauthorized() *DeleteCredentialUnauthorized {
	return &DeleteCredentialUnauthorized{}
}

/*DeleteCredentialUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteCredentialUnauthorized struct {
	Payload *models.DeleteCredentialUnauthorizedBody
}

func (o *DeleteCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key}][%d] deleteCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCredentialUnauthorized) GetPayload() *models.DeleteCredentialUnauthorizedBody {
	return o.Payload
}

func (o *DeleteCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCredentialUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCredentialInternalServerError creates a DeleteCredentialInternalServerError with default headers values
func NewDeleteCredentialInternalServerError() *DeleteCredentialInternalServerError {
	return &DeleteCredentialInternalServerError{}
}

/*DeleteCredentialInternalServerError handles this case with default header values.

Internal server error.
*/
type DeleteCredentialInternalServerError struct {
	Payload *models.DeleteCredentialInternalServerErrorBody
}

func (o *DeleteCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key}][%d] deleteCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCredentialInternalServerError) GetPayload() *models.DeleteCredentialInternalServerErrorBody {
	return o.Payload
}

func (o *DeleteCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCredentialInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCredentialDefault creates a DeleteCredentialDefault with default headers values
func NewDeleteCredentialDefault(code int) *DeleteCredentialDefault {
	return &DeleteCredentialDefault{
		_statusCode: code,
	}
}

/*DeleteCredentialDefault handles this case with default header values.

Default error structure.
*/
type DeleteCredentialDefault struct {
	_statusCode int

	Payload *models.DeleteCredentialDefaultBody
}

// Code gets the status code for the delete credential default response
func (o *DeleteCredentialDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCredentialDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key}][%d] DeleteCredential default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCredentialDefault) GetPayload() *models.DeleteCredentialDefaultBody {
	return o.Payload
}

func (o *DeleteCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCredentialDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
