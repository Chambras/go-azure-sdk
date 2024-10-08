
## `github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2023-12-30/azureadadministrators` Documentation

The `azureadadministrators` SDK allows for interaction with Azure Resource Manager `mysql` (API Version `2023-12-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2023-12-30/azureadadministrators"
```


### Client Initialization

```go
client := azureadadministrators.NewAzureADAdministratorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AzureADAdministratorsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := azureadadministrators.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

payload := azureadadministrators.AzureADAdministrator{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AzureADAdministratorsClient.Delete`

```go
ctx := context.TODO()
id := azureadadministrators.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AzureADAdministratorsClient.Get`

```go
ctx := context.TODO()
id := azureadadministrators.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AzureADAdministratorsClient.ListByServer`

```go
ctx := context.TODO()
id := azureadadministrators.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
