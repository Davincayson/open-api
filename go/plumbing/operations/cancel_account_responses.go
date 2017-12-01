// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// CancelAccountReader is a Reader for the CancelAccount structure.
type CancelAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCancelAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCancelAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelAccountNoContent creates a CancelAccountNoContent with default headers values
func NewCancelAccountNoContent() *CancelAccountNoContent {
	return &CancelAccountNoContent{}
}

/*CancelAccountNoContent handles this case with default header values.

Not Content
*/
type CancelAccountNoContent struct {
}

func (o *CancelAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{account_id}][%d] cancelAccountNoContent ", 204)
}

func (o *CancelAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelAccountDefault creates a CancelAccountDefault with default headers values
func NewCancelAccountDefault(code int) *CancelAccountDefault {
	return &CancelAccountDefault{
		_statusCode: code,
	}
}

/*CancelAccountDefault handles this case with default header values.

error
*/
type CancelAccountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the cancel account default response
func (o *CancelAccountDefault) Code() int {
	return o._statusCode
}

func (o *CancelAccountDefault) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{account_id}][%d] cancelAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CancelAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
