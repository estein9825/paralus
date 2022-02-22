// Code generated by go-swagger; DO NOT EDIT.

package bootstrap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RafaySystems/rcloud-base/components/common/api/def/clients/sentry/models"
)

// BootstrapCreateBootstrapAgentReader is a Reader for the BootstrapCreateBootstrapAgent structure.
type BootstrapCreateBootstrapAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BootstrapCreateBootstrapAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBootstrapCreateBootstrapAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBootstrapCreateBootstrapAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBootstrapCreateBootstrapAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBootstrapCreateBootstrapAgentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBootstrapCreateBootstrapAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBootstrapCreateBootstrapAgentOK creates a BootstrapCreateBootstrapAgentOK with default headers values
func NewBootstrapCreateBootstrapAgentOK() *BootstrapCreateBootstrapAgentOK {
	return &BootstrapCreateBootstrapAgentOK{}
}

/* BootstrapCreateBootstrapAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type BootstrapCreateBootstrapAgentOK struct {
	Payload *models.SentryBootstrapAgent
}

func (o *BootstrapCreateBootstrapAgentOK) Error() string {
	return fmt.Sprintf("[POST /v2/sentry/bootstrap/{spec.templateRef}/agent][%d] bootstrapCreateBootstrapAgentOK  %+v", 200, o.Payload)
}
func (o *BootstrapCreateBootstrapAgentOK) GetPayload() *models.SentryBootstrapAgent {
	return o.Payload
}

func (o *BootstrapCreateBootstrapAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SentryBootstrapAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapCreateBootstrapAgentForbidden creates a BootstrapCreateBootstrapAgentForbidden with default headers values
func NewBootstrapCreateBootstrapAgentForbidden() *BootstrapCreateBootstrapAgentForbidden {
	return &BootstrapCreateBootstrapAgentForbidden{}
}

/* BootstrapCreateBootstrapAgentForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type BootstrapCreateBootstrapAgentForbidden struct {
	Payload interface{}
}

func (o *BootstrapCreateBootstrapAgentForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/sentry/bootstrap/{spec.templateRef}/agent][%d] bootstrapCreateBootstrapAgentForbidden  %+v", 403, o.Payload)
}
func (o *BootstrapCreateBootstrapAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapCreateBootstrapAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapCreateBootstrapAgentNotFound creates a BootstrapCreateBootstrapAgentNotFound with default headers values
func NewBootstrapCreateBootstrapAgentNotFound() *BootstrapCreateBootstrapAgentNotFound {
	return &BootstrapCreateBootstrapAgentNotFound{}
}

/* BootstrapCreateBootstrapAgentNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type BootstrapCreateBootstrapAgentNotFound struct {
	Payload interface{}
}

func (o *BootstrapCreateBootstrapAgentNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/sentry/bootstrap/{spec.templateRef}/agent][%d] bootstrapCreateBootstrapAgentNotFound  %+v", 404, o.Payload)
}
func (o *BootstrapCreateBootstrapAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapCreateBootstrapAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapCreateBootstrapAgentInternalServerError creates a BootstrapCreateBootstrapAgentInternalServerError with default headers values
func NewBootstrapCreateBootstrapAgentInternalServerError() *BootstrapCreateBootstrapAgentInternalServerError {
	return &BootstrapCreateBootstrapAgentInternalServerError{}
}

/* BootstrapCreateBootstrapAgentInternalServerError describes a response with status code 500, with default header values.

Returned for internal server error
*/
type BootstrapCreateBootstrapAgentInternalServerError struct {
	Payload interface{}
}

func (o *BootstrapCreateBootstrapAgentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/sentry/bootstrap/{spec.templateRef}/agent][%d] bootstrapCreateBootstrapAgentInternalServerError  %+v", 500, o.Payload)
}
func (o *BootstrapCreateBootstrapAgentInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapCreateBootstrapAgentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapCreateBootstrapAgentDefault creates a BootstrapCreateBootstrapAgentDefault with default headers values
func NewBootstrapCreateBootstrapAgentDefault(code int) *BootstrapCreateBootstrapAgentDefault {
	return &BootstrapCreateBootstrapAgentDefault{
		_statusCode: code,
	}
}

/* BootstrapCreateBootstrapAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BootstrapCreateBootstrapAgentDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the bootstrap create bootstrap agent default response
func (o *BootstrapCreateBootstrapAgentDefault) Code() int {
	return o._statusCode
}

func (o *BootstrapCreateBootstrapAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v2/sentry/bootstrap/{spec.templateRef}/agent][%d] Bootstrap_CreateBootstrapAgent default  %+v", o._statusCode, o.Payload)
}
func (o *BootstrapCreateBootstrapAgentDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BootstrapCreateBootstrapAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
