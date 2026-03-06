package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main()  {
	var host string
		fmt.Printf("Enter the website or IP to scan: ")
		fmt.Scanln(&host)
	var openPorts[]int

	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)

		go func(j int) {

			port := fmt.Sprintf("%d", j)
			address := host + ":" + port
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err != nil {
				return
			}
			openPorts = append(openPorts, j)
			conn.Close()
			fmt.Printf("Port %d is OPEN on %s\n", j, host)

		}(i)
	}
	wg.Wait()
	fmt.Printf("Scan complete! Open ports: %v\n", openPorts)


}