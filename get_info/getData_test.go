package get_info

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetInfoFromAPI(t *testing.T) {
	testJSONResponse := getArtists()

	// Create a test HTTP server that serves the test JSON response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testJSONResponse))
	}))
	defer server.Close()

	var artists []Artist

	// Call GetInfoFromAPI with the test server URL
	err := GetInfoFromAPI(server.URL, &artists)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Perform assertions on the artists slice
	if len(artists) != 3 {
		t.Errorf("Expected 3 artists, but got %d", len(artists))
	}

	//Check the name of the first artist
	if artists[0].Name != "Queen" {
		t.Errorf("Expected first artist name to be Queen, but got %s", artists[0].Name)
	}
}
