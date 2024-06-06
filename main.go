package main

import (
	"net/http"

	"github.com/HMasataka/ogen/api"
	"github.com/HMasataka/ogen/middleware"
	"github.com/HMasataka/ogen/ogen"
	chi "github.com/go-chi/chi/v5/middleware"
)

func newHTTPServer() *http.Server {
	service := &api.PetService{}

	s, err := ogen.NewServer(service)
	if err != nil {
		panic(err)
	}

	return &http.Server{
		Addr:    ":8080",
		Handler: middleware.Wrap(s, chi.Logger),
	}
}

func main() {
	server := newHTTPServer()
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
