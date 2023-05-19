package sharedgalleries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PirSharedGalleryResource struct {
	Identifier *SharedGalleryIdentifier `json:"identifier,omitempty"`
	Location   *string                  `json:"location,omitempty"`
	Name       *string                  `json:"name,omitempty"`
}
