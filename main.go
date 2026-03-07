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

	results := make(chan int)
	
	go func() {
		for port := range results {
			openPorts = append(openPorts, port)
			fmt.Printf("Port %d is OPEN on %s\n", port, host)
		}
	}()

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			port := fmt.Sprintf("%d", j)
			address := host + ":" + port
			conn, err := net.DialTimeout("tcp", address, 1000*time.Millisecond)
			if err != nil {
				return
			}
			conn.Close()
			results <- j
		}(i)
	}
	wg.Wait()
	time.Sleep(100*time.Millisecond)
	close(results)

	fmt.Printf("Scan complete! Open ports: %v\n", openPorts)
}