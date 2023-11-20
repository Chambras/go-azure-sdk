package eventhubsclustersconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterQuotaConfigurationProperties struct {
	Settings *map[string]string `json:"settings,omitempty"`
}
