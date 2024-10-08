
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/administrators` Documentation

The `administrators` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2023-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/administrators"
```


### Client Initialization

```go
client := administrators.NewAdministratorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdministratorsClient.Create`

```go
ctx := context.TODO()
id := administrators.NewAdministratorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "objectId")

payload := administrators.ActiveDirectoryAdministratorAdd{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AdministratorsClient.Delete`

```go
ctx := context.TODO()
id := administrators.NewAdministratorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "objectId")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AdministratorsClient.Get`

```go
ctx := context.TODO()
id := administrators.NewAdministratorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "objectId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministratorsClient.ListByServer`

```go
ctx := context.TODO()
id := administrators.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
