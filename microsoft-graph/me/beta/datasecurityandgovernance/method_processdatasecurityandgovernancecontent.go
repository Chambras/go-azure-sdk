package datasecurityandgovernance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessDataSecurityAndGovernanceContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ProcessContentResponse
}

type ProcessDataSecurityAndGovernanceContentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultProcessDataSecurityAndGovernanceContentOperationOptions() ProcessDataSecurityAndGovernanceContentOperationOptions {
	return ProcessDataSecurityAndGovernanceContentOperationOptions{}
}

func (o ProcessDataSecurityAndGovernanceContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProcessDataSecurityAndGovernanceContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ProcessDataSecurityAndGovernanceContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ProcessDataSecurityAndGovernanceContent - Invoke action processContent. Process content against data protection
// policies in the context of the current user.
func (c DataSecurityAndGovernanceClient) ProcessDataSecurityAndGovernanceContent(ctx context.Context, input ProcessDataSecurityAndGovernanceContentRequest, options ProcessDataSecurityAndGovernanceContentOperationOptions) (result ProcessDataSecurityAndGovernanceContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/dataSecurityAndGovernance/processContent",
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.ProcessContentResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
