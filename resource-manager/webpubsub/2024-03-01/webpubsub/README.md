
## `github.com/hashicorp/go-azure-sdk/resource-manager/webpubsub/2024-03-01/webpubsub` Documentation

The `webpubsub` SDK allows for interaction with Azure Resource Manager `webpubsub` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/webpubsub/2024-03-01/webpubsub"
```


### Client Initialization

```go
client := webpubsub.NewWebPubSubClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WebPubSubClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := webpubsub.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

payload := webpubsub.NameAvailabilityParameters{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

payload := webpubsub.WebPubSubResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.CustomCertificatesCreateOrUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewCustomCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "certificateName")

payload := webpubsub.CustomCertificate{
	// ...
}


if err := client.CustomCertificatesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.CustomCertificatesDelete`

```go
ctx := context.TODO()
id := webpubsub.NewCustomCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "certificateName")

read, err := client.CustomCertificatesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.CustomCertificatesGet`

```go
ctx := context.TODO()
id := webpubsub.NewCustomCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "certificateName")

read, err := client.CustomCertificatesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.CustomCertificatesList`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.CustomCertificatesList(ctx, id)` can be used to do batched pagination
items, err := client.CustomCertificatesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.CustomDomainsCreateOrUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "name")

payload := webpubsub.CustomDomain{
	// ...
}


if err := client.CustomDomainsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.CustomDomainsDelete`

```go
ctx := context.TODO()
id := webpubsub.NewCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "name")

if err := client.CustomDomainsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.CustomDomainsGet`

```go
ctx := context.TODO()
id := webpubsub.NewCustomDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "name")

read, err := client.CustomDomainsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.CustomDomainsList`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.CustomDomainsList(ctx, id)` can be used to do batched pagination
items, err := client.CustomDomainsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.Delete`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.Get`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.HubsCreateOrUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewHubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "hubName")

payload := webpubsub.WebPubSubHub{
	// ...
}


if err := client.HubsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.HubsDelete`

```go
ctx := context.TODO()
id := webpubsub.NewHubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "hubName")

if err := client.HubsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.HubsGet`

```go
ctx := context.TODO()
id := webpubsub.NewHubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "hubName")

read, err := client.HubsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.HubsList`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.HubsList(ctx, id)` can be used to do batched pagination
items, err := client.HubsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.ListByResourceGroup`

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


### Example Usage: `WebPubSubClient.ListBySubscription`

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


### Example Usage: `WebPubSubClient.ListKeys`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.ListReplicaSkus`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName")

// alternatively `client.ListReplicaSkus(ctx, id)` can be used to do batched pagination
items, err := client.ListReplicaSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.ListSkus`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.ListSkus(ctx, id)` can be used to do batched pagination
items, err := client.ListSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := webpubsub.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "privateEndpointConnectionName")

if err := client.PrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := webpubsub.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "privateEndpointConnectionName")

read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.PrivateEndpointConnectionsList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.PrivateEndpointConnectionsUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "privateEndpointConnectionName")

payload := webpubsub.PrivateEndpointConnection{
	// ...
}


read, err := client.PrivateEndpointConnectionsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.PrivateLinkResourcesList`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.PrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.RegenerateKey`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

payload := webpubsub.RegenerateKeyParameters{
	// ...
}


if err := client.RegenerateKeyThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.ReplicaSharedPrivateLinkResourcesCreateOrUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName", "sharedPrivateLinkResourceName")

payload := webpubsub.SharedPrivateLinkResource{
	// ...
}


if err := client.ReplicaSharedPrivateLinkResourcesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.ReplicaSharedPrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName", "sharedPrivateLinkResourceName")

read, err := client.ReplicaSharedPrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.ReplicaSharedPrivateLinkResourcesList`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName")

// alternatively `client.ReplicaSharedPrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.ReplicaSharedPrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.ReplicasCreateOrUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName")

payload := webpubsub.Replica{
	// ...
}


if err := client.ReplicasCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.ReplicasDelete`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName")

read, err := client.ReplicasDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.ReplicasGet`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName")

read, err := client.ReplicasGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.ReplicasList`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.ReplicasList(ctx, id)` can be used to do batched pagination
items, err := client.ReplicasListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.ReplicasRestart`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName")

if err := client.ReplicasRestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.ReplicasUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewReplicaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "replicaName")

payload := webpubsub.Replica{
	// ...
}


if err := client.ReplicasUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.Restart`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

if err := client.RestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.SharedPrivateLinkResourcesCreateOrUpdate`

```go
ctx := context.TODO()
id := webpubsub.NewSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "sharedPrivateLinkResourceName")

payload := webpubsub.SharedPrivateLinkResource{
	// ...
}


if err := client.SharedPrivateLinkResourcesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.SharedPrivateLinkResourcesDelete`

```go
ctx := context.TODO()
id := webpubsub.NewSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "sharedPrivateLinkResourceName")

if err := client.SharedPrivateLinkResourcesDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.SharedPrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := webpubsub.NewSharedPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "sharedPrivateLinkResourceName")

read, err := client.SharedPrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebPubSubClient.SharedPrivateLinkResourcesList`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

// alternatively `client.SharedPrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.SharedPrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebPubSubClient.Update`

```go
ctx := context.TODO()
id := webpubsub.NewWebPubSubID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

payload := webpubsub.WebPubSubResource{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebPubSubClient.UsagesList`

```go
ctx := context.TODO()
id := webpubsub.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

// alternatively `client.UsagesList(ctx, id)` can be used to do batched pagination
items, err := client.UsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
