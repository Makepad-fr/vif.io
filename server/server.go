package main

import (
	"fmt"
	"net/http"

	"github.com/Makepad-fr/vif.io/server/internal"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve HTML templates
	http.Handle("/", internal.LoggingMiddleWare(internal.HandleRootPath))

	// Start the server
	// TODO: Make port number from envnrionment variables
	fmt.Println("Starting server at :4321")
	err := http.ListenAndServe(":4321", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
