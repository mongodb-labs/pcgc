package opsmanager

import (
	"errors"
	"io"
)

// UpdateAutomationConfig
// https://docs.opsmanager.mongodb.com/master/reference/api/automation-config/#update-the-automation-configuration
func (api opsManagerAPI) UpdateAutomationConfig(projectID string, body io.Reader) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}
