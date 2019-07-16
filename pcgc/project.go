package pcgc

import (
	"context"
	"fmt"
)

// Project represents the structure of a project.
type Project struct {
	ID           string  `json:"id,omitempty"`
	OrgID        string  `json:"orgId,omitempty"`
	Name         string  `json:"name,omitempty"`
	ClusterCount int     `json:"clusterCount,omitempty"`
	Created      string  `json:"created,omitempty"`
	Links        []*Link `json:"links,omitempty"`
}

// Projects represents a array of project
type Projects struct {
	Links      []*Link    `json:"links"`
	Results    []*Project `json:"results"`
	TotalCount int        `json:"totalCount"`
}

// Result is part og TeamsAssigned structure
type Result struct {
	Links     []*Link  `json:"links"`
	RoleNames []string `json:"roleNames"`
	TeamID    string   `json:"teamId"`
}

// RoleName represents the kind of user role in your project
type RoleName struct {
	RoleName string `json:"rolesNames"`
}

// Team reperesents the kind of role that has the team
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

//GetAllProjects gets all project.
//See more: https://docs.atlas.mongodb.com/reference/api/project-get-all/
func (s *Service) GetAllProjects(ctx context.Context) (*Projects, error) {

	var projects Projects
	return &projects, s.Get(ctx, &projects, fmt.Sprintf("/api/public/v1.0/groups"), nil)

}
