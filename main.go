package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
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
	var hostFlag string
	var portsFlag string
	var threadsFlag int

	flag.StringVar(&hostFlag, "host", "127.0.0.1", "Target IP or Domain (e.g., scanme.nmap.org)")
	flag.StringVar(&portsFlag, "ports", "1-1024", "Port range to scan (e.g., 1-1024 or just 80)")
	flag.IntVar(&threadsFlag, "threads", 100, "Number of concurrent workers")
	flag.Parse()

	host := strings.TrimSpace(hostFlag)
	if host == "" {
		fmt.Println("Error: Host cannot be empty.")
		return
	}

	startPort, endPort := parsePorts(portsFlag)
	if startPort == 0 || endPort == 0 || startPort > endPort {
		fmt.Println("Error: Invalid port range specified.")
		return
	}

	fmt.Printf("Scanning %s on ports %d to %d using %d threads...\n", host, startPort, endPort, threadsFlag)

	var openPorts []int
	var wg sync.WaitGroup

	ports := make(chan int, threadsFlag)
	results := make(chan int)

	go func() {
		for port := range results {
			openPorts = append(openPorts, port)
			fmt.Printf("Port %d is OPEN on %s\n", port, host)
		}
	}()

	for i := 0; i < threadsFlag; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(host, ports, results)
		}()
	}

	go func() {
		for i := startPort; i <= endPort; i++ {
			ports <- i
		}
		close(ports)
	}()

	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	close(results)

	fmt.Printf("Scan complete! Open ports: %v\n", openPorts)
}

// parsePorts converts a string like "1-1024" or "80" into two integers.
func parsePorts(p string) (int, int) {
	if strings.Contains(p, "-") {
		parts := strings.Split(p, "-")
		if len(parts) == 2 {
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				return start, end
			}
		}
	} else {
		port, err := strconv.Atoi(p)
		if err == nil {
			return port, port
		}
	}
	return 0, 0
}
