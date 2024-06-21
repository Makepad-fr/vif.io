package main

import (
	"fmt"
	"net/http"

	"github.com/Makepad-fr/vif.io/server/internal"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/", internal.CreateRootHandler())

	// TODO: MAke port number from envnrionment variables
	server := &http.Server{
		Addr:    ":4321",
		Handler: mux,
	}
	// Start the server
	fmt.Println("Starting server at :4321")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
