package securityuserconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserConfigurationPropertiesFormat struct {
	Description       *string            `json:"description,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	ResourceGuid      *string            `json:"resourceGuid,omitempty"`
}
