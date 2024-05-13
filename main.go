package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	pb "google.golang.org/protobuf/proto"
	"goosechase.ai/email-scheduler/proto/proto"
	"goosechase.ai/email-scheduler/services"
	"goosechase.ai/email-scheduler/util/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load .env file")
	}

	ctx := context.Background()

	// Initialize Postgres connection
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	// Initialize Kafka
	services.Kafka.InitializeConfluentKafka()

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

		services.Kafka.ProduceEmail(&emailMessage)
	}
}
