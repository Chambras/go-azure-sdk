
## `github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/linkedworkspace` Documentation

The `linkedworkspace` SDK allows for interaction with Azure Resource Manager `automation` (API Version `2015-10-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/linkedworkspace"
```


### Client Initialization

```go
client := linkedworkspace.NewLinkedWorkspaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LinkedWorkspaceClient.Get`

```go
ctx := context.TODO()
id := linkedworkspace.NewAutomationAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
