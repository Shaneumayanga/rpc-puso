package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

var (
	reply    Pusa
	database []Pusa
)

type Handler struct {
	client     *rpc.Client
	sessions   *Gosession
	Middleware []Middleware
}

func (h *Handler) Use(mw Middleware) {
	h.Middleware = append(h.Middleware, mw)
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.Use(Middleware{Mw: TestMiddleware})
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		values := h.sessions.Get(r, "session-cookie")
		if values["isAuth"] != nil {
			if !values["isAuth"].(bool) {
				http.Redirect(rw, r, "/login", http.StatusSeeOther)
				return
			}
		} else {
			http.Redirect(rw, r, "/login", http.StatusSeeOther)
		}

		if err := h.client.Call("API.DumpDB", "", &database); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("database: %v\n", database)
		json.NewEncoder(rw).Encode(&database)
	}
	if r.Method == http.MethodPost && r.URL.Path == "/postpusa" {
		p := &Pusa{}
		if err := json.NewDecoder(r.Response.Body).Decode(p); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		if err := h.client.Call("API.Save", p, &reply); err != nil {
			log.Fatal(err)
		}
		return
	}
}

func StartClient() {
	options := &Options{
		Path:     "/",
		MaxAge:   200,
		Secure:   true,
		HttpOnly: true,
	}
	gosession := Newgosession("session-cookie", options, []byte("hashKey"), []byte("blockKey"))
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	h := &Handler{
		sessions: gosession,
		client:   client,
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	fmt.Println("Client Started on port :8080")
	log.Fatal(server.ListenAndServe())
}
