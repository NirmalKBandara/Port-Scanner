package main

import (
	"fmt"
	"net"
	"time"
)

func main()  {
	var host string
		fmt.Printf("Enter the website or IP to scan: ")
		fmt.Scanln(&host)
	var openPorts[]int

	for i := 1; i <= 100; i++ {
		port := fmt.Sprintf("%d", i)
		address := host + ":" + port
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			// fmt.Printf("Port %s is CLOSED on %s\n", port, host)
			continue
		}
		conn.Close()
		openPorts = append(openPorts, i)
		// fmt.Printf("Port %s is OPEN on %s\n", port, host)
	}
	fmt.Printf("Scan complete! Open ports: %v\n", openPorts)


}