
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/sites` Documentation

The `sites` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/sites"
```


### Client Initialization

```go
client := sites.NewSitesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SitesClient.ListByMobileNetwork`

```go
ctx := context.TODO()
id := sites.NewMobileNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mobileNetworkName")

// alternatively `client.ListByMobileNetwork(ctx, id)` can be used to do batched pagination
items, err := client.ListByMobileNetworkComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
