package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"alles/boxes/api/apierr"
)

// POST /inbox
func (h handlers) inboxCreate(w http.ResponseWriter, r *http.Request) {
	// parse body
	var body struct {
		Name string
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		apierr.Respond(w, apierr.InvalidBody)
		return
	}

	// session
	session, err := h.auth(r)
	if err != nil {
		apierr.Respond(w, apierr.BadAuthorization)
		return
	}

	// create inbox
	inbox, err := h.db.InboxCreate(r.Context(), session.AccountId, strings.TrimSpace(body.Name))
	if err != nil {
		apierr.Respond(w, apierr.DatabaseError)
	}

	respond(w, struct {
		Id   string `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	}{
		Id:   inbox.Id,
		Code: inbox.Code,
		Name: inbox.Name,
	})
}
