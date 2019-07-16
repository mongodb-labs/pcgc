// Package opsmanager is a HTTP client which abstracts communication with an Ops Manager instance.
//
// To create a new client, you have to call the following code:
//
//		resolver := httpclient.NewURLResolverWithPrefix("http://OPS-MANAGER-INSTANCE", "/api/public/v1.0")
// 		client := opsmanager.NewClient(resolver)
//
// The client can then be used to issue requests such as:
//
//		user := User{Username: ..., Password: ..., ...}
// 		globalOwner := client.CreateFirstUser(user, WhitelistAllowAll)
//
// See the Client interface below for a list of all the support operations.
// If however, you need one that is not currently supported, the _opsManagerApi_ struct extends
// _httpclient.BasicHTTPOperation_, allowing you to issue raw HTTP requests to the specified Ops Manager instance.
//
// 		url := resolver.Of("/path/to/a/resource/%s", id)
//		resp:= client.Get(url)
//		useful.PanicOnUnrecoverableError(resp.Err)
//		defer useful.LogError(resp.Response.Body.Close)
//		var data SomeType
//		decoder := json.NewDecoder(resp.Response.Body)
//		err := decoder.Decode(&result)
//		useful.PanicOnUnrecoverableError(err)
//		// do something with _data_
//
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
	GetAllProjects() (Projects, error)
}

// NewClient builds a new API client for connecting to Ops Manager
func NewClient(resolver httpclient.URLResolver) Client {
	return opsManagerAPI{BasicHTTPOperation: httpclient.NewClient(), resolver: resolver}
}
