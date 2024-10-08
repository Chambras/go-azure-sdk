
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-08-01/scheduledqueryrules` Documentation

The `scheduledqueryrules` SDK allows for interaction with Azure Resource Manager `insights` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-08-01/scheduledqueryrules"
```


### Client Initialization

```go
client := scheduledqueryrules.NewScheduledQueryRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScheduledQueryRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := scheduledqueryrules.NewScheduledQueryRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ruleName")

payload := scheduledqueryrules.ScheduledQueryRuleResource{
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


### Example Usage: `ScheduledQueryRulesClient.Delete`

```go
ctx := context.TODO()
id := scheduledqueryrules.NewScheduledQueryRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ruleName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScheduledQueryRulesClient.Get`

```go
ctx := context.TODO()
id := scheduledqueryrules.NewScheduledQueryRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ruleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScheduledQueryRulesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScheduledQueryRulesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScheduledQueryRulesClient.Update`

```go
ctx := context.TODO()
id := scheduledqueryrules.NewScheduledQueryRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ruleName")

payload := scheduledqueryrules.ScheduledQueryRuleResourcePatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
