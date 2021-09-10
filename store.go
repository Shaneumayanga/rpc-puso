package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

type Values map[interface{}]interface{}

type Store interface {
	Save(r *http.Request, w http.ResponseWriter, session *Gosession)
	Get(r *http.Request, name string) Values
}

type CookieStore struct {
	securecookie *securecookie.SecureCookie
}

func NewCookieStore(securecookie *securecookie.SecureCookie) *CookieStore {
	return &CookieStore{
		securecookie: securecookie,
	}
}

func (cs *CookieStore) Save(r *http.Request, w http.ResponseWriter, session *Gosession) {
	encoded, err := cs.securecookie.Encode(session.Name, session.Value)
	if err == nil {
		cookie := NewCookie(session.Name, encoded, session.Options)
		http.SetCookie(w, cookie)
		return
	}
	fmt.Printf("err: %v\n", err)

}

func (cs *CookieStore) Get(r *http.Request, name string) Values {
	if cookie, err := r.Cookie(name); err == nil {
		values := make(Values)
		if err = cs.securecookie.Decode(name, cookie.Value, &values); err == nil {
			return values
		}
	}
	return nil
}
