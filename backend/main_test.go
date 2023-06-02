package main

import (
	"net/http"
	"testing"
)

func TestFetchStatus(t *testing.T) {
	testCases := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "Google",
			url:     "https://www.google.com",
			wantErr: false,
		},
		{
			name:    "Amazon",
			url:     "https://www.amazon.com",
			wantErr: false,
		},
		{
			name:    "Invalid URL",
			url:     "http://",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := fetchStatus(tc.url)

			if tc.wantErr {
				if err == nil {
					t.Errorf("fetchStatus() for url %v didn't return error", tc.url)
				}
			} else {
				if err != nil {
					t.Errorf("fetchStatus() returned an error: %v", err)
				}
				if status.StatusCode != http.StatusOK {
					t.Errorf("fetchStatus() = %v, want %v", status.StatusCode, http.StatusOK)
				}
				if status.URL != tc.url {
					t.Errorf("fetchStatus() returned status for URL = %v, want %v", status.URL, tc.url)
				}
			}
		})
	}
}
