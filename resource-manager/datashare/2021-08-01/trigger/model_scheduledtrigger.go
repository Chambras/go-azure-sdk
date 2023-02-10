package trigger

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Trigger = ScheduledTrigger{}

type ScheduledTrigger struct {
	Properties ScheduledTriggerProperties `json:"properties"`

	// Fields inherited from Trigger
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = ScheduledTrigger{}

func (s ScheduledTrigger) MarshalJSON() ([]byte, error) {
	type wrapper ScheduledTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScheduledTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScheduledTrigger: %+v", err)
	}
	decoded["kind"] = "ScheduleBased"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScheduledTrigger: %+v", err)
	}

	return encoded, nil
}
