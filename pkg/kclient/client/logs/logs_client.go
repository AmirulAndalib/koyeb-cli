// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new logs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for logs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	LogsTailStackEvents(params *LogsTailStackEventsParams, authInfo runtime.ClientAuthInfoWriter) (*LogsTailStackEventsOK, error)

	LogsTailStackRevisionLogs(params *LogsTailStackRevisionLogsParams, authInfo runtime.ClientAuthInfoWriter) (*LogsTailStackRevisionLogsOK, error)

	LogsTailStackRevisionLogsForFunction(params *LogsTailStackRevisionLogsForFunctionParams, authInfo runtime.ClientAuthInfoWriter) (*LogsTailStackRevisionLogsForFunctionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LogsTailStackEvents logs tail stack events API
*/
func (a *Client) LogsTailStackEvents(params *LogsTailStackEventsParams, authInfo runtime.ClientAuthInfoWriter) (*LogsTailStackEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogsTailStackEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "logs_TailStackEvents",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{stack_id}/events/tail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogsTailStackEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LogsTailStackEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LogsTailStackEventsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  LogsTailStackRevisionLogs logs tail stack revision logs API
*/
func (a *Client) LogsTailStackRevisionLogs(params *LogsTailStackRevisionLogsParams, authInfo runtime.ClientAuthInfoWriter) (*LogsTailStackRevisionLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogsTailStackRevisionLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "logs_TailStackRevisionLogs",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{stack_id}/revisions/{sha}/logs/tail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogsTailStackRevisionLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LogsTailStackRevisionLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LogsTailStackRevisionLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  LogsTailStackRevisionLogsForFunction logs tail stack revision logs for function API
*/
func (a *Client) LogsTailStackRevisionLogsForFunction(params *LogsTailStackRevisionLogsForFunctionParams, authInfo runtime.ClientAuthInfoWriter) (*LogsTailStackRevisionLogsForFunctionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogsTailStackRevisionLogsForFunctionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "logs_TailStackRevisionLogsForFunction",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{stack_id}/revisions/{sha}/functions/{fn_name}/logs/tail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogsTailStackRevisionLogsForFunctionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LogsTailStackRevisionLogsForFunctionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LogsTailStackRevisionLogsForFunctionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}