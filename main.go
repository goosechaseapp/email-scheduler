package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	// query rows
	rows, err := conn.Query(ctx, "SELECT id, email, is_sent FROM scheduled_emails WHERE scheduled_at < now() AND is_sent = FALSE")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var email string
		err := rows.Scan(&id, &email)
		if err != nil {
			panic(err)
		}

		println(id, email)
	}

}
