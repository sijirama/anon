package Receiver

import (
	"fmt"
	"net"
	"strconv"
)

/*
WARN: chatgpt gave me this code, i am too lazy to sit back and understand
i sha know it gives me what exactly i need, we just chilling over here
*/

func getLocalIP() string {

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	for _, iface := range interfaces { //INFO: iterate over all interface
		// Exclude loopback and non-up interfaces
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			// Get the addresses associated with the interface
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// Iterate over the addresses to find an IPv4 address
			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String()
				}
			}
		}
	}

	return ""
}

func findAvailablePort() string {
	startPort := 8000

	for port := startPort; port < startPort+1000; port++ {
		conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			conn.Close()
			startPort = port
			break
		}
	}

	return strconv.Itoa(startPort)
}
