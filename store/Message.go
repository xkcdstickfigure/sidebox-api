package store

import (
	"context"
	"time"
)

type Message struct {
	Id          string
	InboxId     string
	MessageId   string
	FromName    string
	FromAddress string
	Subject     string
	Body        string
	Html        bool
	Date        time.Time
}

func (s Store) MessageCreate(ctx context.Context, data Message) error {
	_, err := s.Conn.Exec(ctx, "insert into message (id, inbox_id, message_id, from_name, from_address, subject, body, html, date) values ($1, $2, $3, $4, $5, $6, $7, $8, now())",
		data.Id, data.InboxId, data.MessageId, data.FromName, data.FromAddress, data.Subject, data.Body, data.Html)
	return err
}
