package main

import (
	"testing"
)

func TestCheckStatus(t *testing.T) {
	type testCase struct {
		url     string
		wantErr bool
	}

	testCases := []testCase{
		{
			url:     "https://www.google.com",
			wantErr: false,
		},
		{
			url:     "https://www.amazon.com",
			wantErr: false,
		},
		{
			url:     "http://",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		_, err := fetchStatus(tc.url)
		if (err != nil) != tc.wantErr {
			t.Errorf("fetchStatus(%v) returned error %v, wantErr %v", tc.url, err, tc.wantErr)
		}
	}
}
