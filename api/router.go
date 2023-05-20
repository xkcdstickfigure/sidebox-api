package api

import (
	"alles/boxes/store"

	"github.com/go-chi/chi/v5"
)

func NewRouter(db store.Store) chi.Router {
	r := chi.NewRouter()
	h := handlers{db}

	r.Post("/login", h.login)
	r.Get("/account", h.account)
	r.Post("/inbox", h.inboxCreate)

	return r
}

type handlers struct {
	db store.Store
}
