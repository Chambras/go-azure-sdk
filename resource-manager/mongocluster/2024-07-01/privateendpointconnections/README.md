
## `github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with Azure Resource Manager `mongocluster` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.Create`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterName", "privateEndpointConnectionName")

payload := privateendpointconnections.PrivateEndpointConnectionResource{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.Delete`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterName", "privateEndpointConnectionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.Get`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterName", "privateEndpointConnectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ListByMongoCluster`

```go
ctx := context.TODO()
id := privateendpointconnections.NewMongoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mongoClusterName")

// alternatively `client.ListByMongoCluster(ctx, id)` can be used to do batched pagination
items, err := client.ListByMongoClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
