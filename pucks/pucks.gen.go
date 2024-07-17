// Package pucks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package pucks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/oapi-codegen/runtime"
)

// Puck defines model for Puck.
type Puck struct {
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *string    `json:"id,omitempty"`
	Location    *string    `json:"location,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Status      *string    `json:"status,omitempty"`
}

// Reading defines model for Reading.
type Reading struct {
	Battery     *float32   `json:"battery,omitempty"`
	Humidity    *float32   `json:"humidity,omitempty"`
	Temperature *float32   `json:"temperature,omitempty"`
	Timestamp   *time.Time `json:"timestamp,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetPucks request
	GetPucks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPuckCurrentReading request
	GetPuckCurrentReading(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetPucks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPucksRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetPuckCurrentReading(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPuckCurrentReadingRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetPucksRequest generates requests for GetPucks
func NewGetPucksRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/pucks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPuckCurrentReadingRequest generates requests for GetPuckCurrentReading
func NewGetPuckCurrentReadingRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/pucks/%s/current-reading", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetPucksWithResponse request
	GetPucksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetPucksResponse, error)

	// GetPuckCurrentReadingWithResponse request
	GetPuckCurrentReadingWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetPuckCurrentReadingResponse, error)
}

type GetPucksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Puck
}

// Status returns HTTPResponse.Status
func (r GetPucksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPucksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPuckCurrentReadingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Reading
}

// Status returns HTTPResponse.Status
func (r GetPuckCurrentReadingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPuckCurrentReadingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetPucksWithResponse request returning *GetPucksResponse
func (c *ClientWithResponses) GetPucksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetPucksResponse, error) {
	rsp, err := c.GetPucks(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPucksResponse(rsp)
}

// GetPuckCurrentReadingWithResponse request returning *GetPuckCurrentReadingResponse
func (c *ClientWithResponses) GetPuckCurrentReadingWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetPuckCurrentReadingResponse, error) {
	rsp, err := c.GetPuckCurrentReading(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPuckCurrentReadingResponse(rsp)
}

// ParseGetPucksResponse parses an HTTP response from a GetPucksWithResponse call
func ParseGetPucksResponse(rsp *http.Response) (*GetPucksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPucksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Puck
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetPuckCurrentReadingResponse parses an HTTP response from a GetPuckCurrentReadingWithResponse call
func ParseGetPuckCurrentReadingResponse(rsp *http.Response) (*GetPuckCurrentReadingResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPuckCurrentReadingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Reading
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}