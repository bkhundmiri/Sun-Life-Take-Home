package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type SiteStatus struct {
	URL        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Duration   int64  `json:"duration"`
	Date       string `json:"date"`
}

func checkStatus(url string) (SiteStatus, error) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching: %s", url)
		return SiteStatus{}, fmt.Errorf("Error fetching: %s", url)
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

func amazonStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := checkStatus("https://www.amazon.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(status)
}

func googleStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := checkStatus("https://www.google.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(status)
}

func allStatusHandler(w http.ResponseWriter, r *http.Request) {
	googleStatus, err1 := checkStatus("https://www.google.com")
	amazonStatus, err2 := checkStatus("https://www.amazon.com")
	if err1 != nil || err2 != nil {
		errMsg := ""
		if err1 != nil {
			errMsg += err1.Error()
		}
		if err2 != nil {
			errMsg += " " + err2.Error()
		}
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	statuses := []SiteStatus{googleStatus, amazonStatus}
	json.NewEncoder(w).Encode(statuses)
}

func main() {
	http.HandleFunc("/v1/amazon-status", amazonStatusHandler)
	http.HandleFunc("/v1/google-status", googleStatusHandler)
	http.HandleFunc("/v1/all-status", allStatusHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", nil)
}
