package main

import (
	"fmt"
	"net/http"

	//pakket voor het maken van HTTP verzoeken
	"time"
	//pakket voor het meten van de tijd bij de ms response time
)

func pingWebsite(url string, ch chan<- string) {
	startTime := time.Now()
	//Hier wordt de huidige tijd vastgesteld dit wordt later gebruikt om de totale tijd te berekenen voor een HTTP-verzoek.

	resp, err := http.Get(url)
	//zorgt ervoor dat er een get verzoek gedaan wordt naar de URL.
	if err != nil {
		ch <- fmt.Sprintf("Failed to ping %s: %s", url, err)
		return
	}
	defer resp.Body.Close()
	//Defer zorgt ervoor dat de aanroep wordt uitgesteld tot het einde van de functie.
	//Als het verzoek niet gesloten wordt krijg je resource leaks. Omdat de aanvraag open blijft.

	elapsed := time.Since(startTime)
	//Dit berekent de verstreken tijd sinds startTime.

	ch <- fmt.Sprintf("%s is online (Response Time: %s)", url, elapsed)
	//Hier wordt geschreven naar de display dat de website online is en de responsetijd komt erachter.
}

func main() {
	websites := []string{"http://www.google.com", "http://www.example.com", "https://www.meter.net/ping-test/"}

	resultChannel := make(chan string)
	//string waar resultaten neer wordt gezet.

	for _, url := range websites {
		go pingWebsite(url, resultChannel)
	}
	//Functie ping website wordt aangeroepen om elke website die aangegeven is te pingen.

	for i := 0; i < len(websites); i++ {
		result := <-resultChannel
		fmt.Println(result)
	}
	//Hier wordt data uit resultChannel ontvangen en afgedrukt naar scherm

	close(resultChannel)
}

//Het kanaal wordt gesloten nadat alle resultaten zijn ontvangen.
