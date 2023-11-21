package main

import (
	"fmt"
	"net/http"
	"time"
)

func pingWebsite(url string, ch chan<- string) {
	startTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Failed to ping %s: %s", url, err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(startTime)

	ch <- fmt.Sprintf("%s is online (Response Time: %s)", url, elapsed)
}

func main() {
	websites := []string{"http://www.google.com", "http://www.example.com", "http://www.nonexistent-website.com"}

	resultChannel := make(chan string)

	for _, url := range websites {
		go pingWebsite(url, resultChannel)
	}

	for i := 0; i < len(websites); i++ {
		result := <-resultChannel
		fmt.Println(result)
	}

	close(resultChannel)
}
