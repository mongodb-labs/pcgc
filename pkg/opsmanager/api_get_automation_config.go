package opsmanager

import (
	"encoding/json"
	"github.com/mongodb-labs/pcgc/pkg/httpclient"
	"github.com/mongodb-labs/pcgc/pkg/useful"
)

// VersionHostnamePair represents a pair of version name and hostname strings
type VersionHostnamePair struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

// Cluster represents a cluster definition within an automation config object
// NOTE: this struct is mutable
type Cluster struct {
	Auth               map[string]interface{}   `json:"auth,omitempty"`
	LDAP               map[string]interface{}   `json:"ldap,omitempty"`
	Processes          []map[string]interface{} `json:"processes,omitempty"`
	ReplicaSets        []map[string]interface{} `json:"replicaSets,omitempty"`
	Roles              map[string]interface{}   `json:"roles,omitempty"`
	MonitoringVersions []*VersionHostnamePair   `json:"monitoringVersions,omitempty"`
	BackupVersions     []*VersionHostnamePair   `json:"backupVersions,omitempty"`
	MongoSQLDs         []map[string]interface{} `json:"mongosqlds,omitempty"`
	AgentVersion       map[string]interface{}   `json:"agentVersion,omitempty"`
	Balancer           map[string]interface{}   `json:"balancer,omitempty"`
	CPSModules         []map[string]interface{} `json:"cpsModules,omitempty"`
	IndexConfigs       []map[string]interface{} `json:"indexConfigs,omitempty"`
	Kerberos           map[string]interface{}   `json:"kerberos,omitempty"`
	MongoTs            map[string]interface{}   `json:"mongots,omitempty"`
	Options            map[string]interface{}   `json:"options,omitempty"`
	SSL                map[string]interface{}   `json:"ssl,omitempty"`
	Version            string                   `json:"version,omitempty"`
	Sharding           []map[string]interface{} `json:"sharding,omitempty"`
}

// AutomationConfig represents a full Ops Manager Automation Config object
// NOTE: this struct is mutable
type AutomationConfig struct {
	IDCounter                        int                    `json:"idCounter,omitempty"`
	State                            string                 `json:"state,omitempty"`
	Version                          int                    `json:"version,omitempty"`
	GroupID                          string                 `json:"groupId,omitempty"`
	UserID                           string                 `json:"userId,omitempty"`
	PublishTimestamp                 int                    `json:"publishTimestamp,omitempty"`
	PublishedVersion                 string                 `json:"publishedVersion,omitempty"`
	SaveTimestamp                    int                    `json:"saveTimestamp,omitempty"`
	LatestAutomationAgentVersionName string                 `json:"latestAutomationAgentVersionName,omitempty"`
	LatestMonitoringAgentVersionName string                 `json:"latestMonitoringAgentVersionName,omitempty"`
	LatestBackupAgentVersionName     string                 `json:"latestBackupAgentVersionName,omitempty"`
	LatestBiConnectorVersionName     string                 `json:"latestBiConnectorVersionName,omitempty"`
	Cluster                          *Cluster               `json:"cluster,omitempty"`
	VersionConfig                    map[string]interface{} `json:"versionConfig,omitempty"`
	LogRotate                        map[string]interface{} `json:"logRotate,omitempty"`
	MonitoringAgentTemplate          map[string]interface{} `json:"monitoringAgentTemplate,omitempty"`
	BackupAgentTemplate              map[string]interface{} `json:"backupAgentTemplate,omitempty"`
	CPSModuleTemplate                map[string]interface{} `json:"cpsModuleTemplate,omitempty"`
	DeploymentJobStatuses            []interface{}          `json:"deploymentJobStatuses,omitempty"`
}

// GetAutomationConfig
// https://docs.opsmanager.mongodb.com/master/reference/api/automation-config/#get-the-automation-configuration
func (client opsManagerClient) GetAutomationConfig(projectID string) (AutomationConfig, error) {
	var result AutomationConfig

	url := client.resolver.Of("/groups/%s/automationConfig", projectID)
	resp := client.GetJSON(url)
	if resp.IsError() {
		return result, resp.Err
	}
	defer httpclient.CloseResponseBodyIfNotNil(resp)

	decoder := json.NewDecoder(resp.Response.Body)
	err := decoder.Decode(&result)
	useful.PanicOnUnrecoverableError(err)

	return result, nil
}
