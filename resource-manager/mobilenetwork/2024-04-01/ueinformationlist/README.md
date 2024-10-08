
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-04-01/ueinformationlist` Documentation

The `ueinformationlist` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-04-01/ueinformationlist"
```


### Client Initialization

```go
client := ueinformationlist.NewUeInformationListClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UeInformationListClient.UeInformationList`

```go
ctx := context.TODO()
id := ueinformationlist.NewPacketCoreControlPlaneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName")

// alternatively `client.UeInformationList(ctx, id)` can be used to do batched pagination
items, err := client.UeInformationListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
