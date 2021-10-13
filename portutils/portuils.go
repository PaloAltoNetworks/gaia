package portutils

import (
	"fmt"
	"strconv"
	"strings"

	"go.aporeto.io/gaia/protocols"
)

// ConvertToPortsRange converts a :-separated string to a min and max port.
func ConvertToPortsRange(ports string) (*PortsRange, error) {

	sp := strings.Split(ports, ":")

	if len(sp) != 2 {
		return nil, fmt.Errorf("%s is not a valid range format. It should be of form fromPort:toPort", ports)
	}

	fromPort, err := ConvertToSinglePort(sp[0])
	if err != nil {
		return nil, err
	}

	toPort, err := ConvertToSinglePort(sp[1])
	if err != nil {
		return nil, err
	}

	if fromPort >= toPort {
		return nil, fmt.Errorf("left bound is greater or equal to right bound")
	}

	return &PortsRange{fromPort, toPort}, nil
}

// ConvertToPortsList convert a , separated string to a list of port.
func ConvertToPortsList(ports string) (*PortsList, error) {

	results := PortsList{}

	p, err := ConvertToSinglePort(ports)
	if err != nil {
		return nil, err
	}

	results = append(results, p)
	return &results, nil

}

// ConvertToSinglePort converts a string to port.
func ConvertToSinglePort(port string) (int64, error) {

	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("%s is not a valid port", port)
	}

	if p < 0 || p > 65535 {
		return -1, fmt.Errorf("%s is not in between 0 and 65535", port)
	}

	return p, nil
}

// ExtractPortsAndProtocol extracts ports and protocol from the service
// NOTE: The protocol should be in uppercase to match our list of protocols here
// https://github.com/aporeto-inc/gaia/blob/master/protocols/protocols.go
func ExtractPortsAndProtocol(service string) (string, string, error) {

	upperService := strings.ToUpper(service)
	proto := ""
	var portSubString string

	if strings.HasPrefix(upperService, protocols.L4ProtocolUDP+"/") {
		proto = protocols.L4ProtocolUDP
		portSubString = upperService[4:]
	} else if strings.HasPrefix(upperService, protocols.L4ProtocolTCP+"/") {
		proto = protocols.L4ProtocolTCP
		portSubString = upperService[4:]
	} else {
		return "", "", fmt.Errorf("invalid protocol/ports: %s", service)
	}

	return portSubString, proto, nil
}

// ExtractPortsAndProtocolFromHostService extracts the port range and the protocol from a host service like tcp/80:100.
// If the prefix is neither tcp or udp it will return an error.
// NOTE: The protocol should be in uppercase to match our list of protocols here
// https://github.com/aporeto-inc/gaia/blob/master/protocols/protocols.go
func ExtractPortsAndProtocolFromHostService(service string) (*PortsRange, string, error) {

	proto := ""
	var portSubString string

	portSubString, proto, err := ExtractPortsAndProtocol(service)
	if err != nil {
		proto = protocols.L4ProtocolTCP
		portSubString = strings.ToUpper(service)
	}

	if strings.Contains(portSubString, ":") {

		parts := strings.SplitN(portSubString, ":", 2)
		if len(parts) != 2 {
			return nil, "", fmt.Errorf("%s does not have a valid port range", service)
		}

		min, err := ConvertToSinglePort(parts[0])
		if err != nil {
			return nil, "", err
		}

		max, err := ConvertToSinglePort(parts[1])
		if err != nil {
			return nil, "", err
		}

		if min > max {
			return nil, "", fmt.Errorf("%s is not a valid port range", portSubString)
		}
		return &PortsRange{FromPort: min, ToPort: max}, proto, nil
	}

	min, err := ConvertToSinglePort(portSubString)
	if err != nil {
		return nil, "", err
	}

	return &PortsRange{FromPort: min, ToPort: min}, proto, nil
}
