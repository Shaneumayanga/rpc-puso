package main

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

type Gosession struct {
	Name    string //cookie name
	Options *Options
	Value   map[interface{}]interface{}
	Store   Store
}

func Newgosession(name string, options *Options, hashKey []byte, blockKey []byte) *Gosession {
	sc := securecookie.New([]byte(hashKey), nil)
	return &Gosession{
		Value:   make(map[interface{}]interface{}),
		Name:    name,
		Options: options,
		Store:   NewCookieStore(sc),
	}
}

func (gosession *Gosession) Save(r *http.Request, w http.ResponseWriter) {
	gosession.Store.Save(r, w, gosession)
}

func (gosession *Gosession) Get(r *http.Request, name string) map[interface{}]interface{} {
	return gosession.Store.Get(r, name)
}
