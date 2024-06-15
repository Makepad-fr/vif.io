package internal

import (
	"log"
	"net/http"
)

func LoggingMiddleWare(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Get as much information as possible from request
		log.Println("Hello from logging middleware")
		next.ServeHTTP(w, r)
	})
}
