package autonomousdatabaseversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AutonomousDbVersionId{}

func TestNewAutonomousDbVersionID(t *testing.T) {
	id := NewAutonomousDbVersionID("12345678-1234-9876-4563-123456789012", "location", "autonomousdbversionsname")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "location" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "location")
	}

	if id.AutonomousDbVersionName != "autonomousdbversionsname" {
		t.Fatalf("Expected %q but got %q for Segment 'AutonomousDbVersionName'", id.AutonomousDbVersionName, "autonomousdbversionsname")
	}
}

func TestFormatAutonomousDbVersionID(t *testing.T) {
	actual := NewAutonomousDbVersionID("12345678-1234-9876-4563-123456789012", "location", "autonomousdbversionsname").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location/autonomousDbVersions/autonomousdbversionsname"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAutonomousDbVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutonomousDbVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location/autonomousDbVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location/autonomousDbVersions/autonomousdbversionsname",
			Expected: &AutonomousDbVersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				LocationName:            "location",
				AutonomousDbVersionName: "autonomousdbversionsname",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location/autonomousDbVersions/autonomousdbversionsname/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutonomousDbVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.AutonomousDbVersionName != v.Expected.AutonomousDbVersionName {
			t.Fatalf("Expected %q but got %q for AutonomousDbVersionName", v.Expected.AutonomousDbVersionName, actual.AutonomousDbVersionName)
		}

	}
}

func TestParseAutonomousDbVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AutonomousDbVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location/autonomousDbVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoN/aUtOnOmOuSdBvErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location/autonomousDbVersions/autonomousdbversionsname",
			Expected: &AutonomousDbVersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				LocationName:            "location",
				AutonomousDbVersionName: "autonomousdbversionsname",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Oracle.Database/locations/location/autonomousDbVersions/autonomousdbversionsname/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoN/aUtOnOmOuSdBvErSiOnS/aUtOnOmOuSdBvErSiOnSnAmE",
			Expected: &AutonomousDbVersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				LocationName:            "lOcAtIoN",
				AutonomousDbVersionName: "aUtOnOmOuSdBvErSiOnSnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/oRaClE.DaTaBaSe/lOcAtIoNs/lOcAtIoN/aUtOnOmOuSdBvErSiOnS/aUtOnOmOuSdBvErSiOnSnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAutonomousDbVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.AutonomousDbVersionName != v.Expected.AutonomousDbVersionName {
			t.Fatalf("Expected %q but got %q for AutonomousDbVersionName", v.Expected.AutonomousDbVersionName, actual.AutonomousDbVersionName)
		}

	}
}

func TestSegmentsForAutonomousDbVersionId(t *testing.T) {
	segments := AutonomousDbVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AutonomousDbVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
