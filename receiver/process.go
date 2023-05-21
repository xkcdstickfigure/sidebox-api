package receiver

import (
	"context"
	"fmt"
	"io"
	"strings"

	"alles/boxes/modules/email"
	"alles/boxes/store"
)

func process(ctx context.Context, db store.Store, body io.Reader) error {
	// get email data
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	bodyString := strings.Join(strings.Split(string(bodyBytes), "\n")[1:], "\n")

	// parse email message
	message, err := email.Parse(strings.NewReader(bodyString))
	if err != nil {
		return err
	}

	// get inbox
	to := message.Header.Get("Delivered-To")
	toUsername := strings.ToLower(strings.Split(to, "@")[0])
	inbox, err := db.InboxGetByCode(ctx, toUsername)
	if err != nil {
		return err
	}

	fmt.Println(inbox.Name)
	return nil
}
