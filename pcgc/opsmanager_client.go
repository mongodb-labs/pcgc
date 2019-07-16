package pcgc

import (
	"bytes"
	"encoding/json"
)

type opsManagerAPI struct {
	BasicHTTPOperation

	resolver URLResolver
}

// OpsManagerUser request object which identifies a user
type OpsManagerUser struct {
	Username     string `json:"username"`
	Password     string `json:"password,omitempty"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

// OpsManagerUserRole denotes a single user role
type OpsManagerUserRole struct {
	GroupID  string `json:"groupId,omitempty"`
	OrgID    string `json:"orgId,omitempty"`
	RoleName string `json:"roleName,omitempty"`
}

// CreateFirstUserResponse API response for the CreateFirstUser() call
type CreateFirstUserResponse struct {
	OpsManagerUser

	ID           string               `json:"id"`
	Links        []interface{}        `json:"links,omitempty"`
	MobileNumber string               `json:"mobileNumber,omitempty"`
	Roles        []OpsManagerUserRole `json:"roles,omitempty"`
}

// OpsManagerClient defines the API actions implemented in this client
type OpsManagerClient interface {
	BasicHTTPOperation

	CreateFirstUser(user OpsManagerUser, whitelistIP string) (CreateFirstUserResponse, error)
}

// NewOpsManagerClient builds a new API client for connecting to Ops Manager
func NewOpsManagerClient(resolver URLResolver) OpsManagerClient {
	return opsManagerAPI{BasicHTTPOperation: NewClient(), resolver: resolver}
}

// CreateFirstUser registers the first ever Ops Manager user
// https://docs.opsmanager.mongodb.com/master/reference/api/user-create-first/
func (api opsManagerAPI) CreateFirstUser(user OpsManagerUser, whitelistIP string) (CreateFirstUserResponse, error) {
	var result CreateFirstUserResponse

	bodyBytes, err := json.Marshal(user)
	if err != nil {
		return result, err
	}

	url := api.resolver.Of("/unauth/users?whitelist=%s", whitelistIP)
	resp := api.PostJSON(url, bytes.NewReader(bodyBytes))
	if resp.IsError() {
		return result, resp.err
	}

	if resp.resp != nil && resp.resp.Body != nil {
		defer logError(resp.resp.Body.Close)
	}

	decoder := json.NewDecoder(resp.resp.Body)
	err2 := decoder.Decode(&result)
	panicOnUnrecoverableError(err2)

	return result, nil
}
