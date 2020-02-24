package main

import (
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/rpc-starter-kit/internal/server/twirp"
	"github.com/pepeunlimited/rpc-starter-kit/pkg/rpc/todo"
	"log"
	"net/http"
)

const (
	Version = "0.0.1"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Printf("Starting the TodoServer... version=[%v]", Version)

	ts := todo.NewTodoServiceServer(twirp.NewTodoServer(), nil)

	mux := http.NewServeMux()
	mux.Handle(ts.PathPrefix(), middleware.Adapt(ts, headers.Username()))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panic(err)
	}
}