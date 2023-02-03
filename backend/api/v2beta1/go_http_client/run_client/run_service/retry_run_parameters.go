// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRetryRunParams creates a new RetryRunParams object
// with the default values initialized.
func NewRetryRunParams() *RetryRunParams {
	var ()
	return &RetryRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetryRunParamsWithTimeout creates a new RetryRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetryRunParamsWithTimeout(timeout time.Duration) *RetryRunParams {
	var ()
	return &RetryRunParams{

		timeout: timeout,
	}
}

// NewRetryRunParamsWithContext creates a new RetryRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetryRunParamsWithContext(ctx context.Context) *RetryRunParams {
	var ()
	return &RetryRunParams{

		Context: ctx,
	}
}

// NewRetryRunParamsWithHTTPClient creates a new RetryRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetryRunParamsWithHTTPClient(client *http.Client) *RetryRunParams {
	var ()
	return &RetryRunParams{
		HTTPClient: client,
	}
}

/*RetryRunParams contains all the parameters to send to the API endpoint
for the retry run operation typically these are written to a http.Request
*/
type RetryRunParams struct {

	/*ExperimentID
	  The ID of the parent experiment.

	*/
	ExperimentID string
	/*RunID
	  The ID of the run to be Retried.

	*/
	RunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retry run params
func (o *RetryRunParams) WithTimeout(timeout time.Duration) *RetryRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retry run params
func (o *RetryRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retry run params
func (o *RetryRunParams) WithContext(ctx context.Context) *RetryRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retry run params
func (o *RetryRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retry run params
func (o *RetryRunParams) WithHTTPClient(client *http.Client) *RetryRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retry run params
func (o *RetryRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExperimentID adds the experimentID to the retry run params
func (o *RetryRunParams) WithExperimentID(experimentID string) *RetryRunParams {
	o.SetExperimentID(experimentID)
	return o
}

// SetExperimentID adds the experimentId to the retry run params
func (o *RetryRunParams) SetExperimentID(experimentID string) {
	o.ExperimentID = experimentID
}

// WithRunID adds the runID to the retry run params
func (o *RetryRunParams) WithRunID(runID string) *RetryRunParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the retry run params
func (o *RetryRunParams) SetRunID(runID string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *RetryRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param experiment_id
	if err := r.SetPathParam("experiment_id", o.ExperimentID); err != nil {
		return err
	}

	// path param run_id
	if err := r.SetPathParam("run_id", o.RunID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
