
## `github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-15-preview/deployments` Documentation

The `deployments` SDK allows for interaction with the Azure Resource Manager Service `apicenter` (API Version `2024-03-15-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-15-preview/deployments"
```


### Client Initialization

```go
client := deployments.NewDeploymentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := deployments.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceValue", "apiValue", "deploymentValue")

payload := deployments.Deployment{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.Delete`

```go
ctx := context.TODO()
id := deployments.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceValue", "apiValue", "deploymentValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.Get`

```go
ctx := context.TODO()
id := deployments.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceValue", "apiValue", "deploymentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.Head`

```go
ctx := context.TODO()
id := deployments.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceValue", "apiValue", "deploymentValue")

read, err := client.Head(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.List`

```go
ctx := context.TODO()
id := deployments.NewApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceValue", "apiValue")

// alternatively `client.List(ctx, id, deployments.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, deployments.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
