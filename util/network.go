package util

import (
	"net"
	"regexp"
	"strings"
)

// ParseSourceRanges validates the sourceRanges string provided by the user.
// Expects valid IPv4/IPv6 CIDR blocks delimited by a comma. i.e. 10.16.16.0/24, 203.19.42.0/25, 2001:db8:fd00:1000:dead:beef:cafe:2/64
func ParseSourceRanges(sourceRanges string) ([]*net.IPNet, error) {
	var sourceIPNetRanges []*net.IPNet
	for _, sourceRange := range strings.Split(sourceRanges, ",") {
		_, ipNet, err := net.ParseCIDR(sourceRange)
		if err != nil {
			return nil, err
		}
		sourceIPNetRanges = append(sourceIPNetRanges, ipNet)
	}
	return sourceIPNetRanges, nil
}

// CheckIPInNetwork checks if the specified ip is contained in any of the provided networks
func CheckIPInNetwork(ip net.IP, networks []*net.IPNet) bool {
	for _, network := range networks {
		if containsIP := network.Contains(ip); containsIP {
			return true
		}
	}
	return false
}

var clientIPregex = regexp.MustCompile(`^\[(.*)].*$`)

// Return the client IP address from the remote address
// This is simple and does not handle the case where a proxy exists
func ClientIPFromRemoteAddr(remoteAddr string) string {
	clientIP := clientIPregex.FindStringSubmatch(remoteAddr)
	return clientIP[1]
}
