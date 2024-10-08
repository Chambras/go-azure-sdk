package integrationruntimes

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomSetupBase = CmdkeySetup{}

type CmdkeySetup struct {
	TypeProperties CmdkeySetupTypeProperties `json:"typeProperties"`

	// Fields inherited from CustomSetupBase

	Type string `json:"type"`
}

func (s CmdkeySetup) CustomSetupBase() BaseCustomSetupBaseImpl {
	return BaseCustomSetupBaseImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = CmdkeySetup{}

func (s CmdkeySetup) MarshalJSON() ([]byte, error) {
	type wrapper CmdkeySetup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CmdkeySetup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CmdkeySetup: %+v", err)
	}

	decoded["type"] = "CmdkeySetup"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CmdkeySetup: %+v", err)
	}

	return encoded, nil
}
