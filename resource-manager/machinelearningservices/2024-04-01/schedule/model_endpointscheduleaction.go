package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScheduleActionBase = EndpointScheduleAction{}

type EndpointScheduleAction struct {
	EndpointInvocationDefinition interface{} `json:"endpointInvocationDefinition"`

	// Fields inherited from ScheduleActionBase

	ActionType ScheduleActionType `json:"actionType"`
}

func (s EndpointScheduleAction) ScheduleActionBase() BaseScheduleActionBaseImpl {
	return BaseScheduleActionBaseImpl{
		ActionType: s.ActionType,
	}
}

var _ json.Marshaler = EndpointScheduleAction{}

func (s EndpointScheduleAction) MarshalJSON() ([]byte, error) {
	type wrapper EndpointScheduleAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EndpointScheduleAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EndpointScheduleAction: %+v", err)
	}

	decoded["actionType"] = "InvokeBatchEndpoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EndpointScheduleAction: %+v", err)
	}

	return encoded, nil
}
