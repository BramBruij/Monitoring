package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func pingHost(hostname string, ch chan<- string) {
	pinger, err := ping.NewPinger(hostname)
	if err != nil {
		ch <- fmt.Sprintf("Failed to create pinger for %s: %s", hostname, err)
		return
	}

	pinger.Count = 3 // Aantal ping-pakketten verzonden
	pinger.Timeout = time.Second * 2

	pinger.OnRecv = func(pkt *ping.Packet) {
		ch <- fmt.Sprintf("%s is online", hostname)
	}

	pinger.OnIdle = func() {
		ch <- fmt.Sprintf("%s is offline", hostname)
	}

	err = pinger.Run()
	if err != nil {
		ch <- fmt.Sprintf("Failed to ping %s: %s", hostname, err)
	}
}

func main() {
	hostnames := []string{"https://www.fontys.nl/Over-Fontys/Fontys-ICT.htm", "google.com", "nonexistent-host.com"}

	resultChannel := make(chan string)

	for _, hostname := range hostnames {
		go pingHost(hostname, resultChannel)
	}

	for i := 0; i < len(hostnames); i++ {
		result := <-resultChannel
		fmt.Println(result)
	}

	close(resultChannel)
}
