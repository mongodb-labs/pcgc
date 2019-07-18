MongoDB Private Cloud Golang Client
===================================

An HTTP client for [Ops Manager](https://docs.opsmanager.mongodb.com/master/reference/api/) 
and [Cloud Manager](https://docs.cloudmanager.mongodb.com/reference/api/) Public API endpoints.

**This project is currently in development and is not yet ready for production use.**

This library is licensed under the terms of the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0).


### Desired feature set for release v0.1.0

- [x] Register the first user
- [x] Implement digest authentication (not available in Go's http client) 
- [x] Create a project
- [x] Create an agent API key
- [x] Retrieve the automation config
- [x] Get all agents in a Project by type
- [x] Patch the automation config: update Deployments
- [ ] Merge an existing automation config with new changes (e.g. `Process`)
- [ ] Wait for goal state
- [ ] Enable monitoring: edit `AutomationCluster` and enable monitoring (add a `VersionHostnamePair`)
  ```json
    {
        "monitoringVersions": [{
            "name": "7.2.0.488-1",
            "hostname": "hostname"
        }]
    }
  ```
- [ ] Deploy a standalone (insert a new `Process` into `AutomationCluster`)
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

### Installing the CLI tool (mpc)

```
make install # mpc will be deployed into $GOPATH/bin

# Run mpc to see its usage
mpc
```
