package sqlpoolsreplicationlinks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsReplicationLinksClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsReplicationLinksClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsReplicationLinksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsreplicationlinks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsReplicationLinksClient: %+v", err)
	}

	return &SqlPoolsReplicationLinksClient{
		Client: client,
	}, nil
}
