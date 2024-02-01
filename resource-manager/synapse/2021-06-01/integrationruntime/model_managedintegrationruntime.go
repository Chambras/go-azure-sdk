package integrationruntime

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IntegrationRuntime = ManagedIntegrationRuntime{}

type ManagedIntegrationRuntime struct {
	ManagedVirtualNetwork *ManagedIntegrationRuntimeManagedVirtualNetworkReference `json:"managedVirtualNetwork,omitempty"`
	ProvisioningState     *IntegrationRuntimeState                                 `json:"provisioningState,omitempty"`
	TypeProperties        ManagedIntegrationRuntimeTypeProperties                  `json:"typeProperties"`

	// Fields inherited from IntegrationRuntime
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = ManagedIntegrationRuntime{}

func (s ManagedIntegrationRuntime) MarshalJSON() ([]byte, error) {
	type wrapper ManagedIntegrationRuntime
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedIntegrationRuntime: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedIntegrationRuntime: %+v", err)
	}
	decoded["type"] = "Managed"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedIntegrationRuntime: %+v", err)
	}

	return encoded, nil
}
