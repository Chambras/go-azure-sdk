package userinsightmonthlymfaregistereduser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlyMfaRegisteredUserClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlyMfaRegisteredUserClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlyMfaRegisteredUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlymfaregistereduser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlyMfaRegisteredUserClient: %+v", err)
	}

	return &UserInsightMonthlyMfaRegisteredUserClient{
		Client: client,
	}, nil
}
