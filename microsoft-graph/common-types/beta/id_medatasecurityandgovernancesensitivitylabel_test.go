package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDataSecurityAndGovernanceSensitivityLabelId{}

func TestNewMeDataSecurityAndGovernanceSensitivityLabelID(t *testing.T) {
	id := NewMeDataSecurityAndGovernanceSensitivityLabelID("sensitivityLabelId")

	if id.SensitivityLabelId != "sensitivityLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelId")
	}
}

func TestFormatMeDataSecurityAndGovernanceSensitivityLabelID(t *testing.T) {
	actual := NewMeDataSecurityAndGovernanceSensitivityLabelID("sensitivityLabelId").ID()
	expected := "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDataSecurityAndGovernanceSensitivityLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDataSecurityAndGovernanceSensitivityLabelId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId",
			Expected: &MeDataSecurityAndGovernanceSensitivityLabelId{
				SensitivityLabelId: "sensitivityLabelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDataSecurityAndGovernanceSensitivityLabelID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

	}
}

func TestParseMeDataSecurityAndGovernanceSensitivityLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDataSecurityAndGovernanceSensitivityLabelId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId",
			Expected: &MeDataSecurityAndGovernanceSensitivityLabelId{
				SensitivityLabelId: "sensitivityLabelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD",
			Expected: &MeDataSecurityAndGovernanceSensitivityLabelId{
				SensitivityLabelId: "sEnSiTiViTyLaBeLiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDataSecurityAndGovernanceSensitivityLabelIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

	}
}

func TestSegmentsForMeDataSecurityAndGovernanceSensitivityLabelId(t *testing.T) {
	segments := MeDataSecurityAndGovernanceSensitivityLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDataSecurityAndGovernanceSensitivityLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
