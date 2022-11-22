package securitymlanalyticssettings

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnomalySecurityMLAnalyticsSettingsProperties struct {
	AnomalySettingsVersion   *int64                                   `json:"anomalySettingsVersion,omitempty"`
	AnomalyVersion           string                                   `json:"anomalyVersion"`
	CustomizableObservations *interface{}                             `json:"customizableObservations,omitempty"`
	Description              *string                                  `json:"description,omitempty"`
	DisplayName              string                                   `json:"displayName"`
	Enabled                  bool                                     `json:"enabled"`
	Frequency                string                                   `json:"frequency"`
	IsDefaultSettings        bool                                     `json:"isDefaultSettings"`
	LastModifiedUtc          *string                                  `json:"lastModifiedUtc,omitempty"`
	RequiredDataConnectors   *[]SecurityMLAnalyticsSettingsDataSource `json:"requiredDataConnectors,omitempty"`
	SettingsDefinitionId     *string                                  `json:"settingsDefinitionId,omitempty"`
	SettingsStatus           SettingsStatus                           `json:"settingsStatus"`
	Tactics                  *[]AttackTactic                          `json:"tactics,omitempty"`
	Techniques               *[]string                                `json:"techniques,omitempty"`
}

func (o *AnomalySecurityMLAnalyticsSettingsProperties) GetLastModifiedUtcAsTime() (*time.Time, error) {
	if o.LastModifiedUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AnomalySecurityMLAnalyticsSettingsProperties) SetLastModifiedUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedUtc = &formatted
}
