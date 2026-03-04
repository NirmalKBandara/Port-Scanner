package main

import (
	"fmt"
	"net"
	"time"
)

func main()  {
	host := "scanme.nmap.org"
	port := "80"
	address := host + ":" + port
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		fmt.Printf("Port %s is CLOSED on %s\n", port, host)
		return
	}

	defer conn.Close()
	fmt.Printf("Port %s is OPEN on %s\n", port, host)
}