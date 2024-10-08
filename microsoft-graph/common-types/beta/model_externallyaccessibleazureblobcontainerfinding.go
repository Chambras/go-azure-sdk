package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = ExternallyAccessibleAzureBlobContainerFinding{}

type ExternallyAccessibleAzureBlobContainerFinding struct {
	Accessibility       *AzureAccessType             `json:"accessibility,omitempty"`
	EncryptionManagedBy *AzureEncryption             `json:"encryptionManagedBy,omitempty"`
	StorageAccount      *AuthorizationSystemResource `json:"storageAccount,omitempty"`

	// Fields inherited from Finding

	// Defines when the finding was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

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

func (s ExternallyAccessibleAzureBlobContainerFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s ExternallyAccessibleAzureBlobContainerFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternallyAccessibleAzureBlobContainerFinding{}

func (s ExternallyAccessibleAzureBlobContainerFinding) MarshalJSON() ([]byte, error) {
	type wrapper ExternallyAccessibleAzureBlobContainerFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternallyAccessibleAzureBlobContainerFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternallyAccessibleAzureBlobContainerFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externallyAccessibleAzureBlobContainerFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternallyAccessibleAzureBlobContainerFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ExternallyAccessibleAzureBlobContainerFinding{}

func (s *ExternallyAccessibleAzureBlobContainerFinding) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Accessibility       *AzureAccessType `json:"accessibility,omitempty"`
		EncryptionManagedBy *AzureEncryption `json:"encryptionManagedBy,omitempty"`
		CreatedDateTime     *string          `json:"createdDateTime,omitempty"`
		Id                  *string          `json:"id,omitempty"`
		ODataId             *string          `json:"@odata.id,omitempty"`
		ODataType           *string          `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Accessibility = decoded.Accessibility
	s.EncryptionManagedBy = decoded.EncryptionManagedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExternallyAccessibleAzureBlobContainerFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["storageAccount"]; ok {
		impl, err := UnmarshalAuthorizationSystemResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StorageAccount' for 'ExternallyAccessibleAzureBlobContainerFinding': %+v", err)
		}
		s.StorageAccount = &impl
	}

	return nil
}
