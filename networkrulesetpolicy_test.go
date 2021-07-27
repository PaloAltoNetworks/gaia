package gaia

import (
	"testing"

	"go.aporeto.io/elemental"
)

func TestNetworkRuleSetPolicy_Validate_custom(t *testing.T) {

	cases := map[string]struct {
		p   *NetworkRuleSetPolicy
		err elemental.Error
	}{

		"valid": {
			p: &NetworkRuleSetPolicy{
				Name: "dummy",
				Subject: [][]string{
					{"tagA=a", "tagB=b"},
					{"tagC=c", "tagD=d"},
				},
				OutgoingRules: []*NetworkRule{
					{
						Action:        NetworkRuleActionAllow,
						Object:        [][]string{{"tag1=a", "tag2=b"}},
						ProtocolPorts: []string{"udp/123", "tcp/80"},
					},
					{
						Action:        NetworkRuleActionAllow,
						Object:        [][]string{{"tag3=c", "tag4=d"}},
						ProtocolPorts: []string{"tcp/90", "udp/456"},
					},
				},
				IncomingRules: []*NetworkRule{
					{
						Action:        NetworkRuleActionAllow,
						Object:        [][]string{{"tag5=e", "tag6=f"}},
						ProtocolPorts: []string{"udp/789", "tcp/100"},
					},
					{
						Action:        NetworkRuleActionAllow,
						Object:        [][]string{{"tag7=g", "tag8=h"}},
						ProtocolPorts: []string{"tcp/110", "udp/101"},
					},
				},
			},
		},

		"empty-subject": {
			p: &NetworkRuleSetPolicy{
				Name:    "dummy",
				Subject: nil,
			},
			err: makeValidationError("subject", "Expression must contain at least one sub-expression"),
		},

		"empty-subexpression-in-subject": {
			p: &NetworkRuleSetPolicy{
				Name:    "dummy",
				Subject: [][]string{{}},
			},
			err: makeValidationError("subject", "Sub-expression must not be empty"),
		},

		"duplicate-tags-in-subject": {
			p: &NetworkRuleSetPolicy{
				Name: "dummy",
				Subject: [][]string{
					{"tag=dup", "tag=dup"},
				},
			},
			err: makeValidationError("subject", "Duplicate tag in a sub-expression: 'tag=dup'"),
		},

		"duplicate-subexpressions-in-subject": {
			p: &NetworkRuleSetPolicy{
				Name: "dummy",
				Subject: [][]string{
					{"tag1=a", "tag2=b"},
					{"tag2=b", "tag1=a"},
				},
			},
			err: makeValidationError("subject", "Duplicate equivalent sub-expressions found"),
		},

		"bad-tag-in-subject": {
			p: &NetworkRuleSetPolicy{
				Name:    "dummy",
				Subject: [][]string{{"troubles"}},
			},
			err: makeValidationError("subject", "'troubles' must contain at least one '=' symbol separating two valid words"),
		},

		"duplicate-incoming-rules": {
			p: &NetworkRuleSetPolicy{
				Name: "dummy",
				Subject: [][]string{
					{"tag=valid"},
				},
				IncomingRules: []*NetworkRule{
					{
						Action:        NetworkRuleActionAllow,
						Object:        [][]string{{"tag1=a", "tag2=b"}},
						ProtocolPorts: []string{"udp/123", "tcp/80"},
					},
					{
						Action:        NetworkRuleActionAllow,
						Object:        [][]string{{"tag2=b", "tag1=a"}},
						ProtocolPorts: []string{"tcp/80", "udp/123"},
					},
				},
			},
			err: makeValidationError("incomingRules", "Duplicate network rules at the following indexes: [1, 2]"),
		},

		"duplicate-outgoing-rules": {
			p: &NetworkRuleSetPolicy{
				Name: "dummy",
				Subject: [][]string{
					{"tag=valid"},
				},
				OutgoingRules: []*NetworkRule{
					{
						Action:        NetworkRuleActionAllow,
						Object:        [][]string{{"tag1=a", "tag2=b"}},
						ProtocolPorts: []string{"udp/123", "tcp/80"},
					},
					{
						Action:        NetworkRuleActionAllow,
						Object:        [][]string{{"tag2=b", "tag1=a"}},
						ProtocolPorts: []string{"tcp/80", "udp/123"},
					},
				},
			},
			err: makeValidationError("outgoingRules", "Duplicate network rules at the following indexes: [1, 2]"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			checkValidationReturnsExpectedErr(c.p, c.err, t)
		})
	}
}
