// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudmanager

import (
	"context"
	"fmt"
	"net/http"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	usersBasePath = APIUnauthPath + "users"
)

// UnAuthService is an interface for interfacing with unauthenticated APIs
type UnauthUsersService interface {
	CreateFirstUser(context.Context, *User, string) (*CreateUserResponse, *atlas.Response, error)
}

// UnAuthServiceOp handles communication with the unauthenticated API
type UnauthUsersServiceOp struct {
	client *Client
}

// See more: https://docs.opsmanager.mongodb.com/master/reference/api/user-create-first/
func (s *UnauthUsersServiceOp) CreateFirstUser(ctx context.Context, user *User, whitelistIP string) (*CreateUserResponse, *atlas.Response, error) {
	// if a whitelist was not specified, do not pass the parameter
	var basePath string
	if whitelistIP == "" {
		basePath = usersBasePath
	} else {
		basePath = fmt.Sprintf("%s?whitelist=%s", usersBasePath, whitelistIP)
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, basePath, user)
	if err != nil {
		return nil, nil, err
	}

	root := new(CreateUserResponse)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// UserResponse wrapper for a user response, augmented with a few extra fields
type User struct {
	Username     string        `json:"username"`
	Password     string        `json:"password,omitempty"`
	FirstName    string        `json:"firstName,omitempty"`
	LastName     string        `json:"lastName,omitempty"`
	EmailAddress string        `json:"emailAddress,omitempty"`
	ID           string        `json:"id,omitempty"`
	Links        []*atlas.Link `json:"links,omitempty"`
	Roles        []*UserRole   `json:"roles,omitempty"`
}

// UserRole denotes a single user role
type UserRole struct {
	RoleName string `json:"roleName"`
	GroupID  string `json:"groupId,omitempty"`
	OrgID    string `json:"orgId,omitempty"`
}

// CreateUserResponse API response for the CreateFirstUser() call
type CreateUserResponse struct {
	APIKey string `json:"apiKey"`
	User   *User  `json:"user"`
}
