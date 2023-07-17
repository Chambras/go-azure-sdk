package eventsubscriptions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = EventSubscriptionId{}

// EventSubscriptionId is a struct representing the Resource ID for a Event Subscription
type EventSubscriptionId struct {
	SubscriptionId        string
	ResourceGroupName     string
	TopicName             string
	EventSubscriptionName string
}

// NewEventSubscriptionID returns a new EventSubscriptionId struct
func NewEventSubscriptionID(subscriptionId string, resourceGroupName string, topicName string, eventSubscriptionName string) EventSubscriptionId {
	return EventSubscriptionId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		TopicName:             topicName,
		EventSubscriptionName: eventSubscriptionName,
	}
}

// ParseEventSubscriptionID parses 'input' into a EventSubscriptionId
func ParseEventSubscriptionID(input string) (*EventSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(EventSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EventSubscriptionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.TopicName, ok = parsed.Parsed["topicName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "topicName", *parsed)
	}

	if id.EventSubscriptionName, ok = parsed.Parsed["eventSubscriptionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventSubscriptionName", *parsed)
	}

	return &id, nil
}

// ParseEventSubscriptionIDInsensitively parses 'input' case-insensitively into a EventSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseEventSubscriptionIDInsensitively(input string) (*EventSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(EventSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EventSubscriptionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.TopicName, ok = parsed.Parsed["topicName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "topicName", *parsed)
	}

	if id.EventSubscriptionName, ok = parsed.Parsed["eventSubscriptionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventSubscriptionName", *parsed)
	}

	return &id, nil
}

// ValidateEventSubscriptionID checks that 'input' can be parsed as a Event Subscription ID
func ValidateEventSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEventSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Event Subscription ID
func (id EventSubscriptionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.EventGrid/topics/%s/eventSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.TopicName, id.EventSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Event Subscription ID
func (id EventSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftEventGrid", "Microsoft.EventGrid", "Microsoft.EventGrid"),
		resourceids.StaticSegment("staticTopics", "topics", "topics"),
		resourceids.UserSpecifiedSegment("topicName", "topicValue"),
		resourceids.StaticSegment("staticEventSubscriptions", "eventSubscriptions", "eventSubscriptions"),
		resourceids.UserSpecifiedSegment("eventSubscriptionName", "eventSubscriptionValue"),
	}
}

// String returns a human-readable description of this Event Subscription ID
func (id EventSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Topic Name: %q", id.TopicName),
		fmt.Sprintf("Event Subscription Name: %q", id.EventSubscriptionName),
	}
	return fmt.Sprintf("Event Subscription (%s)", strings.Join(components, "\n"))
}
