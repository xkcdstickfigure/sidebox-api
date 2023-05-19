package store

import (
	"context"
	"time"
)

type Account struct {
	Id        string
	Name      string
	Email     string
	GoogleId  string
	CreatedAt time.Time
}

func (s Store) AccountGet(ctx context.Context, id string) (Account, error) {
	var account Account
	err := s.Conn.QueryRow(ctx, "select id, name, email, google_id, created_at where id=$1", id).
		Scan(&account.Id, &account.Name, &account.Email, &account.GoogleId, &account.CreatedAt)
	return account, err
}
