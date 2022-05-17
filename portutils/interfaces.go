package portutils

// PortsRange holds a range of ports.
type PortsRange struct {
	FromPort int64
	ToPort   int64
}

// PortsRangeList holds a list of port range.
type PortsRangeList []*PortsRange

// HasOverlapWithPort returns true if the range overlaps the given port.
func (p *PortsRange) HasOverlapWithPort(port int64) bool {
	return port >= p.FromPort && port <= p.ToPort
}

// HasOverlapWithPortsList returns true if the range overlaps one of the given ports.
func (p *PortsRange) HasOverlapWithPortsList(ports *PortsList) bool {

	if len(*ports) == 0 {
		return false
	}

	for _, port := range *ports {
		if p.HasOverlapWithPort(port) {
			return true
		}
	}
	return false
}

// HasOverlapWithPortsRange returns true if the range overlaps the given ports range.
func (p *PortsRange) HasOverlapWithPortsRange(otherRange *PortsRange) bool {

	return p.FromPort >= otherRange.FromPort && p.FromPort <= otherRange.ToPort ||
		p.ToPort >= otherRange.FromPort && p.ToPort <= otherRange.ToPort ||
		otherRange.FromPort >= p.FromPort && otherRange.FromPort <= p.ToPort ||
		otherRange.ToPort >= p.FromPort && otherRange.ToPort <= p.ToPort
}

// ContainsPortsRange returns true if the range contains the given ports range.
func (p *PortsRange) ContainsPortsRange(otherRange *PortsRange) bool {

	return p.FromPort <= otherRange.FromPort && p.ToPort >= otherRange.ToPort
}

// HasOverlapWithPortsRanges returns true if the range has overlaps with the given ports.
func (p *PortsRange) HasOverlapWithPortsRanges(otherRanges *PortsRangeList) bool {

	for _, pr := range *otherRanges {
		if p.HasOverlapWithPortsRange(pr) {
			return true
		}
	}
	return false
}

// PortsList holds a list of ports
type PortsList []int64

// HasOverlapWithPort returns true if the list has one overlap with the given port.
func (p *PortsList) HasOverlapWithPort(port int64) bool {

	for _, i := range *p {
		if i == port {
			return true
		}
	}

	return false
}

// HasOverlapWithPortsList returns true if the list has one overlap with the given ports.
func (p *PortsList) HasOverlapWithPortsList(ports *PortsList) bool {

	for _, p1 := range *ports {
		if p.HasOverlapWithPort(p1) {
			return true
		}
	}
	return false
}

// HasOverlapWithPortsRanges returns true if the list has one overlap with the given ranges.
func (p *PortsList) HasOverlapWithPortsRanges(portRanges *PortsRangeList) bool {

	for _, pr := range *portRanges {
		if pr.HasOverlapWithPortsList(p) {
			return true
		}
	}
	return false
}
