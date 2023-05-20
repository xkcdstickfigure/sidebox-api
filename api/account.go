package api

import (
	"net/http"

	"alles/boxes/api/apierr"
)

// GET /account
func (h handlers) account(w http.ResponseWriter, r *http.Request) {
	// session
	session, err := h.auth(r)
	if err != nil {
		apierr.Respond(w, apierr.BadAuthorization)
		return
	}

	// get account
	account, err := h.db.AccountGet(r.Context(), session.AccountId)
	if err != nil {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// response
	respond(w, struct {
		Id    string `json:"id"`
		Email string `json:"email"`
	}{
		Id:    account.Id,
		Email: account.Email,
	})
}
