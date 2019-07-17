package opsmanager

import "errors"

// GetAllHostsInProject
// https://docs.opsmanager.mongodb.com/current/reference/api/hosts/get-all-hosts-in-group/
func (api opsManagerAPI) GetAllHostsInProject(projectID string, pageNum int, itemsPerPage int) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}
