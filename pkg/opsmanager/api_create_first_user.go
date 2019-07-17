package opsmanager

import (
	"bytes"
	"encoding/json"
	"github.com/mongodb-labs/pcgc/pkg/httpclient"
	"github.com/mongodb-labs/pcgc/pkg/useful"
)

// WhitelistAllowAll allows API access from any IPv4 address
const WhitelistAllowAll = "0.0.0.1/0"

// User request object which identifies a user
type User struct {
	Username     string `json:"username"`
	Password     string `json:"password,omitempty"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

// UserRole denotes a single user role
type UserRole struct {
	RoleName string `json:"roleName"`
	GroupID  string `json:"groupId,omitempty"`
	OrgID    string `json:"orgId,omitempty"`
}

// UserResponse wrapper for a user response, augmented with a few extra fields
type UserResponse struct {
	User

	ID    string     `json:"id"`
	Links []Link     `json:"links,omitempty"`
	Roles []UserRole `json:"roles,omitempty"`
}

// CreateFirstUserResponse API response for the CreateFirstUser() call
type CreateFirstUserResponse struct {
	APIKey string       `json:"apiKey"`
	User   UserResponse `json:"user"`
}

// CreateFirstUser registers the first ever Ops Manager user (global owner)
// https://docs.opsmanager.mongodb.com/master/reference/api/user-create-first/
func (api opsManagerClient) CreateFirstUser(user User, whitelistIP string) (CreateFirstUserResponse, error) {
	var result CreateFirstUserResponse

	bodyBytes, err := json.Marshal(user)
	if err != nil {
		return result, err
	}

	url := api.resolver.Of("/unauth/users?whitelist=%s", whitelistIP)
	resp := api.PostJSON(url, bytes.NewReader(bodyBytes))
	if resp.IsError() {
		return result, resp.Err
	}
	defer httpclient.CloseResponseBodyIfNotNil(resp)

	decoder := json.NewDecoder(resp.Response.Body)
	err2 := decoder.Decode(&result)
	useful.PanicOnUnrecoverableError(err2)

	return result, nil
}
