package main

import (
	"fmt"
	"net/http"

	"alles/boxes/env"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("DatabaseUrl = " + env.DatabaseUrl)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	fmt.Println("starting http server on :3000")
	http.ListenAndServe(":3000", r)
}
