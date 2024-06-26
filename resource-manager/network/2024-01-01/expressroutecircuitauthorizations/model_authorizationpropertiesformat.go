package expressroutecircuitauthorizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationPropertiesFormat struct {
	AuthorizationKey       *string                 `json:"authorizationKey,omitempty"`
	AuthorizationUseStatus *AuthorizationUseStatus `json:"authorizationUseStatus,omitempty"`
	ConnectionResourceUri  *string                 `json:"connectionResourceUri,omitempty"`
	ProvisioningState      *ProvisioningState      `json:"provisioningState,omitempty"`
}
