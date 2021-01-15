# README #
This package contains the system management agent client written in the Go programming language.  
The system management agent client is used by Go services or other Go code to communicate with the EdgeX support-agent microservice (regardless of underlying implementation type) by sending REST requests to the service's API endpoints.

### How To Use ###
To use the management agent client package you first need to import the library into your project:
```
import "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/agent"
```
As an example of use, to find the health of a service using the management agent client:
```
ac := NewAgentClient("localhost:48082")
```
And then use the client to get all value descriptors
```
res, err := ac.Health(context.Background(), []string{"edgex-core-data"})
// check err (not shown)
fmt.Printf("health for core-data service: %s\n", res)
```
