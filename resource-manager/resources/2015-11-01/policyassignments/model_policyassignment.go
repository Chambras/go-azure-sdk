package policyassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignment struct {
	Name       *string                     `json:"name,omitempty"`
	Properties *PolicyAssignmentProperties `json:"properties,omitempty"`
}
