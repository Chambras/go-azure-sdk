
## `github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2023-12-30/backups` Documentation

The `backups` SDK allows for interaction with Azure Resource Manager `mysql` (API Version `2023-12-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2023-12-30/backups"
```


### Client Initialization

```go
client := backups.NewBackupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupsClient.Get`

```go
ctx := context.TODO()
id := backups.NewBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "backupName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupsClient.ListByServer`

```go
ctx := context.TODO()
id := backups.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BackupsClient.Put`

```go
ctx := context.TODO()
id := backups.NewBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "backupName")

read, err := client.Put(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
