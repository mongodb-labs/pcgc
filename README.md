MongoDB Private Cloud Golang Client
===================================

An HTTP client for [Ops Manager](https://docs.opsmanager.mongodb.com/master/reference/api/) 
and [Cloud Manager](https://docs.cloudmanager.mongodb.com/reference/api/) Public API endpoints.

**This project is currently in development and is not yet ready for production use.**

This library is licensed under the terms of the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0).


### Desired feature set for release v0.1.0

- [x] Register the first user
  - send the IP whitelist
  - return: `publicApiKey`

- [x] Implement digest authentication (not available in Go's http client) 

- [ ] Create a project
  - https://docs.opsmanager.mongodb.com/master/reference/api/groups/create-one-group/
  - return: `groupId`

- [ ] Create an agent API key
  https://docs.opsmanager.mongodb.com/master/reference/api/agentapikeys/create-one-agent-api-key/
  - return: `mmsApiKey`

- [ ] Retrieve the monitoring version from the automation config
  https://docs.opsmanager.mongodb.com/master/reference/api/automation-config/#get-the-automation-configuration
  `{"automationAgentVersion":"5.8.0.5546-1","monitoringAgentVersion":"6.7.0.466-1","biConnectorVersion":"2.6.1","backupAgentVersion":"7.1.0.1011-1"}`

- [ ] Get all hosts in a Project
  https://docs.opsmanager.mongodb.com/current/reference/api/hosts/get-all-hosts-in-group/

- [ ] Update the automation config: enable monitoring
  https://docs.opsmanager.mongodb.com/master/reference/api/automation-config/#update-the-automation-configuration
  ```json
    {
        "monitoringVersions": [{
            "name": "7.2.0.488-1",
            "hostname": "hostname"
        }]
    }
  ```

- [ ] Patch the automation config: deploy a standalone (insert an entry into `processes`)
  ```json
    {
          "cluster": {
              "processes": [
                  {
                      "name": "hostname-27017_1",
                      "processType": "mongod",
                      "version": "4.0.10",
                      "authSchemaVersion": 5,
                      "featureCompatibilityVersion": "4.0",
                      "disabled": false,
                      "manualMode": false,
                      "hostname": "hostname",
                      "args2_6": {
                          "net": {
                              "port": 27017
                          },
                          "storage": {
                              "dbPath": "/data"
                          },
                          "systemLog": {
                              "destination": "file",
                              "path": "/data/mongodb.log"
                          }
                      },
                      "logRotate": {
                          "sizeThresholdMB": 1000,
                          "timeThresholdHrs": 24
                      }
                  }
              ],
          "state": "DRAFT"
      }
    }
  ```

- [ ] Publish the automation config and wait for goal state

### Feature Backlog

- [ ] TBD


### Setting up the development environment

Pull requests are always welcome! Please read our [contributor guide](./CONTRIB.md) before starting any work.  

The steps below should help you get started.  They have been tested on MacOS, but should work on Linux systems as well (with minor adaptations.)

1. Install GO (1.12+)
```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

Ensure `$GOROOT/bin` is in your path.

2. Install the git hooks, to automatically fix linting issues and flag any errors 

`make link-git-hooks`
