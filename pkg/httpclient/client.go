// Package httpclient hosts a simple HTTP client which supports sending and receiving JSON data using
// GET, POST, PUT, PATCH, and DELETE requests, with configurable timeouts.
//
// To create a new client, you have to call the following code:
//
// 		client := httpclient.NewClient()
//
// If you want to adjust the timeouts:
//
//		timeouts := httpclient.NewDefaultTimeouts()
//		// adjust any timeouts here
//		client := httpclient.NewClient(httpclient.WithTimeouts(timeouts))
//
// Then, to make a request, call one of the service methods, e.g.:
//		resp := client.GetJSON("http://site/path")
//
// Once you have an user and a corresponding public API key, you can issue authenticated requests,
// by constructing a new client with the appropriate credentials:
//
//		client := httpclient.NewClient(httpclient.WithDigestAuthentication(username, password))
//
package httpclient

import (
	"encoding/json"
	"fmt"
	"github.com/mongodb-labs/pcgc/pkg/useful"
	"gopkg.in/errgo.v1"
	"io"
	"net"
	"net/http"
	"runtime"
	"strings"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
)

const (
	// ContentTypeJSON defines the JSON content type
	ContentTypeJSON = "application/json; charset=UTF-8"
	// PreferJSON signal that we are accepting JSON responses, but do not reject non-JSON data
	PreferJSON = "application/json;q=0.9, */*;q=0.8"
)

var (
	// userAgent holds a user agent string which will be passed along with every request
	userAgent string
	// the version string
	version string
)

func init() {
	ver := version
	if ver == "" {
		// if the version is not passed at build time, flag it as 'unknown'
		ver = "unknown"
	}

	userAgent = fmt.Sprintf("pcgc/httpclient-%s (%s; %s)", ver, runtime.GOOS, runtime.GOARCH)
}

type basicClient struct {
	client                    *http.Client
	auth                      *digest.Transport
	listOfAcceptedStatusCodes []int
}

// HTTPResponse wrapper for HTTP response objects
type HTTPResponse struct {
	Response *http.Response
	Err      error
}

// BasicClient defines a contract for this client's API
type BasicClient interface {
	GetJSON(url string) HTTPResponse
	PostJSON(url string, body io.Reader) HTTPResponse
	PatchJSON(url string, body io.Reader) HTTPResponse
	PutJSON(url string, body io.Reader) HTTPResponse
	Delete(url string) HTTPResponse
}

// Error implementation for error responses
func (r HTTPResponse) Error() string {
	return r.Err.Error()
}

// IsError returns true if the associated error is not nil
func (r HTTPResponse) IsError() bool {
	return r.Err != nil
}

// NewClient builds a new client, allowing for dynamic configuration
// the order of the passed function matters, as they will be applied sequentially
func NewClient(configs ...func(*basicClient)) BasicClient {
	// initialize a bare client
	client := &basicClient{client: &http.Client{}}

	// configure defaults
	WithDefaultTimeouts()(client)
	WithAcceptedStatusCodes([]int{http.StatusOK, http.StatusCreated})(client)

	// apply any other configurations
	for _, configure := range configs {
		configure(client)
	}

	return *client
}

// WithDefaultTimeouts configures a client with default timeouts
func WithDefaultTimeouts() func(*basicClient) {
	return WithTimeouts(NewDefaultTimeouts())
}

// WithAcceptedStatusCodes configures a client with a list of accepted HTTP response status codes
func WithAcceptedStatusCodes(acceptedStatusCodes []int) func(*basicClient) {
	return func(client *basicClient) {
		client.listOfAcceptedStatusCodes = acceptedStatusCodes
	}
}

// WithTimeouts configures a client with the specified timeouts
func WithTimeouts(t *RequestTimeouts) func(*basicClient) {
	return func(client *basicClient) {
		// set global (total) timeout
		client.client.Timeout = t.GlobalTimeout

		// set all other t
		client.client.Transport = &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: t.DialTimeout,
			}).DialContext,
			ExpectContinueTimeout: t.ExpectContinueTimeout,
			IdleConnTimeout:       t.IdleConnectionTimeout,
			ResponseHeaderTimeout: t.ResponseHeaderTimeout,
			TLSHandshakeTimeout:   t.TLSHandshakeTimeout,
		}
	}
}

// WithDigestAuthentication configures a client with digest authentication credentials
func WithDigestAuthentication(username string, password string) func(*basicClient) {
	return func(client *basicClient) {
		client.auth = digest.NewTransport(username, password)
	}
}

// GetJSON retrieves the specified URL
func (c basicClient) GetJSON(url string) HTTPResponse {
	return c.genericJSONRequest(http.MethodGet, url, nil)
}

// PostJson executes a POST request, sending the specified body, encoded as JSON, to the passed URL
func (c basicClient) PostJSON(url string, body io.Reader) HTTPResponse {
	return c.genericJSONRequest(http.MethodPost, url, body)
}

// PutJSON executes a PUT request, sending the specified body, encoded as JSON, to the passed URL
func (c basicClient) PutJSON(url string, body io.Reader) (resp HTTPResponse) {
	return c.genericJSONRequest(http.MethodPut, url, body)
}

// PatchJSON executes a PATCH request, sending the specified body, encoded as JSON, to the passed URL
func (c basicClient) PatchJSON(url string, body io.Reader) (resp HTTPResponse) {
	return c.genericJSONRequest(http.MethodPatch, url, body)
}

// Delete executes a DELETE request
func (c basicClient) Delete(url string) (resp HTTPResponse) {
	return c.genericJSONRequest(http.MethodDelete, url, nil)
}

// CloseResponseBodyIfNotNil simple helper which can ensure a response's body is correctly closed, if one exists
func CloseResponseBodyIfNotNil(resp HTTPResponse) {
	if resp.Response == nil {
		return
	}

	if resp.Response.Body == nil {
		return
	}

	// if a body exists, attempt to close it and log any errors
	useful.LogError(resp.Response.Body.Close)
}

func (c basicClient) genericJSONRequest(verb string, url string, body io.Reader) (resp HTTPResponse) {
	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		resp.Err = err
		return
	}

	req.Header.Add("Accept", PreferJSON)
	req.Header.Add("User-Agent", userAgent)
	if body != nil {
		// only set the request content type if the body is non nil
		req.Header.Add("Content-Type", ContentTypeJSON)
	}

	if c.auth != nil {
		// if we have authentication credentials, use them
		resp.Response, resp.Err = c.auth.RoundTrip(req)
	} else {
		// otherwise issue an unauthenticated request
		resp.Response, resp.Err = c.client.Do(req)
	}

	if !validateStatusCode(&resp, c.listOfAcceptedStatusCodes, verb, url) {
		// if the response code is not expected, stop here
		return
	}

	return
}

func validateStatusCode(r *HTTPResponse, expectedStatuses []int, verb string, url string) bool {
	// no response => not valid
	if r == nil || r.Response == nil {
		return false
	}

	// nothing to check
	if len(expectedStatuses) == 0 {
		return true
	}

	// check if the resulting status is one of the expected ones
	for _, status := range expectedStatuses {
		if r.Response.StatusCode == status {
			return true
		}
	}

	// parse response body
	defer CloseResponseBodyIfNotNil(*r)
	var errorDetails interface{}
	decoder := json.NewDecoder(r.Response.Body)
	err := decoder.Decode(&errorDetails)
	useful.PanicOnUnrecoverableError(err)

	// otherwise augment the error and return false
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("failed to execute %s request to %s\n", verb, url))
	sb.WriteString(fmt.Sprintf("status code: %d\n", r.Response.StatusCode))
	sb.WriteString(fmt.Sprintf("response: %s\n", r.Response.Status))
	sb.WriteString(fmt.Sprintf("details: %s\n", errorDetails))
	r.Err = errgo.Notef(r.Err, sb.String())

	return false
}
