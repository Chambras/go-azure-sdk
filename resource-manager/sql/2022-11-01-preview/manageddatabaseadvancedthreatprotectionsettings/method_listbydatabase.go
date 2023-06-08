package manageddatabaseadvancedthreatprotectionsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDatabaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedDatabaseAdvancedThreatProtection
}

type ListByDatabaseCompleteResult struct {
	Items []ManagedDatabaseAdvancedThreatProtection
}

// ListByDatabase ...
func (c ManagedDatabaseAdvancedThreatProtectionSettingsClient) ListByDatabase(ctx context.Context, id ManagedInstanceDatabaseId) (result ListByDatabaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/advancedThreatProtectionSettings", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]ManagedDatabaseAdvancedThreatProtection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDatabaseComplete retrieves all the results into a single object
func (c ManagedDatabaseAdvancedThreatProtectionSettingsClient) ListByDatabaseComplete(ctx context.Context, id ManagedInstanceDatabaseId) (ListByDatabaseCompleteResult, error) {
	return c.ListByDatabaseCompleteMatchingPredicate(ctx, id, ManagedDatabaseAdvancedThreatProtectionOperationPredicate{})
}

// ListByDatabaseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDatabaseAdvancedThreatProtectionSettingsClient) ListByDatabaseCompleteMatchingPredicate(ctx context.Context, id ManagedInstanceDatabaseId, predicate ManagedDatabaseAdvancedThreatProtectionOperationPredicate) (result ListByDatabaseCompleteResult, err error) {
	items := make([]ManagedDatabaseAdvancedThreatProtection, 0)

	resp, err := c.ListByDatabase(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByDatabaseCompleteResult{
		Items: items,
	}
	return
}