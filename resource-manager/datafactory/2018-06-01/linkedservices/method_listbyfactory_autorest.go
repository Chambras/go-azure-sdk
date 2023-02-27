package linkedservices

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

type ListByFactoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]LinkedServiceResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByFactoryOperationResponse, error)
}

type ListByFactoryCompleteResult struct {
	Items []LinkedServiceResource
}

func (r ListByFactoryOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByFactoryOperationResponse) LoadMore(ctx context.Context) (resp ListByFactoryOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByFactory ...
func (c LinkedServicesClient) ListByFactory(ctx context.Context, id FactoryId) (resp ListByFactoryOperationResponse, err error) {
	req, err := c.preparerForListByFactory(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByFactory", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByFactory", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByFactory(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByFactory", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByFactory prepares the ListByFactory request.
func (c LinkedServicesClient) preparerForListByFactory(ctx context.Context, id FactoryId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/linkedServices", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByFactoryWithNextLink prepares the ListByFactory request with the given nextLink token.
func (c LinkedServicesClient) preparerForListByFactoryWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByFactory handles the response to the ListByFactory request. The method always
// closes the http.Response Body.
func (c LinkedServicesClient) responderForListByFactory(resp *http.Response) (result ListByFactoryOperationResponse, err error) {
	type page struct {
		Values   []LinkedServiceResource `json:"value"`
		NextLink *string                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByFactoryOperationResponse, err error) {
			req, err := c.preparerForListByFactoryWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByFactory", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByFactory", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByFactory(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByFactory", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByFactoryComplete retrieves all of the results into a single object
func (c LinkedServicesClient) ListByFactoryComplete(ctx context.Context, id FactoryId) (ListByFactoryCompleteResult, error) {
	return c.ListByFactoryCompleteMatchingPredicate(ctx, id, LinkedServiceResourceOperationPredicate{})
}

// ListByFactoryCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c LinkedServicesClient) ListByFactoryCompleteMatchingPredicate(ctx context.Context, id FactoryId, predicate LinkedServiceResourceOperationPredicate) (resp ListByFactoryCompleteResult, err error) {
	items := make([]LinkedServiceResource, 0)

	page, err := c.ListByFactory(ctx, id)
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

	out := ListByFactoryCompleteResult{
		Items: items,
	}
	return out, nil
}