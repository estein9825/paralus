// Code generated by go-swagger; DO NOT EDIT.

package bootstrap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RafayLabs/rcloud-base/api/def/clients/sentry/models"
)

// BootstrapGetBootstrapAgentConfigReader is a Reader for the BootstrapGetBootstrapAgentConfig structure.
type BootstrapGetBootstrapAgentConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BootstrapGetBootstrapAgentConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBootstrapGetBootstrapAgentConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBootstrapGetBootstrapAgentConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBootstrapGetBootstrapAgentConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBootstrapGetBootstrapAgentConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBootstrapGetBootstrapAgentConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBootstrapGetBootstrapAgentConfigOK creates a BootstrapGetBootstrapAgentConfigOK with default headers values
func NewBootstrapGetBootstrapAgentConfigOK() *BootstrapGetBootstrapAgentConfigOK {
	return &BootstrapGetBootstrapAgentConfigOK{}
}

/* BootstrapGetBootstrapAgentConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type BootstrapGetBootstrapAgentConfigOK struct {
	Payload *models.V3HTTPBody
}

func (o *BootstrapGetBootstrapAgentConfigOK) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}/config][%d] bootstrapGetBootstrapAgentConfigOK  %+v", 200, o.Payload)
}
func (o *BootstrapGetBootstrapAgentConfigOK) GetPayload() *models.V3HTTPBody {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3HTTPBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentConfigForbidden creates a BootstrapGetBootstrapAgentConfigForbidden with default headers values
func NewBootstrapGetBootstrapAgentConfigForbidden() *BootstrapGetBootstrapAgentConfigForbidden {
	return &BootstrapGetBootstrapAgentConfigForbidden{}
}

/* BootstrapGetBootstrapAgentConfigForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type BootstrapGetBootstrapAgentConfigForbidden struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}/config][%d] bootstrapGetBootstrapAgentConfigForbidden  %+v", 403, o.Payload)
}
func (o *BootstrapGetBootstrapAgentConfigForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentConfigNotFound creates a BootstrapGetBootstrapAgentConfigNotFound with default headers values
func NewBootstrapGetBootstrapAgentConfigNotFound() *BootstrapGetBootstrapAgentConfigNotFound {
	return &BootstrapGetBootstrapAgentConfigNotFound{}
}

/* BootstrapGetBootstrapAgentConfigNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type BootstrapGetBootstrapAgentConfigNotFound struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}/config][%d] bootstrapGetBootstrapAgentConfigNotFound  %+v", 404, o.Payload)
}
func (o *BootstrapGetBootstrapAgentConfigNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentConfigInternalServerError creates a BootstrapGetBootstrapAgentConfigInternalServerError with default headers values
func NewBootstrapGetBootstrapAgentConfigInternalServerError() *BootstrapGetBootstrapAgentConfigInternalServerError {
	return &BootstrapGetBootstrapAgentConfigInternalServerError{}
}

/* BootstrapGetBootstrapAgentConfigInternalServerError describes a response with status code 500, with default header values.

Returned for internal server error
*/
type BootstrapGetBootstrapAgentConfigInternalServerError struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}/config][%d] bootstrapGetBootstrapAgentConfigInternalServerError  %+v", 500, o.Payload)
}
func (o *BootstrapGetBootstrapAgentConfigInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentConfigDefault creates a BootstrapGetBootstrapAgentConfigDefault with default headers values
func NewBootstrapGetBootstrapAgentConfigDefault(code int) *BootstrapGetBootstrapAgentConfigDefault {
	return &BootstrapGetBootstrapAgentConfigDefault{
		_statusCode: code,
	}
}

/* BootstrapGetBootstrapAgentConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BootstrapGetBootstrapAgentConfigDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the bootstrap get bootstrap agent config default response
func (o *BootstrapGetBootstrapAgentConfigDefault) Code() int {
	return o._statusCode
}

func (o *BootstrapGetBootstrapAgentConfigDefault) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}/config][%d] Bootstrap_GetBootstrapAgentConfig default  %+v", o._statusCode, o.Payload)
}
func (o *BootstrapGetBootstrapAgentConfigDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
