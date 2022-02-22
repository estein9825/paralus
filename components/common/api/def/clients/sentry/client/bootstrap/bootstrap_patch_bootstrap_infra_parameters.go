// Code generated by go-swagger; DO NOT EDIT.

package bootstrap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/RafaySystems/rcloud-base/components/common/api/def/clients/sentry/models"
)

// NewBootstrapPatchBootstrapInfraParams creates a new BootstrapPatchBootstrapInfraParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBootstrapPatchBootstrapInfraParams() *BootstrapPatchBootstrapInfraParams {
	return &BootstrapPatchBootstrapInfraParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBootstrapPatchBootstrapInfraParamsWithTimeout creates a new BootstrapPatchBootstrapInfraParams object
// with the ability to set a timeout on a request.
func NewBootstrapPatchBootstrapInfraParamsWithTimeout(timeout time.Duration) *BootstrapPatchBootstrapInfraParams {
	return &BootstrapPatchBootstrapInfraParams{
		timeout: timeout,
	}
}

// NewBootstrapPatchBootstrapInfraParamsWithContext creates a new BootstrapPatchBootstrapInfraParams object
// with the ability to set a context for a request.
func NewBootstrapPatchBootstrapInfraParamsWithContext(ctx context.Context) *BootstrapPatchBootstrapInfraParams {
	return &BootstrapPatchBootstrapInfraParams{
		Context: ctx,
	}
}

// NewBootstrapPatchBootstrapInfraParamsWithHTTPClient creates a new BootstrapPatchBootstrapInfraParams object
// with the ability to set a custom HTTPClient for a request.
func NewBootstrapPatchBootstrapInfraParamsWithHTTPClient(client *http.Client) *BootstrapPatchBootstrapInfraParams {
	return &BootstrapPatchBootstrapInfraParams{
		HTTPClient: client,
	}
}

/* BootstrapPatchBootstrapInfraParams contains all the parameters to send to the API endpoint
   for the bootstrap patch bootstrap infra operation.

   Typically these are written to a http.Request.
*/
type BootstrapPatchBootstrapInfraParams struct {

	// Body.
	Body *models.SentryBootstrapInfra

	/* MetadataName.

	   name of the resource
	*/
	MetadataName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bootstrap patch bootstrap infra params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BootstrapPatchBootstrapInfraParams) WithDefaults() *BootstrapPatchBootstrapInfraParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bootstrap patch bootstrap infra params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BootstrapPatchBootstrapInfraParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) WithTimeout(timeout time.Duration) *BootstrapPatchBootstrapInfraParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) WithContext(ctx context.Context) *BootstrapPatchBootstrapInfraParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) WithHTTPClient(client *http.Client) *BootstrapPatchBootstrapInfraParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) WithBody(body *models.SentryBootstrapInfra) *BootstrapPatchBootstrapInfraParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) SetBody(body *models.SentryBootstrapInfra) {
	o.Body = body
}

// WithMetadataName adds the metadataName to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) WithMetadataName(metadataName string) *BootstrapPatchBootstrapInfraParams {
	o.SetMetadataName(metadataName)
	return o
}

// SetMetadataName adds the metadataName to the bootstrap patch bootstrap infra params
func (o *BootstrapPatchBootstrapInfraParams) SetMetadataName(metadataName string) {
	o.MetadataName = metadataName
}

// WriteToRequest writes these params to a swagger request
func (o *BootstrapPatchBootstrapInfraParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param metadata.name
	if err := r.SetPathParam("metadata.name", o.MetadataName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
