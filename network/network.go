package network

import (
	"net"
	"net/http"
)

func ValidateIPAddress(oRequest *http.Request, aAllowedIPAddresses []string) bool {
	for _, sAllowedIpAddress := range aAllowedIPAddresses {
		_, oCIDR, oErr := net.ParseCIDR(sAllowedIpAddress)
		if oErr != nil {
			panic(oErr)
		}
		sIpAddress, _ := GetIpAddressAndPort(oRequest)
		oIpAddress := net.ParseIP(sIpAddress)
		if oCIDR.Contains(oIpAddress) {
			return true
		}
	}

	return false
}

func GetIpAddressAndPort(oRequest *http.Request) (string, string) {
	sIp, sPort, oErr := net.SplitHostPort(oRequest.RemoteAddr)
	if oErr != nil {
		panic(oErr)
	}
	return sIp, sPort
}
