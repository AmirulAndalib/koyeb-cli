// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCatalogStore(params *GetCatalogStoreParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogStoreOK, error)

	ListCatalogStores(params *ListCatalogStoresParams, authInfo runtime.ClientAuthInfoWriter) (*ListCatalogStoresOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCatalogStore get catalog store API
*/
func (a *Client) GetCatalogStore(params *GetCatalogStoreParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCatalogStore",
		Method:             "GET",
		PathPattern:        "/v1/catalog/stores/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCatalogStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCatalogStores list catalog stores API
*/
func (a *Client) ListCatalogStores(params *ListCatalogStoresParams, authInfo runtime.ClientAuthInfoWriter) (*ListCatalogStoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogStoresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCatalogStores",
		Method:             "GET",
		PathPattern:        "/v1/catalog/stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCatalogStoresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCatalogStoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCatalogStores: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
