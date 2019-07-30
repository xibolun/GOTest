package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func registerRoutes(mux *chi.Mux) {
	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

}
