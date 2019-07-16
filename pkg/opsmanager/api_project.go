package opsmanager

import (
	"encoding/json"
	"github.com/mongodb-labs/pcgc/pkg/httpclient"
	"github.com/mongodb-labs/pcgc/pkg/useful"
)

// Project represents the structure of a project
type Project struct {
	ID           string `json:"id"`
	OrgID        string `json:"orgId"`
	Name         string `json:"name"`
	ClusterCount int    `json:"clusterCount,omitempty"`
	Created      string `json:"created,omitempty"`
	Links        []Link `json:"links,omitempty"`
}

// Projects represents a array of project
type Projects struct {
	Links      []Link    `json:"links"`
	Results    []Project `json:"results"`
	TotalCount int       `json:"totalCount"`
}

// Result is part of TeamsAssigned structure
type Result struct {
	Links     []Link   `json:"links"`
	RoleNames []string `json:"roleNames"`
	TeamID    string   `json:"teamId"`
}

// RoleName represents the kind of user role in your project
type RoleName struct {
	RoleName string `json:"rolesNames"`
}

// Team represents roles assigned to the team
type Team struct {
	TeamID string      `json:"teamId"`
	Roles  []*RoleName `json:"roles"`
}

// TeamsAssigned represents the one team assigned to the project.
type TeamsAssigned struct {
	Links      []*Link   `json:"links"`
	Results    []*Result `json:"results"`
	TotalCount int       `json:"totalCount"`
}

// GetAllProjects registers the first ever Ops Manager user (global owner)
// https://docs.opsmanager.mongodb.com/master/reference/api/groups/get-all-groups-for-current-user/
func (api opsManagerAPI) GetAllProjects() (Projects, error) {
	var result Projects

	url := api.resolver.Of("/groups")
	resp := api.Get(url)
	if resp.IsError() {
		return result, resp.Err
	}
	defer httpclient.CloseResponseBodyIfNotNil(resp)

	decoder := json.NewDecoder(resp.Response.Body)
	err2 := decoder.Decode(&result)
	useful.PanicOnUnrecoverableError(err2)

	return result, nil
}
