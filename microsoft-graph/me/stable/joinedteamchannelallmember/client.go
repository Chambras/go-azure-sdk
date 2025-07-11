package joinedteamchannelallmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelAllMemberClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelAllMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelAllMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelallmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelAllMemberClient: %+v", err)
	}

	return &JoinedTeamChannelAllMemberClient{
		Client: client,
	}, nil
}
