package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Serve static files
func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("static", r.URL.Path))
}

// isTrimmedEmpty returns true if the given string is empty after trimmed
func isTrimmedEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

func createUserProfileHandler(username string) httpHandlerFunc {
	// TODO: Validate username and get user details
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/user-profile.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// TODO: Update the map with a proper structure
		data := map[string]string{
			"Endpoint": username,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Render template
func templateHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[1:] // remove the leading slash
	log.Printf("Username %s", username)
	if isTrimmedEmpty(username) {
		fmt.Fprintf(w, "Landing page is not implemented yet")
		return
	}
	createUserProfileHandler(username)(w, r)
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve HTML templates
	http.HandleFunc("/", templateHandler)

	// Start the server
	// TODO: MAke port number from envnrionment variables
	fmt.Println("Starting server at :4321")
	err := http.ListenAndServe(":4321", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
