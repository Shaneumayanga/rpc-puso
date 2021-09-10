package main

import (
	"fmt"
	"net/http"
)

type CookieStore1 struct{}

func NewCookieStore1() *CookieStore1 {
	return &CookieStore1{}
}

func (cs *CookieStore1) Save(r *http.Request, w http.ResponseWriter, session *Gosession) {
	cookie := NewCookie(session.Name, "somevalue", session.Options)
	http.SetCookie(w, cookie)
}

func (cs *CookieStore1) Get(r *http.Request, name string, session *Gosession) map[interface{}]interface{} {
	cookie, err := r.Cookie(name)
	if err == nil {
		values := cookie.Value
		fmt.Printf("values: %v\n", values)
	}
	return nil
}
