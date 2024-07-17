// Package auth provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// GetAuthorizeParams defines parameters for GetAuthorize.
type GetAuthorizeParams struct {
	ResponseType string `form:"response_type" json:"response_type"`
	ClientId     string `form:"client_id" json:"client_id"`
	RedirectUri  string `form:"redirect_uri" json:"redirect_uri"`
	Scope        string `form:"scope" json:"scope"`
	State        string `form:"state" json:"state"`
}

// PostTokenFormdataBody defines parameters for PostToken.
type PostTokenFormdataBody struct {
	ClientId     *string `form:"client_id,omitempty" json:"client_id,omitempty"`
	ClientSecret *string `form:"client_secret,omitempty" json:"client_secret,omitempty"`
	Code         *string `form:"code,omitempty" json:"code,omitempty"`
	GrantType    *string `form:"grant_type,omitempty" json:"grant_type,omitempty"`
	Password     *string `form:"password,omitempty" json:"password,omitempty"`
	RedirectUri  *string `form:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	Username     *string `form:"username,omitempty" json:"username,omitempty"`
}

// PostTokenFormdataRequestBody defines body for PostToken for application/x-www-form-urlencoded ContentType.
type PostTokenFormdataRequestBody PostTokenFormdataBody

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
	// GetAuthorize request
	GetAuthorize(ctx context.Context, params *GetAuthorizeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostTokenWithBody request with any body
	PostTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostTokenWithFormdataBody(ctx context.Context, body PostTokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAuthorize(ctx context.Context, params *GetAuthorizeParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAuthorizeRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTokenRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTokenWithFormdataBody(ctx context.Context, body PostTokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTokenRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAuthorizeRequest generates requests for GetAuthorize
func NewGetAuthorizeRequest(server string, params *GetAuthorizeParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/authorize")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "response_type", runtime.ParamLocationQuery, params.ResponseType); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "client_id", runtime.ParamLocationQuery, params.ClientId); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "redirect_uri", runtime.ParamLocationQuery, params.RedirectUri); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "scope", runtime.ParamLocationQuery, params.Scope); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "state", runtime.ParamLocationQuery, params.State); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostTokenRequestWithFormdataBody calls the generic PostToken builder with application/x-www-form-urlencoded body
func NewPostTokenRequestWithFormdataBody(server string, body PostTokenFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewPostTokenRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}

// NewPostTokenRequestWithBody generates requests for PostToken with any type of body
func NewPostTokenRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/token")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
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
	// GetAuthorizeWithResponse request
	GetAuthorizeWithResponse(ctx context.Context, params *GetAuthorizeParams, reqEditors ...RequestEditorFn) (*GetAuthorizeResponse, error)

	// PostTokenWithBodyWithResponse request with any body
	PostTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTokenResponse, error)

	PostTokenWithFormdataBodyWithResponse(ctx context.Context, body PostTokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*PostTokenResponse, error)
}

type GetAuthorizeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetAuthorizeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAuthorizeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AccessToken  *string `json:"access_token,omitempty"`
		ExpiresIn    *int    `json:"expires_in,omitempty"`
		RefreshToken *string `json:"refresh_token,omitempty"`
		TokenType    *string `json:"token_type,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r PostTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAuthorizeWithResponse request returning *GetAuthorizeResponse
func (c *ClientWithResponses) GetAuthorizeWithResponse(ctx context.Context, params *GetAuthorizeParams, reqEditors ...RequestEditorFn) (*GetAuthorizeResponse, error) {
	rsp, err := c.GetAuthorize(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAuthorizeResponse(rsp)
}

// PostTokenWithBodyWithResponse request with arbitrary body returning *PostTokenResponse
func (c *ClientWithResponses) PostTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTokenResponse, error) {
	rsp, err := c.PostTokenWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTokenResponse(rsp)
}

func (c *ClientWithResponses) PostTokenWithFormdataBodyWithResponse(ctx context.Context, body PostTokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*PostTokenResponse, error) {
	rsp, err := c.PostTokenWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTokenResponse(rsp)
}

// ParseGetAuthorizeResponse parses an HTTP response from a GetAuthorizeWithResponse call
func ParseGetAuthorizeResponse(rsp *http.Response) (*GetAuthorizeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAuthorizeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParsePostTokenResponse parses an HTTP response from a PostTokenWithResponse call
func ParsePostTokenResponse(rsp *http.Response) (*PostTokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AccessToken  *string `json:"access_token,omitempty"`
			ExpiresIn    *int    `json:"expires_in,omitempty"`
			RefreshToken *string `json:"refresh_token,omitempty"`
			TokenType    *string `json:"token_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}