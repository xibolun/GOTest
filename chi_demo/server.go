package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

// server API s
type server struct {
	handler http.Handler
}

// NewHandler 实例化业务处理器
func NewHandler() (http.Handler, error) {

	r := chi.NewRouter()



	return &server{handler: r}, nil
}

func (server *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.handler.ServeHTTP(w, r)
}
