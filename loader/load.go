package main 

import(
	"log"
	"net/http"
	"time"
)

func fireAndForget(url string) {
	log.Printf("Sending request to %s", url)
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
	}
}

func main() {
	url := "http://localhost:5203/" // Replace with your actual .NET API URL

	// Number of requests you want to send (adjust as needed for your load test)
	numRequests := 100000

	// Create a small delay between requests to avoid overloading instantly (optional)
	delayBetweenRequests := 750 * time.Microsecond
	// Start sending fire-and-forget requests in parallel
	for i := 0; i < numRequests; i++ {
		go fireAndForget(url)
		time.Sleep(delayBetweenRequests)
	}

	// Keep the program running to allow all requests to complete
	// You can adjust the sleep time based on the expected time for requests to finish
	time.Sleep(10 * time.Second)
}