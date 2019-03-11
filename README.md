# go-mod-core-contracts
This module contains the contract models used to describe data as it is passed via Request/Response between various EdgeX Foundry services. It also contains service clients for each service within the [edgex-go](https://github.com/edgexfoundry/edgex-go) repository. The definition of the various models and clients can be found in their respective top-level directories.

The default encoding for the models is JSON, although in at least one case -- [DeviceProfile](https://github.com/edgexfoundry/go-mod-core-contracts/blob/master/models/deviceprofile.go) -- YAML encoding is also supported since a device profile is defined as a YAML document.

### Installation ###
* Make sure you're using at least Go 1.11.1 and have modules enabled, i.e. have an initialized  go.mod file 
* If your code is in your GOPATH then make sure ```GO111MODULE=on``` is set
* Run ```go get github.com/edgexfoundry/go-mod-core-contracts```
    * This will add the go-mod-core-contracts to the go.mod file and download it into the module cache

### How to Use ###
In order to instantiate a service client, you would do the following:

Let's say you want to utilize a client for the core-metadata service, the first thing you want to do is initialize an instance of [types.EndpointParams](https://github.com/edgexfoundry/go-mod-core-contracts/blob/master/clients/types/endpoint_params.go). This simple struct provides all the properties you need to address a given service endpoint. Population of the relevant properties might look like this:

```
params := types.EndpointParams{
        ServiceKey:  internal.CoreMetaDataServiceKey,
	Path:        clients.ApiDeviceRoute,
	UseRegistry: useRegistry,
	Url:         Configuration.Clients["Metadata"].Url() + clients.ApiDeviceRoute,
	Interval:    Configuration.Service.ClientMonitor,
}
```
From there you simply pass the EndpointParams into the initialization function for the client you wish to use. In the above case, we're trying to initialize a client for the Device endpoint of the core-metadata service.
```
mdc := metadata.NewDeviceClient(params, startup.Endpoint{RegistryClient: &registryClient})
```
_More information on the `RegistryClient` can be found [here](https://github.com/edgexfoundry/go-mod-registry). The `RegistryClient` is only used if the `useRegistry` flag provided to the `EndpointParams` is true._ 

Once you have a reference to the client, you simply need to call its methods like so:
`_, err := mdc.CheckForDevice(device, ctx)`

Each client has a `Monitor` goroutine in it. If the registry is being used, the Monitor's job is to refresh the protocol, host and port of your service client at some configured interval. The default interval is 15 seconds.
