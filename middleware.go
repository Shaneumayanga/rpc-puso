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

//lol

// func IsAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		values := sessions.Get(r, "session-cookie")
// 		if values["isAuth"] != nil {
// 			if !values["isAuth"].(bool) {
// 				http.Redirect(rw, r, "/login", http.StatusSeeOther)
// 				return
// 			}
// 		} else {
// 			http.Redirect(rw, r, "/login", http.StatusSeeOther)
// 		}
// 	})
// }
