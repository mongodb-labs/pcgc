package opsmanager

import "errors"

// CreateAgentAPIKEY creates a new Agent API key in the specified project
// https://docs.opsmanager.mongodb.com/master/reference/api/agentapikeys/create-one-agent-api-key/
func (client opsManagerClient) CreateAgentAPIKEY(projectID string, name string) (interface{}, error) {
	// TODO(mihaibojin): Implement
	return nil, errors.New("not implemented yet")
}
