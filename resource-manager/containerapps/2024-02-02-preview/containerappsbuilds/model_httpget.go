package containerappsbuilds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HTTPGet struct {
	FileName *string   `json:"fileName,omitempty"`
	Headers  *[]string `json:"headers,omitempty"`
	Url      string    `json:"url"`
}
