package events

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]EventSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByBillingProfileOperationResponse, error)
}

type ListByBillingProfileCompleteResult struct {
	Items []EventSummary
}

func (r ListByBillingProfileOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByBillingProfileOperationResponse) LoadMore(ctx context.Context) (resp ListByBillingProfileOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByBillingProfileOperationOptions struct {
	EndDate   *string
	StartDate *string
}

func DefaultListByBillingProfileOperationOptions() ListByBillingProfileOperationOptions {
	return ListByBillingProfileOperationOptions{}
}

func (o ListByBillingProfileOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByBillingProfileOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndDate != nil {
		out["endDate"] = *o.EndDate
	}

	if o.StartDate != nil {
		out["startDate"] = *o.StartDate
	}

	return out
}

// ListByBillingProfile ...
func (c EventsClient) ListByBillingProfile(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (resp ListByBillingProfileOperationResponse, err error) {
	req, err := c.preparerForListByBillingProfile(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "events.EventsClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "events.EventsClient", "ListByBillingProfile", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByBillingProfile(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "events.EventsClient", "ListByBillingProfile", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByBillingProfile prepares the ListByBillingProfile request.
func (c EventsClient) preparerForListByBillingProfile(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Consumption/events", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByBillingProfileWithNextLink prepares the ListByBillingProfile request with the given nextLink token.
func (c EventsClient) preparerForListByBillingProfileWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByBillingProfile handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (c EventsClient) responderForListByBillingProfile(resp *http.Response) (result ListByBillingProfileOperationResponse, err error) {
	type page struct {
		Values   []EventSummary `json:"value"`
		NextLink *string        `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByBillingProfileOperationResponse, err error) {
			req, err := c.preparerForListByBillingProfileWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "events.EventsClient", "ListByBillingProfile", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "events.EventsClient", "ListByBillingProfile", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByBillingProfile(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "events.EventsClient", "ListByBillingProfile", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByBillingProfileComplete retrieves all of the results into a single object
func (c EventsClient) ListByBillingProfileComplete(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (ListByBillingProfileCompleteResult, error) {
	return c.ListByBillingProfileCompleteMatchingPredicate(ctx, id, options, EventSummaryOperationPredicate{})
}

// ListByBillingProfileCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c EventsClient) ListByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions, predicate EventSummaryOperationPredicate) (resp ListByBillingProfileCompleteResult, err error) {
	items := make([]EventSummary, 0)

	page, err := c.ListByBillingProfile(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListByBillingProfileCompleteResult{
		Items: items,
	}
	return out, nil
}
