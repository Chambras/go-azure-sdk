package dataproducts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPRules struct {
	Action string  `json:"action"`
	Value  *string `json:"value,omitempty"`
}
