package main

import (
	"fmt"
	"net/http"
)

type Middleware struct {
	Mw func(next http.Handler) http.Handler
}

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware")
		next.ServeHTTP(w, r)
	})
}
