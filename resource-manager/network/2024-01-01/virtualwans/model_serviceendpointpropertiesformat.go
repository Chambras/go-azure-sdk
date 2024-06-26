package virtualwans

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceEndpointPropertiesFormat struct {
	Locations         *[]string          `json:"locations,omitempty"`
	NetworkIdentifier *SubResource       `json:"networkIdentifier,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Service           *string            `json:"service,omitempty"`
}
