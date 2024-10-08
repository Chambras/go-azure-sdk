
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billings` Documentation

The `billings` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billings"
```


### Client Initialization

```go
client := billings.NewBillingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingsClient.BillingSubscriptionsTransfer`

```go
ctx := context.TODO()
id := billings.NewInvoiceSectionBillingSubscriptionID("billingAccountName", "billingProfileName", "invoiceSectionName", "billingSubscriptionName")

payload := billings.TransferBillingSubscriptionRequestProperties{
	// ...
}


if err := client.BillingSubscriptionsTransferThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingsClient.BillingSubscriptionsValidateTransfer`

```go
ctx := context.TODO()
id := billings.NewInvoiceSectionBillingSubscriptionID("billingAccountName", "billingProfileName", "invoiceSectionName", "billingSubscriptionName")

payload := billings.TransferBillingSubscriptionRequestProperties{
	// ...
}


read, err := client.BillingSubscriptionsValidateTransfer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingsClient.ProductsValidateTransfer`

```go
ctx := context.TODO()
id := billings.NewInvoiceSectionProductID("billingAccountName", "billingProfileName", "invoiceSectionName", "productName")

payload := billings.TransferProductRequestProperties{
	// ...
}


read, err := client.ProductsValidateTransfer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
