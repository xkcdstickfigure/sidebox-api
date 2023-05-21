package store

import (
	"context"
	"time"

	"alles/boxes/modules/random"

	"github.com/google/uuid"
)

type Inbox struct {
	Id        string
	AccountId string
	Code      string
	Name      string
	CreatedAt time.Time
}

func (s Store) InboxGet(ctx context.Context, id string) (Inbox, error) {
	var inbox Inbox
	err := s.Conn.QueryRow(ctx, "select id, account_id, code, name, created_at from inbox where id=$1", id).
		Scan(&inbox.Id, &inbox.AccountId, &inbox.Code, &inbox.Name, &inbox.CreatedAt)
	return inbox, err
}

func (s Store) InboxGetByCode(ctx context.Context, code string) (Inbox, error) {
	var inbox Inbox
	err := s.Conn.QueryRow(ctx, "select id, account_id, code, name, created_at from inbox where code=$1", code).
		Scan(&inbox.Id, &inbox.AccountId, &inbox.Code, &inbox.Name, &inbox.CreatedAt)
	return inbox, err
}

func (s Store) InboxCreate(ctx context.Context, accountId string, name string) (Inbox, error) {
	var inbox Inbox
	err := s.Conn.QueryRow(ctx, "insert into inbox (id, account_id, code, name, created_at) "+
		"values ($1, $2, $3, $4, now()) "+
		"returning id, account_id, code, name, created_at",
		uuid.New(), accountId, random.String(16), name).
		Scan(&inbox.Id, &inbox.AccountId, &inbox.Code, &inbox.Name, &inbox.CreatedAt)
	return inbox, err
}

func (s Store) InboxList(ctx context.Context, accountId string) ([]Inbox, error) {
	inboxes := []Inbox{}
	rows, err := s.Conn.Query(ctx, "select id, account_id, code, name, created_at from inbox "+
		"where account_id=$1 "+
		"order by created_at desc",
		accountId)
	if err != nil {
		return inboxes, err
	}
	defer rows.Close()

	for rows.Next() {
		var inbox Inbox
		err = rows.Scan(&inbox.Id, &inbox.AccountId, &inbox.Code, &inbox.Name, &inbox.CreatedAt)
		if err != nil {
			return inboxes, err
		}
		inboxes = append(inboxes, inbox)
	}
	return inboxes, err
}
