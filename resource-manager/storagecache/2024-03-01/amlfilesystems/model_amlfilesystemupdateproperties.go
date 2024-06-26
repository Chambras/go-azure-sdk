package amlfilesystems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AmlFilesystemUpdateProperties struct {
	EncryptionSettings *AmlFilesystemEncryptionSettings                `json:"encryptionSettings,omitempty"`
	MaintenanceWindow  *AmlFilesystemUpdatePropertiesMaintenanceWindow `json:"maintenanceWindow,omitempty"`
	RootSquashSettings *AmlFilesystemRootSquashSettings                `json:"rootSquashSettings,omitempty"`
}
