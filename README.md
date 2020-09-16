# go-mod-core-contracts
This module contains the contract models used to describe data as it is passed via Request/Response between various
EdgeX Foundry services. It also contains service clients for each service within the
[edgex-go](https://github.com/edgexfoundry/edgex-go) repository. The definition of the various models and clients can
be found in their respective top-level directories.

The default encoding for the models is JSON, although in at least one case --
[DeviceProfile](https://github.com/edgexfoundry/go-mod-core-contracts/blob/master/models/deviceprofile.go) --
YAML encoding is also supported since a device profile is defined as a YAML document.

### Installation ###
* Make sure you're using at least Go 1.11.1 (EdgeX currently uses Go 1.15.x) and have modules enabled, i.e. have an initialized  go.mod file 
* If your code is in your GOPATH then make sure ```GO111MODULE=on``` is set
* Run ```go get github.com/edgexfoundry/go-mod-core-contracts```
    * This will add the go-mod-core-contracts to the go.mod file and download it into the module cache
