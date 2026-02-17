package middleware

import (
	"log"
	"net/http"
	"time"
)
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timestamp := time.Now().Format("2006-01-02T15:04:05")
		log.Printf("%s %s %s {%s}", timestamp, r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}
}
func Chain(handler http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}