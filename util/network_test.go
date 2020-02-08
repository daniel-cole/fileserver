package util

import (
	"net"
	"testing"
)

func TestValidCheckIPInNetwork(t *testing.T) {
	_, network1, _ := net.ParseCIDR("10.10.12.0/24")
	_, network2, _ := net.ParseCIDR("172.16.0.0/16")
	networks := []*net.IPNet{network1, network2}
	ip1 := "10.10.12.35"
	if !CheckIPInNetwork(net.ParseIP(ip1), networks) {
		t.Errorf("Expected %s to exist in %v", ip1, networks)
	}

	ip2 := "172.16.10.10"
	if !CheckIPInNetwork(net.ParseIP(ip2), networks) {
		t.Errorf("Expected %s to exist in %v", ip2, networks)
	}

}

func TestInvalidCheckIPInNetwork(t *testing.T) {
	_, network1, _ := net.ParseCIDR("10.10.12.0/24")
	_, network2, _ := net.ParseCIDR("172.16.0.0/16")
	networks := []*net.IPNet{network1, network2}
	ip1 := "10.11.12.35"
	if CheckIPInNetwork(net.ParseIP(ip1), networks) {
		t.Errorf("Expected %s to not exist in %v", ip1, networks)
	}

	ip2 := "172.17.10.10"
	if CheckIPInNetwork(net.ParseIP(ip2), networks) {
		t.Errorf("Expected %s to not exist in %v", ip2, networks)
	}
}

func TestNilCheckIPInNetwork(t *testing.T) {
	_, network1, _ := net.ParseCIDR("10.10.12.0/24")
	_, network2, _ := net.ParseCIDR("172.16.0.0/16")
	networks := []*net.IPNet{network1, network2}
	if CheckIPInNetwork(nil, networks) {
		t.Errorf("Expected nil to not exist in %v", networks)
	}
}

func TestValidParseSourceRanges(t *testing.T) {
	validSourceRanges1 := "10.10.12.0/24,203.19.128.0/25"
	_, err := ParseSourceRanges(validSourceRanges1)
	checkParseSourceRanges(t, validSourceRanges1, err, false)
}

func TestInvalidParseSourceRanges(t *testing.T) {
	invalidSourceRanges1 := "10.12.0.16"
	_, err := ParseSourceRanges(invalidSourceRanges1)
	checkParseSourceRanges(t, invalidSourceRanges1, err, true)

	invalidSourceRanges2 := "10.15.10.12/24,notarealrange"
	_, err = ParseSourceRanges(invalidSourceRanges2)
	checkParseSourceRanges(t, invalidSourceRanges2, err, true)
}

// checkParseSourceRanges to help keep it DRY
func checkParseSourceRanges(t *testing.T, sourceRanges string, err error, expectedError bool) {
	if expectedError && err == nil {
		t.Errorf("Expected error on parsing source ranges: %s", sourceRanges)
	}
	if !expectedError && err != nil {
		t.Errorf("Didn't expect error on parsing source ranges: %s", sourceRanges)
	}
}
