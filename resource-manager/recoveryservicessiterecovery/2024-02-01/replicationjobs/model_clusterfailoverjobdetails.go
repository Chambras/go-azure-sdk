package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobDetails = ClusterFailoverJobDetails{}

type ClusterFailoverJobDetails struct {
	ProtectedItemDetails *[]FailoverReplicationProtectedItemDetails `json:"protectedItemDetails,omitempty"`

	// Fields inherited from JobDetails

	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
	InstanceType          string             `json:"instanceType"`
}

func (s ClusterFailoverJobDetails) JobDetails() BaseJobDetailsImpl {
	return BaseJobDetailsImpl{
		AffectedObjectDetails: s.AffectedObjectDetails,
		InstanceType:          s.InstanceType,
	}
}

var _ json.Marshaler = ClusterFailoverJobDetails{}

func (s ClusterFailoverJobDetails) MarshalJSON() ([]byte, error) {
	type wrapper ClusterFailoverJobDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ClusterFailoverJobDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ClusterFailoverJobDetails: %+v", err)
	}

	decoded["instanceType"] = "ClusterFailoverJobDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ClusterFailoverJobDetails: %+v", err)
	}

	return encoded, nil
}
