package dataset

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSet = KustoClusterDataSet{}

type KustoClusterDataSet struct {
	Properties KustoClusterDataSetProperties `json:"properties"`

	// Fields inherited from DataSet
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = KustoClusterDataSet{}

func (s KustoClusterDataSet) MarshalJSON() ([]byte, error) {
	type wrapper KustoClusterDataSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KustoClusterDataSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KustoClusterDataSet: %+v", err)
	}
	decoded["kind"] = "KustoCluster"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KustoClusterDataSet: %+v", err)
	}

	return encoded, nil
}
