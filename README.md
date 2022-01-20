# go-mod-core-contracts
[![Build Status](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/go-mod-core-contracts/job/main/badge/icon)](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/go-mod-core-contracts/job/main/) [![Code Coverage](https://codecov.io/gh/edgexfoundry/go-mod-core-contracts/branch/main/graph/badge.svg?token=s4Y4L22Bs0)](https://codecov.io/gh/edgexfoundry/go-mod-core-contracts) [![Go Report Card](https://goreportcard.com/badge/github.com/edgexfoundry/go-mod-core-contracts)](https://goreportcard.com/report/github.com/edgexfoundry/go-mod-core-contracts) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/edgexfoundry/go-mod-core-contracts?include_prereleases&sort=semver&label=latest-dev)](https://github.com/edgexfoundry/go-mod-core-contracts/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/edgexfoundry/go-mod-core-contracts?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/edgexfoundry/go-mod-core-contracts)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edgexfoundry/go-mod-core-contracts) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/edgexfoundry/go-mod-core-contracts)](https://github.com/edgexfoundry/go-mod-core-contracts/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/edgexfoundry/go-mod-core-contracts)](https://github.com/edgexfoundry/go-mod-core-contracts/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/edgexfoundry/teams/go-mod-core-contracts-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/edgexfoundry/go-mod-core-contracts)](https://github.com/edgexfoundry/go-mod-core-contracts/commits)

This module contains the contract models used to describe data as it is passed via Request/Response between various
EdgeX Foundry services. It also contains service clients for each service within the
[edgex-go](https://github.com/edgexfoundry/edgex-go) repository. The definition of the various models and clients can
be found in their respective top-level directories.

The default encoding for the models is JSON, although in at least one case --
[DeviceProfile](https://github.com/edgexfoundry/go-mod-core-contracts/blob/master/models/deviceprofile.go) --
YAML encoding is also supported since a device profile is defined as a YAML document.

### Installation ###
* Make sure you're using at least Go 1.11.1 (EdgeX currently uses Go 1.17.x) and have modules enabled, i.e. have an initialized  go.mod file 
* If your code is in your GOPATH then make sure ```GO111MODULE=on``` is set
* Run ```go get github.com/edgexfoundry/go-mod-core-contracts```
    * This will add the go-mod-core-contracts to the go.mod file and download it into the module cache
