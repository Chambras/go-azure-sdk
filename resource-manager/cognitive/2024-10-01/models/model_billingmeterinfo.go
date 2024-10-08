package models

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingMeterInfo struct {
	MeterId *string `json:"meterId,omitempty"`
	Name    *string `json:"name,omitempty"`
	Unit    *string `json:"unit,omitempty"`
}
