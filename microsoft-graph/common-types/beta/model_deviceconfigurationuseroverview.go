package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceConfigurationUserOverview{}

type DeviceConfigurationUserOverview struct {
	// Version of the policy for that overview
	ConfigurationVersion *int64 `json:"configurationVersion,omitempty"`

	// Number of users in conflict
	ConflictCount *int64 `json:"conflictCount,omitempty"`

	// Number of error Users
	ErrorCount *int64 `json:"errorCount,omitempty"`

	// Number of failed Users
	FailedCount *int64 `json:"failedCount,omitempty"`

	// Last update time
	LastUpdateDateTime *string `json:"lastUpdateDateTime,omitempty"`

	// Number of not applicable users
	NotApplicableCount *int64 `json:"notApplicableCount,omitempty"`

	// Number of pending Users
	PendingCount *int64 `json:"pendingCount,omitempty"`

	// Number of succeeded Users
	SuccessCount *int64 `json:"successCount,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceConfigurationUserOverview) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceConfigurationUserOverview{}

func (s DeviceConfigurationUserOverview) MarshalJSON() ([]byte, error) {
	type wrapper DeviceConfigurationUserOverview
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceConfigurationUserOverview: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceConfigurationUserOverview: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceConfigurationUserOverview"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceConfigurationUserOverview: %+v", err)
	}

	return encoded, nil
}
