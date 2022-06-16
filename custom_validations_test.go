package gaia

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/portutils"
)

func TestValidateNoDuplicateNetworkRules(t *testing.T) {
	cases := map[string]struct {
		attribute   string
		rules       []*NetworkRule
		expectedErr elemental.Error
	}{
		"valid-no-rules": {
			rules:     []*NetworkRule{},
			attribute: t.Name(),
		},
		"valid-nil-rules": {
			rules:     nil,
			attribute: t.Name(),
		},
		"valid-rules-with-nil-items-ignored": {
			rules:     []*NetworkRule{nil, nil},
			attribute: t.Name(),
		},
		"valid-multiple-rules": {
			rules: []*NetworkRule{
				{
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello=world"}},
					ProtocolPorts: []string{"tcp/443"},
				},
				{
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello=world"}},
					ProtocolPorts: []string{"udp/334"},
				},
			},
			attribute: t.Name(),
		},
		"valid-multiple-rules-concat-object": {
			rules: []*NetworkRule{
				{
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}},
					ProtocolPorts: []string{"tcp/443"},
				},
				{
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hel", "loworld"}},
					ProtocolPorts: []string{"tcp/443"},
				},
			},
			attribute: t.Name(),
		},
		"valid-multiple-rules-concat-ports": {
			rules: []*NetworkRule{
				{
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}},
					ProtocolPorts: []string{"TCP/44", "33"},
				},
				{
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}},
					ProtocolPorts: []string{"TCP/4", "433"},
				},
			},
			attribute: t.Name(),
		},
		"invalid-duplicates": {
			rules: []*NetworkRule{
				{
					Name:          "bad-rule-1",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}, {"heat", "wave"}},
					ProtocolPorts: []string{"TCP/443", "8080"},
				},
				{
					Name:          "bad-rule-2",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}, {"heat", "wave"}},
					ProtocolPorts: []string{"TCP/443", "8080"},
				},
			},
			attribute:   t.Name(),
			expectedErr: makeValidationError(t.Name(), "Duplicate network rules at the following indexes: [1, 2]"),
		},
		"invalid-duplicates-port-case-insensitive": {
			rules: []*NetworkRule{
				{
					Name:          "bad-rule-1",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}},
					ProtocolPorts: []string{"tcp/443"},
				},
				{
					Name:          "bad-rule-2",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}},
					ProtocolPorts: []string{"TCP/443"},
				},
			},
			attribute:   t.Name(),
			expectedErr: makeValidationError(t.Name(), "Duplicate network rules at the following indexes: [1, 2]"),
		},
		"invalid-duplicates-object-different-order": {
			rules: []*NetworkRule{
				{
					Name:          "bad-rule-1",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}, {"heat", "wave"}},
					ProtocolPorts: []string{"TCP/443"},
				},
				{
					Name:          "bad-rule-2",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"world", "hello"}, {"wave", "heat"}},
					ProtocolPorts: []string{"TCP/443"},
				},
			},
			attribute:   t.Name(),
			expectedErr: makeValidationError(t.Name(), "Duplicate network rules at the following indexes: [1, 2]"),
		},
		"invalid-duplicates-object-different-order-2": {
			rules: []*NetworkRule{
				{
					Name:          "bad-rule-1",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}, {"heat", "wave"}},
					ProtocolPorts: []string{"TCP/443"},
				},
				{
					Name:          "bad-rule-2",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"heat", "wave"}, {"hello", "world"}},
					ProtocolPorts: []string{"TCP/443"},
				},
			},
			attribute:   t.Name(),
			expectedErr: makeValidationError(t.Name(), "Duplicate network rules at the following indexes: [1, 2]"),
		},
		"invalid-duplicates-port-different-order": {
			rules: []*NetworkRule{
				{
					Name:          "bad-rule-1",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}, {"heat", "wave"}},
					ProtocolPorts: []string{"TCP/443", "tcp/8080"},
				},
				{
					Name:          "bad-rule-2",
					Action:        NetworkRuleActionAllow,
					Object:        [][]string{{"hello", "world"}, {"heat", "wave"}},
					ProtocolPorts: []string{"tcp/8080", "TCP/443"},
				},
			},
			attribute:   t.Name(),
			expectedErr: makeValidationError(t.Name(), "Duplicate network rules at the following indexes: [1, 2]"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {

			err := ValidateNoDuplicateNetworkRules(tc.attribute, tc.rules)
			if err == nil && tc.expectedErr == (elemental.Error{}) {
				return
			}

			var actual elemental.Error
			if ok := errors.As(err, &actual); !ok {
				t.Fatalf("unexpected error type\nwant: %T\n got: %T", elemental.Error{}, err)
			}
			if !reflect.DeepEqual(actual, tc.expectedErr) {
				t.Fatalf(
					"actual error does not match expected\nwant: '%s'\n got: '%s'",
					tc.expectedErr, actual,
				)
			}
		})
	}
}

func TestValidatePortString(t *testing.T) {
	type args struct {
		attribute string
		portExp   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// simple
		{
			"simple lower port",
			args{
				"port",
				"1",
			},
			false,
		},
		{
			"simple max port",
			args{
				"port",
				"65535",
			},
			false,
		},
		{
			"port set that is not an int",
			args{
				"port",
				"zero",
			},
			true,
		},
		{
			"port set to 0",
			args{
				"port",
				"0",
			},
			false,
		},
		{
			"port set to 65536",
			args{
				"port",
				"65536",
			},
			true,
		},

		// range
		{
			"port range",
			args{
				"port",
				"1:65535",
			},
			false,
		},
		{
			"port range with left bound set to 0",
			args{
				"port",
				"0:65535",
			},
			false,
		},
		{
			"port range with right bound set to 65536",
			args{
				"port",
				"1:65536",
			},
			true,
		},
		{
			"port range with left bound greater than right bound",
			args{
				"port",
				"65535:1",
			},
			true,
		},
		{
			"port range with right bound that is not a int",
			args{
				"port",
				"1:two",
			},
			true,
		},

		// format
		{
			"broken port expression 1",
			args{
				"port",
				"1-3",
			},
			true,
		},
		{
			"broken port expression 2",
			args{
				"port",
				"1-",
			},
			true,
		},
		{
			"broken port expression 3",
			args{
				"port",
				"-1",
			},
			true,
		},
		{
			"broken port expression 4",
			args{
				"port",
				"-",
			},
			true,
		},
		{
			"broken port expression 5",
			args{
				"port",
				"",
			},
			true,
		},
		{
			"broken port expression 6",
			args{
				"port",
				"1:2:3",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePortString(tt.args.attribute, tt.args.portExp); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePortString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateServicePorts(t *testing.T) {
	type args struct {
		servicePorts []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Service Ports
		{
			"valid serviceports",
			args{
				[]string{"tcp/80", "udp/80", "icmp", "ISIS"},
			},
			false,
		},
		{
			"valid serviceports with only any",
			args{
				[]string{"aNy"},
			},
			false,
		},
		{
			"Any with other protoocls",
			args{
				[]string{"tcp/80", "udp/80", "any", "icmp", "ISIS"},
			},
			true,
		},
		{
			"invalid tcp serviceports",
			args{
				[]string{"tcp", "udp/80", "icmp", "ISIS"},
			},
			true,
		},
		{
			"invalid udp serviceports",
			args{
				[]string{"tcp/80", "udp/", "icmp", "ISIS"},
			},
			true,
		},
		{
			"invalid protocol",
			args{
				[]string{"nope"},
			},
			true,
		},
		{
			"invalid serviceports",
			args{
				[]string{"tcp/80/90"},
			},
			true,
		},
		{
			"serviceports with protocol with ports not supported",
			args{
				[]string{"tcp/90:8000", "udp/90:8000", "icmp/9090"},
			},
			true,
		},
		{
			"one serviceport with protocol with ports not supported",
			args{
				[]string{"isis/9090"},
			},
			true,
		},
		{
			"serviceports with valid port range",
			args{
				[]string{"tcp/90:8000", "udp/90:8000"},
			},
			false,
		},
		{
			"serviceports with protocol numbers",
			args{
				[]string{"udp/90:8000", "6/90:8000"},
			},
			true,
		},
		{
			"serviceports with invalid port range",
			args{
				[]string{"tcp/90:8000", "udp/90:800000"},
			},
			true,
		},
		{
			"serviceports with just port range",
			args{
				[]string{"88", "90:8000"},
			},
			true,
		},
		{
			"empty string serviceports",
			args{
				[]string{""},
			},
			true,
		},
		{
			"empty serviceports",
			args{
				[]string{},
			},
			false,
		},
		{
			"nil serviceports",
			args{
				nil,
			},
			false,
		},
		// ICMP Type/codes
		{
			"icmp alone is valid",
			args{
				[]string{"icmp"},
			},
			false,
		},
		{
			"icmp6 alone is valid",
			args{
				[]string{"icmp6"},
			},
			false,
		},
		{
			"icmp/ is an error",
			args{
				[]string{"icmp/"},
			},
			true,
		},
		{
			"icmp6/ is an error",
			args{
				[]string{"icmp6/"},
			},
			true,
		},
		{
			"icmp/0 is valid",
			args{
				[]string{"icmp/0"},
			},
			false,
		},
		{
			"icmp/255 is valid",
			args{
				[]string{"icmp/255"},
			},
			false,
		},
		{
			"icmp/256 is not valid",
			args{
				[]string{"icmp/256"},
			},
			true,
		},
		{
			"icmp/2,3 is not valid",
			args{
				[]string{"icmp/2,3"},
			},
			true,
		},
		{
			"icmp/2/3:10 is valid",
			args{
				[]string{"icmp/2/3:10"},
			},
			false,
		},
		{
			"icmp/2/3: is not valid",
			args{
				[]string{"icmp/2/3:"},
			},
			true,
		},
		{
			"icmp/2/3/5: is not valid",
			args{
				[]string{"icmp/2/3/5"},
			},
			true,
		},
		{
			"icmp/2/3:1 is not valid",
			args{
				[]string{"icmp/2/3:1"},
			},
			true,
		},
		{
			"icmp/2/3:10:20 is not valid",
			args{
				[]string{"icmp/2/3:10:20"},
			},
			true,
		},
		{
			"icmp/2/3:256 is not valid",
			args{
				[]string{"icmp/2/3:256"},
			},
			true,
		},
		{
			"icmp/2/256:256 is not valid",
			args{
				[]string{"icmp/2/256:256"},
			},
			true,
		},
		{
			"icmp/2/2,3:10 is valid",
			args{
				[]string{"icmp/2/2,3:10"},
			},
			false,
		},
		{
			"icmp/2/3,5,20,40 is valid",
			args{
				[]string{"icmp/2/3,5,20,40"},
			},
			false,
		},
		{
			"duplicte ports",
			args{
				[]string{"tcp/443", "TCP/443"},
			},
			true,
		},
		{
			"serviceports with valid 0 port",
			args{
				[]string{"tcp/0", "udp/0"},
			},
			false,
		},
		{
			"serviceports with valid full port range",
			args{
				[]string{"tcp/0:65535", "udp/0:65535"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateServicePorts("serviceports", tt.args.servicePorts); (err != nil) != tt.wantErr {
				t.Errorf("ValidateServicePorts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateServicePort(t *testing.T) {
	type args struct {
		servicePort string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid serviceport",
			args{
				"tcp/80",
			},
			false,
		},
		{
			"valid serviceport any",
			args{
				"aNy",
			},
			false,
		},
		{
			"invalid tcp serviceport",
			args{
				"tcp",
			},
			true,
		},
		{
			"invalid udp serviceport",
			args{
				"udp/",
			},
			true,
		},
		{
			"invalid protocol",
			args{
				"nope",
			},
			true,
		},
		{
			"invalid serviceport",
			args{
				"tcp/80/90",
			},
			true,
		},
		{
			"serviceport with protocol with ports not supported",
			args{
				"icmp/9090",
			},
			true,
		},
		{
			"one serviceport with protocol with ports not supported",
			args{
				"isis/9090",
			},
			true,
		},
		{
			"serviceport with valid port range",
			args{
				"udp/90:8000",
			},
			false,
		}, {
			"serviceport with protocol numbers",
			args{
				"6/90:8000",
			},
			true,
		},
		{
			"serviceport with invalid port range",
			args{
				"udp/90:800000",
			},
			true,
		},
		{
			"serviceport with just port range",
			args{
				"90:8000",
			},
			true,
		},
		{
			"empty serviceport",
			args{
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateServicePort("serviceport", tt.args.servicePort); (err != nil) != tt.wantErr {
				t.Errorf("ValidateServicePort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePortStringList(t *testing.T) {
	type args struct {
		attribute string
		ports     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"correct list",
			args{
				"ports",
				[]string{"443", "555:666", "80"},
			},
			false,
		},
		{
			"incorrect list",
			args{
				"ports",
				[]string{"443", "555:666", "80", "yo"},
			},
			true,
		},
		{
			"empty list",
			args{
				"ports",
				[]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePortStringList(tt.args.attribute, tt.args.ports); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePortStringList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNetworkOrHostnameList(t *testing.T) {
	type args struct {
		attribute string
		networks  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google.com"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google@com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			true,
		},
		{
			"valid entries with just CIDRs",
			args{
				"nets",
				[]string{"0.0.0.0/0", "!10.0.0.0/8"},
			},
			false,
		},
		{
			"valid entries with except condition operator",
			args{
				"nets",
				[]string{"google.com", "0.0.0.0/0", "!10.0.0.0/8"},
			},
			false,
		},
		{
			"invalid entry with only valid NOT operator",
			args{
				"nets",
				[]string{"!10.0.0.0/8"},
			},
			true,
		},
		{
			"invalid entries with without valid CIDR pairing",
			args{
				"nets",
				[]string{"google.com", "!10.0.0.0/8"},
			},
			true,
		},
		{
			"invalid entry with multiple NOT operators",
			args{
				"nets",
				[]string{"!!10.0.0.0/8"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNetworkOrHostnameList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNetworkOrHostnameList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalCIDR(t *testing.T) {
	type args struct {
		attribute string
		cidr      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"cidr",
				"10.0.0.0/8",
			},
			false,
		},

		// valid empty
		{
			"invalid CIDR",
			args{
				"cidr",
				"",
			},
			false,
		},

		// invalid
		{
			"invalid CIDR",
			args{
				"cidr",
				"foo",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalCIDR(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateOptionalCIDR() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateIPAddress(t *testing.T) {
	type args struct {
		attribute string
		network   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"ip address",
				"10.0.0.1",
			},
			false,
		},

		// invalid
		{
			"invalid CIDR",
			args{
				"cidr",
				"foo",
			},
			true,
		},

		// invalid
		{
			"invalid CIDR",
			args{
				"cidr",
				"10.1.1.0/24",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIPAddress(tt.args.attribute, tt.args.network); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateIPAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalIPAddress(t *testing.T) {
	type args struct {
		attribute string
		network   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"ip address",
				"10.0.0.1",
			},
			false,
		},

		// valid empty
		{
			"valid CIDR",
			args{
				"ip address",
				"",
			},
			false,
		},

		// invalid
		{
			"invalid CIDR",
			args{
				"cidr",
				"foo",
			},
			true,
		},

		// invalid
		{
			"invalid CIDR",
			args{
				"cidr",
				"10.1.1.0/24",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalIPAddress(tt.args.attribute, tt.args.network); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateIPAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCIDR(t *testing.T) {
	type args struct {
		attribute string
		cidr      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"cidr",
				"10.0.0.0/8",
			},
			false,
		},
		{
			"valid DNS name",
			args{
				"cidr",
				"google.com",
			},
			true,
		},
		{
			"valid DNS name",
			args{
				"cidr",
				"*.google.com",
			},
			true,
		},

		// invalid CIDR
		{
			"invalid CIDR",
			args{
				"cidr",
				"",
			},
			true,
		},

		// invalid DNn
		{
			"invalid DNS Name",
			args{
				"cidr",
				"google@com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCIDR(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateCIDR() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCIDRList(t *testing.T) {
	type args struct {
		attribute string
		networks  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"10.0.0.0/8"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google.com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCIDRList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCIDRList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProtocol(t *testing.T) {
	type args struct {
		attribute string
		proto     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid string tcp",
			args{
				"proto",
				"tcp",
			},
			false,
		},
		{
			"valid string TCP",
			args{
				"proto",
				"TCP",
			},
			false,
		},
		{
			"valid string udp",
			args{
				"proto",
				"udp",
			},
			false,
		},
		{
			"valid string UDP",
			args{
				"proto",
				"UDP",
			},
			false,
		},
		{
			"valid string IPIP",
			args{
				"proto",
				"IPIP",
			},
			false,
		},
		{
			"valid string ipip",
			args{
				"proto",
				"IPIP",
			},
			false,
		},
		{
			"valid string icmp",
			args{
				"proto",
				"icmp",
			},
			false,
		},
		{
			"valid string ICMP",
			args{
				"proto",
				"ICMP",
			},
			false,
		},
		{
			"valid string number",
			args{
				"proto",
				"30",
			},
			false,
		},

		// invalid
		{
			"invalid string",
			args{
				"proto",
				"NOT",
			},
			true,
		},
		{
			"invalid number 1",
			args{
				"proto",
				"-1",
			},
			true,
		},
		{
			"invalid number 2",
			args{
				"proto",
				"256",
			},
			true,
		},
		{
			"valid string all",
			args{
				"proto",
				"all",
			},
			true,
		},
		{
			"valid string ALL",
			args{
				"proto",
				"ALL",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateProtocol(tt.args.attribute, tt.args.proto); (err != nil) != tt.wantErr {
				t.Errorf("ValidateProtocol() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProtocolList(t *testing.T) {
	type args struct {
		attribute string
		protocols []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"proto",
				[]string{"udp", "UDP", "TCP", "tcp", "3", "43"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"proto",
				[]string{"udp", "UDP", "NOOOOO", "tcp", "3", "43"},
			},
			true,
		},
		{
			"empty list",
			args{
				"proto",
				[]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateProtocolList(tt.args.attribute, tt.args.protocols); (err != nil) != tt.wantErr {
				t.Errorf("ValidateProtocolList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEnforcerProfile(t *testing.T) {
	tests := []struct {
		name            string
		enforcerprofile *EnforcerProfile
		wantErr         bool
	}{
		// Target Networks
		{
			"valid target network",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"0.0.0.0/0"},
			},
			false,
		},
		{
			"valid target networks with except condition operator",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"0.0.0.0/0", "!10.0.0.0/8"},
			},
			false,
		},
		{
			"invalid target network",
			&EnforcerProfile{
				Name:           "Invalid target network",
				TargetNetworks: []string{"invalid"},
			},
			true,
		},
		{
			"invalid target network with multiple NOT operator",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"!!10.0.0.0/8"},
			},
			true,
		},
		{
			"valid target network with NOT operator",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"!10.0.0.0/8"},
			},
			true,
		},
		// Target UDP Networks
		{
			"valid target UDP network",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"0.0.0.0/0", "!224.0.0.0/4"},
			},
			false,
		},
		{
			"Invalid target UDP network, 0.0.0.0/0",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"0.0.0.0/0", "224.0.0.0/4"},
			},
			true,
		},
		{
			"valid target UDP networks with except condition operator",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"0.0.0.0/0", "!10.0.0.0/8", "!224.0.0.0/4"},
			},
			false,
		},
		{
			"valid target UDP networks with implicitly ignoring multicast range",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"0.0.0.0/0", "!128.0.0.0/1"},
			},
			false,
		},
		{
			"Invalid target UDP networks with a list CIDR containing multicast range",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"0.0.0.0/0", "128.0.0.0/1"},
			},
			true,
		},
		{
			"invalid target UDP network",
			&EnforcerProfile{
				Name:              "Invalid target UDP network",
				TargetUDPNetworks: []string{"invalid"},
			},
			true,
		},
		{
			"invalid target UDP network with multiple NOT operator",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"!!10.0.0.0/8"},
			},
			true,
		},
		{
			"valid target UDP network with NOT operator",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"!10.0.0.0/8"},
			},
			true,
		},
		// Excluded Networks
		{
			"valid excluded network",
			&EnforcerProfile{
				Name:             "Valid excluded network",
				ExcludedNetworks: []string{"0.0.0.0/0"},
			},
			false,
		},
		{
			"valid excluded networks with except condition operator",
			&EnforcerProfile{
				Name:             "Valid excluded network",
				ExcludedNetworks: []string{"0.0.0.0/0", "!10.0.0.0/8"},
			},
			false,
		},
		{
			"invalid excluded network",
			&EnforcerProfile{
				Name:             "Invalid excluded network",
				ExcludedNetworks: []string{"invalid"},
			},
			true,
		},
		{
			"invalid excluded network with multiple NOT operator",
			&EnforcerProfile{
				Name:             "Valid excluded network",
				ExcludedNetworks: []string{"!!10.0.0.0/8"},
			},
			true,
		},
		{
			"valid excluded network with NOT operator",
			&EnforcerProfile{
				Name:             "Valid excluded network",
				ExcludedNetworks: []string{"!10.0.0.0/8"},
			},
			true,
		},
		// Trusted CAs
		{
			"valid trusted CA",
			&EnforcerProfile{
				Name: "Valid CAs",
				TrustedCAs: []string{`-----BEGIN CERTIFICATE-----
MIIBPzCB5aADAgECAhB8oW2KA1BfhUmp6B1TkvcRMAoGCCqGSM49BAMCMA8xDTAL
BgNVBAMTBHRlc3QwHhcNMTgxMDAyMjI1MjUwWhcNMjgwODEwMjI1MjUwWjAPMQ0w
CwYDVQQDEwR0ZXN0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsHxO39rdxGtD
8bP5AW2gqTMxgq4w9nyVbbpetS0rtjxwZ5bMgdS4GmqaifjqNGGt+VkK7gFNEyqG
Hq/uTgtJ56MjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
KoZIzj0EAwIDSQAwRgIhAKRcwnrREJ6EnHOsiUnDfuNFxxALw4kV/ZyRxl1CJcS+
AiEA0epxATHNzheAa8ZuiPeNQL6DhoKYz3B+41J2vgVlGZY=
-----END CERTIFICATE-----`},
			},
			false,
		},
		{
			"invalid trusted CA",
			&EnforcerProfile{
				Name:       "Invalid CAs",
				TrustedCAs: []string{"invalid ca"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEnforcerProfile(tt.enforcerprofile); (err != nil) != tt.wantErr {
				t.Errorf("TestVValidateEnforcerProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEnforcerReport_IDFieldCompatibility(t *testing.T) {

	tests := map[string]struct {
		report  *EnforcerReport
		wantErr bool
	}{
		"only the 'ID' field is specified": {
			report: &EnforcerReport{
				ID: "xxx-xxx-xxx-xxx",
			},
			wantErr: false,
		},
		"only the 'enforcerID' field is specified": {
			report: &EnforcerReport{
				EnforcerID: "xxx-xxx-xxx-xxx",
			},
			wantErr: false,
		},
		"both the 'enforcerID' and 'ID' fields have been left omitted": {
			report: &EnforcerReport{
				EnforcerID: "",
				ID:         "",
			},
			wantErr: true,
		},
		"both the 'enforcerID' and 'ID' fields have been provided": {
			report: &EnforcerReport{
				EnforcerID: "xxx-xxx-xxx-xxx",
				ID:         "xxx-xxx-xxx-xxx",
			},
			wantErr: true,
		},
	}

	for description, scenario := range tests {
		t.Run(description, func(t *testing.T) {
			var err error
			if err = ValidateEnforcerReport(scenario.report); (err != nil) != scenario.wantErr {
				t.Errorf("TestValidateEnforcerReport_IDFieldCompatibility() error = %s, expected an error? %t", err, scenario.wantErr)
			}

			// verify that ONLY the `EnforcerID` field is populated so the backend processors do not need to worry about
			// dealing with the 'ID' field being set
			if err == nil {
				if scenario.report.ID != "" {
					t.Errorf("Validation passed, but the report's 'ID' field was NOT reset. It contained the value: %s", scenario.report.ID)
				}

				if scenario.report.EnforcerID == "" {
					t.Error("Validation passed, but the report's 'enforcerID' field was empty")
				}
			}
		})
	}
}

func TestValidateServiceEntity(t *testing.T) {
	type args struct {
		service *Service
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// OIDC
		{
			"valid OIDC service",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCProviderURL:   "https://url.com",
					OIDCClientID:      "client id",
					OIDCClientSecret:  "client secret",
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"OIDC service with missing OIDCProviderURL",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCClientID:      "client id",
					OIDCClientSecret:  "client secret",
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"OIDC service with bad OIDCProviderURL - no https",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					OIDCProviderURL:   "http://url.com",
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCClientID:      "client id",
					OIDCClientSecret:  "client secret",
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"OIDC service with bad OIDCProviderURL - no https",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					OIDCProviderURL:   "http://!@#!@#/1231",
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCClientID:      "client id",
					OIDCClientSecret:  "client secret",
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"OIDC service with missing OIDCClientID",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCProviderURL:   "http://url.com",
					OIDCClientSecret:  "client secret",
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"OIDC service with missing OIDCClientSecret",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCProviderURL:   "http://url.com",
					OIDCClientID:      "client id",
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},

		// JWT
		{
			"valid JWT service",
			args{
				&Service{
					Hosts:                 []string{"myservice.com"},
					AuthorizationType:     ServiceAuthorizationTypeJWT,
					JWTSigningCertificate: "---pem---",
					Port:                  80,
					ExposedPort:           80,
				},
			},
			false,
		},
		{
			"JWT service with missing JWTSigningCertificate",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeJWT,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},

		// TLS
		{
			"valid service with type TLSType set to Aporeto",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeAporeto,
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"valid service with type TLSType set to LetsEcncrypt",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeLetsEncrypt,
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"valid service with type TLSType set to None",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"valid service with type TLSType set to External",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeExternal,
					TLSCertificate:    "---cert---",
					TLSCertificateKey: "---key---",
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"service with type TLSType set to External with missing cert and key",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeExternal,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service with type TLSType set to External with missing cert",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeExternal,
					TLSCertificateKey: "---key---",
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},

		// Hosts an IP
		{
			"service missing hosts and ips",
			args{
				&Service{
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service with invalid hosts ",
			args{
				&Service{
					Hosts:             []string{"@#$@#$.2234242"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service with both hosts and IPs is valid",
			args{
				&Service{
					Hosts:             []string{"foo.com"},
					IPs:               []string{"10.1.0.0/16"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"service with overlapping hosts",
			args{
				&Service{
					Hosts:             []string{"foo.com", "foo.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service with bad IPs",
			args{
				&Service{
					IPs:               []string{"1234.2.3.3"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service with bad subnets",
			args{
				&Service{
					IPs:               []string{"10.1.1.0/48"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service with overlapping subnets",
			args{
				&Service{
					IPs:               []string{"10.1.0.0/24", "10.1.0.0/16"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service with multiple IPs and multiple error case",
			args{
				&Service{
					IPs:               []string{"10.11.0.1234/24", "10.1.0.0/16", "10.1.0.3/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service with multiple IPs positive case",
			args{
				&Service{
					IPs:               []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		// Listening Port
		{
			"external service must not have a port",
			args{
				&Service{
					IPs:               []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					External:          true,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"external service must not have a port otherwise it fails",
			args{
				&Service{
					IPs:               []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					External:          true,
					Port:              80,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service must have port defined otherwise it fails",
			args{
				&Service{
					IPs:               []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					External:          false,
					ExposedPort:       80,
				},
			},
			true,
		},
		{
			"service must have port defined",
			args{
				&Service{
					IPs:               []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					External:          false,
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"service with exposed port and public port matching fails",
			args{
				&Service{
					IPs:                   []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType:     ServiceAuthorizationTypeMTLS,
					TLSType:               ServiceTLSTypeNone,
					External:              false,
					ExposedPort:           443,
					PublicApplicationPort: 443,
				},
			},
			true,
		},
		// publicApplicationPort
		{
			"external service must not have a publicApplicationPort",
			args{
				&Service{
					IPs:               []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					External:          true,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"external service must not have a publicApplicationPort otherwise it fails",
			args{
				&Service{
					IPs:                   []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType:     ServiceAuthorizationTypeMTLS,
					TLSType:               ServiceTLSTypeNone,
					External:              true,
					PublicApplicationPort: 443,
					ExposedPort:           80,
				},
			},
			true,
		},
		{
			"service might not have a publicApplicationPort",
			args{
				&Service{
					IPs:               []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
					External:          false,
					Port:              80,
					ExposedPort:       80,
				},
			},
			false,
		},
		{
			"service can have a publicApplicationPort",
			args{
				&Service{
					IPs:                   []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType:     ServiceAuthorizationTypeMTLS,
					TLSType:               ServiceTLSTypeNone,
					External:              false,
					Port:                  80,
					PublicApplicationPort: 443,
					ExposedPort:           80,
				},
			},
			false,
		},
		{
			"service with proxy protocol enabled and subnet",
			args{
				&Service{
					IPs:                   []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					ProxyProtocolSubnets:  []string{"1.1.1.0/24"},
					ProxyProtocolEnabled:  true,
					AuthorizationType:     ServiceAuthorizationTypeMTLS,
					TLSType:               ServiceTLSTypeNone,
					External:              false,
					Port:                  80,
					PublicApplicationPort: 443,
					ExposedPort:           80,
				},
			},
			false,
		},
		{
			"service with proxy protocol enabled and no subnet",
			args{
				&Service{
					IPs:                   []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					ProxyProtocolSubnets:  []string{},
					ProxyProtocolEnabled:  true,
					AuthorizationType:     ServiceAuthorizationTypeMTLS,
					TLSType:               ServiceTLSTypeNone,
					External:              false,
					Port:                  80,
					PublicApplicationPort: 443,
					ExposedPort:           80,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateServiceEntity(tt.args.service); (err != nil) != tt.wantErr {
				t.Errorf("ValidateServiceEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateHTTPMethods(t *testing.T) {
	type args struct {
		attribute string
		methods   []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"empty list of methods",
			args{
				attribute: "methods",
				methods:   []string{},
			},
			false,
		},
		{
			"valid list of methods",
			args{
				attribute: "methods",
				methods:   []string{"POST"},
			},
			false,
		},
		{
			"invalid list of methods",
			args{
				attribute: "methods",
				methods:   []string{"BIDULE"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHTTPMethods(tt.args.attribute, tt.args.methods); (err != nil) != tt.wantErr {
				t.Errorf("ValidateHTTPMethods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateHTTPSURL(t *testing.T) {
	type args struct {
		attribute string
		address   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"empty url",
			args{
				attribute: "endpoint",
				address:   "",
			},
			true,
		},
		{
			"valid url",
			args{
				attribute: "endpoint",
				address:   "https://aporeto.com/",
			},
			false,
		},
		{
			"invalid url",
			args{
				attribute: "endpoint",
				address:   "htps:/a",
			},
			true,
		},
		{
			"non https url",
			args{
				attribute: "endpoint",
				address:   "http://hello.com/",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHTTPSURL(tt.args.attribute, tt.args.address); (err != nil) != tt.wantErr {
				t.Errorf("ValidateHTTPSURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalNetworkOrHostnameList(t *testing.T) {
	type args struct {
		attribute string
		networks  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google.com", "10.1.1.2"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google@com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalNetworkOrHostnameList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalNetworkOrHostnameList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAutomation(t *testing.T) {
	testCases := map[string]struct {
		automation    *Automation
		expectedError error
	}{
		"should not return an error if trigger type is not webhook and multiple actions have been defined": {
			automation: &Automation{
				Condition: `
						function when(api, params) {
							return {
								continue: true,
							};
						}
			 	`,
				Trigger: AutomationTriggerRemoteCall,
				Actions: []string{
					"Action 1",
					"Action 2",
				},
			},
			expectedError: nil,
		},
		"should not return an error if trigger type is webhook and one action has been defined": {
			automation: &Automation{
				Trigger: AutomationTriggerWebhook,
				Actions: []string{
					"Action 1",
				},
			},
			expectedError: nil,
		},
		"should return an error if trigger type is set to webhook and more than one action has been defined": {
			automation: &Automation{
				Trigger: AutomationTriggerWebhook,
				Actions: []string{
					"Action 1",
					"Action 2",
				},
			},
			expectedError: makeValidationError("actions", fmt.Sprintf("Only one action can be defined if trigger type is set to %q", AutomationTriggerWebhook)),
		},
		"should return an error if trigger type is set to webhook and no actions have been defined": {
			automation: &Automation{
				Trigger: AutomationTriggerWebhook,
				Actions: nil,
			},
			expectedError: makeValidationError("actions", fmt.Sprintf("Exactly one action must be defined if trigger type is set to %q", AutomationTriggerWebhook)),
		},
	}

	for scenario, tc := range testCases {
		t.Run(scenario, func(t *testing.T) {
			err := ValidateAutomation(tc.automation)
			switch {
			case err != nil && tc.expectedError != nil:
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("\n"+
						"actual error: %+v\n"+
						"did not equal\n"+
						"expected error: %+v", err, tc.expectedError)
				}
			case err != nil && tc.expectedError == nil:
				t.Fatalf("did not expect to get an error, but received: %+v", err)
			case err == nil && tc.expectedError != nil:
				t.Fatalf("expected to get an error, but got nothing")
			}
		})
	}
}

func TestValidateHostService(t *testing.T) {
	tests := []struct {
		name    string
		hs      *HostService
		wantErr bool
	}{
		{
			"valid service",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"tcp/80"},
			},
			false,
		},
		{
			"valid service name < 128",
			&HostService{
				Name:            "asdasdsakljfslahfjashjkdlaksfjaksjfvfsadlkjkljsad",
				HostModeEnabled: false,
				Services:        []string{"tcp/80"},
			},
			false,
		},
		{
			"valid service name > 128",
			&HostService{
				Name:            "asdasdsakljfslahfjashjkdlaksfsdhkjfhsdjkhfsiudfhsdjkfnsdkjfn1342141397asldjasklfjakslfj879797kjasdhjaksjfvfsadlkjkljsaddasfhsajlhdkajshfjksahfjksahdjksahdkjashdkasdhkjsa",
				HostModeEnabled: false,
				Services:        []string{"tcp/80"},
			},
			true,
		},
		{
			"invalid name",
			&HostService{
				Name:            "jboss@!#$!$!",
				HostModeEnabled: false,
				Services:        []string{},
			},
			true,
		},
		{
			"empty name",
			&HostService{
				Name:            "",
				HostModeEnabled: false,
				Services:        []string{},
			},
			true,
		},
		{
			"invalid services",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"foo / 123"},
			},
			true,
		},
		{
			"tcp port overlaps",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"tcp/80:100, tcp/90"},
			},
			true,
		},
		{
			"udp port overlaps",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"udp/80:100, udp/90"},
			},
			true,
		},
		{
			"non overlap with tcp and udp ports",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"udp/90", "tcp/90"},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHostServices(tt.hs); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateHostServicesList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isFQDN(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
	}{
		{
			"valid fqdn",
			args{
				"www.google.com",
			},
			true,
		},
		{
			"xip",
			args{
				"1.2.3.4.xip.io",
			},
			true,
		},
		{
			"trailing dot",
			args{
				"www.google.com.",
			},
			true,
		},
		{
			"bad one",
			args{
				"@#$@#$@.#$@#.#@$2",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isFQDN(tt.args.name) != tt.wantValue {
				t.Errorf("isFQDN() failed at test: %s", tt.name)
			}
		})
	}
}

func TestValidatePEM(t *testing.T) {
	type args struct {
		attribute string
		pemdata   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"nothing set",
			args{
				"pem",
				``,
			},
			false,
		},
		{
			"valid single PEM",
			args{
				"pem",
				`-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5loNahC19hQtExCi64fmI3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----`,
			},
			false,
		},
		{
			"valid single PEM",
			args{
				"pem",
				`-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5loNahC19hQtExCi64fmI3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5loNahC19hQtExCi64fmI3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----
`,
			},
			false,
		},
		{
			"invalid single PEM",
			args{
				"pem",
				`-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoT ----NOT PEM---- I3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----`,
			},
			true,
		},
		{
			"valid single PEM",
			args{
				"pem",
				`-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5loNahC19hQtExCi64fmI3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5     ----NOT PEM----   crEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----
`,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePEM(tt.args.attribute, tt.args.pemdata); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePEM() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// simple
		{
			"simple tag a=a",
			args{
				"a=a",
			},
			false,
		},
		{
			"value is string",
			args{
				"name=value",
			},
			false,
		},
		{
			"value is number",
			args{
				"name=123.12",
			},
			false,
		},
		{
			"value is has space",
			args{
				"name=test 1",
			},
			false,
		},
		{
			"tag value is utf8 character",
			args{
				"name=utf8-_!@#%&\" (*)+.,/$!:;<>=?{}~",
			},
			false,
		},
		{
			"key contains hyphen and underscore",
			args{
				"internal_name:demo-1234=aporeliable_arvind_@",
			},
			false,
		},
		{
			"tag contains 2 equals",
			args{
				"a=a=b",
			},
			false,
		},
		{
			"tag contains 2 touching equals",
			args{
				"a==b",
			},
			false,
		},
		{
			"complex key",
			args{
				`utf8-_!@#%&"(*)+.,/$!:;<>=?{}~=utf8-_!@#%&" (*)+.,/$!:;<>=?{}~`,
			},
			false,
		},

		// Error
		{
			"just a word",
			args{
				"justaword",
			},
			true,
		},
		{
			"key contains spaces",
			args{
				"the key=value",
			},
			true,
		},
		{
			"key starts with =",
			args{
				"=thekey=value",
			},
			true,
		},

		// tag size
		{
			"tags is exactly 1024 bytes",
			args{
				"l=ah6ooPaelaizohkee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			true,
		},
		{
			"tags is exactly 1026 bytes",
			args{
				"l=ah6ooPaelaizohakee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			true,
		},
		{
			"tags is exacty 1023 bytes",
			args{
				"l=ahooPaelaizohkee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateTag("tag", tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateYAMLString(t *testing.T) {
	type args struct {
		attribute string
		data      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid yaml",
			args{
				"prop",
				`
hello: world
things:
- stuff
- trucs
`,
			},
			false,
		},
		{
			"invalid yaml",
			args{
				"prop",
				`not a yaml`,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateYAMLString(tt.args.attribute, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateYAMLString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSAMLProvider(t *testing.T) {
	type args struct {
		provider *SAMLProvider
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"with nothing",
			args{&SAMLProvider{}},
			true,
		},
		{
			"with IDP metadata",
			args{&SAMLProvider{
				IDPMetadata: "m",
			}},
			false,
		},
		{
			"with every other field",
			args{&SAMLProvider{
				IDPURL:         "IDPURL",
				IDPIssuer:      "IDPIssuer",
				IDPCertificate: "IDPCertificate",
			}},
			false,
		},
		{
			"with every other field but IDPURL",
			args{&SAMLProvider{
				IDPIssuer:      "IDPIssuer",
				IDPCertificate: "IDPCertificate",
			}},
			true,
		},
		{
			"with every other field but IDPIssuer",
			args{&SAMLProvider{
				IDPURL:         "IDPURL",
				IDPCertificate: "IDPCertificate",
			}},
			true,
		},
		{
			"with every other field but IDPCertificate",
			args{&SAMLProvider{
				IDPURL:    "IDPURL",
				IDPIssuer: "IDPIssuer",
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateSAMLProvider(tt.args.provider); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSAMLProvider() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAPIAuthorizationPolicySubject(t *testing.T) {
	type args struct {
		attribute string
		subject   [][]string
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		wantErrString string
	}{
		{
			"valid subject",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			false,
			"",
		},
		{
			"missing realm claim",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a"},
					{"@auth:claim=a", "@auth:claim=b"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject line 2 must contain the '@auth:realm' key",
		},
		{
			"2 realm claims",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a", "@auth:realm=vince"},
					{"@auth:claim=a", "@auth:claim=b"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject line 1 must contain only one '@auth:realm' key",
		},
		{
			"single claim line",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a"},
					{"@auth:realm=certificate"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject and line should contain at least 2 claims",
		},
		{
			"missing auth prefix claim",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a"},
					{"@auth:claim=a", "@auth:claim=b", "not:good"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject claims 'not:good' on line 2 must be prefixed by '@auth:'",
		},
		{
			"oidc correct",
			args{
				"subject",
				[][]string{
					{"@auth:realm=oidc", "@auth:claim=a", "@auth:namespace=/a/b"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			false,
			"",
		},
		{
			"oidc missing namespace",
			args{
				"subject",
				[][]string{
					{"@auth:realm=oidc", "@auth:claim=a"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: The realm OIDC mandates to add the '@auth:namespace' key to prevent potential security side effects",
		},
		{
			"saml correct",
			args{
				"subject",
				[][]string{
					{"@auth:realm=saml", "@auth:claim=a", "@auth:namespace=/a/b"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			false,
			"",
		},
		{
			"saml missing namespace",
			args{
				"subject",
				[][]string{
					{"@auth:realm=saml", "@auth:claim=a"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: The realm SAML mandates to add the '@auth:namespace' key to prevent potential security side effects",
		},
		{
			"broken tag with no equal",
			args{
				"subject",
				[][]string{
					{"@auth:realm=saml", "@auth:claim"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject claims '@auth:claim' on line 1 is an invalid tag",
		},
		{
			"broken tag with no value",
			args{
				"subject",
				[][]string{
					{"@auth:realm=saml", "@auth:claim="},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject claims '@auth:claim=' on line 1 has no value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAPIAuthorizationPolicySubject(tt.args.attribute, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAPIAuthorizationPolicySubject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err.Error() != tt.wantErrString {
				t.Errorf("ValidateAPIAuthorizationPolicySubject() error = '%v', wantErrString = '%v'", err, tt.wantErrString)
			}
		})
	}
}

func TestValidateWriteOperations(t *testing.T) {
	type args struct {
		attribute  string
		operations []string
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		wantErrString string
	}{
		{
			"nil",
			args{
				"thing",
				nil,
			},
			false,
			"",
		},
		{
			"empty",
			args{
				"thing",
				[]string{},
			},
			false,
			"",
		},
		{
			"valid list",
			args{
				"thing",
				[]string{
					"create",
					"update",
					"delete",
				},
			},
			false,
			"",
		},
		{
			"invalid list",
			args{
				"thing",
				[]string{
					"Create",
					"Update",
					"NotDelete",
				},
			},
			true,
			"error 422 (gaia): Validation Error: Invalid operation 'Create': must be 'create', 'update' or 'delete'",
		},
		{
			"duplicate create",
			args{
				"thing",
				[]string{
					"create",
					"create",
					"update",
					"delete",
				},
			},
			true,
			"error 422 (gaia): Validation Error: Must not contain the same operation multiple times",
		},
		{
			"duplicate update",
			args{
				"thing",
				[]string{
					"create",
					"update",
					"update",
					"delete",
				},
			},
			true,
			"error 422 (gaia): Validation Error: Must not contain the same operation multiple times",
		},
		{
			"duplicate delete",
			args{
				"thing",
				[]string{
					"create",
					"update",
					"delete",
					"delete",
				},
			},
			true,
			"error 422 (gaia): Validation Error: Must not contain the same operation multiple times",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateWriteOperations(tt.args.attribute, tt.args.operations)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateWriteOperations() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err.Error() != tt.wantErrString {
				t.Errorf("ValidateWriteOperations() error = '%v', wantErrString = '%v'", err, tt.wantErrString)
			}
		})
	}
}

func TestValidateIdentity(t *testing.T) {
	type args struct {
		attribute string
		identity  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				"attr",
				"processingunit",
			},
			false,
		},
		{
			"invalid",
			args{
				"attr",
				"yo",
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIdentity(tt.args.attribute, tt.args.identity); (err != nil) != tt.wantErr {
				t.Errorf("ValidateIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProcessingUnitPolicy(t *testing.T) {
	type args struct {
		policy *ProcessingUnitPolicy
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Empty processing unit policy",
			args{
				NewProcessingUnitPolicy(),
			},
			true,
		},
		{
			"processing unit policy with Action and DatapathType set to Default",
			args{
				&ProcessingUnitPolicy{
					Action:       ProcessingUnitPolicyActionDefault,
					DatapathType: ProcessingUnitPolicyDatapathTypeDefault,
				},
			},
			true,
		},
		{
			"processing unit policy with Action set to Default and DatapathType to something else",
			args{
				&ProcessingUnitPolicy{
					Action:       ProcessingUnitPolicyActionDefault,
					DatapathType: ProcessingUnitPolicyDatapathTypeAporeto,
				},
			},
			false,
		},
		{
			"processing unit policy with DatapathType set to Default and Action to something else",
			args{
				&ProcessingUnitPolicy{
					Action:       ProcessingUnitPolicyActionEnforce,
					DatapathType: ProcessingUnitPolicyDatapathTypeDefault,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateProcessingUnitPolicy(tt.args.policy); (err != nil) != tt.wantErr {
				t.Errorf("ValidateProcessingUnitPolicy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateStringListNotEmpty(t *testing.T) {
	type args struct {
		attribute string
		slice     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"nil",
			args{
				"attr",
				nil,
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				[]string{},
			},
			true,
		},
		{
			"set",
			args{
				"attr",
				[]string{"a=a"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateStringListNotEmpty(tt.args.attribute, tt.args.slice); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStringListNotEmpty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateHookPolicy(t *testing.T) {
	type args struct {
		policy *HookPolicy
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Valid cases
		{
			"Only endpoint specified",
			args{
				&HookPolicy{
					EndpointType: HookPolicyEndpointTypeURL,
					Endpoint:     "https://www.google.com",
				},
			},
			false,
		},
		{
			"Simple selectors specified",
			args{
				&HookPolicy{
					EndpointType: HookPolicyEndpointTypeAutomation,
					Selectors: [][]string{
						{"automation=test"},
					},
				},
			},
			false,
		},
		// Error cases
		{
			"Endpoint trigger with automation selectors",
			args{
				&HookPolicy{
					EndpointType: HookPolicyEndpointTypeURL,
					Selectors: [][]string{
						{"automation=test"},
						{"automation=anotherone"},
					},
				},
			},
			true,
		},
		{
			"Automation Selector trigger with endpoint",
			args{
				&HookPolicy{
					EndpointType: HookPolicyEndpointTypeAutomation,
					Endpoint:     "http://www.google.com",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHookPolicy(tt.args.policy); (err != nil) != tt.wantErr {
				t.Errorf("ValidateHookPolicy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCA(t *testing.T) {

	var (
		notCA = `-----BEGIN CERTIFICATE-----
MIIBKzCB0qADAgECAhBaMjFSVhiDqqG1HMQaPNB3MAoGCCqGSM49BAMCMA8xDTAL
BgNVBAMTBHRvdG8wHhcNMjAwNDEwMTczMjEyWhcNMzAwMjE3MTczMjEyWjAPMQ0w
CwYDVQQDEwR0b3RvMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESFKAZYknsQ/P
u847pk7vkAWkh0hlWCGm8ycTqJtREmfcl3MUYarXjJJLxrg5WmM+xRJuuiTudyLR
MeJFDIVRkqMQMA4wDAYDVR0TAQH/BAIwADAKBggqhkjOPQQDAgNIADBFAiEAr5Sn
5NMpnLPc1yy1A3oyX+iHj2jjPRc/FkTK1ozPA04CIAvBYS8+5Hq+HzmryIP0A57k
dGcCFgI8uzkLKtXLq2n0
-----END CERTIFICATE-----`
		CA = `-----BEGIN CERTIFICATE-----
MIIBPjCB5aADAgECAhAumbVWnomSelU1zBRe9qp4MAoGCCqGSM49BAMCMA8xDTAL
BgNVBAMTBHRvdG8wHhcNMjAwNDEwMTczMjU0WhcNMzAwMjE3MTczMjU0WjAPMQ0w
CwYDVQQDEwR0b3RvMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEUWqBTwhWMrtr
vv2i+nyiLpwZPV+BfJs+W+qpYLxX6A+WqFTNvu/frdm6dfKLc022QIMzcJwsM+Yz
oBYZoyf02KMjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
KoZIzj0EAwIDSAAwRQIhAILlr7jJBS3syxM34zLlwvaoRiCFWnVPHNQeSJtk0h7l
AiBZg9mafabskWu9ekgZm50rKkPeqF94tY6R7tuZBUdcVQ==
-----END CERTIFICATE-----`
		brokenCA = `-----BEGIN CERTIFICATE-----
MDIBPjCB5aADAgECAhAumbVWnomSelU1zBRe9qp4MAoGCCqGSM49BAMCMA8xDTAL
BgNVBAMTBHRvdG8wHhcNMjAwNDEwMTczMjU0WhcNMzAwMjE3MTczMjU0WjAPMQ0w
CwYDVQQDEwR0b3RvMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEUWqBTwhWMrtr
vv2i+nyiLpwZPV+BfJs+W+qpYLxX6A+WqFTNvu/frdm6dfKLc022QIMzcJwsM+Yz
oBYZoyf02KMjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
KoZIzj0EAwIDSAAwRQIhAILlr7jJBS3syxM34zLlwvaoRiCFWnVPHNQeSFtk0h7l
AiBZg9mafabskWu9ekgZm50rKkPeqF94tY6R7tuZBUdcVQ==
-----END CERTIFICATE-----`
	)
	type args struct {
		attribute string
		pemdata   string
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		wantErrString string
	}{
		{
			"empty",
			args{
				"attr",
				``,
			},
			false, "",
		},
		{
			"invalid pem",
			args{
				"attr",
				`oh no`,
			},
			true, "error 422 (gaia): Validation Error: Unable to decode PEM",
		},
		{
			"valid pem / invalid CA",
			args{
				"attr",
				brokenCA,
			},
			true, `error 422 (gaia): Validation Error: Unable to parse x509 certificate:`,
		},
		{
			"valid pem but not CA",
			args{
				"attr",
				notCA,
			},
			true, "error 422 (gaia): Validation Error: Given x509 certificate is not a CA",
		},

		{
			"everything is fine",
			args{
				"attr",
				CA,
			},
			false, "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCA(tt.args.attribute, tt.args.pemdata)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCA() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && !strings.HasPrefix(err.Error(), tt.wantErrString) {
				t.Errorf("ValidateCA() error = '%v', wantErrString = '%v'", err, tt.wantErrString)
			}
		})
	}
}

func TestValidateUIParameters(t *testing.T) {
	type args struct {
		p *UIParameter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Valid cases
		{
			"Test with List and Subtype string",
			args{
				p: &UIParameter{
					Type:    UIParameterTypeList,
					Subtype: "String",
				},
			},
			false,
		},
		{
			"Test with List and Subtype integer",
			args{
				p: &UIParameter{
					Type:    UIParameterTypeList,
					Subtype: "Integer",
				},
			},
			false,
		},
		{
			"Test with List and Subtype endpoint",
			args{
				p: &UIParameter{
					Type:    UIParameterTypeList,
					Subtype: "Endpoint",
				},
			},
			false,
		},
		{
			"Test with List and Subtype automation action",
			args{
				p: &UIParameter{
					Type:    UIParameterTypeList,
					Subtype: "AutomationAction",
				},
			},
			false,
		},
		{
			"Test with List and Subtype enum",
			args{
				p: &UIParameter{
					Type:    UIParameterTypeList,
					Subtype: "Enum",
				},
			},
			false,
		},
		// Error cases
		{
			"Test with List and empty Subtype",
			args{
				p: &UIParameter{
					Type:    UIParameterTypeList,
					Subtype: "",
				},
			},
			true,
		},
		{
			"Test with List and unknown Subtype",
			args{
				p: &UIParameter{
					Type:    UIParameterTypeList,
					Subtype: "Unknown",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateUIParameters(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUIParameters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTimeDuration(t *testing.T) {
	type args struct {
		attribute string
		duration  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				"attr",
				"1m",
			},
			false,
		},
		{
			"invalid",
			args{
				"attr",
				"dog",
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateTimeDuration(tt.args.attribute, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTimeDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalTimeDuration(t *testing.T) {
	type args struct {
		attribute string
		duration  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				"attr",
				"1m",
			},
			false,
		},
		{
			"invalid",
			args{
				"attr",
				"dog",
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				"",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalTimeDuration(tt.args.attribute, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalTimeDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSemVer(t *testing.T) {
	type args struct {
		attribute string
		data      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{

			"valid semver with prefix",
			args{
				"upgradeVersion",
				"v1.2.3",
			},
			false,
		},
		{
			"valid semver without prefix",
			args{
				"upgradeVersion",
				"1.2.3",
			},
			false,
		},
		{
			"empty semver",
			args{
				"upgradeVersion",
				"",
			},
			false,
		},
		{
			"missing patch in version",
			args{
				"upgradeVersion",
				"1.2",
			},
			true,
		},
		{
			"invalid semver",
			args{
				"upgradeVersion",
				"pizza",
			},
			true,
		},
		{
			"aporeto release semver",
			args{
				"upgradeVersion",
				"release-4.0.0",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateSemVer(tt.args.attribute, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSemVer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCachedFlowReport(t *testing.T) {
	type args struct {
		cachedFlowReport *CachedFlowReport
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"invalid-flag",
			args{
				cachedFlowReport: &CachedFlowReport{
					IsLocalDestinationID: false,
					DestinationID:        "5f863169e767b10001b0bb1a",
					DestinationType:      CachedFlowReportDestinationTypeProcessingUnit,
					IsLocalSourceID:      false,
					SourceID:             "5f863169e767b10001b0bb1b",
					SourceType:           CachedFlowReportSourceTypeProcessingUnit,
				},
			},
			true,
		},
		{
			"invalid-type-1",
			args{
				cachedFlowReport: &CachedFlowReport{
					IsLocalDestinationID: true,
					DestinationID:        "L-5f863169e767b10001b0bb1e-1234567890ab-76543210",
					DestinationType:      CachedFlowReportDestinationTypeExternalNetwork,
					IsLocalSourceID:      false,
					SourceID:             "5f863169e767b10001b0bb1b",
					SourceType:           CachedFlowReportSourceTypeProcessingUnit,
				},
			},
			true,
		},
		{
			"invalid-type-2",
			args{
				cachedFlowReport: &CachedFlowReport{
					IsLocalDestinationID: false,
					DestinationID:        "5f863169e767b10001b0bb1a",
					DestinationType:      CachedFlowReportDestinationTypeProcessingUnit,
					IsLocalSourceID:      true,
					SourceID:             "L-5f863169e767b10001b0bb1e-1234567890ab-76543210",
					SourceType:           CachedFlowReportSourceTypeExternalNetwork,
				},
			},
			true,
		},
		{
			"valid-1",
			args{
				cachedFlowReport: &CachedFlowReport{
					IsLocalDestinationID: true,
					DestinationID:        "L-5f863169e767b10001b0bb1e-1234567890ab-76543210",
					DestinationType:      CachedFlowReportDestinationTypeProcessingUnit,
					IsLocalSourceID:      false,
					SourceID:             "5f863169e767b10001b0bb1b",
					SourceType:           CachedFlowReportSourceTypeProcessingUnit,
				},
			},
			false,
		},
		{
			"valid-2",
			args{
				cachedFlowReport: &CachedFlowReport{
					IsLocalDestinationID: false,
					DestinationID:        "5f863169e767b10001b0bb1a",
					DestinationType:      CachedFlowReportDestinationTypeProcessingUnit,
					IsLocalSourceID:      true,
					SourceID:             "L-5f863169e767b10001b0bb1e-1234567890ab-76543210",
					SourceType:           CachedFlowReportSourceTypeProcessingUnit,
				},
			},
			false,
		},
		{
			"valid-3",
			args{
				cachedFlowReport: &CachedFlowReport{
					IsLocalDestinationID: true,
					DestinationID:        "L-5f863169e767b10001b0bb1e-1234567890ab-76543210",
					DestinationType:      CachedFlowReportDestinationTypeProcessingUnit,
					IsLocalSourceID:      true,
					SourceID:             "L-5f863169e767b10001b0bb1e-1234567890ab-76543210",
					SourceType:           CachedFlowReportSourceTypeProcessingUnit,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCachedFlowReport(tt.args.cachedFlowReport); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCachedFlowReport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDNSLookupReport(t *testing.T) {

	tests := map[string]struct {
		report  *DNSLookupReport
		wantErr bool
	}{
		"only the 'namespace' field is specified": {
			report: &DNSLookupReport{
				Namespace: "my/namespace",
			},
			wantErr: false,
		},
		"only the 'processingUnitNamespace' field is specified": {
			report: &DNSLookupReport{
				ProcessingUnitNamespace: "my/namespace",
			},
			wantErr: false,
		},
		"both the 'processingUnitNamespace' and 'namespace' fields have been left omitted": {
			report: &DNSLookupReport{
				Namespace:               "",
				ProcessingUnitNamespace: "",
			},
			wantErr: true,
		},
		"both the 'processingUnitNamespace' and 'namespace' fields have been provided": {
			report: &DNSLookupReport{
				Namespace:               "my/namespace",
				ProcessingUnitNamespace: "my/namespace",
			},
			wantErr: true,
		},
	}

	for description, scenario := range tests {
		t.Run(description, func(t *testing.T) {
			var err error
			if err = ValidateDNSLookupReport(scenario.report); (err != nil) != scenario.wantErr {
				t.Errorf("TestValidateDNSLookupReport() error = %s, expected an error? %t", err, scenario.wantErr)
			}

			// verify that ONLY the `namespace` field is populated so the backend processors do not need to worry about
			// dealing with the 'processingUnitNamespace' field being set
			if err == nil {
				if scenario.report.ProcessingUnitNamespace != "" {
					t.Errorf("Validation passed, but the report's 'processingUnitNamespace' field was NOT reset. It contained the value: %s", scenario.report.ID)
				}

				if scenario.report.Namespace == "" {
					t.Error("Validation passed, but the report's 'namespace' field was empty")
				}
			}
		})
	}
}

func TestValidateConnectionExceptionReport(t *testing.T) {

	tests := map[string]struct {
		report  *ConnectionExceptionReport
		wantErr bool
	}{
		"only the 'namespace' field is specified": {
			report: &ConnectionExceptionReport{
				Namespace: "my/namespace",
			},
			wantErr: false,
		},
		"only the 'processingUnitNamespace' field is specified": {
			report: &ConnectionExceptionReport{
				ProcessingUnitNamespace: "my/namespace",
			},
			wantErr: false,
		},
		"both the 'processingUnitNamespace' and 'namespace' fields have been left omitted": {
			report: &ConnectionExceptionReport{
				Namespace:               "",
				ProcessingUnitNamespace: "",
			},
			wantErr: true,
		},
		"both the 'processingUnitNamespace' and 'namespace' fields have been provided": {
			report: &ConnectionExceptionReport{
				Namespace:               "my/namespace",
				ProcessingUnitNamespace: "my/namespace",
			},
			wantErr: true,
		},
	}

	for description, scenario := range tests {
		t.Run(description, func(t *testing.T) {
			var err error
			if err = ValidateConnectionExceptionReport(scenario.report); (err != nil) != scenario.wantErr {
				t.Errorf("TestValidateConnectionExceptionReport() error = %s, expected an error? %t", err, scenario.wantErr)
			}

			// verify that ONLY the `namespace` field is populated so the backend processors do not need to worry about
			// dealing with the 'processingUnitNamespace' field being set
			if err == nil {
				if scenario.report.ProcessingUnitNamespace != "" {
					t.Errorf("Validation passed, but the report's 'processingUnitNamespace' field was NOT reset. It contained the value: %s", scenario.report.ID)
				}

				if scenario.report.Namespace == "" {
					t.Error("Validation passed, but the report's 'namespace' field was empty")
				}
			}
		})
	}
}

func TestValidateCounterReport(t *testing.T) {

	tests := map[string]struct {
		report  *CounterReport
		wantErr bool
	}{
		"only the 'namespace' field is specified": {
			report: &CounterReport{
				Namespace: "my/namespace",
			},
			wantErr: false,
		},
		"only the 'enforcerNamespace' field is specified": {
			report: &CounterReport{
				EnforcerNamespace: "my/namespace",
			},
			wantErr: false,
		},
		"both the 'enforcerNamespace' and 'namespace' fields have been left omitted": {
			report: &CounterReport{
				Namespace:         "",
				EnforcerNamespace: "",
			},
			wantErr: true,
		},
		"both the 'enforcerNamespace' and 'namespace' fields have been provided": {
			report: &CounterReport{
				Namespace:         "my/namespace",
				EnforcerNamespace: "my/namespace",
			},
			wantErr: true,
		},
	}

	for description, scenario := range tests {
		t.Run(description, func(t *testing.T) {
			var err error
			if err = ValidateCounterReport(scenario.report); (err != nil) != scenario.wantErr {
				t.Errorf("TestValidateCounterReport() error = %s, expected an error? %t", err, scenario.wantErr)
			}

			// verify that ONLY the `namespace` field is populated so the backend processors do not need to worry about
			// dealing with the 'enforcerNamespace' field being set
			if err == nil {
				if scenario.report.EnforcerNamespace != "" {
					t.Errorf("Validation passed, but the report's 'enforcerNamespace' field was NOT reset. It contained the value: %s", scenario.report.ID)
				}

				if scenario.report.Namespace == "" {
					t.Error("Validation passed, but the report's 'namespace' field was empty")
				}
			}
		})
	}
}

func TestValidateCIDROrIP(t *testing.T) {
	type args struct {
		attribute string
		cidr      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"cidr",
				"10.0.0.0/8",
			},
			false,
		},
		{
			"valid IP",
			args{
				"cidr",
				"10.1.1.2",
			},
			false,
		},

		// invalid CIDR
		{
			"invalid CIDR",
			args{
				"cidr",
				"",
			},
			true,
		},

		// invalid DNn
		{
			"invalid DNS Name",
			args{
				"cidr",
				"google@com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCIDROrIP(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateCIDR() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalCIDRorIP(t *testing.T) {
	type args struct {
		attribute string
		cidr      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"cidr",
				"10.0.0.0/8",
			},
			false,
		},
		{
			"valid IP",
			args{
				"cidr",
				"10.1.1.2",
			},
			false,
		},
		{
			"valid IP",
			args{
				"cidr",
				"",
			},
			false,
		},

		// invalid CIDR
		{
			"invalid DNS Name",
			args{
				"cidr",
				"google@com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalCIDRorIP(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateCIDR() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCIDRLorIPList(t *testing.T) {
	type args struct {
		attribute string
		networks  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"10.0.0.0/8"},
			},
			false,
		},
		{
			"valid mixed list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "10.1.1.1"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google.com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCIDROrIPList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCIDRList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalCIDRLorIPList(t *testing.T) {
	type args struct {
		attribute string
		networks  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"10.0.0.0/8"},
			},
			false,
		},
		{
			"valid mixed list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "10.1.1.1"},
			},
			false,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google.com"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalCIDRorIPList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCIDRList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCloudTagsExpression(t *testing.T) {
	type args struct {
		attribute string
		tags      [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"valid",
				[][]string{{"a=b"}},
			},
			false,
		},
		{
			"valid list",
			args{
				"valid",
				[][]string{{"sg-123"}},
			},
			false,
		},
		{
			"valid list",
			args{
				"valid",
				[][]string{{"com.aporeto.io"}},
			},
			false,
		},
		{
			"valid list",
			args{
				"valid",
				[][]string{{"test:data:test"}},
			},
			false,
		},
		{
			"valid list",
			args{
				"valid",
				[][]string{{"test_data"}},
			},
			false,
		},
		{
			"valid list",
			args{
				"valid",
				[][]string{{"d@data"}},
			},
			false,
		},
		{
			"invalid-list",
			args{
				"invalid",
				[][]string{{"'"}},
			},
			true,
		},
		{
			"invalid-list",
			args{
				"invalid",
				[][]string{{"%"}},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCloudTagsExpression(tt.args.attribute, tt.args.tags); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCloudTagsExpression() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateExpressionNotEmpty(t *testing.T) {

	cases := map[string]struct {
		attribute   string
		expression  [][]string
		expectedErr elemental.Error
	}{
		"valid-single-subexpression": {
			attribute:  t.Name(),
			expression: [][]string{{"a=b"}},
		},
		"valid-multiple-subexpressions": {
			attribute:  t.Name(),
			expression: [][]string{{"a=b"}, {"c=d"}},
		},
		"invalid-nil-expression": {
			attribute:   t.Name(),
			expression:  nil,
			expectedErr: makeValidationError(t.Name(), "Expression must contain at least one sub-expression"),
		},
		"invalid-empty-expression": {
			attribute:   t.Name(),
			expression:  [][]string{},
			expectedErr: makeValidationError(t.Name(), "Expression must contain at least one sub-expression"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {

			err := ValidateExpressionNotEmpty(tc.attribute, tc.expression)
			if err == nil && tc.expectedErr == (elemental.Error{}) {
				return
			}

			var actual elemental.Error
			if ok := errors.As(err, &actual); !ok {
				t.Fatalf("unexpected error type\nwant: %T\n got: %T", elemental.Error{}, err)
			}
			if !reflect.DeepEqual(actual, tc.expectedErr) {
				t.Fatalf(
					"actual error does not match expected\nwant: '%s'\n got: '%s'",
					tc.expectedErr, actual,
				)
			}
		})
	}
}

func TestValidateSubExpressionsNotEmpty(t *testing.T) {

	cases := map[string]struct {
		attribute   string
		expression  [][]string
		expectedErr elemental.Error
	}{
		"valid-subexpression": {
			attribute:  t.Name(),
			expression: [][]string{{"a=b"}},
		},
		"valid-nil-expression": {
			attribute:  t.Name(),
			expression: nil,
		},
		"valid-empty-expression": {
			attribute:  t.Name(),
			expression: [][]string{},
		},
		"invalid-empty-expression": {
			attribute:   t.Name(),
			expression:  [][]string{{"a=b"}, {}},
			expectedErr: makeValidationError(t.Name(), "Sub-expression must not be empty"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {

			err := ValidateSubExpressionsNotEmpty(tc.attribute, tc.expression)
			if err == nil && tc.expectedErr == (elemental.Error{}) {
				return
			}

			var actual elemental.Error
			if ok := errors.As(err, &actual); !ok {
				t.Fatalf("unexpected error type\nwant: %T\n got: %T", elemental.Error{}, err)
			}
			if !reflect.DeepEqual(actual, tc.expectedErr) {
				t.Fatalf(
					"actual error does not match expected\nwant: '%s'\n got: '%s'",
					tc.expectedErr, actual,
				)
			}
		})
	}
}

func TestValidateEachSubExpressionHasNoDuplicateTags(t *testing.T) {

	cases := map[string]struct {
		attribute   string
		expression  [][]string
		expectedErr elemental.Error
	}{
		"valid-subexpressions": {
			attribute:  t.Name(),
			expression: [][]string{{"a=b"}, {"a=b"}, {"c=d"}},
		},
		"valid-nil-expression": {
			attribute:  t.Name(),
			expression: nil,
		},
		"valid-empty-expression": {
			attribute:  t.Name(),
			expression: [][]string{},
		},
		"invalid-subexpression": {
			attribute:   t.Name(),
			expression:  [][]string{{"a=b"}, {"c=d", "c=d"}},
			expectedErr: makeValidationError(t.Name(), "Duplicate tag in a sub-expression: 'c=d'"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {

			err := ValidateEachSubExpressionHasNoDuplicateTags(tc.attribute, tc.expression)
			if err == nil && tc.expectedErr == (elemental.Error{}) {
				return
			}

			var actual elemental.Error
			if ok := errors.As(err, &actual); !ok {
				t.Fatalf("unexpected error type\nwant: %T\n got: %T", elemental.Error{}, err)
			}
			if !reflect.DeepEqual(actual, tc.expectedErr) {
				t.Fatalf(
					"actual error does not match expected\nwant: '%s'\n got: '%s'",
					tc.expectedErr, actual,
				)
			}
		})
	}
}

func TestValidateNoDuplicateSubExpressions(t *testing.T) {

	cases := map[string]struct {
		attribute   string
		expression  [][]string
		expectedErr elemental.Error
	}{
		"valid-subexpressions": {
			attribute:  t.Name(),
			expression: [][]string{{"a=b", "c=d"}, {"a=b"}, {"c=d"}},
		},
		"valid-nil-expression": {
			attribute:  t.Name(),
			expression: nil,
		},
		"valid-empty-expression": {
			attribute:  t.Name(),
			expression: [][]string{},
		},
		"invalid-expression": {
			attribute:   t.Name(),
			expression:  [][]string{{"a=b", "c=d"}, {"c=d"}, {"a=b"}, {"c=d", "a=b"}},
			expectedErr: makeValidationError(t.Name(), "Duplicate equivalent sub-expressions found"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {

			err := ValidateNoDuplicateSubExpressions(tc.attribute, tc.expression)
			if err == nil && tc.expectedErr == (elemental.Error{}) {
				return
			}

			var actual elemental.Error
			if ok := errors.As(err, &actual); !ok {
				t.Fatalf("unexpected error type\nwant: %T\n got: %T", elemental.Error{}, err)
			}
			if !reflect.DeepEqual(actual, tc.expectedErr) {
				t.Fatalf(
					"actual error does not match expected\nwant: '%s'\n got: '%s'",
					tc.expectedErr, actual,
				)
			}
		})
	}
}

func TestValidatePortsList(t *testing.T) {
	type args struct {
		attribute string
		ports     []*portutils.PortsRange
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"valid",
				[]*portutils.PortsRange{
					{
						FromPort: 22,
						ToPort:   22,
					},
				},
			},
			false,
		},
		{
			"valid list",
			args{
				"valid",
				[]*portutils.PortsRange{
					{
						FromPort: -1,
						ToPort:   65535,
					},
				},
			},
			false,
		},

		{
			"valid empty",
			args{
				"valid",
				[]*portutils.PortsRange{},
			},
			false,
		},
		{
			"invalid-list",
			args{
				"invalid",
				[]*portutils.PortsRange{
					{
						FromPort: -2,
						ToPort:   10,
					},
				},
			},
			true,
		},
		{
			"invalid-list",
			args{
				"invalid",
				[]*portutils.PortsRange{
					{
						FromPort: 10,
						ToPort:   70000,
					},
				},
			},
			true,
		},
		{
			"invalid-list",
			args{
				"invalid",
				[]*portutils.PortsRange{
					{
						FromPort: 10,
						ToPort:   2,
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePortsList(tt.args.attribute, tt.args.ports); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePortsList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNativeID(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			"valid",
			"sg-123",
			false,
		},
		{
			"valid",
			"sg.123.3.3",
			false,
		},
		{
			"valid",
			"sg:2",
			false,
		},
		{
			"invalid",
			"a b",
			true,
		},
		{
			"invalid",
			"",
			true,
		},
		{
			"valid",
			"arn:aws:elasticloadbalancing:us-east-2:345744466724:loadbalancer/app/abhi-aws-v2/eb7edb2d68b71a0f",
			false,
		},
		{
			"valid",
			"/subscriptions/e3e57e4f-0ceb-449c-a8b4-23d61bd60554/resourceGroups/someaccount/providers/Microsoft.Network/virtualNetworks/somevnet/subnets/default",
			false,
		},
		{
			"valid",
			"{version:\"2020.2\",id:\"/subscriptions/e3e57e4f-0ceb-449c-a8b4-23d61bd60554/resourceGroups/someaccount/providers/Microsoft.Network/virtualNetworks/somevnet/subnets/default\"}",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNativeID(tt.name, tt.id); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNativeID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestValidateCloudGraphQuery(t *testing.T) {
	type args struct {
		attribute string
		query     *CloudNetworkQuery
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				"valid",
				&CloudNetworkQuery{
					SourceIP: "0.0.0.0/0",
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
				},
			},
			false,
		},
		{
			"source and dest ip set",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceIP:      "0.0.0.0/0",
					DestinationIP: "0.0.0.0/0",
				},
			},
			true,
		},
		{
			"private source IP set",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceIP: "10.1.1.0/24",
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
				},
			},
			true,
		},
		{
			"east/west with no source VPC",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						AccountIDs: []string{"account1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
				},
			},
			true,
		},
		{
			"east/west with no destination VPC",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						AccountIDs: []string{"account1"},
					},
				},
			},
			true,
		},
		{
			"east/west with multiple source VPCs",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1", "vpc2"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
				},
			},
			true,
		},
		{
			"east/west with multiple destination VPCs",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1", "vpc2"},
					},
				},
			},
			true,
		},
		{
			"east/west valid",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc2"},
					},
				},
			},
			false,
		},
		{
			"source ip and selector set",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
				},
			},
			true,
		},
		{
			"neither source ip or selector set",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector:      &CloudNetworkQueryFilter{},
					DestinationSelector: &CloudNetworkQueryFilter{},
				},
			},
			true,
		},
		{
			"destination ip and selector set",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceIP:      "0.0.0.0/0",
					DestinationIP: "0.0.0.0/0",
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
				},
			},
			true,
		},
		{
			"neither destination IP nor selector set",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceIP: "0.0.0.0/0",
				},
			},
			true,
		},
		{
			"private destination IP",
			args{
				"invalid",
				&CloudNetworkQuery{
					DestinationIP: "10.1.1.0/24",
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
				},
			},
			true,
		},
		{
			"bad destination selector with security tags",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceIP: "0.0.0.0/0",
					DestinationSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInstance,
						SecurityTags: []string{"a=b"},
					},
				},
			},
			true,
		},
		{
			"bad source selector service owners",
			args{
				"invalid",
				&CloudNetworkQuery{
					DestinationIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType:  CloudNetworkQueryFilterResourceTypeInstance,
						ServiceOwners: []string{"a=b"},
					},
				},
			},
			true,
		},
		{
			"bad source selector service owners",
			args{
				"invalid",
				&CloudNetworkQuery{
					DestinationIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInstance,
						ServiceTypes: []string{"a=b"},
					},
				},
			},
			true,
		},
		{
			"product info type filter for interface",
			args{
				"invalid",
				&CloudNetworkQuery{
					DestinationIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType:    CloudNetworkQueryFilterResourceTypeInterface,
						ProductInfoType: "marketplace",
					},
				},
			},
			true,
		},
		{
			"product info value filter for interface",
			args{
				"invalid",
				&CloudNetworkQuery{
					DestinationIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType:     CloudNetworkQueryFilterResourceTypeInterface,
						ProductInfoValue: "vmseries",
					},
				},
			},
			true,
		},
		{
			"image as filter for interface",
			args{
				"invalid",
				&CloudNetworkQuery{
					DestinationIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						ImageIDs:     []string{"image1"},
					},
				},
			},
			true,
		},
		{
			"service name as the selector",
			args{
				"valid",
				&CloudNetworkQuery{
					DestinationIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeService,
						ServiceNames: []string{"service1"},
					},
				},
			},
			false,
		},
		{
			"service name as the selector when resource type is interface",
			args{
				"valid",
				&CloudNetworkQuery{
					DestinationIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						ServiceNames: []string{"service1"},
					},
				},
			},
			true,
		},
		{
			"excluded network validation with valid source",
			args{
				"valid",
				&CloudNetworkQuery{
					DestinationIP: "0.0.0.0/0",
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeService,
						ServiceNames: []string{"service1"},
					},
					ExcludedNetworks: []string{"30.1.1.0/24"},
				},
			},
			false,
		},
		{
			"excluded network validation without IP definition",
			args{
				"valid",
				&CloudNetworkQuery{
					DestinationSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeService,
						ServiceNames: []string{"service1"},
					},
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeService,
						ServiceNames: []string{"service1"},
					},
					ExcludedNetworks: []string{"30.1.1.0/24"},
				},
			},
			true,
		},
		{
			"cloud network path with source without object",
			args{
				"invalid",
				&CloudNetworkQuery{
					Type: CloudNetworkQueryTypeNetworkPath,
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs:    []string{"vpc3"},
						ObjectIDs: []string{"object1"},
					},
				},
			},
			true,
		},
		{
			"cloud network path with destination without object",
			args{
				"invalid",
				&CloudNetworkQuery{
					Type: CloudNetworkQueryTypeNetworkPath,
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs:    []string{"vpc3"},
						ObjectIDs: []string{"object1"},
					},
				},
			},
			true,
		},
		{
			"cloud network path with destination without object and source IP",
			args{
				"invalid",
				&CloudNetworkQuery{
					Type: CloudNetworkQueryTypeNetworkPath,
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
					SourceIP: "1.1.1.0/24",
				},
			},
			true,
		},
		{
			"cloud network path with source without object and destination IP",
			args{
				"invalid",
				&CloudNetworkQuery{
					Type: CloudNetworkQueryTypeNetworkPath,
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs: []string{"vpc1"},
					},
					DestinationIP: "1.1.1.0/24",
				},
			},
			true,
		},
		{
			"cloud network path with source without object",
			args{
				"invalid",
				&CloudNetworkQuery{
					Type: CloudNetworkQueryTypeNetworkPath,
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs:    []string{"vpc1"},
						ObjectIDs: []string{"object1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs:    []string{"vpc3"},
						ObjectIDs: []string{"object1"},
					},
				},
			},
			false,
		},
		{
			"both source and dest IP empty but address match is set",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						VPCIDs:    []string{"vpc1"},
						ObjectIDs: []string{"object1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						VPCIDs:    []string{"vpc3"},
						ObjectIDs: []string{"object1"},
					},
					AddressMatchCriteria: CloudNetworkQueryAddressMatchCriteriaPartialMatch,
				},
			},
			true,
		},
		{
			"source and dest have same cloud types",
			args{
				"valid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						CloudTypes:   []string{"AWS"},
						VPCIDs:       []string{"vpc1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						CloudTypes:   []string{"AWS"},
						VPCIDs:       []string{"vpc1"},
					},
				},
			},
			false,
		},
		{
			"source and dest have different cloud types",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						CloudTypes:   []string{"GCP"},
						VPCIDs:       []string{"vpc1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						CloudTypes:   []string{"AWS"},
						VPCIDs:       []string{"vpc1"},
					},
				},
			},
			true,
		},
		{
			"source and dest have different cloud types",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						CloudTypes:   []string{"GCP", "AWS"},
						VPCIDs:       []string{"vpc1"},
					},
					DestinationSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						CloudTypes:   []string{"AWS"},
						VPCIDs:       []string{"vpc1"},
					},
				},
			},
			true,
		},
		{
			"paas filter is set for source selector",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInterface,
						CloudTypes:   []string{"Azure"},
						PaasTypes:    []string{"sql"},
					},
					DestinationIP: "1.1.1.0/24",
				},
			},
			true,
		},
		{
			"paas filter is not empty but resource type is not paas",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceIP: "1.1.1.0/24",
					DestinationSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypeInstance,
						CloudTypes:   []string{"Azure"},
						PaasTypes:    []string{"sql"},
					},
				},
			},
			true,
		},
		{
			"paas filter is set for AWS",
			args{
				"invalid",
				&CloudNetworkQuery{
					SourceIP: "1.1.1.0/24",
					DestinationSelector: &CloudNetworkQueryFilter{
						ResourceType: CloudNetworkQueryFilterResourceTypePaaS,
						CloudTypes:   []string{"Aws"},
						PaasTypes:    []string{"sql"},
					},
				},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCloudNetworkQueryEntity(tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCloudNetworkQueryEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateIPAddressList(t *testing.T) {
	type args struct {
		attribute string
		ips       []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid ips",
			args{
				"attr",
				[]string{"1.2.3.4", "2.3.4.5"},
			},
			false,
		},
		{
			"on invalid ips",
			args{
				"attr",
				[]string{"1.2.3.4", "oh no"},
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				[]string{},
			},
			true,
		},
		{
			"nil",
			args{
				"attr",
				nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIPAddressList(tt.args.attribute, tt.args.ips); (err != nil) != tt.wantErr {
				t.Errorf("ValidateIPAddressList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalIPAddressList(t *testing.T) {
	type args struct {
		attribute string
		ips       []string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid ips",
			args{
				"attr",
				[]string{"1.2.3.4", "2.3.4.5"},
			},
			false,
		},
		{
			"on invalid ips",
			args{
				"attr",
				[]string{"1.2.3.4", "oh no"},
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				[]string{},
			},
			false,
		},
		{
			"nil",
			args{
				"attr",
				nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalIPAddressList(tt.args.attribute, tt.args.ips); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalIPAddressList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsAddressPrivate(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    bool
		wantErr bool
	}{
		{
			"private address",
			"127.0.0.0/10",
			true,
			false,
		},
		{
			"public address",
			"135.20.0.0/24",
			false,
			false,
		},
		{
			"public address",
			"badIP",
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if private, err := IsAddressPrivate(tt.args); (private != tt.want) || (err == nil) == tt.wantErr {
				t.Errorf("IsAddressPrivate() got wrong response %v for %+v with error %s", private, tt.args, err)
			}
		})
	}
}

func TestValidateAPIServerServiceName(t *testing.T) {
	type args struct {
		attribute   string
		serviceName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				"attr",
				"abc.abcns.svc",
			},
			false,
		},
		{
			"empty name",
			args{
				"attr",
				".abcns.svc",
			},
			true,
		},
		{
			"empty namespace",
			args{
				"attr",
				"abc..svc",
			},
			true,
		},
		{
			"empty name and namespace",
			args{
				"attr",
				"..svc",
			},
			true,
		},
		{
			"empty name, namespace and svc",
			args{
				"attr",
				"..",
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAPIServerServiceName(tt.args.attribute, tt.args.serviceName); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAPIServerServiceName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSyslogEndpoint(t *testing.T) {
	type args struct {
		attribute string
		endpoint  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty endpoint",
			args: args{
				"syslogEndpoint",
				"",
			},
			wantErr: false,
		},
		{
			name: "valid udp",
			args: args{
				"syslogEndpoint",
				"udp://example.com",
			},
			wantErr: false,
		},
		{
			name: "valid tcp",
			args: args{
				"syslogEndpoint",
				"tcp://example.com",
			},
			wantErr: false,
		},
		{
			name: "valid tls",
			args: args{
				"syslogEndpoint",
				"tls://example.com",
			},
		},
		{
			name: "invalid tcp",
			args: args{
				"syslogEndpoint",
				"tcp:/example.com",
			},
			wantErr: true,
		},
		{
			name: "invalid protocol",
			args: args{
				"syslogEndpoint",
				"ftp://example.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateSyslogEndpoint(tt.args.attribute, tt.args.endpoint); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSyslogEndpoint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
