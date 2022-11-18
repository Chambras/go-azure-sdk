package fileshares

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnabledProtocols string

const (
	EnabledProtocolsNFS EnabledProtocols = "NFS"
	EnabledProtocolsSMB EnabledProtocols = "SMB"
)

func PossibleValuesForEnabledProtocols() []string {
	return []string{
		string(EnabledProtocolsNFS),
		string(EnabledProtocolsSMB),
	}
}

func parseEnabledProtocols(input string) (*EnabledProtocols, error) {
	vals := map[string]EnabledProtocols{
		"nfs": EnabledProtocolsNFS,
		"smb": EnabledProtocolsSMB,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnabledProtocols(input)
	return &out, nil
}

type LeaseDuration string

const (
	LeaseDurationFixed    LeaseDuration = "Fixed"
	LeaseDurationInfinite LeaseDuration = "Infinite"
)

func PossibleValuesForLeaseDuration() []string {
	return []string{
		string(LeaseDurationFixed),
		string(LeaseDurationInfinite),
	}
}

func parseLeaseDuration(input string) (*LeaseDuration, error) {
	vals := map[string]LeaseDuration{
		"fixed":    LeaseDurationFixed,
		"infinite": LeaseDurationInfinite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LeaseDuration(input)
	return &out, nil
}

type LeaseShareAction string

const (
	LeaseShareActionAcquire LeaseShareAction = "Acquire"
	LeaseShareActionBreak   LeaseShareAction = "Break"
	LeaseShareActionChange  LeaseShareAction = "Change"
	LeaseShareActionRelease LeaseShareAction = "Release"
	LeaseShareActionRenew   LeaseShareAction = "Renew"
)

func PossibleValuesForLeaseShareAction() []string {
	return []string{
		string(LeaseShareActionAcquire),
		string(LeaseShareActionBreak),
		string(LeaseShareActionChange),
		string(LeaseShareActionRelease),
		string(LeaseShareActionRenew),
	}
}

func parseLeaseShareAction(input string) (*LeaseShareAction, error) {
	vals := map[string]LeaseShareAction{
		"acquire": LeaseShareActionAcquire,
		"break":   LeaseShareActionBreak,
		"change":  LeaseShareActionChange,
		"release": LeaseShareActionRelease,
		"renew":   LeaseShareActionRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LeaseShareAction(input)
	return &out, nil
}

type LeaseState string

const (
	LeaseStateAvailable LeaseState = "Available"
	LeaseStateBreaking  LeaseState = "Breaking"
	LeaseStateBroken    LeaseState = "Broken"
	LeaseStateExpired   LeaseState = "Expired"
	LeaseStateLeased    LeaseState = "Leased"
)

func PossibleValuesForLeaseState() []string {
	return []string{
		string(LeaseStateAvailable),
		string(LeaseStateBreaking),
		string(LeaseStateBroken),
		string(LeaseStateExpired),
		string(LeaseStateLeased),
	}
}

func parseLeaseState(input string) (*LeaseState, error) {
	vals := map[string]LeaseState{
		"available": LeaseStateAvailable,
		"breaking":  LeaseStateBreaking,
		"broken":    LeaseStateBroken,
		"expired":   LeaseStateExpired,
		"leased":    LeaseStateLeased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LeaseState(input)
	return &out, nil
}

type LeaseStatus string

const (
	LeaseStatusLocked   LeaseStatus = "Locked"
	LeaseStatusUnlocked LeaseStatus = "Unlocked"
)

func PossibleValuesForLeaseStatus() []string {
	return []string{
		string(LeaseStatusLocked),
		string(LeaseStatusUnlocked),
	}
}

func parseLeaseStatus(input string) (*LeaseStatus, error) {
	vals := map[string]LeaseStatus{
		"locked":   LeaseStatusLocked,
		"unlocked": LeaseStatusUnlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LeaseStatus(input)
	return &out, nil
}

type RootSquashType string

const (
	RootSquashTypeAllSquash    RootSquashType = "AllSquash"
	RootSquashTypeNoRootSquash RootSquashType = "NoRootSquash"
	RootSquashTypeRootSquash   RootSquashType = "RootSquash"
)

func PossibleValuesForRootSquashType() []string {
	return []string{
		string(RootSquashTypeAllSquash),
		string(RootSquashTypeNoRootSquash),
		string(RootSquashTypeRootSquash),
	}
}

func parseRootSquashType(input string) (*RootSquashType, error) {
	vals := map[string]RootSquashType{
		"allsquash":    RootSquashTypeAllSquash,
		"norootsquash": RootSquashTypeNoRootSquash,
		"rootsquash":   RootSquashTypeRootSquash,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RootSquashType(input)
	return &out, nil
}

type ShareAccessTier string

const (
	ShareAccessTierCool                 ShareAccessTier = "Cool"
	ShareAccessTierHot                  ShareAccessTier = "Hot"
	ShareAccessTierPremium              ShareAccessTier = "Premium"
	ShareAccessTierTransactionOptimized ShareAccessTier = "TransactionOptimized"
)

func PossibleValuesForShareAccessTier() []string {
	return []string{
		string(ShareAccessTierCool),
		string(ShareAccessTierHot),
		string(ShareAccessTierPremium),
		string(ShareAccessTierTransactionOptimized),
	}
}

func parseShareAccessTier(input string) (*ShareAccessTier, error) {
	vals := map[string]ShareAccessTier{
		"cool":                 ShareAccessTierCool,
		"hot":                  ShareAccessTierHot,
		"premium":              ShareAccessTierPremium,
		"transactionoptimized": ShareAccessTierTransactionOptimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShareAccessTier(input)
	return &out, nil
}
