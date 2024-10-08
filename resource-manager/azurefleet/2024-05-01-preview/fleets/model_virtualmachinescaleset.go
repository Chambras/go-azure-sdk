package fleets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSet struct {
	Error           *ApiError         `json:"error,omitempty"`
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	OperationStatus ProvisioningState `json:"operationStatus"`
	Type            *string           `json:"type,omitempty"`
}
