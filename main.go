package main

import (
	"bufio"
	"fmt"
	"launchdarklytest/db"
	"launchdarklytest/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting LaunchDarkly test app")
	// curl https://live-test-scores.herokuapp.com/scores
	url := "https://live-test-scores.herokuapp.com/scores"

	updates := make(chan string)

	// Start a goroutine to continuously fetch score updates
	go fetchLiveScores(url, updates)

	// Listen for updates
	go func() {
		for {
			select {
			case update := <-updates:
				db.AddScore(update)
			}
		}
	}()

	// setup request handler
	router := mux.NewRouter()
	handler.SetRoutes(router)
	go func() {
		err := http.ListenAndServe(":80", router)
		if err != nil {
			fmt.Println("Error starting HTTP server:", err)
		}
	}()

	// Keep the main goroutine alive
	select {}
}

func fetchLiveScores(url string, updates chan string) {
	// Create an HTTP client
	client := &http.Client{}

	// Create a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Perform the streaming request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Create a buffered reader to read data incrementally
	reader := bufio.NewReader(resp.Body)

	// Main loop to continuously read and process data
	for {
		// Read a line of data until new line character
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading data:", err)
			break
		}

		// Process the score data. Only pass back the data line. It looks like this:
		// event: score
		// data: {"studentId":"Hannah.Herman76","exam":14102,"score":0.6815787969197457}
		if len(line) > 6 {
			if line[:5] == "data:" {
				updates <- line[6:]
			}
		}
	}
}
