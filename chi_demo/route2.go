package main

import (
	"github.com/go-chi/chi"
	"fmt"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		w.Write([]byte(fmt.Sprintf("welcome, %s", id)))
	})
	http.ListenAndServe(":3000", r)
}
