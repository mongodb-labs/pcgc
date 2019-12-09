package cloudmanager

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestProject_GetAllProjects(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
			"links": [{
				"href": "https://cloud.mongodb.com/api/public/v1.0/groups",
				"rel": "self"
			}],
			"results": [{
				"activeAgentCount": 0,
				"hostCounts": {
					"arbiter": 0,	
					"config": 0,
					"master": 0,
					"mongos": 0,
					"primary": 0,
					"secondary": 0,
					"slave": 0
				},
				"id": "56a10a80e4b0fd3b9a9bb0c2",
				"lastActiveAgent": "2016-03-09T18:19:37Z",
				"links": [{
					"href": "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
					"rel": "self"
				}],
				"name": "012i3091203jioawjioej",
				"orgId": "5980cfdf0b6d97029d82f86e",
				"publicApiEnabled": true,
				"replicaSetCount": 0,
				"shardCount": 0,
				"tags": []
			}, {
				"activeAgentCount": 0,
				"hostCounts": {
					"arbiter": 0,
					"config": 0,
					"master": 0,
					"mongos": 0,
					"primary": 0,
					"secondary": 0,
					"slave": 0
				},
				"id": "56aa691ce4b0a0e8c4be51f7",
				"lastActiveAgent": "2016-01-29T19:02:56Z",
				"links": [{
					"href": "https://cloud.mongodb.com/api/public/v1.0/groups/56aa691ce4b0a0e8c4be51f7",
					"rel": "self"
				}],
				"name": "1454008603036",
				"orgId": "5980d0040b6d97029d831798",
				"publicApiEnabled": true,
				"replicaSetCount": 0,
				"shardCount": 0,
				"tags": []
			}],
			"totalCount": 2
		}`)
	})

	projects, _, err := client.Projects.GetAllProjects(ctx)
	if err != nil {
		t.Errorf("Projects.GetAllProjects returned error: %v", err)
	}

	expected := &Projects{
		Links: []*mongodbatlas.Link{
			{
				Href: "https://cloud.mongodb.com/api/public/v1.0/groups",
				Rel:  "self",
			},
		},
		Results: []*Project{
			{
				ActiveAgentCount: 0,
				HostCounts: &HostCount{
					Arbiter:   0,
					Config:    0,
					Master:    0,
					Mongos:    0,
					Primary:   0,
					Secondary: 0,
					Slave:     0,
				},
				ID:              "56a10a80e4b0fd3b9a9bb0c2",
				LastActiveAgent: "2016-03-09T18:19:37Z",
				Links: []*mongodbatlas.Link{
					{
						Href: "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
						Rel:  "self",
					},
				},
				Name:             "012i3091203jioawjioej",
				OrgID:            "5980cfdf0b6d97029d82f86e",
				PublicAPIEnabled: true,
				ReplicaSetCount:  0,
				ShardCount:       0,
				Tags:             []*string{},
			},
			{
				ActiveAgentCount: 0,
				HostCounts: &HostCount{
					Arbiter:   0,
					Config:    0,
					Master:    0,
					Mongos:    0,
					Primary:   0,
					Secondary: 0,
					Slave:     0,
				},
				ID:              "56aa691ce4b0a0e8c4be51f7",
				LastActiveAgent: "2016-01-29T19:02:56Z",
				Links: []*mongodbatlas.Link{
					{
						Href: "https://cloud.mongodb.com/api/public/v1.0/groups/56aa691ce4b0a0e8c4be51f7",
						Rel:  "self",
					},
				},
				Name:             "1454008603036",
				OrgID:            "5980d0040b6d97029d831798",
				PublicAPIEnabled: true,
				ReplicaSetCount:  0,
				ShardCount:       0,
				Tags:             []*string{},
			},
		},
		TotalCount: 2,
	}

	if !reflect.DeepEqual(projects, expected) {
		t.Errorf("Projects.GetAllProjects\n got=%#v\nwant=%#v", projects, expected)
	}
}

func TestProject_GetOneProject(t *testing.T) {
	setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"

	mux.HandleFunc(fmt.Sprintf("/%s/%s", projectBasePath, projectID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
		"activeAgentCount": 0,
		"hostCounts": {
			"arbiter": 0,
			"config": 0,
			"master": 0,
			"mongos": 0,
			"primary": 0,
			"secondary": 0,
			"slave": 0
		},
		"id": "56a10a80e4b0fd3b9a9bb0c2",
		"lastActiveAgent": "2016-03-09T18:19:37Z",
		"links": [{
			"href": "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
			"rel": "self"
		}],
		"name": "012i3091203jioawjioej",
		"orgId": "5980cfdf0b6d97029d82f86e",
		"publicApiEnabled": true,
		"replicaSetCount": 0,
		"shardCount": 0,
		"tags": []
	  }`)
	})

	projectResponse, _, err := client.Projects.GetOneProject(ctx, projectID)
	if err != nil {
		t.Errorf("Projects.GetOneProject returned error: %v", err)
	}

	expected := &Project{
		ActiveAgentCount: 0,
		HostCounts: &HostCount{
			Arbiter:   0,
			Config:    0,
			Master:    0,
			Mongos:    0,
			Primary:   0,
			Secondary: 0,
			Slave:     0,
		},
		ID:              "56a10a80e4b0fd3b9a9bb0c2",
		LastActiveAgent: "2016-03-09T18:19:37Z",
		Links: []*mongodbatlas.Link{
			{
				Href: "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
				Rel:  "self",
			},
		},
		Name:             "012i3091203jioawjioej",
		OrgID:            "5980cfdf0b6d97029d82f86e",
		PublicAPIEnabled: true,
		ReplicaSetCount:  0,
		ShardCount:       0,
		Tags:             []*string{},
	}

	if !reflect.DeepEqual(projectResponse, expected) {
		t.Errorf("Projects.GetOneProject\n got=%#v\nwant=%#v", projectResponse, expected)
	}
}

func TestProject_GetOneProjectByName(t *testing.T) {
	setup()
	defer teardown()

	projectName := "012i3091203jioawjioej"

	mux.HandleFunc(fmt.Sprintf("/%s/byName/%s", projectBasePath, projectName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
			"activeAgentCount": 0,
			"hostCounts": {
				"arbiter": 0,
				"config": 0,
				"master": 0,
				"mongos": 0,
				"primary": 0,
				"secondary": 0,
				"slave": 0
			},
			"id": "56a10a80e4b0fd3b9a9bb0c2",
			"lastActiveAgent": "2016-03-09T18:19:37Z",
			"links": [{
				"href": "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
				"rel": "self"
			}],
			"name": "012i3091203jioawjioej",
			"orgId": "5980cfdf0b6d97029d82f86e",
			"publicApiEnabled": true,
			"replicaSetCount": 0,
			"shardCount": 0,
			"tags": []
		}`)
	})

	projectResponse, _, err := client.Projects.GetOneProjectByName(ctx, projectName)
	if err != nil {
		t.Errorf("Projects.GetOneProject returned error: %v", err)
	}

	expected := &Project{
		ActiveAgentCount: 0,
		HostCounts: &HostCount{
			Arbiter:   0,
			Config:    0,
			Master:    0,
			Mongos:    0,
			Primary:   0,
			Secondary: 0,
			Slave:     0,
		},
		ID:              "56a10a80e4b0fd3b9a9bb0c2",
		LastActiveAgent: "2016-03-09T18:19:37Z",
		Links: []*mongodbatlas.Link{
			{
				Href: "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
				Rel:  "self",
			},
		},
		Name:             "012i3091203jioawjioej",
		OrgID:            "5980cfdf0b6d97029d82f86e",
		PublicAPIEnabled: true,
		ReplicaSetCount:  0,
		ShardCount:       0,
		Tags:             []*string{},
	}

	if diff := deep.Equal(projectResponse, expected); diff != nil {
		t.Error(diff)
	}
	if !reflect.DeepEqual(projectResponse, expected) {
		t.Errorf("Projects.GetOneProject\n got=%#v\nwant=%#v", projectResponse, expected)
	}
}

func TestProject_Create(t *testing.T) {
	setup()
	defer teardown()

	createRequest := &Project{
		OrgID: "5a0a1e7e0f2912c554080adc",
		Name:  "ProjectFoobar",
	}

	mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, `{
			"activeAgentCount": 0,
			"hostCounts": {
				"arbiter": 0,
				"config": 0,
				"master": 0,
				"mongos": 0,
				"primary": 0,
				"secondary": 0,
				"slave": 0
			},
			"id": "56a10a80e4b0fd3b9a9bb0c2",
			"lastActiveAgent": "2016-03-09T18:19:37Z",
			"links": [{
				"href": "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
				"rel": "self"
			}],
			"name": "ProjectFoobar",
			"orgId": "5a0a1e7e0f2912c554080adc",
			"publicApiEnabled": true,
			"replicaSetCount": 0,
			"shardCount": 0,
			"tags": []
		}`)
	})

	project, _, err := client.Projects.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("Projects.Create returned error: %v", err)
	}

	expected := &Project{
		ActiveAgentCount: 0,
		HostCounts: &HostCount{
			Arbiter:   0,
			Config:    0,
			Master:    0,
			Mongos:    0,
			Primary:   0,
			Secondary: 0,
			Slave:     0,
		},
		ID:              "56a10a80e4b0fd3b9a9bb0c2",
		LastActiveAgent: "2016-03-09T18:19:37Z",
		Links: []*mongodbatlas.Link{
			{
				Href: "https://cloud.mongodb.com/api/public/v1.0/groups/56a10a80e4b0fd3b9a9bb0c2",
				Rel:  "self",
			},
		},
		Name:             "ProjectFoobar",
		OrgID:            "5a0a1e7e0f2912c554080adc",
		PublicAPIEnabled: true,
		ReplicaSetCount:  0,
		ShardCount:       0,
		Tags:             []*string{},
	}

	if diff := deep.Equal(project, expected); diff != nil {
		t.Error(diff)
	}
	if !reflect.DeepEqual(project, expected) {
		t.Errorf("DatabaseUsers.Get\n got=%#v\nwant=%#v", project, expected)
	}
}

func TestProject_Delete(t *testing.T) {
	setup()
	defer teardown()

	projectID := "5a0a1e7e0f2912c554080adc"

	mux.HandleFunc(fmt.Sprintf("/groups/%s", projectID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Projects.Delete(ctx, projectID)
	if err != nil {
		t.Errorf("Projects.Delete returned error: %v", err)
	}
}
