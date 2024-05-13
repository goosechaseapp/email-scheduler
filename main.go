package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	pb "google.golang.org/protobuf/proto"
	"goosechase.ai/email-scheduler/proto/proto"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	// query rows
	rows, err := conn.Query(ctx, "SELECT id, msg FROM scheduled_emails WHERE scheduled_at < now() AND is_sent = FALSE")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var email []byte
		err := rows.Scan(&id, &email)
		if err != nil {
			panic(err)
		}

		var emailMessage proto.SendEmailDocument
		err = pb.Unmarshal(email, &emailMessage)

		if err != nil {
			panic(err)
		}

		println("Sending email to", emailMessage.Subject)
	}

}
