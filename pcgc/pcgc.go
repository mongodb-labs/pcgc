// Package pcgc is a service client for Cloud Manager API.
//
// To be able to interact with this API, you have to
// create a new service:
//
//     s := pcgc.NewService(nil)
//
// The Service struct has all the methods you need
// to interact with Cloud Manager API.
//
package pcgc

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/bobziuchkovski/digest"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"reflect"
	"runtime"
)

const (
	// Version Library version
	Version = "v1"
	// DefaultUserAgent User agent
	DefaultUserAgent = "pcgc/" + Version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
	// DefaultURL api base url
	DefaultURL = "http://localhost:8080"
	// DefaultMediaType default media type
	DefaultMediaType = "application/json"
)

// Service represents your API.
type Service struct {
	client   *http.Client
	URL      string
	user     string
	password string
}

// NewService creates a Service using the given client, if none is provided
// it uses http.DefaultClient.
func NewService(c *http.Client, u string, p string) *Service {
	if c == nil {
		c = http.DefaultClient
	}
	return &Service{
		client:   c,
		URL:      DefaultURL,
		user:     u,
		password: p,
	}
}

// NewRequest generates an HTTP request, but does not perform the request.
func (s *Service) NewRequest(ctx context.Context, method, path string, body interface{}, q interface{}) (*http.Request, error) {

	var rbody io.Reader
	switch t := body.(type) {
	case nil:
	case string:
		rbody = bytes.NewBufferString(t)
	case io.Reader:
		rbody = t
	default:
		v := reflect.ValueOf(body)
		if !v.IsValid() {
			break
		}
		if v.Type().Kind() == reflect.Ptr {
			v = reflect.Indirect(v)
			if !v.IsValid() {
				break
			}
		}
		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rbody = bytes.NewReader(j)
	}
	req, err := http.NewRequest(method, s.URL+path, rbody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if q != nil {
		v, err := query.Values(q)
		if err != nil {
			return nil, err
		}
		query2 := v.Encode()
		if req.URL.RawQuery != "" && query2 != "" {
			req.URL.RawQuery += "&"
		}
		req.URL.RawQuery += query2
	}
	req.Header.Set("Accept", DefaultMediaType)
	req.Header.Set("User-Agent", DefaultUserAgent)

	return req, nil
}

// Do sends a request and decodes the response into v.
func (s *Service) Do(ctx context.Context, v interface{}, method, path string, body interface{}, q interface{}) error {
	req, err := s.NewRequest(ctx, method, path, body, q)
	if err != nil {
		return err
	}
	t := digest.NewTransport(s.user, s.password)

	resp, err := t.RoundTrip(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch t := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(t, resp.Body)
	default:
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return err
}

// Get sends a GET request and decodes the response into v.
func (s *Service) Get(ctx context.Context, v interface{}, path string, query interface{}) error {
	return s.Do(ctx, v, "GET", path, nil, query)
}

// Patch sends a Path request and decodes the response into v.
func (s *Service) Patch(ctx context.Context, v interface{}, path string, body interface{}) error {
	return s.Do(ctx, v, "PATCH", path, body, nil)
}

// Post sends a POST request and decodes the response into v.
func (s *Service) Post(ctx context.Context, v interface{}, path string, body interface{}) error {
	return s.Do(ctx, v, "POST", path, body, nil)
}

// Put sends a PUT request and decodes the response into v.
func (s *Service) Put(ctx context.Context, v interface{}, path string, body interface{}) error {
	return s.Do(ctx, v, "PUT", path, body, nil)
}

// Delete sends a DELETE request.
func (s *Service) Delete(ctx context.Context, v interface{}, path string) error {
	return s.Do(ctx, v, "DELETE", path, nil, nil)
}
