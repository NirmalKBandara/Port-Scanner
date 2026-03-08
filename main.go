package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	"strings"
)

func worker(host string, ports, results chan int) {
	for p := range ports {
		address := net.JoinHostPort(host, strconv.Itoa(p))
		conn, err := net.DialTimeout("tcp", address, 1000*time.Millisecond)

		if err != nil {
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	var host string
	fmt.Printf("Enter the website or IP to scan: ")
	fmt.Scanln(&host)
	host = strings.TrimSpace(host)

	var openPorts []int
	var wg sync.WaitGroup

	ports := make(chan int, 100)
	results := make(chan int)

	go func() {
		for port := range results {
			openPorts = append(openPorts, port)
			fmt.Printf("Port %d is OPEN on %s\n", port, host)
		}
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(host, ports, results)
		}()
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
		close(ports)
	}()

	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	close(results)

	fmt.Printf("Scan complete! Open ports: %v\n", openPorts)
}
