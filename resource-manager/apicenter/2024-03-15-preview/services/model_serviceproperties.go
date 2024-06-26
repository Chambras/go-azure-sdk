package services

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProperties struct {
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Restore           *bool              `json:"restore,omitempty"`
}
