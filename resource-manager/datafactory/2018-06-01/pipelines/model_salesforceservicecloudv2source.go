package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CopySource = SalesforceServiceCloudV2Source{}

type SalesforceServiceCloudV2Source struct {
	AdditionalColumns     *interface{} `json:"additionalColumns,omitempty"`
	IncludeDeletedObjects *interface{} `json:"includeDeletedObjects,omitempty"`
	Query                 *interface{} `json:"query,omitempty"`
	SOQLQuery             *interface{} `json:"SOQLQuery,omitempty"`

	// Fields inherited from CopySource
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	SourceRetryCount         *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{} `json:"sourceRetryWait,omitempty"`
}

var _ json.Marshaler = SalesforceServiceCloudV2Source{}

func (s SalesforceServiceCloudV2Source) MarshalJSON() ([]byte, error) {
	type wrapper SalesforceServiceCloudV2Source
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SalesforceServiceCloudV2Source: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SalesforceServiceCloudV2Source: %+v", err)
	}
	decoded["type"] = "SalesforceServiceCloudV2Source"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SalesforceServiceCloudV2Source: %+v", err)
	}

	return encoded, nil
}
