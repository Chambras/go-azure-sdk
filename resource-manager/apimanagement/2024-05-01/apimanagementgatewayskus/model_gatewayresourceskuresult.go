package apimanagementgatewayskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayResourceSkuResult struct {
	Capacity     *GatewaySkuCapacity `json:"capacity,omitempty"`
	ResourceType *string             `json:"resourceType,omitempty"`
	Sku          *GatewaySku         `json:"sku,omitempty"`
}
