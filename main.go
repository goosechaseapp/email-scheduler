package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	pb "google.golang.org/protobuf/proto"
	"goosechase.ai/email-scheduler/config"
	"goosechase.ai/email-scheduler/proto/proto"
	"goosechase.ai/email-scheduler/services"
	"goosechase.ai/email-scheduler/util/log"
)

func main() {
	log.Initialize()

	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load .env file")
	}

	ctx := context.Background()

	// Initialize Postgres connection
	conn, err := pgxpool.New(ctx, config.Env("DB_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	defer conn.Close()

	log.Info().Msg("Connected to database")

	if err := conn.Ping(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

	log.Info().Msg("Pinged database")

	// Initialize Kafka
	services.Kafka.InitializeConfluentKafka()

	log.Info().Msg("Initialized Kafka")

	// query rows
	rows, err := conn.Query(ctx, "SELECT id, msg FROM scheduled_emails WHERE scheduled_at < now() AND is_sent = FALSE")

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to query rows")
	}

	log.Info().Msg("Queried rows")

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

		err = services.Kafka.ProduceEmail(&emailMessage)
		if err != nil {
			log.Error().Err(err).Msg("Job ID " + id + " failed to produce email")
			continue
		}

		// set is_sent to true
		_, err = conn.Exec(ctx, "UPDATE scheduled_emails SET is_sent = TRUE, sent_at = now() where ID=$1 ", id)
		if err != nil {
			log.Error().Err(err).Msg("Failed to update is_sent id " + id)
			continue
		}
	}
}
