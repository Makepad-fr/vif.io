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

// Middleware function
func checkIfUserExistsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check if the user exists
		username := r.URL.Path[1:] // remove the leading slash
		log.Printf("Username %s", username)
		if isTrimmedEmpty(username) || username != "kaanyagci" {
			log.Println("Landing page is not implemented yet")
			// TODO: If the username does not exists in the path render the landing page template
			// TODO: If the username does not exists but present in the path: Redirect to the a page that says "the username is available create yours
			// http.Redirect(w, r, "/", http.StatusFound)
			fmt.Fprintf(w, "hello world")
			return
		}
		// Call the next handler if the condition is met
		next.ServeHTTP(w, r)
	})
}

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

// Render template
func existingUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Username should exists here
	username := r.URL.Path[1:] // remove the leading slash
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

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	existingUserProfile := http.HandlerFunc(existingUserProfileHandler)
	userProfile := checkIfUserExistsMiddleware(existingUserProfile)

	// Serve HTML templates
	http.Handle("/", userProfile)

	// Start the server
	// TODO: MAke port number from envnrionment variables
	fmt.Println("Starting server at :4321")
	err := http.ListenAndServe(":4321", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
