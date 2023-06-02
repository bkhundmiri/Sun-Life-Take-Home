package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type SiteStatus struct {
	URL        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Duration   int64  `json:"duration"`
	Date       string `json:"date"`
}

var statuses map[string]SiteStatus
var statusMutex sync.Mutex
var urls = []string{"https://www.google.com", "https://www.amazon.com"}

func fetchStatus(url string) (SiteStatus, error) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching: %s", err)
		return SiteStatus{}, err
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	date := time.Now().Format("2006-01-02T15:04:05Z07:00")

	return SiteStatus{
		URL:        url,
		StatusCode: resp.StatusCode,
		Duration:   duration.Milliseconds(),
		Date:       date,
	}, nil
}

func checkAllStatuses() {
	for {
		newStatuses := make(map[string]SiteStatus)

		for _, url := range urls {
			status, err := fetchStatus(url)
			if err != nil {
				log.Printf("Error fetching: %s", err)
				continue
			}
			newStatuses[url] = status
		}

		// Update the shared variable
		statusMutex.Lock()
		statuses = newStatuses
		statusMutex.Unlock()

		// Sleep for 1 minute
		time.Sleep(1 * time.Minute)
	}
}

func amazonStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := fetchStatus("https://www.amazon.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(status)
}

func googleStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := fetchStatus("https://www.google.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(status)
}

func allStatusHandler(w http.ResponseWriter, r *http.Request) {
	statusMutex.Lock()
	defer statusMutex.Unlock()

	var allStatuses []SiteStatus
	for _, status := range statuses {
		allStatuses = append(allStatuses, status)
	}

	json.NewEncoder(w).Encode(allStatuses)
}

func main() {
	go checkAllStatuses()
	http.HandleFunc("/v1/amazon-status", amazonStatusHandler)
	http.HandleFunc("/v1/google-status", googleStatusHandler)
	http.HandleFunc("/v1/all-status", allStatusHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", nil)
}
