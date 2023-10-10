package views

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	// fake HTTP response writer for testing
	rw := httptest.NewRecorder()

	// fake HTTP request for testing
	req, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the ErrorHandler function with a custom error message
	errorMessage := "Test error message, from the test case"
	ErrorHandler(rw, req, errorMessage)

	// Check the HTTP response status code
	if rw.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, rw.Code)
	}

	// Check the response body for the error message
	if !strings.Contains(rw.Body.String(), errorMessage) {
		t.Errorf("Expected response body:\n%s\nbut got:\n%s", errorMessage, rw.Body.String())
	}
}
