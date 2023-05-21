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

	// get inboxes
	inboxes, err := h.db.InboxList(r.Context(), account.Id)
	if err != nil {
		apierr.Respond(w, apierr.DatabaseError)
		return
	}

	// response
	type resInbox struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	}

	resInboxes := []resInbox{}
	for _, i := range inboxes {
		resInboxes = append(resInboxes, resInbox{
			Id:   i.Id,
			Name: i.Name,
			Code: i.Code,
		})
	}

	respond(w, struct {
		Id      string     `json:"id"`
		Email   string     `json:"email"`
		Inboxes []resInbox `json:"inboxes"`
	}{
		Id:      account.Id,
		Email:   account.Email,
		Inboxes: resInboxes,
	})
}
