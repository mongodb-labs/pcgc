package opsmanager

import (
	"github.com/mongodb-labs/pcgc/pkg/httpclient"
)

type opsManagerAPI struct {
	httpclient.BasicHTTPOperation

	resolver httpclient.URLResolver
}

// Client defines the API actions implemented in this client
type Client interface {
	httpclient.BasicHTTPOperation

	CreateFirstUser(user User, whitelistIP string) (CreateFirstUserResponse, error)
}

// NewClient builds a new API client for connecting to Ops Manager
func NewClient(resolver httpclient.URLResolver) Client {
	return opsManagerAPI{BasicHTTPOperation: httpclient.NewClient(), resolver: resolver}
}
