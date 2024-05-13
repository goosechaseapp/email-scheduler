package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	pb "google.golang.org/protobuf/proto"
	"goosechase.ai/email-scheduler/config"
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
	conn, err := pgx.Connect(ctx, config.Env("DB_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

	// Initialize Kafka
	services.Kafka.InitializeConfluentKafka()

	// query rows
	rows, err := conn.Query(ctx, "SELECT id, msg FROM scheduled_emails WHERE scheduled_at < now() AND is_sent = FALSE")

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to query rows")
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var email []byte
		err := rows.Scan(&id, &email)
		if err != nil {
			log.Error().Err(err).Msg("Failed to scan row")
			continue
		}

		log.Info().Msg("Processing job id " + id)

		var emailMessage proto.SendEmailDocument
		err = pb.Unmarshal(email, &emailMessage)

		// reset schedule time
		emailMessage.ScheduledTime = 0

		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal email")
			continue
		}

		services.Kafka.ProduceEmail(&emailMessage)
	}
}
