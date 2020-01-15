package main

import (
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/rpc-starter-kit/internal/app/app1/server"
	"github.com/pepeunlimited/rpc-starter-kit/rpctodo"
	"log"
	"net/http"
)

const (
	Version = "0.1"
)

func main() {
	log.Printf("Starting the TodoServer... version=[%v]", Version)

	ts := rpctodo.NewTodoServiceServer(server.NewTodoServer(), nil)

	mux := http.NewServeMux()
	mux.Handle(ts.PathPrefix(), middleware.Adapt(ts, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}