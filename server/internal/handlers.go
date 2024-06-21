package internal

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/Makepad-fr/gam"
	"github.com/Makepad-fr/vif.io/model"
	usernameverifierservice "github.com/Makepad-fr/vif.io/services/username-verifier-service"
)

// trimTrailingSlashes remove the trailing slash characters from the beginning and the end of the string
func trimTrailingSlashes(s string) string {
	const slash = "/"
	return strings.TrimSuffix(strings.TrimPrefix(s, slash), slash)
}

func CreateRootHandler() http.Handler {
	log.Println("Creating root handelr")
	// TODO: Move username, passoword, hostname, port, database name and table name to environment variables
	g, err := gam.Init("custom_user", "custom_password", "clickhouse", "9000", "custom_database", "usage_analytics", true, true)
	if err != nil {
		log.Println("Error while intiaing gam")
		log.Fatalf("Error while initialising analytics middleware instance: %v", err)
	}
	return g.Middleware(LoggingMiddleWare(http.HandlerFunc(handleRootPath)))
}

// Middleware function
func handleRootPath(w http.ResponseWriter, r *http.Request) {
	parsedPath := strings.Split(trimTrailingSlashes(r.URL.Path), "/")
	l := len(parsedPath)
	switch {
	case l == 0:
		landingPageHandler(w, r)
	case l == 1:
		// TODO: Handle favicon.ico
		// TODO: Handle shortened URL
		// If there's only one item it's either a user's profile or a shortened link
		username := parsedPath[0]
		log.Printf("Username %s", username)
		user, exists := usernameverifierservice.GetUserDetails(username)
		if exists {
			links := usernameverifierservice.GetUserLinks(username)
			// If the user exists call the next handler
			existingUserProfileHandler(w, r, user, links)
			return
		}
		createYoursHandler(w, r, username)
	case l == 2:
		// A user's shortened URL /kaanyagci/gh
		// TODO: Get the associated link and redirect it
	default:
		notFoundPageHandler(w, r)
	}
}

func landingPageHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createYoursHandler(w http.ResponseWriter, _ *http.Request, username string) {
	// Username should exists here
	tmpl, err := template.ParseFiles("templates/create-yours.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, map[string]string{
		"Username": username,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func notFoundPageHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type userProfileTemplate struct {
	Details model.UserDetails
	Links   []model.Link
}

// Render template
func existingUserProfileHandler(w http.ResponseWriter, _ *http.Request, user model.UserDetails, links []model.Link) {
	tmpl, err := template.ParseFiles("templates/user-profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := userProfileTemplate{
		Details: user,
		Links:   links,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
