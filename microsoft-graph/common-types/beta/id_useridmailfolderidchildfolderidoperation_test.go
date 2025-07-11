package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdOperationId{}

func TestNewUserIdMailFolderIdChildFolderIdOperationID(t *testing.T) {
	id := NewUserIdMailFolderIdChildFolderIdOperationID("userId", "mailFolderId", "mailFolderId1", "mailFolderOperationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.MailFolderId != "mailFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderId")
	}

	if id.MailFolderId1 != "mailFolderId1" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId1'", id.MailFolderId1, "mailFolderId1")
	}

	if id.MailFolderOperationId != "mailFolderOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderOperationId'", id.MailFolderOperationId, "mailFolderOperationId")
	}
}

func TestFormatUserIdMailFolderIdChildFolderIdOperationID(t *testing.T) {
	actual := NewUserIdMailFolderIdChildFolderIdOperationID("userId", "mailFolderId", "mailFolderId1", "mailFolderOperationId").ID()
	expected := "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMailFolderIdChildFolderIdOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdChildFolderIdOperationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId",
			Expected: &UserIdMailFolderIdChildFolderIdOperationId{
				UserId:                "userId",
				MailFolderId:          "mailFolderId",
				MailFolderId1:         "mailFolderId1",
				MailFolderOperationId: "mailFolderOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdChildFolderIdOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.MailFolderOperationId != v.Expected.MailFolderOperationId {
			t.Fatalf("Expected %q but got %q for MailFolderOperationId", v.Expected.MailFolderOperationId, actual.MailFolderOperationId)
		}

	}
}

func TestParseUserIdMailFolderIdChildFolderIdOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdChildFolderIdOperationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId",
			Expected: &UserIdMailFolderIdChildFolderIdOperationId{
				UserId:                "userId",
				MailFolderId:          "mailFolderId",
				MailFolderId1:         "mailFolderId1",
				MailFolderOperationId: "mailFolderOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/operations/mailFolderOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/oPeRaTiOnS/mAiLfOlDeRoPeRaTiOnId",
			Expected: &UserIdMailFolderIdChildFolderIdOperationId{
				UserId:                "uSeRiD",
				MailFolderId:          "mAiLfOlDeRiD",
				MailFolderId1:         "mAiLfOlDeRiD1",
				MailFolderOperationId: "mAiLfOlDeRoPeRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/oPeRaTiOnS/mAiLfOlDeRoPeRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdChildFolderIdOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.MailFolderOperationId != v.Expected.MailFolderOperationId {
			t.Fatalf("Expected %q but got %q for MailFolderOperationId", v.Expected.MailFolderOperationId, actual.MailFolderOperationId)
		}

	}
}

func TestSegmentsForUserIdMailFolderIdChildFolderIdOperationId(t *testing.T) {
	segments := UserIdMailFolderIdChildFolderIdOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMailFolderIdChildFolderIdOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
