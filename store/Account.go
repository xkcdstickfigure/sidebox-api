package store

import (
	"context"
	"time"

	"github.com/google/uuid"
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

func (s Store) AccountCreate(ctx context.Context, name string, email string, googleId string) (Account, error) {
	var account Account
	err := s.Conn.QueryRow(ctx, "insert into account (id, name, email, google_id, created_at) "+
		"values ($1, $2, $3, $4, now()) "+
		"on conflict (google_id) do update set name=$2, email=$3 "+
		"returning id, name, email google_id, created_at",
		uuid.New(), name, email, googleId).
		Scan(&account.Id, &account.Name, &account.Email, &account.CreatedAt)
	return account, err
}
