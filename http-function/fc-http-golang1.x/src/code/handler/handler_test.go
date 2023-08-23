package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProcessHandler(t *testing.T) {
	testCases := []struct {
		name           string
		method         string
		payload        string
		expectedStatus int
	}{
		{
			name:           "Valid JSON",
			method:         http.MethodPost,
			payload:        `{"id": "recMzWmyCs", "address": "test address", "package": {"pkgNo": "73506545530230"}}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			payload:        `{"id": recMzWmyCs, "address": test address, "package": {"pkgNo": 73506545530230}}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/", strings.NewReader(tc.payload))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				err := HandleHttpRequest(context.Background(), w, req)
				if err != nil {
					http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
				}
			})

			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, but got %d", tc.expectedStatus, rr.Code)
			}

			if tc.expectedStatus == http.StatusOK {
				if !json.Valid(rr.Body.Bytes()) {
					t.Errorf("Error decoding response JSON: %v", err)
				}
			}
		})
	}
}
