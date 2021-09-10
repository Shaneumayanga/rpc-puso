package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func StartServer() {
	db := NewDataBase()
	api := NewAPI(db)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("Serving rpc on port :1234")
	http.Serve(l, nil)
}

func main() {
	go StartServer()
	StartClient()
}
