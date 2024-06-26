
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2023-09-02-preview/machines` Documentation

The `machines` SDK allows for interaction with the Azure Resource Manager Service `containerservice` (API Version `2023-09-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2023-09-02-preview/machines"
```


### Client Initialization

```go
client := machines.NewMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MachinesClient.Get`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterValue", "agentPoolValue", "machineValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinesClient.List`

```go
ctx := context.TODO()
id := machines.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterValue", "agentPoolValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
