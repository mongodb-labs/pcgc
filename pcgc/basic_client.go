package pcgc

import (
	"io"
	"net/http"
	"time"

	"gopkg.in/errgo.v1"
)

type basicHTTPClient struct {
	client *http.Client
}

// HTTPResponse wrapper for HTTP response objects
type HTTPResponse struct {
	resp *http.Response
	err  error
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
	return resp.err.Error()
}

// IsError returns true if the associated error is not nil
func (resp HTTPResponse) IsError() bool {
	return resp.err != nil
}

// NewClient build a new HTTP client
func NewClient() BasicHTTPOperation {
	return basicHTTPClient{client: &http.Client{
		Transport: &http.Transport{
			ExpectContinueTimeout: 1 * time.Second,
			IdleConnTimeout:       RequestTimeout,
			ResponseHeaderTimeout: ResponseHeaderTimeout,
			TLSHandshakeTimeout:   RequestTimeout,
		},
		Timeout: HTTPRequestResponseTimeout,
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
		resp.err = err
		return
	}

	if body != nil {
		// only set the content type if the body is non nil
		req.Header.Add("Content-Type", ContentTypeJSON)
	}

	resp.resp, resp.err = cl.client.Do(req)
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
		if resp.resp.StatusCode == status {
			return true
		}
	}

	// otherwise augment the error and return false
	resp.err = errgo.Notef(resp.err, "Failed to execute %s request to %s; got status code %d (%v)", verb, url, resp.resp.StatusCode, resp.resp.Status)
	return false
}
