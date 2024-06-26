package dataproducts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworkRule struct {
	Action *string `json:"action,omitempty"`
	Id     string  `json:"id"`
	State  *string `json:"state,omitempty"`
}
