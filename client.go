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
	client *rpc.Client
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		p := Pusa{
			Name:     "Bole",
			Laziness: 12,
		}
		if err := h.client.Call("API.Save", p, &reply); err != nil {
			log.Fatal(err)
		}

		if err := h.client.Call("API.DumpDB", "", &database); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("database: %v\n", database)
		json.NewEncoder(rw).Encode(&database)
	}
}

func StartClient() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	h := &Handler{
		client: client,
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	fmt.Println("Client Started on port :8080")
	log.Fatal(server.ListenAndServe())
}
