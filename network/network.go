package network

import (
	"net"
)

func ValidateIPAddress(sIpAddress string, aAllowedIPAddresses []string) bool {
	oIpAddress := net.ParseIP(sIpAddress)
	for _, sAllowedIpAddress := range aAllowedIPAddresses {
		_, oCIDR, oErr := net.ParseCIDR(sAllowedIpAddress)
		if oErr != nil {
			panic(oErr)
		}
		if oCIDR.Contains(oIpAddress) {
			return true
		}
	}

	return false
}

func GetIpAddressAndPort(sHostPort string) (string, string) {
	sIp, sPort, oErr := net.SplitHostPort(sHostPort)
	if oErr != nil {
		panic(oErr)
	}
	return sIp, sPort
}
