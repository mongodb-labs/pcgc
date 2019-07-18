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

// Build a MongoDB build
type Build struct {
	Architecture       string   `json:"architecture"`
	Bits               int      `json:"bits"`
	Flavor             string   `json:"flavor,omitempty"`
	GitVersion         string   `json:"gitVersion,omitempty"`
	MaxOsVersion       string   `json:"maxOsVersion,omitempty"`
	MinOsVersion       string   `json:"minOsVersion,omitempty"`
	Platform           string   `json:"platform,omitempty"`
	URL                string   `json:"url,omitempty"`
	Modules            []string `json:"modules,omitempty"`
	Win2008plus        bool     `json:"win2008plus,omitempty"`
	WinVCRedistDll     string   `json:"winVCRedistDll,omitempty"`
	WinVCRedistOptions []string `json:"winVCRedistOptions,omitempty"`
	WinVCRedistURL     string   `json:"winVCRedistUrl,omitempty"`
	WinVCRedistVersion string   `json:"winVCRedistVersion,omitempty"`
}

// MongoDBVersion ways to install MongoDB
type MongoDBVersion struct {
	Name   string  `json:"name,omitempty"`
	Builds []Build `json:"builds,omitempty"`
}

// AutomationConfigResponse represents a cluster definition within an automation config object
// NOTE: this struct is mutable
type AutomationConfig struct {
	Auth               map[string]interface{}   `json:"auth,omitempty"`
	LDAP               map[string]interface{}   `json:"ldap,omitempty"`
	Processes          []*Process               `json:"processes,omitempty"`
	ReplicaSets        []map[string]interface{} `json:"replicaSets,omitempty"`
	Roles              []map[string]interface{} `json:"roles,omitempty"`
	MonitoringVersions []*VersionHostnamePair   `json:"monitoringVersions,omitempty"`
	BackupVersions     []*VersionHostnamePair   `json:"backupVersions,omitempty"`
	MongoSQLDs         []map[string]interface{} `json:"mongosqlds,omitempty"`
	MongoDBVersions    []*MongoDBVersion        `json:"mongoDbVersions,omitempty"`
	AgentVersion       map[string]interface{}   `json:"agentVersion,omitempty"`
	Balancer           map[string]interface{}   `json:"balancer,omitempty"`
	CPSModules         []map[string]interface{} `json:"cpsModules,omitempty"`
	IndexConfigs       []map[string]interface{} `json:"indexConfigs,omitempty"`
	Kerberos           map[string]interface{}   `json:"kerberos,omitempty"`
	MongoTs            []map[string]interface{} `json:"mongots,omitempty"`
	Options            map[string]interface{}   `json:"options,omitempty"`
	SSL                map[string]interface{}   `json:"ssl,omitempty"`
	Version            int                      `json:"version,omitempty"`
	Sharding           []map[string]interface{} `json:"sharding,omitempty"`
	UIBaseURL          string                   `json:"uiBaseUrl,omitempty"`
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
