package get_info

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetArtistInfoFromAPI(t *testing.T) {
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

func TestGetRelationInfoFromAPI(t *testing.T) {
	testJSONResponse := getRelations()

	// Create a test HTTP server that serves the test JSON response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testJSONResponse))
	}))
	defer server.Close()

	var relations []Relation

	// Call GetInfoFromAPI with the test server URL
	err := GetInfoFromAPI(server.URL, &relations)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Perform assertions on the relations slice
	if len(relations) != 3 {
		t.Errorf("Expected 3 relations, but got %d", len(relations))
	}

	// Check the value of the first relation
	expectedDate := "10-02-2020"
	actualDate := relations[0].DatesLocation["dunedin-new_zealand"][0]
	if actualDate != expectedDate {
		t.Errorf("Unexpected JSON parsing result, got %s, expected %s", actualDate, expectedDate)
	}
}
