package opsmanager

import (
	"bytes"
	"encoding/json"
	"github.com/mongodb-labs/pcgc/pkg/httpclient"
	"github.com/mongodb-labs/pcgc/pkg/useful"
)

// LDAPGroupMapping holds a single mapping of role to LDAP groups
type LDAPGroupMapping struct {
	RoleName   string   `json:"roleName,omitempty"`
	LdapGroups []string `json:"ldapGroups,omitempty"`
}

// CreateOneProjectResponse API response for the CreateOneProject() call
type CreateOneProjectResponse struct {
	ID                string             `json:"id"`
	Name              string             `json:"name,omitempty"`
	AgentAPIKey       string             `json:"agentApiKey,omitempty"`
	ActiveAgentCount  int                `json:"activeAgentCount,omitempty"`
	HostCounts        map[string]int     `json:"hostCounts,omitempty"`
	LastActiveAgent   string             `json:"lastActiveAgent,omitempty"`
	LDAPGroupMappings []LDAPGroupMapping `json:"ldapGroupMappings,omitempty"`
	Links             []Link             `json:"links,omitempty"`
	OrgID             string             `json:"orgId,omitempty"`
	PublicAPIEnabled  bool               `json:"publicApiEnabled,omitempty"`
	ReplicaSetCount   int                `json:"replicaSetCount,omitempty"`
	ShardCount        int                `json:"shardCount,omitempty"`
	Tags              []string           `json:"tags,omitempty"`
}

// CreateOneProject
// https://docs.opsmanager.mongodb.com/master/reference/api/groups/create-one-group/
func (client opsManagerClient) CreateOneProject(name string, orgID string) (CreateOneProjectResponse, error) {
	var result CreateOneProjectResponse

	// create request object
	request := make(map[string]string)
	request["name"] = name
	if orgID != "" {
		// optional parameter
		request["orgId"] = orgID
	}

	bodyBytes, err := json.Marshal(request)
	if err != nil {
		return result, err
	}

	url := client.resolver.Of("/groups")
	resp := client.PostJSON(url, bytes.NewReader(bodyBytes))
	if resp.IsError() {
		return result, resp.Err
	}
	defer httpclient.CloseResponseBodyIfNotNil(resp)

	decoder := json.NewDecoder(resp.Response.Body)
	err = decoder.Decode(&result)
	useful.PanicOnUnrecoverableError(err)

	return result, nil
}
