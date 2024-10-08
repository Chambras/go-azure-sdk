package securityuserrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RuleCollectionRuleId{})
}

var _ resourceids.ResourceId = &RuleCollectionRuleId{}

// RuleCollectionRuleId is a struct representing the Resource ID for a Rule Collection Rule
type RuleCollectionRuleId struct {
	SubscriptionId                string
	ResourceGroupName             string
	NetworkManagerName            string
	SecurityUserConfigurationName string
	RuleCollectionName            string
	RuleName                      string
}

// NewRuleCollectionRuleID returns a new RuleCollectionRuleId struct
func NewRuleCollectionRuleID(subscriptionId string, resourceGroupName string, networkManagerName string, securityUserConfigurationName string, ruleCollectionName string, ruleName string) RuleCollectionRuleId {
	return RuleCollectionRuleId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		NetworkManagerName:            networkManagerName,
		SecurityUserConfigurationName: securityUserConfigurationName,
		RuleCollectionName:            ruleCollectionName,
		RuleName:                      ruleName,
	}
}

// ParseRuleCollectionRuleID parses 'input' into a RuleCollectionRuleId
func ParseRuleCollectionRuleID(input string) (*RuleCollectionRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RuleCollectionRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RuleCollectionRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRuleCollectionRuleIDInsensitively parses 'input' case-insensitively into a RuleCollectionRuleId
// note: this method should only be used for API response data and not user input
func ParseRuleCollectionRuleIDInsensitively(input string) (*RuleCollectionRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RuleCollectionRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RuleCollectionRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RuleCollectionRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NetworkManagerName, ok = input.Parsed["networkManagerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "networkManagerName", input)
	}

	if id.SecurityUserConfigurationName, ok = input.Parsed["securityUserConfigurationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityUserConfigurationName", input)
	}

	if id.RuleCollectionName, ok = input.Parsed["ruleCollectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "ruleCollectionName", input)
	}

	if id.RuleName, ok = input.Parsed["ruleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "ruleName", input)
	}

	return nil
}

// ValidateRuleCollectionRuleID checks that 'input' can be parsed as a Rule Collection Rule ID
func ValidateRuleCollectionRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRuleCollectionRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rule Collection Rule ID
func (id RuleCollectionRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkManagers/%s/securityUserConfigurations/%s/ruleCollections/%s/rules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetworkManagerName, id.SecurityUserConfigurationName, id.RuleCollectionName, id.RuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Rule Collection Rule ID
func (id RuleCollectionRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticNetworkManagers", "networkManagers", "networkManagers"),
		resourceids.UserSpecifiedSegment("networkManagerName", "networkManagerName"),
		resourceids.StaticSegment("staticSecurityUserConfigurations", "securityUserConfigurations", "securityUserConfigurations"),
		resourceids.UserSpecifiedSegment("securityUserConfigurationName", "configurationName"),
		resourceids.StaticSegment("staticRuleCollections", "ruleCollections", "ruleCollections"),
		resourceids.UserSpecifiedSegment("ruleCollectionName", "ruleCollectionName"),
		resourceids.StaticSegment("staticRules", "rules", "rules"),
		resourceids.UserSpecifiedSegment("ruleName", "ruleName"),
	}
}

// String returns a human-readable description of this Rule Collection Rule ID
func (id RuleCollectionRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Network Manager Name: %q", id.NetworkManagerName),
		fmt.Sprintf("Security User Configuration Name: %q", id.SecurityUserConfigurationName),
		fmt.Sprintf("Rule Collection Name: %q", id.RuleCollectionName),
		fmt.Sprintf("Rule Name: %q", id.RuleName),
	}
	return fmt.Sprintf("Rule Collection Rule (%s)", strings.Join(components, "\n"))
}
