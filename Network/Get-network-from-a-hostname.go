package main

import (
	"fmt"
	"net"
)

func main() {
	hostname := "www.boston.com"

	IPAddr, _ := net.ResolveIPAddr("ip", hostname)
	addr := net.ParseIP(IPAddr.String())

	// DefaultMask returns the default IP mask for the IP address ip.
	mask := addr.DefaultMask()

	// Mask returns the result of masking the IP address ip with mask.
	network := addr.Mask(mask)

	fmt.Println("Hostname:", hostname)
	fmt.Println("Address :", addr.String())
	fmt.Println("Network :", network)
}
