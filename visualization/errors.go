package visualization

import (
	"html/template"
	"net/http"
	"os"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, errorMessage string) {
	// Read and serve the error page HTML
	errorPageHTML, err := os.ReadFile("errorPage/error.html")
	if err != nil {
		http.Error(w, "Error page not found", http.StatusInternalServerError)
		return
	}

	// Create a template with the error message
	tmpl, err := template.New("errorPage").Parse(string(errorPageHTML))
	if err != nil {
		http.Error(w, "Error creating template", http.StatusInternalServerError)
		return
	}

	// Define a struct to hold the error message
	data := struct {
		ErrorMessage string
	}{
		ErrorMessage: errorMessage,
	}

	// Execute the template with the data and write it to the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
