package netutils

import (
	"fmt"
	"math/big"
	"net"
	"strings"

	"github.com/yl2chen/cidranger"
)

var (
	opInclude         = "include" // default operation to include cidr in networks
	opExclude         = "exclude" // not include cidr when cidr string has the ! prefix
	multicastv4Subnet = "224.0.0.0/4"
	multicastv6Subnet = "ff00::/8"
)

// cidr represents the cidr network to be included or excluded
type cidr struct {
	// op is the opreation to perform with the cidr: include or exclude
	op string

	// ipNet is the IPNet object that contains ip and mask of the cidr
	ipNet net.IPNet

	// str is the original cidr string
	str string
}

func checkExcPfxContainedInc(entries []cidranger.RangerEntry, mask net.IPMask, _ net.IPNet) bool { // nolint:unparam

	for _, e := range entries {
		cidr := e.(*cidr)

		// we only are interested if the excluded pfx is present in the included.
		if cidr.op == opExclude {
			continue
		}

		includedOnes, size1 := e.Network().Mask.Size() //  included CIDR subnet
		excludedOnes, size2 := mask.Size()             // exclude CIDR subnet, basically the ip under check.

		if size1 != size2 {
			continue
		}
		// basically, check here is that the,
		// excluded subnet should always be greater than included, if its less, then return ERROR/FALSE.
		if includedOnes <= excludedOnes {
			return true
		}

	}
	return false
}

// parseCIDR converts the given string to cidr. Returns an error if it wasnt able to parse a CIDR
func parseCIDR(s string) (*cidr, error) {

	c := &cidr{op: opInclude, str: s}
	_, network, err := net.ParseCIDR(strings.TrimPrefix(s, "!"))
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid CIDR", s)
	}
	c.ipNet = *network
	if strings.HasPrefix(s, "!") {
		c.op = opExclude
	}

	return c, nil
}

// get function for network
func (b *cidr) Network() net.IPNet {
	return b.ipNet
}

// create customRangerEntry object using net and cidr
func newCustomRangerEntry(c *cidr) cidranger.RangerEntry {
	return c
}

func addressv6RangeCount(ip net.IPNet) *big.Int {
	ones, bits := ip.Mask.Size()
	pow := bits - ones
	var base, p = big.NewInt(2), big.NewInt(int64(pow))
	return base.Exp(base, p, nil)
}

func addressv4RangeCount(ip net.IPNet) uint64 {
	ones, bits := ip.Mask.Size()
	return 1 << (uint64(bits) - uint64(ones))
}

// ValidateUDPCIDRs validates that the list of string provided as a set is a valid CIDR set
func ValidateUDPCIDRs(ss []string) error {

	cidrmap := make(map[string]*cidr)
	// instantiate NewPCTrieRanger
	ranger := cidranger.NewPCTrieRanger()
	multicastv4SubnetPresent := false
	multicastv6SubnetPresent := false

	for _, s := range ss {
		cidr, err := parseCIDR(s)
		if err != nil {
			return err
		}
		if _, ok := cidrmap[cidr.ipNet.String()]; !ok {
			cidrmap[cidr.ipNet.String()] = cidr
		} else {
			return fmt.Errorf("CIDR subnet parsed from %s is duplicated", cidr.str)
		}
		if cidr.ipNet.String() == multicastv4Subnet {
			multicastv4SubnetPresent = true
		}
		if cidr.ipNet.String() == multicastv6Subnet {
			multicastv6SubnetPresent = true
		}
		if err := ranger.Insert(newCustomRangerEntry(cidr)); err != nil {
			return fmt.Errorf("Error adding CIDR %s", cidr.str)
		}
	}
	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, c := range cidrmap {

		if c.op == opExclude {
			// get all the networks which are containing the network,
			// basically getting all the parent prefixes.
			entries, err := ranger.ContainingNetworks(c.ipNet.IP)
			if err != nil {
				return fmt.Errorf("Cannot find the CIDR: %s", err)
			}

			mask := c.ipNet.Mask
			// make sure the NOT(!) CIDR is part/contained of/by a included CIDR
			present := checkExcPfxContainedInc(entries, mask, c.ipNet)

			// if the excluded CIDR is not contained in the included CIDR then return error
			if !present {
				return fmt.Errorf("%s is not contained in any CIDR", c.str)
			}
		}
	}

	// now if all the CIDR make sense, check for multicast subnet,
	//  check if we have a 224/4 subnet in the pfx tree and if its not excluded,
	// if yes then return error

	_, v4network, err := net.ParseCIDR(multicastv4Subnet)
	if err != nil {
		return fmt.Errorf("%s is not a valid CIDR", v4network)
	}
	_, v6network, err := net.ParseCIDR(multicastv6Subnet)
	if err != nil {
		return fmt.Errorf("%s is not a valid CIDR", v6network)
	}

	// get all the networks that have 224/4 and if anyone is included in the 224/4 subnet
	// then we report error.

	// ipv4 checks
	multicastv4Entries, _ := ranger.CoveredNetworks(*v4network)

	expectedmulticastv4AddrRange := addressv4RangeCount(*v4network)
	actualmulticastv4AddrRange := uint64(0)

	multicastv4SubnetInc := false
	wrongEntryv4CIDR := make([]string, 0, len(multicastv4Entries))

	for _, entry := range multicastv4Entries {
		cidr := entry.(*cidr)
		if cidr.op == opInclude {
			multicastv4SubnetInc = true
			wrongEntryv4CIDR = append(wrongEntryv4CIDR, cidr.str)
		}
		// we count the actual host addresses when the pfx tree has 224/4 and all of them are excluded.
		actualmulticastv4AddrRange += addressv4RangeCount(entry.Network())
		ranger.Remove(cidr.ipNet) //nolint
	}
	if len(multicastv4Entries) > 0 && multicastv4SubnetInc {
		return fmt.Errorf("The CIDR %s are multicast subnets, should not be included in UDP target networks", wrongEntryv4CIDR)
	}
	var lastMultiContainedv4CIDR, lastMultiContainedv6CIDR *cidr
	// now check here if multicast subnet is contained in the pfx-tree

	MulticastContainedv4Subnets, _ := ranger.ContainingNetworks(v4network.IP)
	if len(MulticastContainedv4Subnets) > 0 {
		lastMultiContainedv4CIDR = MulticastContainedv4Subnets[len(MulticastContainedv4Subnets)-1].(*cidr)
	}
	// now we come only in 3 cases:
	// 1. When the pfx tree has 224/4, and itself and all the network under them are excluded,
	// 		so, we cal the actual multicast addresses count to be sure that really everything is excluded.
	// 2. When there is no 224/4 network in tree, ex: 0/0, 128/0 so get the last contained subnet and if
	// 		the last subnet is excluded then we are good else return error.
	// 3. When the pfx tree doesn't have the exact 224/4 but has 224/5, 232/5, we need to check for the actual multicastrange
	if lastMultiContainedv4CIDR != nil && !multicastv4SubnetPresent && actualmulticastv4AddrRange != expectedmulticastv4AddrRange {
		if lastMultiContainedv4CIDR.op == opInclude {
			return fmt.Errorf("The CIDR %s includes multicast subnets, should not be included in UDP target networks", lastMultiContainedv4CIDR.str)
		}
	}
	multicastv6Entries, _ := ranger.CoveredNetworks(*v6network)

	expectedmulticastv6AddrRange := addressv6RangeCount(*v6network)
	actualmulticastv6AddrRange := big.NewInt(0)

	multicastv6SubnetInc := false
	wrongEntryv6CIDR := make([]string, 0, len(multicastv6Entries))
	for _, entry := range multicastv6Entries {
		cidr := entry.(*cidr)
		if cidr.op == opInclude {
			multicastv6SubnetInc = true
			wrongEntryv6CIDR = append(wrongEntryv6CIDR, cidr.str)
		}

		actualmulticastv6AddrRange.Add(addressv6RangeCount(entry.Network()), actualmulticastv6AddrRange)
		ranger.Remove(cidr.ipNet) //nolint
	}
	if len(multicastv6Entries) > 0 && multicastv6SubnetInc {
		return fmt.Errorf("The CIDR %s are multicast subnets, should not be included in UDP target networks", wrongEntryv6CIDR)
	}

	MulticastContainedv6Subnets, _ := ranger.ContainingNetworks(v6network.IP)
	if len(MulticastContainedv6Subnets) > 0 {
		lastMultiContainedv6CIDR = MulticastContainedv6Subnets[len(MulticastContainedv6Subnets)-1].(*cidr)
	}
	if lastMultiContainedv6CIDR != nil && !multicastv6SubnetPresent && actualmulticastv6AddrRange != expectedmulticastv6AddrRange {
		if lastMultiContainedv6CIDR.op == opInclude {
			return fmt.Errorf("The CIDR %s includes multicast subnets, should not be included in UDP target networks", lastMultiContainedv6CIDR.str)
		}
	}

	return nil
}

// ValidateCIDRs validates that the list of string provided as a set is a valid CIDR set
func ValidateCIDRs(ss []string) error {

	cidrmap := make(map[string]*cidr)
	// instantiate NewPCTrieRanger
	ranger := cidranger.NewPCTrieRanger()

	for _, s := range ss {
		cidr, err := parseCIDR(s)
		if err != nil {
			return err
		}
		if _, ok := cidrmap[cidr.ipNet.String()]; !ok {
			cidrmap[cidr.ipNet.String()] = cidr
		} else {
			return fmt.Errorf("CIDR subnet parsed from %s is duplicated", cidr.str)
		}
		if err := ranger.Insert(newCustomRangerEntry(cidr)); err != nil {
			return fmt.Errorf("Error adding CIDR %s", cidr.str)
		}
	}

	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, c := range cidrmap {

		if c.op == opExclude {

			entries, err := ranger.ContainingNetworks(c.ipNet.IP)
			if err != nil {
				return fmt.Errorf("Cannot find the CIDR: %s", err)
			}

			mask := c.ipNet.Mask
			present := checkExcPfxContainedInc(entries, mask, c.ipNet)

			// if the excluded CIDR is not contained in the included CIDR then return error
			if !present {
				return fmt.Errorf("%s is not contained in any CIDR", c.str)
			}
		}
	}

	return nil
}
