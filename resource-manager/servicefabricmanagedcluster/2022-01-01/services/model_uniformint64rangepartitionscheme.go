package services

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Partition = UniformInt64RangePartitionScheme{}

type UniformInt64RangePartitionScheme struct {
	Count   int64 `json:"count"`
	HighKey int64 `json:"highKey"`
	LowKey  int64 `json:"lowKey"`

	// Fields inherited from Partition

	PartitionScheme PartitionScheme `json:"partitionScheme"`
}

func (s UniformInt64RangePartitionScheme) Partition() BasePartitionImpl {
	return BasePartitionImpl{
		PartitionScheme: s.PartitionScheme,
	}
}

var _ json.Marshaler = UniformInt64RangePartitionScheme{}

func (s UniformInt64RangePartitionScheme) MarshalJSON() ([]byte, error) {
	type wrapper UniformInt64RangePartitionScheme
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UniformInt64RangePartitionScheme: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UniformInt64RangePartitionScheme: %+v", err)
	}

	decoded["partitionScheme"] = "UniformInt64Range"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UniformInt64RangePartitionScheme: %+v", err)
	}

	return encoded, nil
}
