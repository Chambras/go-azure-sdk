package datasetmapping

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = SqlDBTableDataSetMapping{}

type SqlDBTableDataSetMapping struct {
	Properties SqlDBTableDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping

	Id         *string                `json:"id,omitempty"`
	Kind       DataSetMappingKind     `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s SqlDBTableDataSetMapping) DataSetMapping() BaseDataSetMappingImpl {
	return BaseDataSetMappingImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = SqlDBTableDataSetMapping{}

func (s SqlDBTableDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper SqlDBTableDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SqlDBTableDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SqlDBTableDataSetMapping: %+v", err)
	}

	decoded["kind"] = "SqlDBTable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SqlDBTableDataSetMapping: %+v", err)
	}

	return encoded, nil
}
