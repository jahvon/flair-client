// Package flair provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package flair

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

// Bridge defines model for Bridge.
type Bridge struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Location  *string    `json:"location,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Status    *string    `json:"status,omitempty"`
}

// BridgeReading defines model for BridgeReading.
type BridgeReading struct {
	Humidity       *float32   `json:"humidity,omitempty"`
	SignalStrength *float32   `json:"signalStrength,omitempty"`
	Temperature    *float32   `json:"temperature,omitempty"`
	Timestamp      *time.Time `json:"timestamp,omitempty"`
}

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

// Room defines model for Room.
type Room struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Location  *string    `json:"location,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Status    *string    `json:"status,omitempty"`
}

// Structure defines model for Structure.
type Structure struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Location  *string    `json:"location,omitempty"`
	Name      *string    `json:"name,omitempty"`
}

// PatchRoomSetPointJSONBody defines parameters for PatchRoomSetPoint.
type PatchRoomSetPointJSONBody struct {
	SetPoint *float32 `json:"setPoint,omitempty"`
}

// PatchStructureJSONBody defines parameters for PatchStructure.
type PatchStructureJSONBody struct {
	Location *string `json:"location,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// PatchRoomSetPointJSONRequestBody defines body for PatchRoomSetPoint for application/json ContentType.
type PatchRoomSetPointJSONRequestBody PatchRoomSetPointJSONBody

// PatchStructureJSONRequestBody defines body for PatchStructure for application/json ContentType.
type PatchStructureJSONRequestBody PatchStructureJSONBody

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
	// GetBridges request
	GetBridges(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetBridgeCurrentReading request
	GetBridgeCurrentReading(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPucks request
	GetPucks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPuckCurrentReading request
	GetPuckCurrentReading(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetRooms request
	GetRooms(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchRoomSetPointWithBody request with any body
	PatchRoomSetPointWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchRoomSetPoint(ctx context.Context, id string, body PatchRoomSetPointJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetStructures request
	GetStructures(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchStructureWithBody request with any body
	PatchStructureWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchStructure(ctx context.Context, id string, body PatchStructureJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetBridges(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetBridgesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetBridgeCurrentReading(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetBridgeCurrentReadingRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
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

func (c *Client) GetStructures(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetStructuresRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchStructureWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchStructureRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchStructure(ctx context.Context, id string, body PatchStructureJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchStructureRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetBridgesRequest generates requests for GetBridges
func NewGetBridgesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/bridges")
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

// NewGetBridgeCurrentReadingRequest generates requests for GetBridgeCurrentReading
func NewGetBridgeCurrentReadingRequest(server string, id string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v1/bridges/%s/current-reading", pathParam0)
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

// NewGetRoomsRequest generates requests for GetRooms
func NewGetRoomsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/rooms")
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

	operationPath := fmt.Sprintf("/api/v1/rooms/%s/set-point", pathParam0)
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

// NewGetStructuresRequest generates requests for GetStructures
func NewGetStructuresRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/structures")
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

// NewPatchStructureRequest calls the generic PatchStructure builder with application/json body
func NewPatchStructureRequest(server string, id string, body PatchStructureJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPatchStructureRequestWithBody(server, id, "application/json", bodyReader)
}

// NewPatchStructureRequestWithBody generates requests for PatchStructure with any type of body
func NewPatchStructureRequestWithBody(server string, id string, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v1/structures/%s", pathParam0)
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
	// GetBridgesWithResponse request
	GetBridgesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBridgesResponse, error)

	// GetBridgeCurrentReadingWithResponse request
	GetBridgeCurrentReadingWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetBridgeCurrentReadingResponse, error)

	// GetPucksWithResponse request
	GetPucksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetPucksResponse, error)

	// GetPuckCurrentReadingWithResponse request
	GetPuckCurrentReadingWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetPuckCurrentReadingResponse, error)

	// GetRoomsWithResponse request
	GetRoomsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRoomsResponse, error)

	// PatchRoomSetPointWithBodyWithResponse request with any body
	PatchRoomSetPointWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchRoomSetPointResponse, error)

	PatchRoomSetPointWithResponse(ctx context.Context, id string, body PatchRoomSetPointJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchRoomSetPointResponse, error)

	// GetStructuresWithResponse request
	GetStructuresWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetStructuresResponse, error)

	// PatchStructureWithBodyWithResponse request with any body
	PatchStructureWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchStructureResponse, error)

	PatchStructureWithResponse(ctx context.Context, id string, body PatchStructureJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchStructureResponse, error)
}

type GetBridgesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Bridge
}

// Status returns HTTPResponse.Status
func (r GetBridgesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetBridgesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetBridgeCurrentReadingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BridgeReading
}

// Status returns HTTPResponse.Status
func (r GetBridgeCurrentReadingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetBridgeCurrentReadingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
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

type GetRoomsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Room
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
	JSON200      *Room
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

type GetStructuresResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Structure
}

// Status returns HTTPResponse.Status
func (r GetStructuresResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetStructuresResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchStructureResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Structure
}

// Status returns HTTPResponse.Status
func (r PatchStructureResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchStructureResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetBridgesWithResponse request returning *GetBridgesResponse
func (c *ClientWithResponses) GetBridgesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBridgesResponse, error) {
	rsp, err := c.GetBridges(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetBridgesResponse(rsp)
}

// GetBridgeCurrentReadingWithResponse request returning *GetBridgeCurrentReadingResponse
func (c *ClientWithResponses) GetBridgeCurrentReadingWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetBridgeCurrentReadingResponse, error) {
	rsp, err := c.GetBridgeCurrentReading(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetBridgeCurrentReadingResponse(rsp)
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

// GetStructuresWithResponse request returning *GetStructuresResponse
func (c *ClientWithResponses) GetStructuresWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetStructuresResponse, error) {
	rsp, err := c.GetStructures(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetStructuresResponse(rsp)
}

// PatchStructureWithBodyWithResponse request with arbitrary body returning *PatchStructureResponse
func (c *ClientWithResponses) PatchStructureWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchStructureResponse, error) {
	rsp, err := c.PatchStructureWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchStructureResponse(rsp)
}

func (c *ClientWithResponses) PatchStructureWithResponse(ctx context.Context, id string, body PatchStructureJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchStructureResponse, error) {
	rsp, err := c.PatchStructure(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchStructureResponse(rsp)
}

// ParseGetBridgesResponse parses an HTTP response from a GetBridgesWithResponse call
func ParseGetBridgesResponse(rsp *http.Response) (*GetBridgesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetBridgesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Bridge
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetBridgeCurrentReadingResponse parses an HTTP response from a GetBridgeCurrentReadingWithResponse call
func ParseGetBridgeCurrentReadingResponse(rsp *http.Response) (*GetBridgeCurrentReadingResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetBridgeCurrentReadingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BridgeReading
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
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
		var dest []Room
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
		var dest Room
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetStructuresResponse parses an HTTP response from a GetStructuresWithResponse call
func ParseGetStructuresResponse(rsp *http.Response) (*GetStructuresResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetStructuresResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Structure
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePatchStructureResponse parses an HTTP response from a PatchStructureWithResponse call
func ParsePatchStructureResponse(rsp *http.Response) (*PatchStructureResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PatchStructureResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Structure
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}