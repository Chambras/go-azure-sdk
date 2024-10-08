package checknameavailabilitywithsubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityInput struct {
	Name string       `json:"name"`
	Type ResourceType `json:"type"`
}
