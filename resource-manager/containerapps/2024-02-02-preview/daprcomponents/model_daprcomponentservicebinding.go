package daprcomponents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprComponentServiceBinding struct {
	Metadata  *DaprServiceBindMetadata `json:"metadata,omitempty"`
	Name      *string                  `json:"name,omitempty"`
	ServiceId *string                  `json:"serviceId,omitempty"`
}
