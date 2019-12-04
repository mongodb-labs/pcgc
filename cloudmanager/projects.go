package cloudmanager

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	projectBasePath = "groups"
)

// ProjectsService is an interface for interfacing with the Projects
// endpoints of the MongoDB Atlas API.
// See more: https://docs.atlas.mongodb.com/reference/api/projects/
type ProjectsService interface {
	GetAllProjects(context.Context) (*Projects, *Response, error)
	GetOneProject(context.Context, string) (*Project, *Response, error)
	GetOneProjectByName(context.Context, string) (*Project, *Response, error)
	Create(context.Context, *Project) (*Project, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

//ProjectsServiceOp handles communication with the Projects related methos of the
//MongoDB Atlas API
type ProjectsServiceOp struct {
	client *Client
}

var _ ProjectsService = &ProjectsServiceOp{}

// HostCount number of processes per project.
type HostCount struct {
	Arbiter   int `json:"arbiter"`
	Config    int `json:"config"`
	Master    int `json:"master"`
	Mongos    int `json:"mongos"`
	Primary   int `json:"primary"`
	Secondary int `json:"secondary"`
	Slave     int `json:"slave"`
}

// Project represents the structure of a project.
type Project struct {
	ActiveAgentCount int                  `json:"activeAgentCount,omitempty"`
	HostCounts       *HostCount           `json:"hostCounts,omitempty"`
	ID               string               `json:"id,omitempty"`
	LastActiveAgent  string               `json:"lastActiveAgent,omitempty"`
	Links            []*mongodbatlas.Link `json:"links,omitempty"`
	Name             string               `json:"name,omitempty"`
	OrgID            string               `json:"orgId,omitempty"`
	PublicAPIEnabled bool                 `json:"publicApiEnabled,omitempty"`
	ReplicaSetCount  int                  `json:"replicaSetCount,omitempty"`
	ShardCount       int                  `json:"shardCount,omitempty"`
	Tags             []*string            `json:"tags,omitempty"`
}

// Projects represents a array of project
type Projects struct {
	Links      []*mongodbatlas.Link `json:"links"`
	Results    []*Project           `json:"results"`
	TotalCount int                  `json:"totalCount"`
}

//GetAllProjects gets all project.
//See more: https://docs.cloudmanager.mongodb.com/reference/api/groups/get-all-groups-for-current-user/
func (s *ProjectsServiceOp) GetAllProjects(ctx context.Context) (*Projects, *Response, error) {

	req, err := s.client.NewRequest(ctx, http.MethodGet, projectBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Projects)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}

//GetOneProject gets a single project.
//See more: https://docs.cloudmanager.mongodb.com/reference/api/groups/get-one-group-by-id/
func (s *ProjectsServiceOp) GetOneProject(ctx context.Context, projectID string) (*Project, *Response, error) {
	if projectID == "" {
		return nil, nil, mongodbatlas.NewArgError("projectID", "must be set")
	}

	path := fmt.Sprintf("%s/%s", projectBasePath, projectID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Project)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//GetOneProjectByName gets a single project by its name.
//See more: https://docs.cloudmanager.mongodb.com/reference/api/groups/get-one-group-by-name/
func (s *ProjectsServiceOp) GetOneProjectByName(ctx context.Context, projectName string) (*Project, *Response, error) {
	if projectName == "" {
		return nil, nil, mongodbatlas.NewArgError("projectName", "must be set")
	}

	path := fmt.Sprintf("%s/byName/%s", projectBasePath, projectName)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Project)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//Create creates a project.
//See more: https://docs.cloudmanager.mongodb.com/reference/api/groups/create-one-group/
func (s *ProjectsServiceOp) Create(ctx context.Context, createRequest *Project) (*Project, *Response, error) {
	if createRequest == nil {
		return nil, nil, mongodbatlas.NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, projectBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(Project)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//Delete deletes a project.
// See more: https://docs.cloudmanager.mongodb.com/reference/api/groups/delete-one-group/
func (s *ProjectsServiceOp) Delete(ctx context.Context, projectID string) (*Response, error) {
	if projectID == "" {
		return nil, mongodbatlas.NewArgError("projectID", "must be set")
	}

	basePath := fmt.Sprintf("%s/%s", projectBasePath, projectID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, basePath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}
