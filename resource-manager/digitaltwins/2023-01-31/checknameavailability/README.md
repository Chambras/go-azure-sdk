
## `github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2023-01-31/checknameavailability` Documentation

The `checknameavailability` SDK allows for interaction with Azure Resource Manager `digitaltwins` (API Version `2023-01-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2023-01-31/checknameavailability"
```


### Client Initialization

```go
client := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckNameAvailabilityClient.DigitalTwinsCheckNameAvailability`

```go
ctx := context.TODO()
id := checknameavailability.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

payload := checknameavailability.CheckNameRequest{
	// ...
}


read, err := client.DigitalTwinsCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
