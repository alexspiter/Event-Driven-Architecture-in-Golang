// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DisableCustomer(params *DisableCustomerParams, opts ...ClientOption) (*DisableCustomerOK, error)

	ChangeSmsNumber(params *ChangeSmsNumberParams, opts ...ClientOption) (*ChangeSmsNumberOK, error)

	EnableCustomer(params *EnableCustomerParams, opts ...ClientOption) (*EnableCustomerOK, error)

	GetCustomer(params *GetCustomerParams, opts ...ClientOption) (*GetCustomerOK, error)

	RegisterCustomer(params *RegisterCustomerParams, opts ...ClientOption) (*RegisterCustomerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DisableCustomer disables a customer
*/
func (a *Client) DisableCustomer(params *DisableCustomerParams, opts ...ClientOption) (*DisableCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableCustomerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DisableCustomer",
		Method:             "PUT",
		PathPattern:        "/api/customers/{id}/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableCustomerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DisableCustomerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeSmsNumber changes a customers s m s number
*/
func (a *Client) ChangeSmsNumber(params *ChangeSmsNumberParams, opts ...ClientOption) (*ChangeSmsNumberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeSmsNumberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changeSmsNumber",
		Method:             "PUT",
		PathPattern:        "/api/customers/{id}/change-sms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangeSmsNumberReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeSmsNumberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeSmsNumberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EnableCustomer enables a customer
*/
func (a *Client) EnableCustomer(params *EnableCustomerParams, opts ...ClientOption) (*EnableCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableCustomerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enableCustomer",
		Method:             "PUT",
		PathPattern:        "/api/customers/{id}/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnableCustomerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EnableCustomerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCustomer gets a customer
*/
func (a *Client) GetCustomer(params *GetCustomerParams, opts ...ClientOption) (*GetCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCustomer",
		Method:             "GET",
		PathPattern:        "/api/customers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCustomerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCustomerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RegisterCustomer registers a new customer
*/
func (a *Client) RegisterCustomer(params *RegisterCustomerParams, opts ...ClientOption) (*RegisterCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterCustomerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "registerCustomer",
		Method:             "POST",
		PathPattern:        "/api/customers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterCustomerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegisterCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RegisterCustomerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
