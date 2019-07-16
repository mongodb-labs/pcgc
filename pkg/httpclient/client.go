// Package httpclient is a simple HTTP client which supports sending and receiving JSON strings using
// GET, POST, PUT, PATCH, and DELETE requests, with configurable timeouts.
//
// To create a new client, you have to call the following code:
//
// 		client := httpclient.NewClient()
//
// If you want to adjust the timeouts:
//
//		timeouts := InitTimeouts()
//		// adjust any timeouts here
//		client := httpclient.NewClientWithTimeouts(timeouts)
//
// Then, to make a request, call one of the service methods, e.g.:
//		resp := client.Get("http://site/path")
//
package httpclient

import (
	"fmt"
	"gopkg.in/errgo.v1"
	"io"
	"net"
	"net/http"
	"runtime"
)

// ContentTypeJSON defines the JSON content type
const ContentTypeJSON = "application/json; charset=UTF-8"

// PreferJSON signal that we are accepting JSON responses, but do not reject non-JSON data
const PreferJSON = "application/json;q=0.9, */*;q=0.8"

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

type basicHTTPClient struct {
	client *http.Client
}

// HTTPResponse wrapper for HTTP response objects
type HTTPResponse struct {
	Response *http.Response
	Err      error
}

// BasicHTTPOperation defines a contract for this client's API
type BasicHTTPOperation interface {
	Get(url string) HTTPResponse
	PostJSON(url string, body io.Reader) HTTPResponse
	PatchJSON(url string, body io.Reader) HTTPResponse
	PutJSON(url string, body io.Reader) HTTPResponse
	Delete(url string) HTTPResponse
}

// Error implementation for error responses
func (resp HTTPResponse) Error() string {
	return resp.Err.Error()
}

// IsError returns true if the associated error is not nil
func (resp HTTPResponse) IsError() bool {
	return resp.Err != nil
}

// NewClient build a new HTTP client with default timeouts
func NewClient() BasicHTTPOperation {
	return NewClientWithTimeouts(InitTimeouts())
}

// NewClientWithTimeouts build a new HTTP client with specified timeouts
func NewClientWithTimeouts(timeouts *RequestTimeouts) BasicHTTPOperation {
	return basicHTTPClient{client: &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: timeouts.DialTimeout,
			}).DialContext,
			ExpectContinueTimeout: timeouts.ExpectContinueTimeout,
			IdleConnTimeout:       timeouts.IdleConnectionTimeout,
			ResponseHeaderTimeout: timeouts.ResponseHeaderTimeout,
			TLSHandshakeTimeout:   timeouts.TLSHandshakeTimeout,
		},
		Timeout: timeouts.GlobalTimeout,
	}}
}

// Get retrieves the specified URL
func (cl basicHTTPClient) Get(url string) HTTPResponse {
	return cl.genericJSONRequest("GET", url, nil, []int{http.StatusOK})
}

// PostJson executes a POST request, sending the specified body, encoded as JSON, to the passed URL
func (cl basicHTTPClient) PostJSON(url string, body io.Reader) HTTPResponse {
	return cl.genericJSONRequest("POST", url, body, []int{http.StatusOK})
}

// PutJSON executes a PUT request, sending the specified body, encoded as JSON, to the passed URL
func (cl basicHTTPClient) PutJSON(url string, body io.Reader) (resp HTTPResponse) {
	return cl.genericJSONRequest("PUT", url, body, []int{http.StatusOK})
}

// PatchJSON executes a PATCH request, sending the specified body, encoded as JSON, to the passed URL
func (cl basicHTTPClient) PatchJSON(url string, body io.Reader) (resp HTTPResponse) {
	return cl.genericJSONRequest("PATCH", url, body, []int{http.StatusOK})
}

// Delete executes a DELETE request
func (cl basicHTTPClient) Delete(url string) (resp HTTPResponse) {
	return cl.genericJSONRequest("DELETE", url, nil, []int{http.StatusOK})
}

func (cl basicHTTPClient) genericJSONRequest(verb string, url string, body io.Reader, expectedStatuses []int) (resp HTTPResponse) {
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

	resp.Response, resp.Err = cl.client.Do(req)
	if !validateStatusCode(&resp, expectedStatuses, verb, url) {
		return
	}

	return
}

func validateStatusCode(resp *HTTPResponse, expectedStatuses []int, verb string, url string) bool {
	// nothing to check
	if len(expectedStatuses) == 0 {
		return true
	}

	// check if the resulting status is one of the expected ones
	for _, status := range expectedStatuses {
		if resp.Response.StatusCode == status {
			return true
		}
	}

	// otherwise augment the error and return false
	resp.Err = errgo.Notef(resp.Err, "Failed to execute %s request to %s; got status code %d (%v)", verb, url, resp.Response.StatusCode, resp.Response.Status)
	return false
}
