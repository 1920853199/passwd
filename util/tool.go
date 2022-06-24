package util

import (
	"fmt"
	"net"
)

func GetMacAddrs() string {
	var macAddrs = make([]string, 0)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return "00.00.00.00:00:00"
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	if len(macAddrs) > 0 {
		return macAddrs[0]
	}
	return "00.00.00.00:00:00"
}
