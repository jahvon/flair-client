// Package rooms provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package rooms

import (
	"bytes"
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

// Room defines model for Room.
type Room struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Location  *string    `json:"location,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Status    *string    `json:"status,omitempty"`
}

// PatchRoomSetPointJSONBody defines parameters for PatchRoomSetPoint.
type PatchRoomSetPointJSONBody struct {
	SetPoint *float32 `json:"setPoint,omitempty"`
}

// PatchRoomSetPointJSONRequestBody defines body for PatchRoomSetPoint for application/json ContentType.
type PatchRoomSetPointJSONRequestBody PatchRoomSetPointJSONBody

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
	// GetRooms request
	GetRooms(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchRoomSetPointWithBody request with any body
	PatchRoomSetPointWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchRoomSetPoint(ctx context.Context, id string, body PatchRoomSetPointJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetRooms(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRoomsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchRoomSetPointWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchRoomSetPointRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchRoomSetPoint(ctx context.Context, id string, body PatchRoomSetPointJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchRoomSetPointRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetRoomsRequest generates requests for GetRooms
func NewGetRoomsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/rooms")
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

// NewPatchRoomSetPointRequest calls the generic PatchRoomSetPoint builder with application/json body
func NewPatchRoomSetPointRequest(server string, id string, body PatchRoomSetPointJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPatchRoomSetPointRequestWithBody(server, id, "application/json", bodyReader)
}

// NewPatchRoomSetPointRequestWithBody generates requests for PatchRoomSetPoint with any type of body
func NewPatchRoomSetPointRequestWithBody(server string, id string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/rooms/%s/set-point", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// GetRoomsWithResponse request
	GetRoomsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRoomsResponse, error)

	// PatchRoomSetPointWithBodyWithResponse request with any body
	PatchRoomSetPointWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchRoomSetPointResponse, error)

	PatchRoomSetPointWithResponse(ctx context.Context, id string, body PatchRoomSetPointJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchRoomSetPointResponse, error)
}

type GetRoomsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *[]Room `json:"data,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetRoomsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRoomsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchRoomSetPointResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *Room `json:"data,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r PatchRoomSetPointResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchRoomSetPointResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetRoomsWithResponse request returning *GetRoomsResponse
func (c *ClientWithResponses) GetRoomsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRoomsResponse, error) {
	rsp, err := c.GetRooms(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRoomsResponse(rsp)
}

// PatchRoomSetPointWithBodyWithResponse request with arbitrary body returning *PatchRoomSetPointResponse
func (c *ClientWithResponses) PatchRoomSetPointWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchRoomSetPointResponse, error) {
	rsp, err := c.PatchRoomSetPointWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchRoomSetPointResponse(rsp)
}

func (c *ClientWithResponses) PatchRoomSetPointWithResponse(ctx context.Context, id string, body PatchRoomSetPointJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchRoomSetPointResponse, error) {
	rsp, err := c.PatchRoomSetPoint(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchRoomSetPointResponse(rsp)
}

// ParseGetRoomsResponse parses an HTTP response from a GetRoomsWithResponse call
func ParseGetRoomsResponse(rsp *http.Response) (*GetRoomsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRoomsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *[]Room `json:"data,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePatchRoomSetPointResponse parses an HTTP response from a PatchRoomSetPointWithResponse call
func ParsePatchRoomSetPointResponse(rsp *http.Response) (*PatchRoomSetPointResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PatchRoomSetPointResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *Room `json:"data,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
