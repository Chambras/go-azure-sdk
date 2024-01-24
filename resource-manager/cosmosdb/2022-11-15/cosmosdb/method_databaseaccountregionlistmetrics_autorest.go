package cosmosdb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseAccountRegionListMetricsOperationResponse struct {
	HttpResponse *http.Response
	Model        *MetricListResult
}

type DatabaseAccountRegionListMetricsOperationOptions struct {
	Filter *string
}

func DefaultDatabaseAccountRegionListMetricsOperationOptions() DatabaseAccountRegionListMetricsOperationOptions {
	return DatabaseAccountRegionListMetricsOperationOptions{}
}

func (o DatabaseAccountRegionListMetricsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DatabaseAccountRegionListMetricsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// DatabaseAccountRegionListMetrics ...
func (c CosmosDBClient) DatabaseAccountRegionListMetrics(ctx context.Context, id RegionId, options DatabaseAccountRegionListMetricsOperationOptions) (result DatabaseAccountRegionListMetricsOperationResponse, err error) {
	req, err := c.preparerForDatabaseAccountRegionListMetrics(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cosmosdb.CosmosDBClient", "DatabaseAccountRegionListMetrics", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "cosmosdb.CosmosDBClient", "DatabaseAccountRegionListMetrics", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDatabaseAccountRegionListMetrics(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cosmosdb.CosmosDBClient", "DatabaseAccountRegionListMetrics", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDatabaseAccountRegionListMetrics prepares the DatabaseAccountRegionListMetrics request.
func (c CosmosDBClient) preparerForDatabaseAccountRegionListMetrics(ctx context.Context, id RegionId, options DatabaseAccountRegionListMetricsOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/metrics", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDatabaseAccountRegionListMetrics handles the response to the DatabaseAccountRegionListMetrics request. The method always
// closes the http.Response Body.
func (c CosmosDBClient) responderForDatabaseAccountRegionListMetrics(resp *http.Response) (result DatabaseAccountRegionListMetricsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
