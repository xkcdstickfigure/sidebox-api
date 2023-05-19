package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"alles/boxes/env"
	"alles/boxes/store"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// connect to database
	conn, err := pgxpool.New(context.Background(), env.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}
	_ = store.Store{Conn: conn}

	// http server
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	fmt.Println("starting http server on :3000")
	http.ListenAndServe(":3000", r)
}
