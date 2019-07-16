MongoDB Private Cloud Golang Client
===================================

An HTTP client for [Ops Manager](https://docs.opsmanager.mongodb.com/master/reference/api/) 
and [Cloud Manager](https://docs.cloudmanager.mongodb.com/reference/api/) Public API endpoints.

**This project is currently in development and is not yet ready for production use.**

This library is licensed under the terms of the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0).


### Desired feature set for release v0.1.0

- [ ] Register the first user
  - send the IP whitelist
  - return: `groupId`, `publicApiKey`

- [ ] Implement digest authentication (not available in Go's http client), see: 
  - https://godoc.org/github.com/bobziuchkovski/digest
  - https://en.wikipedia.org/wiki/Digest_access_authentication

- [ ] Create an agent API key
  - return: `mmsApiKey`

- [ ] Retrieve the monitoring version from the automation config

- [ ] Patch the automation config: enable monitoring
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
