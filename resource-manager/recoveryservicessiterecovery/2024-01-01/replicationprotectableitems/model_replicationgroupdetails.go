package replicationprotectableitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConfigurationSettings = ReplicationGroupDetails{}

type ReplicationGroupDetails struct {

	// Fields inherited from ConfigurationSettings

	InstanceType string `json:"instanceType"`
}

func (s ReplicationGroupDetails) ConfigurationSettings() BaseConfigurationSettingsImpl {
	return BaseConfigurationSettingsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = ReplicationGroupDetails{}

func (s ReplicationGroupDetails) MarshalJSON() ([]byte, error) {
	type wrapper ReplicationGroupDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ReplicationGroupDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ReplicationGroupDetails: %+v", err)
	}

	decoded["instanceType"] = "ReplicationGroupDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ReplicationGroupDetails: %+v", err)
	}

	return encoded, nil
}
